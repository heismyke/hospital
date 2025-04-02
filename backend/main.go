package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/heismyke/backend_hospital/internal/app"
	"github.com/heismyke/backend_hospital/routes"
)

func main(){
  var port int
  flag.IntVar(&port, "port", 8080, "go backend server")
  flag.Parse()
  myApp, err := app.NewApplication()
  if err != nil{
    panic(err)
  }
  defer myApp.DB.Close()
  r := routes.SetupRoutes(myApp)
  server := &http.Server{
    Addr: fmt.Sprintf(":%d", port), 
    Handler: r,
    IdleTimeout: time.Minute,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 30 * time.Second,
  }
  myApp.Logger.Printf("server running on port: %v", port)
  err = server.ListenAndServe()
  if err != nil{
    myApp.Logger.Fatalf("error starting server: %v", err)
  }
}
