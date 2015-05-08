package main

import (
	"github.com/00Soul/oxpit"
	"github.com/00Soul/oxpit/json"
	"strconv"
	"time"
)

const timeLayout string = "2006-01-02T15:04:05-07:00"

func setupJsonMappings() {
	mapping := json.mapping.New(time.Time)
	mapping.FlattenFunc(fromTime)
	mapping.UnflattenFunc(toTime)

	mapping = json.mapping.New(oxpit.AccountState)
	mapping.FlattenFunc(flattenAccountState)
	mapping.UnflattenFunc(unflattenAccountState)

	mapping = json.mapping.New(oxpit.User)
	mapping.Field(oxpit.User.Id).Name("id")
	mapping.Field(oxpit.User.State).Name("state")
	mapping.Field(oxpit.User.CreatedWhen).Name("created-when")
}

func flattenAccountState(goObject interface{}) interface{} {
	var jsonString string

	state, ok := v.(oxpit.AccountState)
	if ok {
		switch state {
		case oxpit.AccountActive:
			jsonString = "active"
		case oxpit.AccountInactive:
			jsonString = "inactive"
		case oxpit.AccountClosed:
			jsonString = "closed"
		}
	}

	return jsonString
}

func unflattenAccountState(jsonString interface{}) interface{} {
	var goObject oxpit.AccountState

	switch jsonString {
	case "active":
		goObject = oxpit.AccountActive
	case "inactive":
		goObject = oxpit.AccountInactive
	case "closed":
		goObject = oxpit.AccountClosed
	}

	return goObject
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

func jsonFromSlice(slice []interface{}) []bytes {
}

func toSliceFrom(jsonArray interface{}, convertFunction func(map[string]interface{}) interface{}) []interface{} {
	// First, using a type assertion, convert the generic empty interface
	//    to an array of generic empty interfaces (which is what we are
	//    expecting.
	goSlice, goSliceOk := jsonArray.([]interface{})
	if !goSliceOk {
		// What happens now?
	}

	// Create a new slice to hold the converted elements.
	slice := make([]interface{}, 0, len(goSlice))
	for _, item := range goSlice {
		goMap, goMapOk := item.(map[string]interface{})
		if goMapOk {
			slice = append(slice, convertFunction(goMap))
		}
	}

	return slice
}

func toTimeFrom(jsonTime interface{}) time.Time {
	when, err := time.Parse(timeLayout, jsonTime.(string))
	if err != nil {
		when = time.Now().UTC()
	}

	return when
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

func toProfileFromString(jsonString string) oxpit.Profile {
	return toProfileFromBytes([]byte(jsonObject))
}

func toProfileFromBytes(jsonBytes []bytes) oxpit.Profile {
	var jsonObject map[string]interface{}
	json.Unmarshal(jsonBytes, &jsonObject)

	return toProfile(jsonObject)
}

func toProfile(jsonObject map[string]interface{}) oxpit.Profile {
	var profile oxpit.Profile

	for key, value := range jsonObject {
		switch key {
		case "email":
			profile[key] = value.(string)
		case "username":
			profile[key] = value.(string)
		case "alias":
			profile[key] = value.(string)
		}
	}

	return profile
}

func jsonFromBoard(board oxpit.Board) string {
	bytes, _ := json.Marshal(map[string]interface{}{
		"id":           board.Id,
		"title":        board.Title,
		"columns":      board.Columns,
		"cards":        board.Cards,
		"owned-by":     board.OwnedBy,
		"created-by":   board.CreatedBy,
		"created-when": board.CreatedWhen,
	})

	return string(bytes)
}

func toBoardFromString(jsonString string) oxpit.Board {
	return toBoardFromBytes([]byte(jsonObject))
}

func toBoardFromBytes(bytes []byte) oxpit.Board {
	return toStructFromBytes(bytes, toBoard).(oxpit.Board)
}

func toBoard(jsonObject map[string]interface{}) oxpit.Board {
	var board oxpit.Board

	for key, value := range jsonObject {
		switch key {
		case "id", "owned-by", "created-by":
			board[key] = value.(int)
		case "title":
			board[key] = value.(string)
		case "created-when":
			board[key] = toTime(value)
		case "columns":
			board[key] = toTime(value)
		}
	}

	return board
}

func toCardFromString(jsonString string) oxpit.Card {
	return toCardFromBytes([]byte(jsonObject))
}

func toCardFromBytes(jsonBytes []byte) oxpit.Card {
	return toStructFromBytes(bytes, toCard).(oxpit.Card)
}

func toCard(jsonObject map[string]interface{}) oxpit.Card {
	var column oxpit.Card

	for key, value := range jsonObject {
		switch key {
		case "card-type", "state", "date":
			board[key] = value.(string)
		case "id":
			board[key] = value.(int)
		}
	}

	return column
}

func toColumnFromString(jsonString string) oxpit.Column {
	return toColumnFromBytes([]byte(jsonString))
}

func toColumnFromBytes(jsonBytes []byte) oxpit.Column {
	return toStructFromBytes(bytes, toColumn).(oxpit.Column)
}

func toColumn(jsonObject map[string]interface{}) oxpit.Column {
	var column oxpit.Column

	for key, value := range jsonObject {
		switch key {
		case "title":
			board[key] = value.(string)
		case "wiplimit":
			board[key] = value.(int)
		}
	}

	return column
}

func toStructFromBytes(jsonBytes []byte, convertFunction func(map[string]interface{}) interface{}) interface{} {
	var object map[string]interface{}
	json.Unmarshal(bytes, &object)

	return convertFunction(object)
}
