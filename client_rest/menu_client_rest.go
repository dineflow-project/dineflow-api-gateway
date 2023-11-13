package client_rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"dineflow-api-gateway/configs"
	"dineflow-api-gateway/model"
)

type MenuClient struct {
	client *http.Client
}

type MenuClientRest interface {
	//menu
	GetAllMenus(canteenName, vendorName, minprice, maxprice string) ([]model.Menu, error)
	GetMenuByID(id string) (model.Menu, error)
	GetMenuByVendorID(vendor_id string) ([]model.Menu, error)
	CreateMenu(menu model.Menu) (model.Menu, error)
	UpdateMenuByID(id string, Menu model.Menu) error
	DeleteMenuByID(id string) error

	//vendor
	GetAllVendors() ([]model.Vendor, error)
	GetVendorByID(id string) (model.Vendor, error)
	GetVendorByOwnerID(id string) (model.Vendor, error)
	GetAllVendorsByCanteenID(id string) ([]model.Vendor, error)
	CreateVendor(menu model.Vendor) error
	UpdateVendorByID(id string, Vendor model.Vendor) error
	DeleteVendorByID(id string) error

	//canteen
	GetAllCanteens() ([]model.Canteen, error)
	GetCanteenByID(id string) (model.Canteen, error)
	CreateCanteen(menu model.Canteen) error
	UpdateCanteenByID(id string, Canteen model.Canteen) error
	DeleteCanteenByID(id string) error
}

// menu ---------------------------------------------------------------------------------------------------------------------------------------
func (s *MenuClient) GetAllMenus(canteenId, vendorId, minprice, maxprice string) ([]model.Menu, error) {
	// path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus" + "?canteen=" + canteenName + "&vendor=" + vendorName + "&minprice=" + minprice + "&maxprice=" + maxprice
	queryParams := url.Values{}
	queryParams.Set("canteenId", canteenId)
	queryParams.Set("vendorId", vendorId)
	queryParams.Set("minprice", minprice)
	queryParams.Set("maxprice", maxprice)

	baseURL := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus"

	path := baseURL + "?" + queryParams.Encode()
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp []model.Menu
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MenuClient) GetMenuByID(id string) (model.Menu, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.Menu{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Menu{}, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return model.Menu{}, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	// read response
	var resp model.Menu
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return model.Menu{}, err
	}
	return resp, nil
}

func (s *MenuClient) GetMenuByVendorID(vendorId string) ([]model.Menu, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus/byVendor/" + vendorId
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return nil, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	// read response
	var resp []model.Menu
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MenuClient) CreateMenu(Menu model.Menu) (model.Menu, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus"

	//prepare request body
	byteData, err := json.Marshal(Menu)
	if err != nil {
		return model.Menu{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	fmt.Println(response)
	if err != nil {
		return model.Menu{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Menu{}, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return model.Menu{}, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}

	// read response
	var resp model.Menu
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Menu{}, err
	}

	return resp, nil
}

func (s *MenuClient) UpdateMenuByID(id string, Menu model.Menu) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus/" + id

	// prepare request body
	byteData, err := json.Marshal(Menu)
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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return nil
}

func (s *MenuClient) DeleteMenuByID(id string) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/menus/" + id

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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return err
}

// vendor ---------------------------------------------------------------------------------------------------------------------------------------
func (s *MenuClient) GetAllVendors() ([]model.Vendor, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp []model.Vendor
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MenuClient) GetVendorByID(id string) (model.Vendor, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	fmt.Println(err)
	if err != nil {
		return model.Vendor{}, err
	}
	defer response.Body.Close()

	// print response body
	// bodyBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(bodyBytes))
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Vendor{}, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return model.Vendor{}, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}

	// read response

	var resp model.Vendor
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return model.Vendor{}, err
	}
	return resp, nil
}

func (s *MenuClient) GetVendorByOwnerID(id string) (model.Vendor, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors/byOwner/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	fmt.Println(err)
	if err != nil {
		return model.Vendor{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Vendor{}, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return model.Vendor{}, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}

	// read response

	var resp model.Vendor
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return model.Vendor{}, err
	}
	return resp, nil
}

func (s *MenuClient) GetAllVendorsByCanteenID(id string) ([]model.Vendor, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors/canteens/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return nil, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}

	// read response
	var resp []model.Vendor
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MenuClient) CreateVendor(Menu model.Vendor) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors"

	//prepare request body
	byteData, err := json.Marshal(Menu)
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

func (s *MenuClient) UpdateVendorByID(id string, Vendor model.Vendor) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors/" + id

	// prepare request body
	byteData, err := json.Marshal(Vendor)
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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return nil
}

func (s *MenuClient) DeleteVendorByID(id string) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/vendors/" + id

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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return err
}

// canteen ---------------------------------------------------------------------------------------------------------------------------------------
func (s *MenuClient) GetAllCanteens() ([]model.Canteen, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/canteens"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp []model.Canteen
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MenuClient) GetCanteenByID(id string) (model.Canteen, error) {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/canteens/" + id
	fmt.Println(path)
	// send request
	response, err := s.client.Get(path)
	fmt.Println(err)
	if err != nil {
		return model.Canteen{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Canteen{}, fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return model.Canteen{}, fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}

	// read response
	var resp model.Canteen
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return model.Canteen{}, err
	}
	return resp, nil
}

func (s *MenuClient) CreateCanteen(Menu model.Canteen) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/canteens"

	//prepare request body
	byteData, err := json.Marshal(Menu)
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

func (s *MenuClient) UpdateCanteenByID(id string, Canteen model.Canteen) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/canteens/" + id

	// prepare request body
	byteData, err := json.Marshal(Canteen)
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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return nil
}

func (s *MenuClient) DeleteCanteenByID(id string) error {
	path := configs.EnvMenuServiceHost() + ":" + configs.EnvMenuServicePort() + "/canteens/" + id

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
	if response.StatusCode != http.StatusOK {
		// Read the error response from the service
		errorResponse, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %s", err.Error())
		}

		return fmt.Errorf("HTTP Error: %s", string(errorResponse))
	}
	return err
}

func ProvideMenuClientRest(client *http.Client) *MenuClient {
	return &MenuClient{client: client}
}
