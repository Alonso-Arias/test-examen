package context

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alonso/test/log"
	"github.com/alonso/test/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "context")

var urlBase = os.Getenv("BASE_URL")

// se rea la request por detras hacia la url para consumir el api
func GetLocals() (model.GetLocales, error) {
	log := loggerf.WithField("func", "GetLocals")

	req, err := http.NewRequest("GET", urlBase, nil)
	if err != nil {
		return model.GetLocales{}, err
	}
	req.Header.Set("Accept", "application/json")

	client := getClientCofiguration()

	resp, err := client.Do(req)
	if err != nil {
		return model.GetLocales{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.GetLocales{}, err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf")) // evitar problemas de formato utf

	var localsResponse model.GetLocales
	// se parsea el json a la structura declarada
	err = json.Unmarshal(bodyBytes, &localsResponse)
	if err != nil {
		return model.GetLocales{}, err
	}

	log.Debugf("Body: %s", string(bodyBytes))

	return localsResponse, err

}

// se configuran los deadines de peticiones hacia la api
func getClientCofiguration() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Duration(5) * time.Second,
				KeepAlive: time.Duration(5),
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return client
}
