package go_tests

import (
	"errors"
	"fmt"
	"github.com/cloudevents/sdk-go/v2"
	"github.com/imroc/req"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/kubernetes-utils/pkg"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const defaultKeptnNamespace = "keptn"

const KeptnSpecVersion = "0.2.0"

type APIEventSender struct {
}

func (sender *APIEventSender) SendEvent(event v2.Event) error {
	_, err := ApiPOSTRequest("/v1/event", event)
	return err
}

func ApiGETRequest(path string) (*req.Resp, error) {
	apiToken, keptnAPIURL, err := GetApiCredentials()
	if err != nil {
		return nil, err
	}

	authHeader := req.Header{
		"Accept":  "application/json",
		"x-token": apiToken,
	}

	r, err := req.Get(keptnAPIURL+path, authHeader)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func ApiPOSTRequest(path string, payload interface{}) (*req.Resp, error) {
	apiToken, keptnAPIURL, err := GetApiCredentials()
	if err != nil {
		return nil, err
	}

	authHeader := req.Header{
		"Accept":  "application/json",
		"x-token": apiToken,
	}

	r, err := req.Post(keptnAPIURL+path, authHeader, req.BodyJSON(payload))
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetApiCredentials() (string, string, error) {
	apiToken, err := keptnkubeutils.GetKeptnAPITokenFromSecret(false, os.Getenv("KEPTN_NAMESPACE"), "keptn-api-token")
	if err != nil {
		return "", "", err
	}
	keptnAPIURL := os.Getenv("KEPTN_ENDPOINT")
	if keptnAPIURL == "" {
		serviceIP, err := keptnkubeutils.GetKeptnEndpointFromService(false, os.Getenv("KEPTN_NAMESPACE"), "api-gateway-nginx")
		if err != nil {
			return "", "", err
		}
		keptnAPIURL = "http://" + serviceIP + "/api"
	}
	return apiToken, keptnAPIURL, nil
}

func ScaleDownUniform(deployments []string) error {
	for _, deployment := range deployments {
		if err := keptnkubeutils.ScaleDeployment(false, deployment, os.Getenv("KEPTN_NAMESPACE"), 0); err != nil {
			// log the error but continue
			fmt.Println("could not scale down deployment: " + err.Error())
		}
	}
	return nil
}

func ScaleUpUniform(deployments []string) error {
	for _, deployment := range deployments {
		if err := keptnkubeutils.ScaleDeployment(false, deployment, os.Getenv("KEPTN_NAMESPACE"), 1); err != nil {
			// log the error but continue
			fmt.Println("could not scale up deployment: " + err.Error())
		}
	}
	return nil
}

func CreateTmpShipyardFile(shipyardContent string) (string, error) {
	file, err := ioutil.TempFile(".", "shipyard-*.yaml")
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(file.Name(), []byte(shipyardContent), os.ModeAppend); err != nil {
		os.Remove(file.Name())
		return "", err
	}
	return file.Name(), nil
}

func ExecuteCommand(cmd string) (string, error) {
	split := strings.Split(cmd, " ")
	if len(split) == 0 {
		return "", errors.New("invalid command")
	}
	return keptnkubeutils.ExecuteCommand(split[0], split[1:])
}

func PrepareEnvVars() {
	if os.Getenv("KEPTN_NAMESPACE") == "" {
		os.Setenv("KEPTN_NAMESPACE", defaultKeptnNamespace)
	}
}

func GetLatestEventOfType(keptnContext, projectName, stage, eventType string) (*models.KeptnContextExtendedCE, error) {
	resp, err := ApiGETRequest("/mongodb-datastore/event?project=" + projectName + "&keptnContext=" + keptnContext + "&stage=" + stage + "&type=" + eventType)
	if err != nil {
		return nil, err
	}
	events := &models.Events{}
	if err := resp.ToJSON(events); err != nil {
		return nil, err
	}
	if len(events.Events) > 0 {
		return events.Events[0], nil
	}
	return nil, nil
}

func IsEqual(t *testing.T, property string, expected, actual interface{}) bool {
	if expected != actual {
		t.Logf("%s: expected %v, got %v", property, expected, actual)
		return false
	}
	return true
}

