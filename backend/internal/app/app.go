package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heismyke/backend_hospital/api"
	"github.com/heismyke/backend_hospital/migrations"
	"github.com/heismyke/backend_hospital/store"
)


type Application struct{
  Logger *log.Logger
  UserHandler *api.UserHandler
  DB *sql.DB
}


func NewApplication() (*Application, error){
  logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile) 
  pgDB,err := store.Open()
  if err != nil{
    return nil, err
  }
  err = store.MigrateFS(pgDB, migrations.FS, "migrations")
  if err != nil{
    panic(err)
  }
  userStore := store.NewPostgresUserStore(pgDB)
  userHandler := api.NewUserHandler(userStore, logger)
  return &Application{
    Logger : logger,
    UserHandler: userHandler,
    DB: pgDB,
  },nil
}


func (a *Application) CheckHealthStatus(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "status us live") 
}