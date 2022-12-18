package ryanair

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetFlightSample(origin, destination string, dateOut string) *FlightResponce {
	c := http.DefaultClient
	url := "https://www.ryanair.com/api/booking/v4/uk-ua/availability"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	q := req.URL.Query()
	q.Add("ADT", "1")
	q.Add("DateOut", dateOut)
	q.Add("Origin", origin)
	q.Add("Destination", destination)
	q.Add("FlexDaysOut", "0") // max 6 (in total 7 days)
	q.Add("ToUs", "AGREED")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var ryanairResp FlightResponce
	json.Unmarshal(body, &ryanairResp)
	return &ryanairResp
}

func GetDestinations() *DestinationsResponce {
	c := http.DefaultClient
	url := "https://www.ryanair.com/api/views/locate/3/airports/uk/active"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var ryanairResp DestinationsResponce
	json.Unmarshal(body, &ryanairResp)
	return &ryanairResp
}

func GEtAvailabilities(origin, destination string) *Availabilities {
	c := http.DefaultClient
	url, _ := url.JoinPath("https://www.ryanair.com/api/farfnd/3/oneWayFares/", origin, destination, "availabilities")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var ryanairResp Availabilities
	json.Unmarshal(body, &ryanairResp)
	return &ryanairResp
}
