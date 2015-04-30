package system

import (
	"time"
)

type Card struct {
	Id       int
	CardType string
	Stage    string
	Data     string
}

type CardTransition struct {
	CardId    int
	MovedBy   int
	MovedWhen time.Time
}

type Column struct {
	Title    string
	WipLimit int
}

type Board struct {
	Id          int
	Title       string
	Columns     []Column
	Cards       map[int]Card
	OwnedBy     int
	CreatedBy   int
	CreatedWhen time.Time
}
