package main

import (
	"flag"
	"log"

	"github.com/bruceadowns/gostubhub/lib"
)

const (
	cliTypeEvents = iota
	cliTypeListings
)

func main() {
	//
	// initialize cli arguments, context, cookies, env
	// and merge into a single config
	//

	var contextsFileName string
	var cookiesFileName string
	var context string
	var debug bool
	flag.StringVar(&contextsFileName, "contexts", "contexts.yaml", "configuration yaml file name")
	flag.StringVar(&cookiesFileName, "cookies", ".cookies.yaml", "cookies yaml file name")
	flag.StringVar(&context, "context", "gostubhub", "which stubhub config context to use")
	flag.BoolVar(&debug, "debug", false, "verbose debugging (default false)")
	flag.Parse()
	log.Print(contextsFileName)
	log.Print(cookiesFileName)
	log.Print(context)
	log.Print(debug)

	if len(flag.Args()) != 2 {
		log.Fatal("Missing action and/or filter. Expect 'events' or 'listings' plus 'filter'")
	}
	action := flag.Arg(0)
	if action == "" {
		log.Fatal("Action is empty. Expect 'events' or 'listings'")
	}
	log.Print(action)
	apiType := cliTypeListings
	switch action {
	case "events":
		apiType = cliTypeEvents
	case "listings":
	default:
		log.Fatal("Invalid action. Expect 'events' or 'listings'")
	}

	filter := flag.Arg(1)
	if filter == "" {
		log.Fatal("Action filter is empty")
	}
	log.Print(filter)

	// initialize contexts from config file
	contexts, err := lib.DeserContexts(contextsFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(contexts)

	// initialize cookies if exists
	cookies, err := lib.DeserCookies(cookiesFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cookies)

	// initialize environment if exists
	env, err := lib.InitEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(env)

	// coalesce context + cookies + environment
	// where the environment may override
	config := lib.Coalesce(contexts[context], cookies[context], env)
	log.Print(config)

	var isUserAPI bool
	switch apiType {
	case cliTypeEvents:
		isUserAPI = false
	case cliTypeListings:
		isUserAPI = true
	default:
		log.Fatal("Invalid api type")
	}

	// obtain application or user token
	token, err := lib.ObtainToken(config, isUserAPI, debug)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(token)

	// retry once on http error 401
	retry := 0
	for retry < 2 {
		switch apiType {
		case cliTypeListings:
			if errCode, err := lib.Listings(filter, token, config, debug); err == nil {
				// success; break
				retry += 2
			} else {
				if errCode == 401 && retry == 0 {
					if token, err = lib.RefreshToken(config, debug); err != nil {
						log.Fatal(err)
					}

					retry++
				} else {
					log.Fatal(err)
				}
			}

		case cliTypeEvents:
			if errCode, err := lib.Events(filter, token, config, debug); err == nil {
				// success; break
				retry += 2
			} else {
				if errCode == 401 && retry == 0 {
					if token, err = lib.RefreshToken(config, debug); err != nil {
						log.Fatal(err)
					}

					retry++
				} else {
					log.Fatal(err)
				}
			}

		default:
			log.Fatal("Invalid api type")
		}
	}

	// save new cookie file
	if err := lib.SerCookies(cookiesFileName, context, cookies, config); err != nil {
		log.Fatal(err)
	}

	log.Print("Done")
}
