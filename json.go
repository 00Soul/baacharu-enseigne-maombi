package oxpit

import (
	"encoding/json"
	"strconv"
	"time"
)

const timeLayout string = "2006-01-02T15:04:05-07:00"

func jsonAccountState(jsonString string) AccountState {
	var state string
	json.Unmarshal([]byte(jsonString), &state)

	switch state {
	case "active":
		return AccountActive
	case "inactive":
		return AccountInactive
	case "closed":
		return AccountClosed
	}
}

func (state AccountState) json() string {
	switch state {
	case AccountActive:
		return "active"
	case AccountInactive:
		return "inactive"
	case AccountClosed:
		return "closed"
	}
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
