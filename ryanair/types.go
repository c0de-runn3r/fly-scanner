package ryanair

type FlightResponce struct {
	// TermsOfUse    string `json:"termsOfUse"`
	// Currency      string `json:"currency"`
	// CurrPrecision int    `json:"currPrecision"`
	// RouteGroup    string `json:"routeGroup"`
	// TripType      string `json:"tripType"`
	// UpgradeType   string `json:"upgradeType"`
	Trips []struct {
		Origin          string `json:"origin"`
		OriginName      string `json:"originName"`
		Destination     string `json:"destination"`
		DestinationName string `json:"destinationName"`
		// RouteGroup      string `json:"routeGroup"`
		// TripType        string `json:"tripType"`
		// UpgradeType     string `json:"upgradeType"`
		Dates []struct {
			DateOut string `json:"dateOut"`
			Flights []struct {
				FaresLeft int `json:"faresLeft"`
				// FlightKey   string `json:"flightKey"`
				// InfantsLeft int    `json:"infantsLeft"`
				RegularFare struct {
					// FareKey   string `json:"fareKey"`
					// FareClass string `json:"fareClass"`
					Fares []struct {
						// Type              string  `json:"type"`
						Amount float64 `json:"amount"`
						// Count             int     `json:"count"`
						// HasDiscount       bool    `json:"hasDiscount"`
						// PublishedFare float64 `json:"publishedFare"`
						// DiscountInPercent int     `json:"discountInPercent"`
						// HasPromoDiscount  bool    `json:"hasPromoDiscount"`
						// DiscountAmount    float64 `json:"discountAmount"`
						// HasBogof          bool    `json:"hasBogof"`
					} `json:"fares"`
				} `json:"regularFare"`
				// OperatedBy string `json:"operatedBy"`
				// Segments []struct {
				// SegmentNr    int         `json:"segmentNr"`
				// Origin       string `json:"origin"`
				// Destination  string `json:"destination"`
				// FlightNumber string `json:"flightNumber"`
				// Time         []string    `json:"time"`
				// TimeUTC      []time.Time `json:"timeUTC"`
				// Duration     string      `json:"duration"`
				// } `json:"segments"`
				// FlightNumber string      `json:"flightNumber"`
				// Time         []string    `json:"time"`
				// TimeUTC      []time.Time `json:"timeUTC"`
				// Duration     string      `json:"duration"`
			} `json:"flights"`
		} `json:"dates"`
	} `json:"trips"`
	// ServerTimeUTC time.Time `json:"serverTimeUTC"`
}

type DestinationsResponce []struct {
	IataCode string `json:"iataCode"`
	Name     string `json:"name"`
	SeoName  string `json:"seoName"`
	// Aliases     []interface{} `json:"aliases"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	// Base           bool          `json:"base"`
	CountryCode  string   `json:"countryCode"`
	RegionCode   string   `json:"regionCode"`
	CityCode     string   `json:"cityCode"`
	CurrencyCode string   `json:"currencyCode"`
	Routes       []string `json:"routes"`
	// SeasonalRoutes []interface{} `json:"seasonalRoutes"`
	Categories []string `json:"categories"`
	// Priority       int           `json:"priority"`
	// TimeZone       string        `json:"timeZone"`
	// MacCityCode    string        `json:"macCityCode,omitempty"`
}

type Availabilities []string
