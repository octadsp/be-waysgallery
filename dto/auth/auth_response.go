package authdto

type RegisterResponse struct {
	Email    string `json:"email" gorm:"type: varchar(255)"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}
