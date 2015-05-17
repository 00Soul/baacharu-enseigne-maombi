package main

import (
	"fmt"
	"github.com/00Soul/mappings/json"
	"github.com/00Soul/oxpit"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func readJson(request *http.Request, object interface{}) error {
	return json.DecodeWithContext(request.Body, object, jsonMappingContext)
}

func writeJson(writer http.ResponseWriter, object interface{}) error {
	return writeJsonWithCode(writer, object, http.StatusOK)
}

func writeJsonWithCode(writer http.ResponseWriter, object interface{}, code int) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	return json.EncodeWithContext(writer, object, jsonMappingContext)
}

func createAccessToken(writer http.ResponseWriter, request *http.Request) {
	object, err := json.DecodeMap(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	} else if keyString, found := object["accesskey"]; found {
		accessKey, _ := oxpit.NewIdentityTokenFromBase32(keyString)

		if keyString, found = object["secretkey"]; found {
			secretKey := keyString
			// Let's assume authentication succeeds
			//writer.WriteHeader(http.StatusForbidden)

			accessToken, _ := oxpit.NewAccessTokenWithIdentity(accessKey)
			tokenString := accessToken.ToBase32()

			route := GetServiceContext().router.Get("token")
			url, muxErr := route.URL("id-token", tokenString)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
			} else {
				dtFormat := "2006-01-02T15:04:05-07:00"
				now := time.Now().UTC()
				later := now.AddDate(0, 0, 7)

				lease := make(map[string]string)
				lease["token"] = tokenString
				lease["issued-to"] = accessKey.ToBase32()
				lease["issued-when"] = now.Format(dtFormat)
				lease["expires-when"] = later.Format(dtFormat)

				writer.Header().Set("Location", url.String())
				writeJsonWithCode(writer, lease, http.StatusCreated)
			}
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func retrieveAccessToken(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	lease := make(map[string]string)
	lease["token"] = vars["id-token"]

	writeJson(writer, lease)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	user := oxpit.NewUser()

	route := GetServiceContext().router.Get("user")
	url, err := route.URL("user-id", strconv.Itoa(user.Id))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Header().Set("Location", url.String())
		writeJsonWithCode(writer, user, http.StatusCreated)
	}
}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	userList := oxpit.GetSystem().GetUsers()

	writeJson(writer, userList)
}

func retrieveUser(writer http.ResponseWriter, request *http.Request) (oxpit.User, bool) {
	var user oxpit.User

	vars := mux.Vars(request)
	found := false

	if userId, err := strconv.Atoi(vars["user-id"]); err == nil {
		if user, found = oxpit.GetSystem().GetUserById(userId); !found {
			writer.WriteHeader(http.StatusNotFound)
		}
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	return user, found
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		writeJson(writer, user)
	}
}

func createProfile(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		var profile oxpit.Profile

		readJson(request, &profile)
		user.SetProfile(profile)

		route := GetServiceContext().router.Get("profile")
		url, err := route.URL("user-id", strconv.Itoa(user.Id))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			//url, _ = url.Parse("http://localhost:8088/try/this")
		} else {
			writer.Header().Set("Location", url.String())
			writeJsonWithCode(writer, profile, http.StatusCreated)
		}
	}
}

func updateProfile(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		var profile oxpit.Profile

		readJson(request, &profile)
		user.SetProfile(profile)

		writeJson(writer, profile)
	}
}

func getProfile(writer http.ResponseWriter, request *http.Request) {
	if user, userFound := retrieveUser(writer, request); userFound {
		if profile, profileFound := user.GetProfile(); profileFound {
			writeJson(writer, profile)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}

func createBoard(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		var board oxpit.Board

		readJson(request, &board)
		board = user.CreateBoard(board.Title)

		route := GetServiceContext().router.Get("board")
		url, err := route.URL(
			"user-id", strconv.Itoa(user.Id),
			"board-id", strconv.Itoa(board.Id))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			//url, _ = url.Parse("http://localhost:8088/try/this")
		} else {
			writer.Header().Set("Location", url.String())
			writeJsonWithCode(writer, profile, http.StatusCreated)
		}
	}
}

func listBoards(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		boards := user.GetBoards()
		writeJson(writer, boards)
	}
}

func viewBoard(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	if user, userFound := retrieveUser(writer, request); userFound {
		if boardId, err := strconv.Atoi(vars["board-id"]); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			if board, boardFound := user.GetBoardById(); boardFound {
				writeJson(writer, board)
			} else {
				writer.WriteHeader(http.StatusNotFound)
			}
		}
	}
}

func modifyBoard(writer http.ResponseWriter, request *http.Request) {
}

func deleteBoard(writer http.ResponseWriter, request *http.Request) {
	var userId, boardId int
	var err error

	vars := mux.Vars(request)

	userId, err = strconv.Atoi(vars["user-id"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		boardId, err = strconv.Atoi(vars["board-id"])
	}

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		if !accountService.UserExists(userId) {
			err = fmt.Errorf("User (%d) cannot be found", userId)
		}
	}

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		if !boardService.BoardExists(boardId) {
			err = fmt.Errorf("Board (%d) does not exist", boardId)
		}
	}

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		err = boardService.Delete(userId, boardId)
	}

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func createCard(writer http.ResponseWriter, request *http.Request) {
}

func listCards(writer http.ResponseWriter, request *http.Request) {
	if user, found := retrieveUser(writer, request); found {
		boards := user.GetBoards()
		writeJson(writer, boards)
	}
}

func inspectCard(writer http.ResponseWriter, request *http.Request) {
}

func modifyCard(writer http.ResponseWriter, request *http.Request) {
}

func deleteCard(writer http.ResponseWriter, request *http.Request) {
}
