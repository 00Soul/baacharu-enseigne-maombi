package oxpit

import (
	"fmt"
	"github.com/gorilla/mux"
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
	user := NewUser()

	context := GetServiceContext()
	url, _ := context.router.Get("user").URL("user-id", strconv.Itoa(user.Id))

	header := writer.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Location", url.String())

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "%s", user.json())
}

func SetupRoutes() {
	context := GetServiceContext()

	apiRouter := context.router.PathPrefix("/api").Subrouter()

	usersRouter := apiRouter.Path("/users").Name("users").Subrouter()
	usersRouter.Methods("POST").HandlerFunc(createUser)

	//userRouter := usersRouter.Path("/{user-id}").Subrouter()

	//boardsRouter := userRouter.Path("/boards").Subrouter()
	//boardRouter := boardsRouter.Path("/{board-id}").Subrouter()

	http.ListenAndServe(":8088", context.router)
}
