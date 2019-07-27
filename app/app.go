package app

import (
	"fmt"
	"log"
	"net/http"

	"goToken/app/sync"
	"goToken/config"
	
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Not connect db")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {

	a.Get("/hashes", a.handleRequest(sync.GetAllTokens))
	a.Post("/hashes", a.handleRequest(sync.CreateTokens))
	a.Get("/hashes/{hash}", a.handleRequest(sync.GetToken))
	a.Delete("/hashes/{hash}", a.handleRequest(sync.DeleteToken))

}

// Router for GET method
func (a *App) Get(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Router for POST method
func (a *App) Post(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Router for PUT method
func (a *App) Put(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Router for DELETE method
func (a *App) Delete(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, response http.ResponseWriter, request *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		handler(a.DB, response, request)
	}
}
