package models

type Venue struct {
	ID          int
	City        string `json:"city"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Event struct {
	ID        int
	Title     string `json:"title"`
	Type      string `json:"type"`
	City      string `json:"city"`
	StartTime string `json:"start_time"`
}

type User struct {
	ID    int
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Watcher struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}

type Hoster struct {
	EventID int `json:"event_id"`
	VenueID int `json:"venue_id"`
}

type Interest struct {
	Type   string `json:"type"`
	City   string `json:"city"`
	UserID int    `json:"user_id"`
}
