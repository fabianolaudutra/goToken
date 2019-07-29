package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"	
	"github.com/fabianolaudutra/goToken/app/sync"
	"github.com/fabianolaudutra/goToken/config"
	"github.com/fabianolaudutra/goToken/app/model"	
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
		log.Fatal("Nao conectou db")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {

	a.Get("/hashes", a.handleRequest(sync.GetAllTokens))
	a.Post("/hashe", a.handleRequest(sync.CreateTokens))
	a.Get("/hashes/{token}", a.handleRequest(sync.GetToken))
	//a.Delete("/hashe/{token}", a.handleRequest(sync.DeleteToken))

}


func (a *App) Get(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}


func (a *App) Post(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}


func (a *App) Put(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}


func (a *App) Delete(path string, f func(response http.ResponseWriter, request *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}


func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, response http.ResponseWriter, request *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		handler(a.DB, response, request)
	}
}
