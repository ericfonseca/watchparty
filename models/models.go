package models

type Venue struct {
	City        string `json:"city"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Event struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	City      string `json:"city"`
	StartTime string `json:"start_time"`
}

type User struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	TypeInterests string `json:"type_interests"`
	CityInterests string `json:"city_interests"`
}
