package lib

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmcvetta/napping"
)

const (
	listingsURLPath = "/accountmanagement/listings/v1/seller"
)

// Listings ...
// https://developer.stubhub.com/store/site/pages/doc-viewer.jag?category=AccountManagement&api=AccountManagementAPI&endpoint=getlisting&version=v1
func Listings(filter string, token string, c *Config, debug bool) (int, error) {
	if token == "" {
		return 0, fmt.Errorf("token is empty")
	}
	if c.UserID == "" {
		return 0, fmt.Errorf("user id is empty")
	}

	authHeader := fmt.Sprintf("Authorization: Bearer %s", token)
	log.Print(authHeader)

	s := &napping.Session{
		Header: &http.Header{"Authorization": []string{fmt.Sprintf("Bearer %s", token)}},
		Log:    debug,
	}

	url := fmt.Sprintf("%s%s/%s", c.Server, listingsURLPath, c.UserID)
	log.Print(url)

	np := napping.Params{
		"filters": filter, // ""
		//"sort":    "SALEENDDATE Desc", // EVENTDATE Asc
		//"start":   "0",                // 0
		"rows": "5", // 200
	}
	params := np.AsUrlValues()

	jResponse := &ListingsJSON{}
	jError := &ListingsErrorJSON{}
	resp, err := s.Get(url, &params, &jResponse, &jError)
	if err != nil {
		return 0, err
	}
	log.Print(resp.Header)
	log.Print(resp.HttpResponse().Header)

	if resp.Status() != 200 {
		log.Print(jError)
		return resp.Status(), fmt.Errorf("Error occurred retrieving listings [%d]", resp.Status())
	}

	log.Print(fmt.Sprintf("Retrieved [%d] user listings [%d]", len(jResponse.Listings), resp.Status()))
	for _, v := range jResponse.Listings {
		log.Printf("%s: $%d %s", v.ID, v.PricePerTicket.Amount, v.VenueDescription)
	}

	return 0, nil
}
