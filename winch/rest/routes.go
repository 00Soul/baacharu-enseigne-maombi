package main

import (
	"github.com/gorilla/mux"
	"net/http"
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

func setupRoutes() {
	context := GetServiceContext()

	tokensRouter := context.router.Path("/api/tokens").Name("tokens").Subrouter()
	tokensRouter.Methods("POST").HandlerFunc(createAccessToken)

	tokenRouter := context.router.Path("/api/tokens/{id-token}").Name("token").Subrouter()
	tokenRouter.Methods("GET").HandlerFunc(retrieveAccessToken)

	usersRouter := context.router.Path("/api/users").Name("users").Subrouter()
	usersRouter.Methods("POST").HandlerFunc(createUser)
	usersRouter.Methods("GET").HandlerFunc(listUsers)

	userRouter := context.router.Path("/api/users/{user-id}").Name("user").Subrouter()
	userRouter.Methods("GET").HandlerFunc(getUser)

	userProfileRouter := context.router.Path("/api/users/{user-id}/profile").Name("profile").Subrouter()
	userProfileRouter.Methods("POST").HandlerFunc(createProfile)
	userProfileRouter.Methods("PUT").HandlerFunc(updateProfile)
	userProfileRouter.Methods("GET").HandlerFunc(getProfile)

	boardsRouter := context.router.Path("/api/users/{user-id}/boards").Name("boards").Subrouter()
	boardsRouter.Methods("POST").HandlerFunc(createBoard)
	boardsRouter.Methods("GET").HandlerFunc(listBoards)

	boardRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}").Name("board").Subrouter()
	boardRouter.Methods("GET").HandlerFunc(viewBoard)
	boardRouter.Methods("PUT").HandlerFunc(modifyBoard)
	boardRouter.Methods("DELETE").HandlerFunc(deleteBoard)

	cardsRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}/cards").Name("cards").Subrouter()
	cardsRouter.Methods("POST").HandlerFunc(createCard)
	cardsRouter.Methods("GET").HandlerFunc(listCards)

	cardRouter := context.router.Path("/api/users/{user-id}/boards/{board-id}/cards/{card-id}").Name("card").Subrouter()
	cardRouter.Methods("GET").HandlerFunc(inspectCard)
	cardRouter.Methods("PUT").HandlerFunc(modifyCard)
	cardRouter.Methods("DELETE").HandlerFunc(deleteCard)

	http.ListenAndServe(":8088", context.router)
}
