package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Cookie ...
type Cookie struct {
	AccessToken  string `yaml:"access_token"`
	ExpiresIn    int    `yaml:"expires_in"`
	RefreshToken string `yaml:"refresh_token"`
	TokenType    string `yaml:"token_type"`
	UserID       string `yaml:"X-Stubhub-User-Guid"`
}

// Cookies ...
type Cookies map[string]*Cookie

func (c Cookies) String() string {
	buf := &bytes.Buffer{}

	for k, v := range c {
		if buf.Len() > 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(fmt.Sprintf("%s: %s %d %s %s %s",
			k,
			v.AccessToken, v.ExpiresIn, v.RefreshToken, v.TokenType, v.UserID))
	}

	return buf.String()
}

// DeserCookies ...
func DeserCookies(f string) (Cookies, error) {
	var res Cookies

	if fi, err := os.Stat(f); err == nil {
		log.Printf("Found cookies file: %s [%d]", fi.Name(), fi.Size())

		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}

		in, err := ioutil.ReadAll(bufio.NewReader(file))
		if err != nil {
			return nil, err
		}

		res = Cookies{}
		if err := yaml.Unmarshal(in, &res); err != nil {
			return nil, err
		}
	} else {
		res = Cookies{}
		log.Printf("cookies file %s not found", f)
	}

	return res, nil
}

// SerCookies ...
func SerCookies(f string, context string, cookies Cookies, config *Config) error {
	if fi, err := os.Stat(f); err == nil {
		log.Printf("Found cookies file: %s [%d]", fi.Name(), fi.Size())
	} else {
		log.Printf("cookies file %s not found", f)
	}
	log.Print(config)

	if cookie, ok := cookies[context]; ok {
		log.Print("Found cookie")
		cookie.AccessToken = config.AccessToken
		cookie.ExpiresIn = config.ExpiresIn
		cookie.RefreshToken = config.RefreshToken
		cookie.TokenType = config.TokenType
		cookie.UserID = config.UserID
	} else {
		log.Print("Cookie not found")
		cookies[context] = &Cookie{
			AccessToken:  config.AccessToken,
			ExpiresIn:    config.ExpiresIn,
			RefreshToken: config.RefreshToken,
			TokenType:    config.TokenType,
			UserID:       config.UserID,
		}
	}
	log.Print(context)
	log.Print(cookies)

	bCookies, err := yaml.Marshal(cookies)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(f, bCookies, os.FileMode(0600)); err != nil {
		return err
	}

	if fi, err := os.Stat(f); err == nil {
		log.Printf("New cookies file: %s [%d]", fi.Name(), fi.Size())
	} else {
		log.Printf("cookies file %s not found", f)
	}

	return nil
}
