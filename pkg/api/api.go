package api

import (
	"io/ioutil"
	"net/http"

	log "github.com/gookit/slog"
	"github.com/gorilla/mux"
	"github.com/koyote/pkg/events"
)

func receiveEventJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	events.EventMatcher(body, vars["chat_id"])
}

func StartPolling() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/notify/{chat_id}", receiveEventJSON).Methods("POST")

	log.Info("Starting API on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
