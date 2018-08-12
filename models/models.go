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
	EventID string `json:"event_id`
	UserID  string `json:"user_id`
}

type Hoster struct {
	EventID string `json:"event_id`
	VenueID string `json:"venue_id`
}

type Interest struct {
	Type   string `json:"type"`
	City   string `json:"city"`
	UserID string `json:"user_id`
}
