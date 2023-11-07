package client_rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-gateway/configs"
	"api-gateway/model"
)

type NotificationClient struct {
	client *http.Client
}

type NotificationClientRest interface {
	GetUnreadNotification(recipientID string) (model.UnreadNotificationResponseBody, error)
	GetAllNotifiactions(recipientID string) ([]model.Notification, error)
}

func (s *NotificationClient) GetUnreadNotification(recipientID string) (model.UnreadNotificationResponseBody, error) {
	path := configs.EnvHost() + ":" + configs.EnvNotificationServicePort() + "/notifications/unread/" + recipientID
	fmt.Println(path)

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.UnreadNotificationResponseBody{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UnreadNotificationResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.UnreadNotificationResponseBody{}, err
	}
	return resp, nil
}

func (s *NotificationClient) GetAllNotifiactions(recipientID string) ([]model.Notification, error) {
	path := configs.EnvHost() + ":" + configs.EnvNotificationServicePort() + "/notifications/" + recipientID

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp model.NotificationResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func ProvideNotificationClientRest(client *http.Client) *NotificationClient {
	return &NotificationClient{client: client}
}
