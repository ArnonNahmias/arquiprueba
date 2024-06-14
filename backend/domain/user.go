package domain

type User struct {
	ID       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username"`
	Password string `json:"password"`
}
