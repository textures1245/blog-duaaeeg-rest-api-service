package entities

type UserProfileRes struct {
	UUID           string `json:"uuid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_pic"`
	CreatedAt      string `json:"created_at"`
	UpdateAt       string `json:"update_at"`
}
