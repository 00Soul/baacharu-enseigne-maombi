package main

import (
	"encoding/json"
	"github.com/00Soul/oxpit"
	"strconv"
	"time"
)

const timeLayout string = "2006-01-02T15:04:05-07:00"

func toAccountState(jsonString string) oxpit.AccountState {
	var stateObject oxpit.AccountState
	var stateString string

	json.Unmarshal([]byte(jsonString), &stateString)

	switch stateString {
	case "active":
		stateObject = oxpit.AccountActive
	case "inactive":
		stateObject = oxpit.AccountInactive
	case "closed":
		stateObject = oxpit.AccountClosed
	}

	return stateObject
}

func jsonFromAccountState(state oxpit.AccountState) string {
	var jsonString string

	switch state {
	case oxpit.AccountActive:
		jsonString = "active"
	case oxpit.AccountInactive:
		jsonString = "inactive"
	case oxpit.AccountClosed:
		jsonString = "closed"
	}

	return jsonString
}

func toUser(jsonString string) oxpit.User {
	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObject)

	id, err := strconv.Atoi(jsonObject["id"].(string))
	if err != nil {
		id = 0
	}

	when, err := time.Parse(timeLayout, jsonObject["created-when"].(string))
	if err != nil {
		when = time.Now().UTC()
	}

	var user = oxpit.User{
		Id:          id,
		State:       toAccountState(jsonObject["state"].(string)),
		CreatedWhen: when,
	}

	return user
}

func interfaceFromUser(user oxpit.User) interface{} {
	return map[string]string{
		"id":           strconv.Itoa(user.Id),
		"state":        jsonFromAccountState(user.State),
		"created-when": user.CreatedWhen.Format(timeLayout),
	}
}

func jsonFromUser(user oxpit.User) string {
	bytes, _ := json.Marshal(interfaceFromUser(user))

	return string(bytes)
}

func jsonFromProfile(profile oxpit.Profile) string {
	bytes, _ := json.Marshal(map[string]string{
		"email":    profile.Email,
		"username": profile.Username,
		"alias":    profile.Alias,
	})

	return string(bytes)
}

func toProfile(jsonString string) oxpit.Profile {
	var jsonObject map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObject)

	return oxpit.Profile{
		Email:    jsonObject["email"].(string),
		Username: jsonObject["username"].(string),
		Alias:    jsonObject["alias"].(string),
	}
}
