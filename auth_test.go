package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testID = 1
var testEmail = "gopher@golang.org"

var testID2 = 2
var testEmail2 = "michael.lawson@regres.in"

func SetupAuthServer(t *testing.T) *httptest.Server {
	// setup http test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pages, ok := r.URL.Query()["page"]
		if !ok || pages[0] == "1" {
			JSONResponse(w, UserList{
				Page:       1,
				TotalPages: 2,
				Users: []User{
					User{
						UserID: testID,
						Email:  testEmail,
					},
				},
			})
		} else if pages[0] == "2" {
			JSONResponse(w, UserList{
				Page:       2,
				TotalPages: 2,
				Users: []User{
					User{
						UserID: testID2,
						Email:  testEmail2,
					},
				},
			})
		}
	}))
	assert.NotNil(t, ts)

	return ts
}

func TestAuthLogin(t *testing.T) {
	ts := SetupAuthServer(t)
	defer ts.Close()

	client := ts.Client()

	// test loging in a user using an external api
	auth, err := CreateAuthenticator(client, ts.URL)
	assert.Nil(t, err)
	assert.NotNil(t, auth)

	user, err := auth.Login(testEmail2, "password")
	assert.Nil(t, err)
	assert.Equal(t, testEmail2, user.Email)
	assert.Equal(t, testID2, user.UserID)
}

func TestAuthLoginUserNotExist(t *testing.T) {
	ts := SetupAuthServer(t)
	defer ts.Close()

	client := ts.Client()

	// test loging in a user using an external api
	auth, err := CreateAuthenticator(client, ts.URL)
	assert.Nil(t, err)
	assert.NotNil(t, auth)

	_, err = auth.Login("foo", "password")
	assert.NotNil(t, err)
}
