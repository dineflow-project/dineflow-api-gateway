package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gatewaygRPCClient "api-gateway/client_grpc"
	restClient "api-gateway/client_rest"
	"api-gateway/configs"
	gatewayHTTPHandler "api-gateway/http_handler"
	pb "api-gateway/pb"
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

func main() {
	flag.Parse()

	defer gracefulShutdown()

	// initConfig()

	r := gin.Default()

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
	userClientRest := restClient.ProvideUserClientRest(&http.Client{})
	userHandler := gatewayHTTPHandler.ProvideUserHandler(userClientRest)
	menuClientRest := restClient.ProvideMenuClientRest(&http.Client{})
	menuHandler := gatewayHTTPHandler.ProvideMenuHandler(menuClientRest)

	r.Use(cors.Default())
	gatewayHTTPHandler.ProvideRouter(r,
		orderHandler,
		reviewHandler,
		userHandler,
		notificationHandler,
		menuHandler,
	)

	// r.Run(":" + viper.GetString("api-gateway.port"))
	r.Run(":" + configs.EnvPort())
}

// Read Config file
// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath("./config")
// 	viper.SetConfigType("yaml")

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("fatal error config file: %s", err))
// 	}
// }

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
