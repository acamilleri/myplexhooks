package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const annotationsPath = "/api/annotations"

type Annotations []Annotation
type Annotation struct {
	ID          int      `json:"id"`
	AlertID     int      `json:"alertId"`
	DashboardID int      `json:"dashboardId"`
	PanelID     int      `json:"panelId"`
	UserID      int      `json:"userId"`
	UserName    string   `json:"userName"`
	NewState    string   `json:"newState"`
	PrevState   string   `json:"prevState"`
	Time        int64    `json:"time"`
	TimeEnd     int64    `json:"timeEnd"`
	Text        string   `json:"text"`
	Metric      string   `json:"metric"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	Data        struct{} `json:"data"`
}

func (c *Client) FindAnnotation(req FindAnnotationRequest) (*Annotation, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s?%s", c.BaseURL, annotationsPath, v.Encode())

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[%d]: failed to find the first annotation", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var listAnnotations Annotations
	err = json.Unmarshal(data, &listAnnotations)
	if err != nil {
		return nil, err
	}

	if len(listAnnotations) == 0 {
		return nil, fmt.Errorf("failed to find the media start annotation")
	}

	return &listAnnotations[0], nil
}

func (c *Client) CreateAnnotation(req CreateAnnotationRequest) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, annotationsPath)

	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}

	resp, err := c.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[%d]: failed to create annotation", resp.StatusCode)
	}
	return nil
}

func (c *Client) UpdateAnnotation(req UpdateAnnotationRequest) error {
	url := fmt.Sprintf("%s%s/%d", c.BaseURL, annotationsPath, req.ID)

	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}

	resp, err := c.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[%d]: failed to patch annotation %d", resp.StatusCode, req.ID)
	}

	return nil
}

type CreateAnnotationRequest struct {
	Time    int64    `json:"time"`
	TimeEnd int64    `json:"timeEnd"`
	Tags    []string `json:"tags"`
	Text    string   `json:"text"`
}

type FindAnnotationRequest struct {
	From  int64    `url:"from"`
	To    int64    `url:"to"`
	Tags  []string `url:"tags"`
	Limit int      `url:"limit"`
}

type UpdateAnnotationRequest struct {
	ID      int
	Time    int64    `json:"time"`
	TimeEnd int64    `json:"timeEnd"`
	Tags    []string `json:"tags"`
	Text    string   `json:"text"`
}
