package domain

type Booking struct {
	ID         int    `json:"id"`
	RoomName   string `json:"room_name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	ReservedBy string `json:"reserved_by"`
}
