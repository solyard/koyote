package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/gookit/slog"
	"github.com/gorilla/mux"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/events"
)

func returnError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	http.Error(w, "Please add chatID like in example: /notify/123123123123", http.StatusNotFound)
}

func receiveEventJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Error while read payload from request to Koyote. Error: ", err)
	}

	events.EventMatcher(body, vars["chat_id"])
}

func StartPolling() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/notify/{chat_id}", receiveEventJSON).Methods("POST")
	router.HandleFunc("/notify", returnError)

	log.Info("Starting API on port", config.GlobalAppConfig.Global.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.GlobalAppConfig.Global.ListenPort), router))
}
