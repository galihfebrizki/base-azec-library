package eureka

import (
	"base-azec-library/com/azec/base/baseModel"
	"crypto/tls"
	"github.com/ian-kent/go-log/log"
	"net/http"
	"strings"
)

// Accepts a Httpaction and a one-way channel to write the results to.
func DoHttpRequest(httpAction baseModel.HttpAction) bool {
	req := buildHttpRequest(httpAction)

	var DefaultTransport http.RoundTripper = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp, err := DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Println("HTTP request failed: %s", err)
		return false
	} else {
		log.Println("httpAction.Url = "+httpAction.Url)
		log.Println("method = "+httpAction.Method)
		log.Println(resp.StatusCode)
		return true
		defer resp.Body.Close()
	}
	return false
}

func buildHttpRequest(httpAction baseModel.HttpAction) *http.Request {
	var req *http.Request
	var err error
	if httpAction.Body != "" {
		reader := strings.NewReader(httpAction.Body)
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
	} else if httpAction.Template != "" {
		reader := strings.NewReader(httpAction.Template)
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
	} else {
		req, err = http.NewRequest(httpAction.Method, httpAction.Url, nil)
	}
	if err != nil {
		log.Error(err.Error())
	}

	// Add headers
	req.Header.Add("Accept", httpAction.Accept)
	if (httpAction.ContentType != "") {
		req.Header.Add("Content-Type", httpAction.ContentType)
	}
	return req
}