package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/e0/teacherMe/go_app/back_end/model"
)

// UpdateUser updates a user in the Auth0 database.
func UpdateUser(courseID string, baseURL string, authToken string) int {
	tokenInfoURL := baseURL + "tokeninfo"
	user := getUser(authToken, tokenInfoURL)

	updateUserURL := baseURL + "api/v2/users/" + user.UserID
	courses := append(user.Metadata.Courses, courseID)
	coursesJSONData, _ := json.Marshal(courses)

	jsonStr := "{\"user_metadata\":{\"courses\":" + string(coursesJSONData) + "}}"
	status, _ := makeAuth0Request([]byte(jsonStr), updateUserURL, "PATCH", authToken)

	return status
}

// GetUser gets the user from Auth0 database.
func getUser(authToken string, url string) model.User {
	idToken := strings.Fields(authToken)[1]
	jsonStr := "{\"id_token\":\"" + idToken + "\"}"

	_, resp := makeAuth0Request([]byte(jsonStr), url, "POST", "")
	var user model.User

	err := json.Unmarshal(resp, &user)

	if err != nil {
		panic(err)
	}

	return user
}

func makeAuth0Request(data []byte, url string, httpMethod string, authToken string) (int, []byte) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(data))
	// TODO: handle this err
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}
