package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type RouteService interface {
	CreateAndModifyRoute(input *RouteInfo) (*CreateRouteResponse, error)
	DeleteRoute(routeId string) (*DeleteResponse, error)
	QueryRouteById(routeId string) (*CreateRouteResponse, error)
	QueryRoutesByIds(routeIds []string) (*QueryMultiResponse, error)
	QueryAllRoutes() (*QueryMultiResponse, error)
	QueryRoutesByStartAndEnd(start, end string) (*QueryMultiResponse, error)
}
type QueryMultiResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id           string   `json:"id"`
		Stations     []string `json:"stations"`
		Distances    []int    `json:"distances"`
		StartStation string   `json:"startStation"`
		EndStation   string   `json:"endStation"`
	} `json:"data"`
}

type DeleteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type RouteInfo struct {
	ID           string `json:"id"`
	StartStation string `json:"startStation"`
	EndStation   string `json:"endStation"`
	StationList  string `json:"stationList"`
	DistanceList string `json:"distanceList"`
}

type CreateRouteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id           string   `json:"id"`
		Stations     []string `json:"stations"`
		Distances    []int    `json:"distances"`
		StartStation string   `json:"startStation"`
		EndStation   string   `json:"endStation"`
	} `json:"data"`
}

type RouteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id           string   `json:"id"`
		Stations     []string `json:"stations"`
		Distances    []int    `json:"distances"`
		StartStation string   `json:"startStation"`
		EndStation   string   `json:"endStation"`
	} `json:"data"`
}

func (s *SvcImpl) CreateAndModifyRoute(input *RouteInfo) (*CreateRouteResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/routeservice/routes", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result CreateRouteResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) DeleteRoute(routeId string) (*DeleteResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/routeservice/routes/%s", routeId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DeleteResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryRouteById(routeId string) (*CreateRouteResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/routeservice/routes/%s", routeId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result CreateRouteResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryRoutesByIds(routeIds []string) (*QueryMultiResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/routeservice/routes/byIds", routeIds)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryMultiResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryAllRoutes() (*QueryMultiResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/routeservice/routes", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryMultiResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryRoutesByStartAndEnd(start, end string) (*QueryMultiResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/routeservice/routes/%s/%s", start, end), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryMultiResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
