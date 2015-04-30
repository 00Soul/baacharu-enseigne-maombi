package system

import (
	"time"
)

type System struct {
	counter int
	users   map[int]*User
	boards  map[int]*Board
}

func NewSystem() *System {
	system := new(System)
	system.counter = 0
	system.boards = make(map[int]*Board)
	system.users = make(map[int]*User)

	return system
}

var cachedSystem *System = nil

func GetSystem() *System {
	if cachedSystem == nil {
		cachedSystem = NewSystem()
	}

	return cachedSystem
}

func (system *System) createUser() *User {
	user := new(User)
	user.id = system.counter
	user.state = AccountActive
	user.createdWhen = time.Now().UTC()

	system.counter++
	system.users[user.id] = &user

	return &user
}

func (system *System) GetUsers() []*User {
	userList := make([]*User, len(system.users))
	i := 0
	for _, user := range system.users {
		userList[i] = user
	}

	return userList
}

func (system *System) GetUser(userId int) (User, bool) {
	return system.users[userId]
}

func (system *System) createBoard(user *User, title string) *Board {
	board := new(Board)
	board.id = system.counter
	board.title = title
	board.columns = make([]Column, 0, 3)
	board.cards = make(map[int]Card)
	board.ownedBy = user.id
	board.createdBy = user.id
	board.createdWhen = time.Now().UTC()

	system.counter++

	return board
}
