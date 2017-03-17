package lib

import (
	"fmt"
	"log"
)

// Config ...
type Config struct {
	AccessToken      string
	ApplicationToken string
	CustomerKey      string
	CustomerSecret   string
	ExpiresIn        int
	Password         string
	RefreshToken     string
	Scope            string
	Server           string
	TokenType        string
	UserName         string
	UserID           string
}

func (c *Config) String() string {
	return fmt.Sprintf("%s %s %s %s %d %s %s %s %s %s %s %s",
		c.AccessToken, c.ApplicationToken, c.CustomerKey, c.CustomerSecret,
		c.ExpiresIn, c.Password, c.RefreshToken, c.Scope,
		c.Server, c.TokenType, c.UserName, c.UserID)
}

// Coalesce ...
func Coalesce(context *Context, cookie *Cookie, env *Env) *Config {
	res := &Config{}

	// add context
	if context == nil {
		log.Print("Context not found")
	} else {
		res.ApplicationToken = context.ApplicationToken
		res.CustomerKey = context.CustomerKey
		res.CustomerSecret = context.CustomerSecret
		res.Password = context.Password
		res.Scope = context.Scope
		res.Server = context.Server
		res.UserName = context.UserName
	}

	// add cookies
	if cookie == nil {
		log.Print("Cookie not found")
	} else {
		res.AccessToken = cookie.AccessToken
		res.ExpiresIn = cookie.ExpiresIn
		res.RefreshToken = cookie.RefreshToken
		res.TokenType = cookie.TokenType
		res.UserID = cookie.UserID
	}

	// override env
	if len(env.ApplicationToken) > 0 {
		res.ApplicationToken = env.ApplicationToken
	}
	if len(env.CustomerKey) > 0 {
		res.CustomerKey = env.CustomerKey
	}
	if len(env.CustomerSecret) > 0 {
		res.CustomerSecret = env.CustomerSecret
	}
	if len(env.Password) > 0 {
		res.Password = env.Password
	}
	if len(env.UserName) > 0 {
		res.UserName = env.UserName
	}

	return res
}

// Freshen ...
func (c *Config) Freshen(j *LoginJSON, user string) {
	c.AccessToken = j.AccessToken
	c.ExpiresIn = j.ExpiresIn
	c.RefreshToken = j.RefreshToken
	c.TokenType = j.TokenType

	if len(user) > 0 {
		c.UserID = user
	}
}
