package main

import (
	"github.com/acamilleri/go-plexhooks/plex"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/acamilleri/go-plexhooks"

	"github.com/acamilleri/myplexhooks/internal/events"
	"github.com/acamilleri/myplexhooks/pkg/grafana"
)

var (
	version string

	listenAddrFlag = kingpin.Flag("listen.addr", "listen address for http server").
		Envar("PLEXHOOKS_LISTEN_ADDR").
		Default("0.0.0.0:8080").
		TCP()

	logLevelFlag = kingpin.Flag("log.level", "log level").
		Envar("PLEXHOOKS_LOG_LEVEL").
		Default(logrus.InfoLevel.String()).
		Enum(logrus.DebugLevel.String(), logrus.InfoLevel.String(), logrus.ErrorLevel.String())

	logFormatterFlag = kingpin.Flag("log.formatter", "log formatter").
		Envar("PLEXHOOKS_LOG_FORMATTER").
		Default("text").
		Enum("text", "json")

	grafanaURLFlag = kingpin.Flag("grafana.url", "grafana url (eg: http://grafana.local)").
		Envar("PLEXHOOKS_GRAFANA_URL").
		Required().
		String()

	grafanaTokenFlag = kingpin.Flag("grafana.token", "grafana api token").
		Envar("PLEXHOOKS_GRAFANA_TOKEN").
		Required().
		String()
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	grafanaClient, err := grafana.New(*grafanaURLFlag, *grafanaTokenFlag)
	if err != nil {
		panic(err)
	}

	actions := plexhooks.NewActions()
	// Grafana actions
	actions.Add(plex.MediaPlay, &events.AddGrafanaAnnotationOnMediaPlay{Client: grafanaClient})
	actions.Add(plex.MediaPause, &events.UpdateGrafanaAnnotationOnMediaPause{Client: grafanaClient})
	actions.Add(plex.MediaResume, &events.UpdateGrafanaAnnotationOnMediaResume{Client: grafanaClient})
	actions.Add(plex.MediaStop, &events.UpdateGrafanaAnnotationOnMediaStop{Client: grafanaClient})

	logLevel, _ := logrus.ParseLevel(*logLevelFlag)
	var logFormatter logrus.Formatter = &logrus.TextFormatter{}
	if *logFormatterFlag == "json" {
		logFormatter = &logrus.JSONFormatter{}
	}

	app := plexhooks.New(plexhooks.Definition{
		ListenAddr: *listenAddrFlag,
		Actions:    actions,
		Logger: plexhooks.LoggerDefinition{
			Level:     logLevel,
			Formatter: logFormatter,
		},
	})

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
