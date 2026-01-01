package models

type EventList struct {
	Events []*Event `json:"items,omitempty"`
}

type Event struct {
	EventName   string `json:"summary,omitempty"`
	DestiNation string `json:"location,omitempty"`
}

type DisplayName struct {
	Text string `json:"text"`
}

type Place struct {
	Rating      float64     `json:"rating"`
	MapUri      string      `json:"googleMapsUri"`
	DisplayName DisplayName `json:"displayName"`
}

type PlacesResponse struct {
	Places []Place `json:"places"`
}

type PlacesAPIResponse struct {
	Results []struct {
		Name             string  `json:"name"`
		Rating           float64 `json:"rating"`
		FormattedAddress string  `json:"formatted_address"`
		PlaceID          string  `json:"place_id"`
	} `json:"results"`
	Status string `json:"status"`
}
