package entity

type Admin struct {
	BaseEntity

	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	RefreshToken string `json:"-"`
}
