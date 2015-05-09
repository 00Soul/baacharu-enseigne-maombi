package main

import (
	"github.com/00Soul/mappings"
	"github.com/00Soul/oxpit"
	"time"
)

const timeLayout string = "2006-01-02T15:04:05-07:00"

func setupMappings() {
	mapping := mappings.New(time.Time)
	mapping.FlattenFunc(fromTime)
	mapping.UnflattenFunc(toTime)

	mappings.New(oxpit.AccountState).FlattenFunc(flattenAccountState).UnflattenFunc(unflattenAccountState)

	mapping = mappings.New(oxpit.User)
	mapping.Field(oxpit.User.Id).Name("id")
	mapping.Field(oxpit.User.State).Name("state")
	mapping.Field(oxpit.User.CreatedWhen).Name("created-when")

	mapping = mappings.New(oxpit.Profile)
	mapping.Field(oxpit.Profile.Email).Name("email")
	mapping.Field(oxpit.Profile.Username).Name("username")
	mapping.Field(oxpit.Profile.Alias).Name("alias")

	mapping = mappings.New(oxpit.Board)
	mapping.Field(oxpit.Board.Id).Name("id")
	mapping.Field(oxpit.Board.Title).Name("title")
	mapping.Field(oxpit.Board.Columns).Name("columns")
	mapping.Field(oxpit.Board.Cards).Name("cards")
	mapping.Field(oxpit.Board.OwnedBy).Name("owned-by")
	mapping.Field(oxpit.Board.CreatedBy).Name("created-by")
	mapping.Field(oxpit.Board.CreatedWhen).Name("created-when")

	mapping = mappings.New(oxpit.Card)
	mapping.Field(oxpit.Card.Id).Name("id")
	mapping.Field(oxpit.Card.Stage).Name("stage")
	mapping.Field(oxpit.Card.CardType).Name("card-type")
	mapping.Field(oxpit.Card.Data).Name("data")

	mapping = mappings.New(oxpit.Column)
	mapping.Field(oxpit.Column.Title).Name("title")
	mapping.Field(oxpit.Column.WipLimit).Name("wiplimit")
}

func fromTime(i interface{}) interface{} {
	if when, ok := i.(time.Time); ok {
		return when.Format(timeLayout)
	} else {
		return "0000-01-01T00:00:00+00:00"
	}
}

func toTime(i interface{}) interface{} {
	if str, ok := i.(string); ok {
		if when, err := time.Parse(timeLayout, str); err == nil {
			return when
		} else {
			return time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
		}
	} else {
		return time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	}
}

func flattenAccountState(i interface{}) interface{} {
	var jsonString string

	if state, ok := i.(oxpit.AccountState); ok {
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

func unflattenAccountState(i interface{}) interface{} {
	var state oxpit.AccountState

	if str, ok := i.(string); ok {
		switch str {
		case "active":
			state = oxpit.AccountActive
		case "inactive":
			state = oxpit.AccountInactive
		case "closed":
			state = oxpit.AccountClosed
		}
	}

	return state
}
