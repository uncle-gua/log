package log_test

import (
	"log"
	"net/http"

	"github.com/uncle-gua/log"
)

func ExampleLogger_Writer_httpServer() {
	logger := log.New()
	w := logger.Writer()
	defer w.Close()

	srv := http.Server{
		// create a stdlib log.Logger that writes to
		// log.Logger.
		ErrorLog: log.New(w, "", 0),
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ExampleLogger_Writer_stdlib() {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}

	// Use log for standard log output
	// Note that `log` here references stdlib's log
	// Not log imported under the name `log`.
	log.SetOutput(logger.Writer())
}
