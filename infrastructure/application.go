package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type application struct {
	config   *config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func Application() *application {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("failed to retrieve env variable", err)
	}
	return &application{
		config: &config{
			Port: port,
		},
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
	}
}

func (app application) Serve(mux http.Handler) error {

	server := http.Server{
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
		Addr:              fmt.Sprintf(":%d", app.config.Port),
	}

	err := server.ListenAndServe()
	return err
}
