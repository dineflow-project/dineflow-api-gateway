package client_rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"api-gateway/configs"
	"api-gateway/model"
)

type ReviewClient struct {
	client *http.Client
}

type ReviewClientRest interface {
	GetAllReviews() ([]model.Review, error)
	GetReviewByID(id string) (model.Review, error)
	GetReviewByVendorID(vendor_id string) ([]model.Review, error)
	CreateReview(review model.Review) error
	UpdateReviewByID(id string, Review model.Review) error
	DeleteReviewByID(id string) error
}

func (s *ReviewClient) GetAllReviews() ([]model.Review, error) {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp []model.Review
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ReviewClient) GetReviewByID(id string) (model.Review, error) {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.Review{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.Review
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return model.Review{}, err
	}
	return resp, nil
}

func (s *ReviewClient) GetReviewByVendorID(vendorId string) ([]model.Review, error) {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews/byVendor/" + vendorId
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp []model.Review
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ReviewClient) CreateReview(Review model.Review) error {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews"

	//prepare request body
	byteData, err := json.Marshal(Review)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (s *ReviewClient) UpdateReviewByID(id string, Review model.Review) error {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews/" + id

	// prepare request body
	byteData, err := json.Marshal(Review)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	req, err := http.NewRequest(http.MethodPut, path, bodyReader)
	if err != nil {
		return err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (s *ReviewClient) DeleteReviewByID(id string) error {
	path := configs.EnvHost() + ":" + configs.EnvReviewServicePort() + "/reviews/" + id

	// send request
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return err
}

func ProvideReviewClientRest(client *http.Client) *ReviewClient {
	return &ReviewClient{client: client}
}
