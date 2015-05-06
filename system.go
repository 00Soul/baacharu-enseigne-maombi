package oxpit

import (
	"time"
)

type System struct {
	counter int
	users   map[int]User
	boards  map[int]Board
}

func NewSystem() *System {
	system := new(System)
	system.counter = 0
	system.boards = make(map[int]Board)
	system.users = make(map[int]User)

	return system
}

var cachedSystem *System = nil

func GetSystem() *System {
	if cachedSystem == nil {
		cachedSystem = NewSystem()
	}

	return cachedSystem
}

func (system *System) createUser() User {
	user := new(User)
	user.Id = system.counter
	user.State = AccountActive
	user.CreatedWhen = time.Now().UTC()

	system.counter++
	system.users[user.Id] = *user

	return *user
}

func (system *System) GetUsers() []User {
	userList := make([]User, len(system.users))
	i := 0
	for _, user := range system.users {
		userList[i] = user
	}

	return userList
}

func (system System) GetUser(userId int) (User, bool) {
	var user User
	var ok bool

	user, ok = system.users[userId]

	return user, ok
}

func (system *System) createBoard(user User, title string) Board {
	board := new(Board)
	board.Id = system.counter
	board.Title = title
	board.Columns = make([]Column, 0, 3)
	board.Cards = make(map[int]Card)
	board.OwnedBy = user.Id
	board.CreatedBy = user.Id
	board.CreatedWhen = time.Now().UTC()

	system.counter++
	system.boards[board.Id] = *board

	return *board
}
