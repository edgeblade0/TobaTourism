package models

type Accounts struct {
	AccountID int64  `json:"account_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Token     string `json:"token";sql:"-"`
}

type Profile struct {
	ProfileID   int64  `json:"profile_id,omitempty"`
	AccountID   int64  `json:"account_id,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	NoHP        string `json:"nohp"`
	Alamat      string `json:"alamat"`
	DOB         string `json:"dob"`
	SocialMedia string `json:"social_media"`
}
