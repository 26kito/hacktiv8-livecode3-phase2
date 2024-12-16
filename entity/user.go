package entity

type User struct {
	UserID   uint   `gorm:"primaryKey;column:user_id" json:"user_id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	JwtToken string `gorm:"type:varchar(255)" json:"jwt_token"`
}

type UserRegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
