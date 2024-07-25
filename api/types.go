package api

type SubscribeUserRequest struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type Notification struct {
	DestinationID string `json:"destination_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
}
