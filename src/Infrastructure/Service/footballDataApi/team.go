package footballdataapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Team struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	Tla         string `json:"tla": `
	Crest       string `json:"crest"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Founded     int16  `json:"founded"`
	ClubColors  string `json:"clubColors"`
	Venue       string `json:"venue"`
	LastUpdated string `json:"lastUpdated"`
}

type TeamsList struct {
	Count   string `json:"count"`
	Filters string `json:"filters"`
	Teams   []Team `json:"teams"`
}

func (rcv *httpClient) GetTeamsList() (*TeamsList, error) {

	response, err := rcv.client.Get("teams/44")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := &TeamsList{}

	switch response.StatusCode {
	case http.StatusOK:
		if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
			return nil, err
		}

		return result, nil
	case http.StatusNotFound:
		return result, nil
	default:
	}

	return nil, errors.New(fmt.Sprintf("Unexpected status code: %d", response.StatusCode))
}
