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
	Id          int
	State       AccountState
	CreatedWhen time.Time
}

type Profile struct {
	Email    string
	Username string
	Alias    string
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
		if board.OwnedBy == user.Id {
			boardList = append(boardList, board)
		}
	}

	return boardList
}
