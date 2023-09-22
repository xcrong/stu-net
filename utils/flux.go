package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"stu-net/tools"
)

func Flux(sessionID string) (result *tools.FluxResult, err error) {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar: jar,
	}

	// Set-Cookie:[AUTHSESSID=ed726c757d9e; HttpOnly;Secure;]
	cookie := &http.Cookie{
		Name:  "AUTHSESSID",
		Value: sessionID,
	}

	req, err := http.NewRequest("POST", FluxUrl, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Can't not to Close the ")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fluxInfo, err := tools.ParseFluxInfo(string(body))
	if err != nil {
		return nil, err
	}
	return fluxInfo, nil
}
