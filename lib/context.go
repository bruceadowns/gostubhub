package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Context ...
type Context struct {
	ApplicationToken string `yaml:"application_token"`
	CustomerKey      string `yaml:"customer_key"`
	CustomerSecret   string `yaml:"customer_secret"`
	Password         string `yaml:"password"`
	Scope            string `yaml:"scope"`
	Server           string `yaml:"server"`
	UserName         string `yaml:"user_name"`
}

// Contexts ...
type Contexts map[string]*Context

func (c Contexts) String() string {
	buf := &bytes.Buffer{}

	for k, v := range c {
		if buf.Len() > 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(fmt.Sprintf("%s: %s %s %s %s %s %s %s",
			k,
			v.ApplicationToken, v.CustomerKey, v.CustomerSecret,
			v.Password, v.Scope, v.Server, v.UserName))
	}

	return buf.String()
}

// DeserContexts ...
func DeserContexts(f string) (Contexts, error) {
	res := Contexts{}

	// read config file name default to config.yaml
	// a dash indicates to read from stdin
	var r io.Reader = os.Stdin
	if !strings.EqualFold(f, "-") {
		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}

		r = bufio.NewReader(file)
	}

	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading config file")
	}

	if err := yaml.Unmarshal(in, &res); err != nil {
		return nil, fmt.Errorf("error decoding context yaml: %s", err)
	}

	return res, nil
}
