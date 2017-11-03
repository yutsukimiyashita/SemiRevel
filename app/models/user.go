package models

type User struct {
    Id          string `gorm:"primary_key" json:"id"`
    Name        string `json:"name"`
    Password    string `json:"password"`
    Grade       uint64 `json:"grade"`
}