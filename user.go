package oxpit

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
	return accountStates[state]
}

type User struct {
	Id          IdentityToken
	State       AccountState
	CreatedWhen time.Time
}

type Profile struct {
	Email    string
	Username string
	Alias    string
}

/*func NewUser() User {
	return GetSystem().createUser()
}

func (user User) CreateBoard(title string) Board {
	return GetSystem().createBoard(user, title)
}

func (user User) GetBoards() []Board {
	boardList := make([]Board, 0, 5)
	for _, board := range GetSystem().boards {
		if board.OwnedBy == user.Id {
			boardList = append(boardList, board)
		}
	}

	return boardList
}

func (user User) GetBoardById(id int) (Board, bool) {
	board, found := GetSystem().boards[id]
	if found && (board.OwnedBy != user.Id) {
		board = Board{}
		found = false
	}

	return board, found
}

func (user User) GetBoardByTitle(title string) (Board, bool) {
	var board Board
	boardList := user.GetBoards()
	found := false

	title = strings.ToLower(title)
	for i := 0; (i < len(boardList)) && !found; i++ {
		if strings.ToLower(boardList[i].Title) == title {
			board = boardList[i]
			found = true
		}
	}

	return board, found
}

func (user User) SetProfile(profile Profile) {
	_, found := user.GetProfile()
	if !found {
		GetSystem().createProfile(user, profile)
	} else {
		GetSystem().updateProfile(user, profile)
	}
}

func (user User) GetProfile() (Profile, bool) {
	return GetSystem().getProfile(user)
}*/
