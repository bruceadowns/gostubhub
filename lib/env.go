package lib

import (
	"fmt"
	"os"
)

// Env ...
type Env struct {
	ApplicationToken string
	CustomerKey      string
	CustomerSecret   string
	Password         string
	UserName         string
}

func (e *Env) String() string {
	return fmt.Sprintf("%s %s %s %s %s",
		e.ApplicationToken, e.CustomerKey, e.CustomerSecret, e.Password, e.UserName)
}

// InitEnv ...
func InitEnv() (*Env, error) {
	res := &Env{}

	res.ApplicationToken = GetEnvStr("STUBHUB_APPLICATION_TOKEN", "")
	res.CustomerKey = GetEnvStr("STUBHUB_CUSTOMER_KEY", "")
	res.CustomerSecret = GetEnvStr("STUBHUB_CUSTOMER_SECRET", "")
	res.Password = GetEnvStr("STUBHUB_PASSWORD", "")
	res.UserName = GetEnvStr("STUBHUB_USER_NAME", "")

	return res, nil
}

// GetEnvStr returns string for environment string with default
func GetEnvStr(name, def string) (res string) {
	res = def

	s := os.Getenv(name)
	if len(s) > 0 {
		res = s
	}

	return
}
