package oxpit

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

func System(actor User, app Application) *System {
}

func CreateBoard(actor User, app Application, user User, title string) *Board {

}

func (user *User) CreateBoard(title string) *Board {
	user.system.kanbanService.CreateBoard
}

func (board Board) GetAllCards() []Card {
	var cardList []Card

	for _, card := range board.Cards {
		cardList = append(cardList, card)
	}

	return cardList
}

func (board Board) GetCardsByColumn(column int) []Card {
	var cardList []Card

	return cardList
}

func (service *BoardService) CreateBoard(title string) int {
}

func (service *BoardService) AddCard() {
}

func (service *BoardService) MoveCard() {
}
