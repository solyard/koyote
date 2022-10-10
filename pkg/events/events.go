package events

import (
	"encoding/json"
	"fmt"

	log "github.com/gookit/slog"
	"github.com/koyote/pkg/telegram"
)

type Event interface {
	TemplateMessage() string
}

func (e GitlabJobEvent) TemplateMessage() string {
	result, err := templateJobEventMessage(e, "job.tpl")
	if err != nil {
		log.Error("Error while receiving message from templator. Error: ", err)
	}

	return result
}

func (e GitlabMergeRequestEvent) TemplateMessage() string {
	result, err := templateMREventMessage(e, "merge_request.tpl")
	if err != nil {
		log.Error("Error while receiving message from templator. Error: ", err)
	}

	return result
}

func (e GitlabNoteEvent) TemplateMessage() string {
	result, err := templateNoteEventMessage(e, "note.tpl")
	if err != nil {
		log.Error("Error while receiving message from templator. Error: ", err)
	}

	return result
}

func (e GitlabPipelineEvent) TemplateMessage() string {
	result, err := templatePipelineEventMessage(e, "pipeline.tpl")
	if err != nil {
		log.Error("Error while receiving message from templator. Error: ", err)
	}

	return result
}

func (e GitlabPushEvent) TemplateMessage() string {
	result, err := templatePushEventMessage(e, "push.tpl")
	if err != nil {
		log.Error("Error while receiving message from templator. Error: ", err)
	}

	return result
}

func EventMatcher(eventJSON []byte) {
	var receivedEventType GitlabEventTypeDetector
	err := json.Unmarshal(eventJSON, &receivedEventType)
	if err != nil {
		log.Error("Cannot unmarshal received event to GitlabTypeDetector structure. Error: ", err)
	}

	event, err := eventComparator(receivedEventType.ObjectKind, eventJSON)
	if err != nil {
		log.Error("Error while compare event with struct", err)
		return
	}
	//log.Info(event.TemplateMessage())
	telegram.SendEventMessage(129913666, event.TemplateMessage())
}

func eventComparator(eventType string, data []byte) (Event, error) {
	switch eventType {
	case "build":
		var gitlabEvent GitlabJobEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "merge_request":
		var gitlabEvent GitlabMergeRequestEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "note":
		var gitlabEvent GitlabNoteEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "pipeline":
		var gitlabEvent GitlabPipelineEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "push":
		var gitlabEvent GitlabPushEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	default:
		return nil, fmt.Errorf("Unknown event type: %s", eventType)
	}
}
