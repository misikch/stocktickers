package app

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/misikch/stocktickers/src/internal/config"
)

func Run(cfg *config.Config) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		io.WriteString(writer, `{"code":0, "status":"ok"}`)
	})

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	return nil
}
