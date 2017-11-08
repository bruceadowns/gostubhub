package lib

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"syscall"

	//"gopkg.in/jmcvetta/napping.v3.2.0"
	"github.com/jmcvetta/napping"
	"golang.org/x/crypto/ssh/terminal"
)

func promptPassword() (password string, err error) {
	fmt.Print("Enter Password: ")
	bPassword, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println()

	if err == nil {
		password = string(bPassword)
	}

	return
}

// RefreshToken ...
// https://developer.stubhub.com/store/site/pages/doc-tutorials.jag
func RefreshToken(config *Config, debug bool) (string, error) {
	if config.RefreshToken == "" {
		return "", fmt.Errorf("refresh token is empty")
	}

	basicAuthToken := fmt.Sprintf("%s:%s",
		config.CustomerKey, config.CustomerSecret)
	log.Print(basicAuthToken)
	basicAuthTokenEncoded := base64.StdEncoding.EncodeToString([]byte(basicAuthToken))
	log.Print(basicAuthTokenEncoded)

	url := fmt.Sprintf("%s/login", config.Server)
	log.Print(url)

	s := &napping.Session{
		Log: debug,
	}
	log.Print(s)

	loginRequestBody := fmt.Sprintf("grant_type=%s&refresh_token=%s&scope=%s",
		"refresh_token",
		config.RefreshToken,
		config.Scope)
	log.Print(loginRequestBody)

	headers := []struct {
		name, value string
	}{
		{"Content-Type", "application/x-www-form-urlencoded"},
		{"Authorization", fmt.Sprintf("Basic %s", basicAuthTokenEncoded)},
	}
	log.Print(headers)

	s.Header = &http.Header{}
	for _, v := range headers {
		log.Printf("%s: %s", v.name, v.value)
		s.Header.Set(v.name, v.value)
	}

	payload := bytes.NewBufferString(loginRequestBody)
	jLoginResponse := &LoginJSON{}
	jLoginError := &LoginErrorJSON{}
	r := napping.Request{
		Method:     "POST",
		Url:        url,
		Payload:    payload,
		Result:     &jLoginResponse,
		Error:      &jLoginError,
		RawPayload: true,
	}
	resp, err := s.Send(&r)
	if err != nil {
		return "", err
	}
	log.Print(resp.Header)
	log.Print(resp.HttpResponse().Header)

	if resp.Status() != http.StatusOK {
		log.Print(jLoginError)
		return "", fmt.Errorf("invalid status code: %d", resp.Status())
	}

	log.Printf("Response: %d type: %s",
		resp.Status(),
		resp.HttpResponse().Header.Get("Content-Type"))

	config.Freshen(jLoginResponse, "")
	return config.AccessToken, nil
}

// ObtainToken ...
// https://developer.stubhub.com/store/site/pages/guides.jag?type=developersguide
func ObtainToken(config *Config, userAPI, debug bool) (string, error) {
	if !userAPI {
		if config.ApplicationToken == "" {
			return "", fmt.Errorf("application token is empty")
		}

		return config.ApplicationToken, nil
	}

	if len(config.AccessToken) > 0 {
		return config.AccessToken, nil
	}

	if config.UserName == "" {
		return "", fmt.Errorf("user name is empty")
	}
	if config.Scope == "" {
		return "", fmt.Errorf("scope is empty")
	}
	if config.CustomerKey == "" {
		return "", fmt.Errorf("customer key is empty")
	}
	if config.CustomerSecret == "" {
		return "", fmt.Errorf("customer secret is empty")
	}

	if config.Password == "" {
		password, err := promptPassword()
		if err != nil {
			return "", err
		}

		config.Password = password
	}

	basicAuthToken := fmt.Sprintf("%s:%s",
		config.CustomerKey, config.CustomerSecret)
	log.Print(basicAuthToken)
	basicAuthTokenEncoded := base64.StdEncoding.EncodeToString([]byte(basicAuthToken))
	log.Print(basicAuthTokenEncoded)

	url := fmt.Sprintf("%s/login", config.Server)
	log.Print(url)

	s := &napping.Session{
		Log: debug,
	}
	log.Print(s)

	loginRequestBody := fmt.Sprintf("grant_type=%s&username=%s&password=%s&scope=%s",
		"password",
		config.UserName,
		config.Password,
		config.Scope)
	log.Print(loginRequestBody)

	headers := []struct {
		name, value string
	}{
		{"Content-Type", "application/x-www-form-urlencoded"},
		{"Authorization", fmt.Sprintf("Basic %s", basicAuthTokenEncoded)},
	}
	log.Print(headers)

	s.Header = &http.Header{}
	for _, v := range headers {
		log.Printf("%s: %s", v.name, v.value)
		s.Header.Set(v.name, v.value)
	}

	payload := bytes.NewBufferString(loginRequestBody)
	jLoginResponse := &LoginJSON{}
	jLoginError := &LoginErrorJSON{}
	r := napping.Request{
		Method:     "POST",
		Url:        url,
		Payload:    payload,
		Result:     &jLoginResponse,
		Error:      &jLoginError,
		RawPayload: true,
	}
	resp, err := s.Send(&r)
	if err != nil {
		return "", err
	}
	log.Print(resp.Header)
	log.Print(resp.HttpResponse().Header)

	if resp.Status() != http.StatusOK {
		log.Print(jLoginError)
		return "", fmt.Errorf("invalid status code: %d", resp.Status())
	}

	userID := resp.HttpResponse().Header.Get("X-Stubhub-User-Guid")
	log.Printf("Response: %d type: %s id: %s",
		resp.Status(),
		resp.HttpResponse().Header.Get("Content-Type"),
		userID)

	config.Freshen(jLoginResponse, userID)
	return config.AccessToken, nil
}
