package spi

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

type User struct {
	Id IdentityToken
}

type Application struct {
	Id IdentityToken
}

type Actor struct {
	User *User
	App  *Application
}

func NewUser(provider ServiceProvider) {
}

func (user User) CreateBoard(title string) *Board {
	locator.Kanban().CreateBoard(app.Id, user.Id)
}

// One implementation (the best) of the SPI

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
}

func Identities(request *http.Request) (AccessToken, AccessToken, bool) {
	var userToken, appToken AccessToken
	var authorized bool = false

	if utStr, atStr, ok := request.BasicAuth(); ok {
		service := provider.Services().Authorization()

		at = AccessTokenFromString(atStr)
		authorized := service.ApplicationAuthorized(at)

		if authorized {
			ut = AccessTokenFromString(utStr)
			authorized = authorizationService.UserAuthorized(ut)
		}
	}

	return userToken, appToken, authorized
}

func authenticateApp(writer http.ResponseWriter, request *http.Request) {
	if accessKey, secretKey, ok := request.BasicAuth(); ok {
		if accessToken, authenticated := authenticationService.AuthenicateApplication(accessKey, secretKey); authenticated {
			writer.Header().Set("Location", url.String())
			writeJsonWithCode(writer, profile, http.StatusCreated)
		}
	}
}

func createBoard(writer http.ResponseWriter, request *http.Request) {
	if ut, at, authorized := Identitites(request); authorized {
		user := provider.Services(at).User(ut)
		user.CreateBoard(title)
	} else {
		writer.WriterHeader(http.StatusUnauthorized)
	}
}
