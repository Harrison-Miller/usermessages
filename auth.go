package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
}

type UserList struct {
	Page       int    `json:"page"`
	TotalPages int    `json:"total_pages"`
	Users      []User `json:"data"`
}

type Authenticator struct {
	Client *http.Client
	URL    string
}

func CreateAuthenticator(client *http.Client, url string) (*Authenticator, error) {
	return &Authenticator{
		Client: client,
		URL:    url,
	}, nil
}

func (auth *Authenticator) GetUsers(page int) (UserList, error) {
	req, err := http.NewRequest("GET", auth.URL, nil)
	values := req.URL.Query()
	values.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = values.Encode()

	res, err := auth.Client.Do(req)
	if err != nil {
		return UserList{}, err
	}
	defer res.Body.Close()

	var users UserList
	err = json.NewDecoder(res.Body).Decode(&users)
	return users, err
}

func (auth *Authenticator) Login(email string, password string) (User, error) {
	var totalPages = 2
	for page := 1; page <= totalPages; page++ {
		users, err := auth.GetUsers(page)
		if err != nil {
			return User{}, err
		}

		for _, u := range users.Users {
			if u.Email == email {
				return u, nil
			}
		}

		totalPages = users.TotalPages
	}

	return User{}, fmt.Errorf("user %s not found", email)
}
