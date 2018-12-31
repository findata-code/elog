package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"time"
)

type elasticLog struct {
	client       *elastic.Client
	elasticUrl   string
	elasticIndex string
	elasticType  string
	app          string
	logBuffer    []logMessage
}

type logMessage struct {
	Severity    string
	Message     string
	Timestamp   time.Time
	Application string
}

var (
	_panic   = "panic"
	_fatal   = "fatal"
	_error   = "error"
	_warning = "warning"
	_info    = "info"
	_debug   = "debug"
)

var instance elasticLog

func Init(elasticUrl, elasticIndex, elasticType, app string) error {

	client, err := elastic.NewClient(
		elastic.SetURL(elasticUrl),
		elastic.SetSniff(false),
	)

	client.CreateIndex(elasticIndex)

	if err != nil {
		return err
	}

	instance = elasticLog{
		client:       client,
		elasticUrl:   elasticUrl,
		elasticIndex: elasticIndex,
		elasticType:  elasticType,
		app:app,
	}

	log.SetOutput(instance)

	return nil
}

func Panic(msg string) {
	write(_panic, msg)
}

func Fatal(msg string) {
	write(_fatal, msg)
}

func Error(msg string) {
	write(_error, msg)
}

func Warning(msg string) {
	write(_warning, msg)
}

func Info(msg string) {
	write(_info, msg)
}

func Debug(msg string) {
	write(_debug, msg)
}

func write(severity, msg string) {
	lm := logMessage{
		Severity:    severity,
		Message:     msg,
		Timestamp:   time.Now(),
		Application: instance.app,
	}

	b, err := json.Marshal(lm)

	if err != nil {
		Error(err.Error())
	}

	log.SetFlags(0)
	log.Print(string(b))
}

func (e elasticLog) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	ctx := context.Background()
	_, err := e.client.Index().
		Index(e.elasticIndex).
		Type(e.elasticType).
		BodyString(string(p)).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return len(p), nil
}
