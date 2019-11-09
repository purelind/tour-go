package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	ID    int
	Name  string
}

type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	u, err := retrieveUser("Hoanh")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile
	// Since the returned u is an address, use * to get the value
	fmt.Printf("%+v %T\n", *u, *u)

	// Update user name. Don't care about update stats
	// This _ is called blank identifier.
	// Since we don't need anything outside the scope of if, we can use the compact syntax
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Updated user record for ID", u.ID)
}

func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Goal: Unmarshal the json document into a value of the user struct type
	// Create a value type user
	var u user

	err = json.Unmarshal([]byte(r), &u)

	return &u, err
}

// GetUser simulate a web call that returns a json
// document for the specified user.
func getUser(name string) (string, error) {
	response := `{"ID":102, "Name":"Hoanh"}`
	return response, nil
}

// updateUser updates the specified user document
func updateUser(u *user) (*updateStats, error) {
	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success": true, "Message": "updated"}`

	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Check the update status to verify the update is successful.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
