# gostubhub

## Overview

An exploratory client cli written in golang that exercises the stubhub api. See details at https://github.com/bruceadowns/talks/tree/master/2017/stubhub-client.slide.

It does the following:

* login to stubhub using ssl/oauth
* issue events or listings query
* display results
* use via cli or http server

## Dev Notes

```
# view help
go run main.go --help
# or
gostubhub --help

# search for Wynton
go run main.go -context gostubhub -debug events 'Wynton Marsalis'

# search for active listings
go run main.go -context gostubhub -debug listings 'STATUS:ACTIVE'
```

## Config

#### context.yaml

Defines potential contexts:

* scope
* server
* application_token - optional
* customer_key - optional
* customer_secret - optional
* user_name - optional
* password - optional

#### .cookies.yaml

Generated file storing transient login info:

* AccessToken
* ExpiresIn
* RefreshToken
* TokenType
* UserID

#### environment

Overridable context variables:

* STUBHUB_APPLICATION_TOKEN
* STUBHUB_CUSTOMER_KEY
* STUBHUB_CUSTOMER_SECRET
* STUBHUB_PASSWORD
* STUBHUB_USER_NAME

#### Password Note

The stubhub password may be specified in contexts.yaml or the STUBHUB_PASSWORD environment variable. If a password is not provided and a user api is utilitized, the cli will prompt for the user password. Non-user apis use the application token, therefore a password is not needed.

## stubhub apis

* Account Management - AccountManagementAPI, AccountManagementSalesAPI
* Catalog - EventsAPI, VenuesAPI
* Listings - ListingAPI
* Localization - I18nLocalizationAPI
* Search - InventorySearchAPI, EventSearchAPI, InventorySearchAPIv2
* User Management - UserCustomerAPI

## References

* https://golang.org
* https://developer.stubhub.com
* https://myaccount.stubhubsandbox.com

## Dependencies

* rest - https://github.com/jmcvetta/napping
* yaml - https://github.com/go-yaml/yaml
