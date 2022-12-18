package main

import (
	"fmt"
	"time"

	"github.com/c0de_runn3r/fly-scanner/ryanair"
)

func main() {
	resp := *ryanair.GEtAvailabilities("CGN", "CPH")
	for i := 0; i < len(resp); i++ {
		result := ryanair.GetFlightSample("CGN", "CPH", resp[i])
		if result.Trips[0].Dates[0].Flights[0].FaresLeft != 0 {
			fmt.Printf("Flight %s - %s [%s] Price: %v\n", result.Trips[0].OriginName, result.Trips[0].DestinationName, result.Trips[0].Dates[0].DateOut[0:10], result.Trips[0].Dates[0].Flights[0].RegularFare.Fares[0].Amount)
		}
		time.Sleep(time.Second / 10)
	}

}
