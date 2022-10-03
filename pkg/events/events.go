package events

import (
	"encoding/json"
	"fmt"

	log "github.com/gookit/slog"
)

var eventType GitlabEventTypeDetector

type Event interface {
	TemplateMessage() string
}

func (e GitlabJobEvent) TemplateMessage() string {
	return "JOB EVENT"
}

func (e GitlabMergeRequestEvent) TemplateMessage() string {
	return "MERGE REQUEST EVENT"
}

func (e GitlabNoteEvent) TemplateMessage() string {
	return "NOTE EVENT"
}

func (e GitlabPipelineEvent) TemplateMessage() string {
	return "PIPELINE EVENT"
}

func (e GitlabPushEvent) TemplateMessage() string {
	return "PUSH EVENT OR TAG PUSH EVENT"
}

func EventMatcher(eventJSON []byte) {
	err := json.Unmarshal(eventJSON, &eventType)
	if err != nil {
		log.Error("Cannot unmarshal received event to GitlabTypeDetector structure. Error: ", err)
	}

	log.Info("Received event with type: ", eventType.ObjectKind)
	event, err := eventComparator(eventType.ObjectKind, eventJSON)
	if err != nil {
		log.Error("Error while compare event with struct", err)
	}
	log.Info(event.TemplateMessage())
}

func eventComparator(eventType string, data []byte) (Event, error) {
	switch eventType {
	case "build":
		var gitlabEvent GitlabJobEvent
		return gitlabEvent, json.Unmarshal(data, &gitlabEvent)
	case "merge_request":
		var gitlabEvent GitlabMergeRequestEvent
		return gitlabEvent, json.Unmarshal(data, &gitlabEvent)
	case "note":
		var gitlabEvent GitlabNoteEvent
		return gitlabEvent, json.Unmarshal(data, &gitlabEvent)
	case "pipeline":
		var gitlabEvent GitlabPipelineEvent
		return gitlabEvent, json.Unmarshal(data, &gitlabEvent)
	case "push":
		var gitlabEvent GitlabPushEvent
		return gitlabEvent, json.Unmarshal(data, &gitlabEvent)
	default:
		return nil, fmt.Errorf("Unknown event type: %s", eventType)
	}
}
