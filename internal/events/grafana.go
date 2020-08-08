package events

import (
	"fmt"
	"time"

	"github.com/acamilleri/go-plexhooks/plex"

	"github.com/acamilleri/myplexhooks/pkg/grafana"
)

type AddGrafanaAnnotationOnMediaPlay struct {
	*grafana.Client
}

func (a AddGrafanaAnnotationOnMediaPlay) Name() string {
	return "add Grafana annotation on play event"
}

func (a AddGrafanaAnnotationOnMediaPlay) Execute(event plex.Event) error {
	var annotationMsg string

	switch event.Metadata.Type {
	case plex.EpisodeMetadataType:
		annotationMsg = fmt.Sprintf("[%s]\n %s:S%d:E%d - %s(%s) by %s on %s",
			event.Name,
			event.Metadata.GrandparentTitle,
			event.Metadata.ParentIndex,
			event.Metadata.Index,
			event.Metadata.Title,
			event.Metadata.Type,
			event.Account.Title,
			event.Player.Title,
		)
	case plex.MovieMetadataType:
		annotationMsg = fmt.Sprintf("[%s]\n %s(%s) by %s on %s",
			event.Name,
			event.Metadata.Title,
			event.Metadata.Type,
			event.Account.Title,
			event.Player.Title,
		)
	default:
		return fmt.Errorf("unknown metadata type: %s", event.Metadata.Type)

	}

	return a.CreateAnnotation(grafana.CreateAnnotationRequest{
		Time:    time.Now().UTC().UnixNano() / 1000000,
		TimeEnd: time.Now().UTC().UnixNano() / 1000000,
		Tags: []string{
			"plex",
			event.Account.Title,
			event.Metadata.GUID,
		},
		Text: annotationMsg,
	})
}

type UpdateGrafanaAnnotationOnMediaPause struct {
	*grafana.Client
}

func (a UpdateGrafanaAnnotationOnMediaPause) Name() string {
	return "update Grafana annotation on pause event"
}

func (a UpdateGrafanaAnnotationOnMediaPause) Execute(event plex.Event) error {
	annotation, err := a.FindAnnotation(grafana.FindAnnotationRequest{
		From: time.Now().UTC().Add(time.Duration(-4)*time.Hour).UnixNano() / 1000000,
		Tags: []string{
			"plex",
			event.Account.Title,
			event.Metadata.GUID,
		},
		Limit: 1,
	})
	if err != nil {
		return err
	}

	return a.UpdateAnnotation(grafana.UpdateAnnotationRequest{
		ID:      annotation.ID,
		TimeEnd: time.Now().UTC().UnixNano() / 1000000,
		Text:    fmt.Sprintf("%s \n Pause: UTC(%s)", annotation.Text, time.Now().UTC().Format(time.RFC3339)),
	})
}

type UpdateGrafanaAnnotationOnMediaResume struct {
	*grafana.Client
}

func (a UpdateGrafanaAnnotationOnMediaResume) Name() string {
	return "update Grafana annotation on pause event"
}

func (a UpdateGrafanaAnnotationOnMediaResume) Execute(event plex.Event) error {
	annotation, err := a.FindAnnotation(grafana.FindAnnotationRequest{
		From: time.Now().UTC().Add(time.Duration(-4)*time.Hour).UnixNano() / 1000000,
		Tags: []string{
			"plex",
			event.Account.Title,
			event.Metadata.GUID,
		},
		Limit: 1,
	})
	if err != nil {
		return err
	}

	return a.UpdateAnnotation(grafana.UpdateAnnotationRequest{
		ID:      annotation.ID,
		TimeEnd: time.Now().UTC().UnixNano() / 1000000,
		Text:    fmt.Sprintf("%s \n Resume: UTC(%s)", annotation.Text, time.Now().UTC().Format(time.RFC3339)),
	})
}

type UpdateGrafanaAnnotationOnMediaStop struct {
	*grafana.Client
}

func (a UpdateGrafanaAnnotationOnMediaStop) Name() string {
	return "update Grafana annotation on pause event"
}

func (a UpdateGrafanaAnnotationOnMediaStop) Execute(event plex.Event) error {
	annotation, err := a.FindAnnotation(grafana.FindAnnotationRequest{
		From: time.Now().UTC().Add(time.Duration(-4)*time.Hour).UnixNano() / 1000000,
		Tags: []string{
			"plex",
			event.Account.Title,
			event.Metadata.GUID,
		},
		Limit: 1,
	})
	if err != nil {
		return err
	}

	return a.UpdateAnnotation(grafana.UpdateAnnotationRequest{
		ID:      annotation.ID,
		TimeEnd: time.Now().UTC().UnixNano() / 1000000,
		Text:    fmt.Sprintf("%s \n Stop: UTC(%s)", annotation.Text, time.Now().UTC().Format(time.RFC3339)),
	})
}
