package oxpit

import (
	"net/http"
)

type IdentityToken struct {
	Token
}

type AccessToken struct {
	Token
}

/*type User struct {
	Id IdentityToken
}*/

type Application struct {
	Id IdentityToken
}

type Actor struct {
	User *User
	App  *Application
}

// The service provider interface (SPI)
type AccountService interface {
	NewUser() *User
}

type IdentityService interface {
	GetUser(user IdentityToken) *User
}

type AuthenticationService interface {
	User(appToken AccessToken, accessKey string, secretKey string) AccessToken
	Application(accessKey IdentityToken, secretKey string) AccessToken
}

type AuthorizationService interface {
	Application(appToken AccessToken) bool
	User(userToken AccessToken) bool
}

type KanbanService interface {
	CreateBoard(actor Actor, user User, title string) *Board
}

type ServiceLocator interface {
	Accounts() AccountService
	Identity() IdentityService
	Authentication() AuthenticationService
	Authorization() AuthorizationService
	Kanban() KanbanService
}

type ServiceProvider interface {
	Services(actor Actor) *ServiceLocator
}

// One implementation (the best) of the SPI
/*
type BestServiceProvider struct {
}

func NewBestServiceProvider() *BestServiceProvider {
	return new(BestServiceProvider)
}

func (sp *BestServiceProvider) Services(userToken AccessToken, appToken AccessToken) *ServiceLocator {
	locator := new(BestServiceLocatorInstance)
	locator.user = GetUser(userToken)
	locator.app = GetApplication(appToken)

	return locator
}

func System(actor AccessToken, app AccessToken) *SystemContext {
}*/

func Identities(request *http.Request) (AccessToken, AccessToken, bool) {
	var userToken, appToken AccessToken
	var authorized bool = false

	if atStr, utStr, ok := request.BasicAuth(); ok {
		//service := provider.Services().Authorization()

		appToken, _ = NewAccessTokenFromBase32(atStr)
		//authorized := service.ApplicationAuthorized(at)

		if authorized {
			userToken, _ = NewAccessTokenFromBase32(utStr)
			//authorized = authorizationService.UserAuthorized(ut)
		}
	}

	return userToken, appToken, authorized
}

/*func authenticateApp(writer http.ResponseWriter, request *http.Request) {
	if accessKey, secretKey, ok := request.BasicAuth(); ok {
		if accessToken, authenticated := authenticationService.AuthenicateApplication(accessKey, secretKey); authenticated {
			writer.Header().Set("Location", url.String())
			writeJsonWithCode(writer, profile, http.StatusCreated)
		}
	}
}

func createBoard(writer http.ResponseWriter, request *http.Request) {
	if ut, at, authorized := Identities(request); authorized {
		user := provider.Services(at).User(ut)
		user.CreateBoard(title)
	} else {
		writer.WriterHeader(http.StatusUnauthorized)
	}
}*/
