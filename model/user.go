package model

import "fmt"

type User struct {
	ID     int `json:"id"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Age    int `json:"age"`
}

func (u *User) String() string {
	return fmt.Sprintf("User body: height: %d, weight: %d, and age: %d",
		u.Height, u.Weight, u.Age)
}
