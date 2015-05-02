package oxpit

import (
	"encoding/json"
	"strconv"
	"time"
)

const timeLayout string = "2006-01-02T15:04:05-07:00"

func jsonAccountState(jsonString string) AccountState {
	var stateObject AccountState
	var stateString string

	json.Unmarshal([]byte(jsonString), &stateString)

	switch stateString {
	case "active":
		stateObject = AccountActive
	case "inactive":
		stateObject = AccountInactive
	case "closed":
		stateObject = AccountClosed
	}

	return stateObject
}

func (state AccountState) json() string {
	var jsonString string

	switch state {
	case AccountActive:
		jsonString = "active"
	case AccountInactive:
		jsonString = "inactive"
	case AccountClosed:
		jsonString = "closed"
	}

	return jsonString
}

func jsonUser(jsonString string) User {
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

	var user = User{
		Id:          id,
		State:       jsonAccountState(jsonObject["state"].(string)),
		CreatedWhen: when,
	}

	return user
}

func (user User) json() string {
	jsonObject := map[string]string{
		"id":           strconv.Itoa(user.Id),
		"state":        user.State.json(),
		"created-when": user.CreatedWhen.Format(timeLayout),
	}

	bytes, _ := json.Marshal(jsonObject)

	return string(bytes)
}
