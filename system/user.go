package system

import (
	"time"
)

type AccountState int

const (
	AccountActive AccountState = iota
	AccountInactive
	AccountClosed
)

var accountStates = [...]string{
	"AccountActive",
	"AccountInactive",
	"AccountClosed",
}

func (state AccountState) String() string {
	return accountStates[s]
}

type User struct {
	id          int
	state       AccountState
	createdWhen time.Time
}

type Profile struct {
	email    string
	username string
	alias    string
}

func CreateUser() *User {
	return getSystem().createUser()
}

func (user *User) CreateBoard(title string) *Board {
	return getSystem().createBoard(user, title)
}

func (user *User) GetBoards() []*Board {
	boardList := make([]*Board, 0, 5)
	for _, board := range getSystem().boards {
		if board.ownedBy == user.id {
			boardList = append(boardList, board)
		}
	}

	return boardList
}
