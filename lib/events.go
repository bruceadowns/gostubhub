package lib

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmcvetta/napping"
)

const (
	eventsURLPath = "/search/catalog/events/v3"
)

// Events ...
// https://developer.stubhub.com/store/site/pages/doc-viewer.jag?category=Search&api=EventSearchAPI&endpoint=searchforeventsv3&version=v3
func Events(filter string, token string, c *Config, debug bool) (int, error) {
	authHeader := fmt.Sprintf("Authorization: Bearer %s", token)
	log.Print(authHeader)

	s := &napping.Session{
		Header: &http.Header{"Authorization": []string{fmt.Sprintf("Bearer %s", token)}},
		Log:    debug,
	}

	url := c.Server + eventsURLPath
	log.Print(url)

	np := napping.Params{
		//"id":           "1234567890",
		//"status":       "active |contingent",
		"status": "active",
		//"name":         "-mirror",
		"start": "0",
		"limit": "10",
		//"rows":         "20",
		//"geoExpansion": "false",
		"sort": "eventDateLocal asc",
		//"radius":       "200",
		//"point":        "25.77427,-80.19366",
		"q": filter,
		//"venueId": "43580",
	}
	params := np.AsUrlValues()

	jResponse := &SearchJSON{}
	jError := &SearchErrorJSON{}
	resp, err := s.Get(url, &params, &jResponse, &jError)
	if err != nil {
		return 0, err
	}
	log.Print(resp.Header)
	log.Print(resp.HttpResponse().Header)

	if resp.Status() != 200 {
		log.Print(jError)
		return resp.Status(), fmt.Errorf("Error occurred retrieving events [%d]", resp.Status())
	}

	log.Print(fmt.Sprintf("Retrieved %d events [%d]", len(jResponse.Events), resp.Status()))
	log.Printf("%d", jResponse.NumFound)
	for _, v := range jResponse.Events {
		log.Printf("%s AT %s ON %s", v.Name, v.Venue.Name, v.EventDateLocal)
	}

	return 0, nil
}
