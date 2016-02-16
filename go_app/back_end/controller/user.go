package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetUserID gets the user id from Auth0 database.
func GetUserID(authToken string, url string) string {
	idToken := strings.Fields(authToken)[1]
	jsonStr := "{\"id_token\":\"" + idToken + "\"}"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(body)))
	bodyContent := map[string]string{}
	decoder.Decode(&bodyContent)

	return bodyContent["user_id"]
}

// UpdateUser updates a user in the Auth0 database.
func UpdateUser(userID string, courseID string, url string, authToken string) int {
	jsonStr := "{\"user_metadata\":{\"courses\":[\"" + courseID + "\"]}}"

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return resp.StatusCode
}
