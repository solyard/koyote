package api

import (
	"io/ioutil"
	"net/http"

	log "github.com/gookit/slog"
	"github.com/gorilla/mux"
)

func receiveEventJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Info(body)
}

func InitialiseAPI() {
	log.Info("Launching API and Loading configuration")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/notify", receiveEventJSON).Methods("POST")

	log.Info("Configuration Loaded! Starting API ...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
