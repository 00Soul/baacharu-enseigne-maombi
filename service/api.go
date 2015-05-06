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
	//header.Set("Location", "http://localhost:8088/try/this")

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "%s", jsonFromUser(user))
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["user-id"])
	user, found := oxpit.GetSystem().GetUser(userId)

	log.Printf("DEBUG: Found user -")
	log.Printf("DEBUG:    user.Id: %d", user.Id)
	log.Printf("DEBUG:    user.State: %d", jsonFromAccountState(user.State))
	log.Printf("DEBUG:    user.CreatedWhen: %d", user.CreatedWhen.Format(timeLayout))

	if !found {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		header := writer.Header()
		header.Set("Content-Type", "application/json")

		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "%s", jsonFromUser(user))
	}
}

func getUserList(writer http.ResponseWriter, request *http.Request) {
	userList := oxpit.GetSystem().GetUsers()

	log.Printf("DEBUG: length of userList is %d", len(userList))

	users := make([]int, len(userList))
	//users := []int{0, 1, 2, 3}
	for i, u := range userList {
		//users = append(users, interfaceFromUser(*u))
		//log.Printf("DEBUG: user=%s", interfaceFromUser(*u))
		//uid := strconv.Itoa(u.Id)
		log.Printf("DEBUG: user.Id=%s", u.Id)
		//users = append(users, u.Id)
		users[i] = u.Id
	}
	log.Printf("DEBUG: ...outside range loop")

	jsonString, err := json.Marshal(users)
	if err != nil {
		jsonString = []byte("{\"error\":\"json.Marshal() failed\"}")
	}
	log.Printf("DEBUG: ...just Marshalled users slice")
	log.Printf("DEBUG: jsonString=\"%s\"", jsonString)

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%s", jsonString)
}

func setupRoutes() {
	context := GetServiceContext()

	apiRouter := context.router.PathPrefix("/api").Subrouter()

	usersRouter := apiRouter.Path("/users").Name("users").Subrouter()
	usersRouter.Methods("POST").HandlerFunc(createUser)
	usersRouter.Methods("GET").HandlerFunc(getUserList)

	userRouter := usersRouter.Path("/{user-id}").Name("user").Subrouter()
	userRouter.Methods("GET").HandlerFunc(getUser)

	//boardsRouter := userRouter.Path("/boards").Subrouter()
	//boardRouter := boardsRouter.Path("/{board-id}").Subrouter()

	http.ListenAndServe(":8088", context.router)
}
