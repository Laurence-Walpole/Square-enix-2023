package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var db *sql.DB
var e *echo.Echo

func main() {
	//Create echo object
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	_ = createDBConnection()

	//Define routes
	routes(e)

	//Start server
	err := e.Start(getVar("SERVER_HOST" + ":" + getVar("SERVER_PORT")))
	if err != nil {
		return
	}
}

func startServer() {
	go func() {
		if err := e.Start(getVar("SERVER_HOST" + ":" + getVar("SERVER_PORT"))); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer closeDB()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("h" + err.Error())
	}
}

func createDBConnection() *sql.DB {
	cfg := mysql.Config{
		User:   getVar("DB_USER"),
		Passwd: getVar("DB_PASS"),
		Net:    getVar("DB_NET"),
		Addr:   getVar("DB_HOST") + ":" + getVar("DB_PORT"),
		DBName: getVar("DB_NAME"),
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db
}

func closeDB() bool {
	return db.Close() == nil
}

func getVar(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
