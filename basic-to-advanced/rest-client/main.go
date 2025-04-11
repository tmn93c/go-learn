package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	resp, err := resty.New().R().
		SetResult(&User{}).
		Get("https://api.example.com/user/123")

	if err != nil {
		log.Fatal(err)
	}

	user := resp.Result().(*User)
	fmt.Println(user.Name)

}
