package main

import (
	"encoding/json"
	"fmt"
	"github.com/00Soul/oxpit"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ServiceContext struct {
	router *mux.Router
}

var persistentServiceContext *ServiceContext

func GetServiceContext() *ServiceContext {
	if persistentServiceContext == nil {
		persistentServiceContext = &ServiceContext{mux.NewRouter()}
	}

	return persistentServiceContext
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	user := oxpit.NewUser()

	context := GetServiceContext()
	route := context.router.Get("user")
	url, err := route.URL("user-id", strconv.Itoa(user.Id))
	if err != nil {
		url, _ = url.Parse("http://localhost:8088/try/this")
	}

	header := writer.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Location", url.String())

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "%s", jsonFromUser(user))
}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	userList := oxpit.GetSystem().GetUsers()

	users := make([]interface{}, 0, 1)
	for _, u := range userList {
		users = append(users, interfaceFromUser(u))
	}

	jsonString, err := json.Marshal(users)
	if err != nil {
		jsonString = []byte("{\"error\":\"json.Marshal() failed\"}")
	}

	header := writer.Header()
	header.Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", jsonString)

	mapping := oxpit.json.NewMapping(oxpit.User)
	mapping.Field("Id").Name("id")
	mapping.Field("State").Name("state").MarshalFunc(acctState2Json)
	mapping.Field("State").UnmarshalFunc(func(data []byte) (interface{}, error) {
		jsonString, ok := data.(string)
		when, err := time.Parse(timeLayout, jsonString)
	})
	mapping.Field("CreatedWhen").Name("created-when")
	mapping.Field("CreatedWhen").MarshalFunc(toTime)
	mapping.Field("CreatedWhen").UnmarshalFunc(fromTime)
	fmt.Fprintf(writer, "%s", oxpit.json.toJson(users))
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["user-id"])
	user, found := oxpit.GetSystem().GetUser(userId)

	if !found {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		header := writer.Header()
		header.Set("Content-Type", "application/json")

		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "%s", jsonFromUser(user))
	}
}

func createProfile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	userId, _ := strconv.Atoi(vars["user-id"])
	user, found := oxpit.GetSystem().GetUser(userId)

	if !found {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		bytes := make([]byte, request.ContentLength)
		_, err := request.Read(bytes)
		if err == nil {
			var profile = toProfileFromBytes(bytes)

			user.SetProfile(profile)

			header := writer.Header()
			header.Set("Content-Type", "application/json")

			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, "%s", jsonFromProfile(user.GetProfile()))
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func updateProfile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	userId, _ := strconv.Atoi(vars["user-id"])
	user, found := oxpit.GetSystem().GetUser(userId)

	if !found {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		bytes := make([]byte, request.ContentLength)
		_, err := request.Read(bytes)
		if err == nil {
			var profile = toProfileFromBytes(bytes)

			user.SetProfile(profile)

			header := writer.Header()
			header.Set("Content-Type", "application/json")

			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, "%s", jsonFromProfile(user.GetProfile()))
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getProfile(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	userId, _ := strconv.Atoi(vars["user-id"])
	user, userFound := oxpit.GetSystem().GetUser(userId)

	if !userFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		profile, profileFound := user.GetProfile()
		if !profileFound {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			header := writer.Header()
			header.Set("Content-Type", "application/json")

			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, "%s", jsonFromProfile(profile))
		}
	}
}

func createBoard(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	userId, _ := strconv.Atoi(vars["user-id"])
	user, userFound := oxpit.GetSystem().GetUser(userId)

	if !userFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		bytes := make([]byte, request.ContentLength)
		_, err := request.Read(bytes)
		if err == nil {
			receivedBoard := toBoardFromBytes(bytes)
			newBoard := user.CreateBoard(receivedBoard.Title)

			header := writer.Header()
			header.Set("Content-Type", "application/json")

			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, "%s", jsonFromBoard(newBoard))
		}
	}
}

func listBoards(writer http.ResponseWriter, request *http.Request) {
}

func viewBoard(writer http.ResponseWriter, request *http.Request) {
}

func modifyBoard(writer http.ResponseWriter, request *http.Request) {
}

func deleteBoard(writer http.ResponseWriter, request *http.Request) {
}

func createCard(writer http.ResponseWriter, request *http.Request) {
}

func listCards(writer http.ResponseWriter, request *http.Request) {
}

func inspectCard(writer http.ResponseWriter, request *http.Request) {
}

func modifyCard(writer http.ResponseWriter, request *http.Request) {
}

func deleteCard(writer http.ResponseWriter, request *http.Request) {
}

func setupRoutes() {
	context := GetServiceContext()

	usersRouter := context.router.Path("/api/users").Name("users").Subrouter()
	usersRouter.Methods("POST").HandlerFunc(createUser)
	usersRouter.Methods("GET").HandlerFunc(listUsers)

	userRouter := context.router.Path("/api/users/{user-id}").Name("user").Subrouter()
	userRouter.Methods("GET").HandlerFunc(getUser)

	userProfileRouter := context.router.Path("/api/users/{user-id}/profile").Name("user").Subrouter()
	userProfileRouter.Methods("POST").HandlerFunc(createProfile)
	userProfileRouter.Methods("PUT").HandlerFunc(updateProfile)
	userProfileRouter.Methods("GET").HandlerFunc(getProfile)

	boardsRouter := context.router.Path("/api/users/{user-id}/boards").Subrouter()
	boardsRouter.Methods("POST").HandlerFunc(createBoard)
	boardsRouter.Methods("GET").HandlerFunc(listBoards)

	boardRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}").Subrouter()
	boardRouter.Methods("GET").HandlerFunc(viewBoard)
	boardRouter.Methods("PUT").HandlerFunc(modifyBoard)
	boardRouter.Methods("DELETE").HandlerFunc(deleteBoard)

	cardsRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}/cards").Subrouter()
	cardsRouter.Methods("POST").HandlerFunc(createCard)
	cardsRouter.Methods("GET").HandlerFunc(listCards)

	cardRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}/cards/{card-id}").Subrouter()
	cardRouter.Methods("GET").HandlerFunc(inspectCard)
	cardRouter.Methods("PUT").HandlerFunc(modifyCard)
	cardRouter.Methods("DELETE").HandlerFunc(deleteCard)

	http.ListenAndServe(":8088", context.router)
}
