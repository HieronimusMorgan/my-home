package out

type RegisterResponse struct {
	UserID         uint    `json:"user_id"`
	Username       string  `json:"username"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	PhoneNumber    string  `json:"phone_number"`
	ProfilePicture string  `json:"profile_picture"`
	Balance        float64 `json:"balance"`
	Token          string  `json:"token"`
}
