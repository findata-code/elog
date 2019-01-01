package elog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"time"
)

type elasticLog struct {
	client       *elastic.Client
	elasticUrl   string
	elasticIndex string
	elasticType  string
	app          string
	version      string
}

type logMessage struct {
	Severity    string
	Message     string
	Timestamp   time.Time
	Application string
	Version     string
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

func Init(elasticUrl, elasticIndex, elasticType, app, version string) error {

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
		app:          app,
		version:      version,
	}

	log.SetOutput(instance)

	return nil
}

func Panic(msg string) {
	write(_panic, msg)
	os.Exit(1)
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
	fmt.Println(fmt.Sprintf("[%s]\t[%s]\t:%s", time.Now().Format("2006-01-02 15:04:05"), severity, msg))
	lm := logMessage{
		Severity:    severity,
		Message:     msg,
		Timestamp:   time.Now(),
		Application: instance.app,
		Version: instance.version,
	}

	b, err := json.Marshal(lm)

	if err != nil {
		Error(err.Error())
	}

	log.SetFlags(0)
	log.Print(string(b))
}

func (e elasticLog) Write(p []byte) (int, error) {
	ctx := context.Background()
	_, err := e.client.Index().
		Index(e.elasticIndex).
		Type(e.elasticType).
		BodyString(string(p)).
		Do(ctx)

	if err != nil {
		return 0, err
	}

	return len(p), nil
}
