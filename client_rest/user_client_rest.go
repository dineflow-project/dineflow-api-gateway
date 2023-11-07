package client_rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"api-gateway/configs"
	"api-gateway/model"
)

type UserClient struct {
	client *http.Client
}

type UserClientRest interface {
	GetMe() (model.User, error)
	SignUp(user model.User) (model.User, error)
	Login(user model.User) (model.User, error)
}

func (s *UserClient) GetMe() (model.User, error) {
	path := configs.EnvHost() + ":" + configs.EnvUserServicePort() + "/api/users/me"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		fmt.Println('1', err)
		return model.User{}, err
	}
	defer response.Body.Close()
	fmt.Println(response.Body)

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		fmt.Println('2', err)
		return model.User{}, err
	}
	if resp.Status != "success" {
		return model.User{}, errors.New(resp.Message)
	}
	return resp.Data.User, nil
}

func (s *UserClient) SignUp(user model.User) (model.User, error) {
	path := configs.EnvHost() + ":" + configs.EnvUserServicePort() + "/api/auth/register"

	// prepare request body
	byteData, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Status != "success" {
		return model.User{}, errors.New(resp.Message)
	}

	return resp.Data.User, nil
}

func (s *UserClient) Login(user model.User) (model.User, error) {
	path := configs.EnvHost() + ":" + configs.EnvUserServicePort() + "/api/auth/login"

	// prepare request body
	byteData, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Status != "success" {
		return model.User{}, errors.New(resp.Status)
	}
	resp.Data.User.Token = resp.Token

	return resp.Data.User, nil
}

func ProvideUserClientRest(client *http.Client) *UserClient {
	return &UserClient{client: client}
}
