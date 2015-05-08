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
	system.users = make(map[int]User)
	system.boards = make(map[int]Board)
	system.profiles = make(map[int]Profile)

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
		i++
	}

	return userList
}

func (system System) GetUserById(userId int) (User, bool) {
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

func (system *System) createProfile(user User, profile Profile) {
	system.profiles[user.Id] = profile
}

func (system *System) updateProfile(user User, profile Profile) {
	existingProfile, found := system.getProfile(user)
	if found {
		if profile.Email != "" {
			existingProfile.Email = profile.Email
		}

		if profile.Username != "" {
			existingProfile.Username = profile.Username
		}

		if profile.Alias != "" {
			existingProfile.Alias = profile.Alias
		}
	}
}

func (system System) getProfile(user User) {
	return system.profiles[user.Id]
}
