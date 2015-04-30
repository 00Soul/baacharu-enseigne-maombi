package system

import (
	"time"
)

type Card struct {
	id       int
	cardType string
	stage    string
	data     string
}

type CardTransition struct {
	cardId    int
	movedBy   int
	movedwhen time.Time
}

type Column struct {
	title    string
	wipLimit int
}

type Board struct {
	id          int
	title       string
	columns     []Column
	cards       map[int]Card
	ownedBy     int
	createdBy   int
	createdWhen time.Time
}
