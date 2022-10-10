package events

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	log "github.com/gookit/slog"
)

func prepareTemplate(eventType, templateFilePath string) (*template.Template, error) {
	templatefuncMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	tplfile, err := os.ReadFile(fmt.Sprintf("config/event_templates/%v", templateFilePath))
	if err != nil {
		return nil, err
	}

	template, err := template.New(eventType).Funcs(templatefuncMap).Parse(string(tplfile))
	if err != nil {
		return nil, err
	}

	return template, nil
}

func templateJobEventMessage(gitlabEvent GitlabJobEvent, fileName string) (string, error) {
	template, err := prepareTemplate(gitlabEvent.ObjectKind, fileName)
	if err != nil {
		log.Error("Error while preparing template for event. Error: ", err)
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, gitlabEvent)
	if err != nil {
		log.Error("Error while executing template. Error: ", err)
		return "", err
	}

	response := fmt.Sprintf("%v", &message)
	return response, nil
}

func templateMREventMessage(gitlabEvent GitlabMergeRequestEvent, fileName string) (string, error) {
	template, err := prepareTemplate(gitlabEvent.ObjectKind, fileName)
	if err != nil {
		log.Error("Error while preparing template for event. Error: ", err)
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, gitlabEvent)
	if err != nil {
		log.Error("Error while executing template. Error: ", err)
		return "", err
	}

	response := fmt.Sprintf("%v", &message)
	return response, nil
}

func templateNoteEventMessage(gitlabEvent GitlabNoteEvent, fileName string) (string, error) {
	template, err := prepareTemplate(gitlabEvent.ObjectKind, fileName)
	if err != nil {
		log.Error("Error while preparing template for event. Error: ", err)
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, gitlabEvent)
	if err != nil {
		log.Error("Error while executing template. Error: ", err)
		return "", err
	}

	response := fmt.Sprintf("%v", &message)
	return response, nil
}

func templatePipelineEventMessage(gitlabEvent GitlabPipelineEvent, fileName string) (string, error) {
	template, err := prepareTemplate(gitlabEvent.ObjectKind, fileName)
	if err != nil {
		log.Error("Error while preparing template for event. Error: ", err)
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, gitlabEvent)
	if err != nil {
		log.Error("Error while executing template. Error: ", err)
		return "", err
	}

	response := fmt.Sprintf("%v", &message)
	return response, nil
}

func templatePushEventMessage(gitlabEvent GitlabPushEvent, fileName string) (string, error) {
	template, err := prepareTemplate(gitlabEvent.ObjectKind, fileName)
	if err != nil {
		log.Error("Error while preparing template for event. Error: ", err)
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, gitlabEvent)
	if err != nil {
		log.Error("Error while executing template. Error: ", err)
		return "", err
	}

	response := fmt.Sprintf("%v", &message)
	return response, nil
}
