package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gatewaygRPCClient "api-gateway/client_grpc"
	restClient "api-gateway/client_rest"
	"api-gateway/configs"
	"api-gateway/controllers"
	gatewayHTTPHandler "api-gateway/http_handler"
	pb "api-gateway/pb"
	"api-gateway/routes"
	"api-gateway/services"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Handling CORS :D")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func init() {
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(configs.EnvMongoURI())
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	authCollection = mongoclient.Database("user_db").Collection("users")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService, ctx, authCollection)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:" + configs.EnvServicePort()}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// Use the main engine instance directly
	router := server.Group("/api")

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)

	flag.Parse()
	defer gracefulShutdown()

	// Initialize gRPC connections
	ordergRPCConn := initOrdergRPCConnection()
	ordergRPCClienter := pb.NewOrderServiceClient(ordergRPCConn)
	defer ordergRPCConn.Close()

	// Dependency Injection
	ordergRPCClient := gatewaygRPCClient.ProvideOrderClient(&ordergRPCClienter)
	orderHandler := gatewayHTTPHandler.ProvideOrderHandler(ordergRPCClient)
	reviewClientRest := restClient.ProvideReviewClientRest(&http.Client{})
	reviewHandler := gatewayHTTPHandler.ProvideReviewHandler(reviewClientRest)
	notificationClientRest := restClient.ProvideNotificationClientRest(&http.Client{})
	notificationHandler := gatewayHTTPHandler.ProvideNotificationHandler(notificationClientRest)
	menuClientRest := restClient.ProvideMenuClientRest(&http.Client{})
	menuHandler := gatewayHTTPHandler.ProvideMenuHandler(menuClientRest)

	// Use the main engine instance directly
	gatewayHTTPHandler.ProvideRouter(server,
		userService,
		orderHandler,
		reviewHandler,
		notificationHandler,
		menuHandler,
	)

	log.Fatal(server.Run(":" + configs.EnvPort()))
}
func initOrdergRPCConnection() *grpc.ClientConn {
	// dest := fmt.Sprintf("%s:%s", viper.GetString("order-service.grpc-host"), viper.GetString("order-service.grpc-port"))
	dest := fmt.Sprintf("localhost:%s", configs.EnvOrderServicePort())
	// Set up a connection to the server.
	conn, err := grpc.Dial(dest, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Default().Println("Connected to Order gRPC Service")
	return conn
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
}
