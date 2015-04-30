package restapi

import (
	"bem/system"
	"encoding/json"
	"gorilla/mux"
	"net/http"
)

func createUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write(system.GetSystem().CreateUser().json())
}

func setupRoutes() {
	router := mux.NewRouter()

	apiRoute := router.PathPrefix("/api").Subrouter()

	usersRoute := apiRoute.Path("/users").Subrouter()
	usersRoute.Methods("POST").HandleFunc(createUser)

	userRoute := usersRoute.Path("/{user-id}").Subrouter()

	boardsRoute := userRoute.Path("/boards").Subrouter()
	boardRoute := boardRoute.Path("/{board-id}").Subrouter()

	http.ListenAndServe(":8088", router)
}
