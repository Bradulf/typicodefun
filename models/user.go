package models

type User struct {
    ID  int `json:"id"`
    Name  string `json:"name"`
    UserName  string  `json:"username"`
    Email  string  `json:"email"`
}
