package lib

// LoginJSON ...
type LoginJSON struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// LoginErrorJSON ...
type LoginErrorJSON struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// PointJSON ...
type PointJSON struct {
	Point string `json:"point"`
}

// VenueJSON ...
type VenueJSON struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	URL           string  `json:"url"`
	WebURI        string  `json:"webURI"`
	SeoURI        string  `json:"seoURI"`
	VenueURL      string  `json:"venueUrl"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	Timezone      string  `json:"timezone"`
	JdkTimezone   string  `json:"jdkTimezone"`
	Address1      string  `json:"address1"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	PostalCode    string  `json:"postalCode"`
	Country       string  `json:"country"`
	VenueConfigID int     `json:"venueConfigId"`
}

// VenueConfigurationJSON ...
type VenueConfigurationJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AncestorsJSON ...
type AncestorsJSON struct {
	Categories []PropertyExJSON `json:"categories"`
	Groupings  []PropertyExJSON `json:"groupings"`
	Performers []PropertyExJSON `json:"performers"`
	Geos       []PropertyExJSON `json:"geos"`
}

// PropertyExJSON ...
type PropertyExJSON struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	WebURI string `json:"webURI"`
	SeoURI string `json:"seoURI"`
}

// ImageJSON ...
type ImageJSON struct {
	URL         string `json:"url"`
	IsResizable bool   `json:"isResizable"`
	URLSSL      string `json:"urlSsl"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Source      string `json:"source"`
	Credit      string `json:"credit"`
}

// AttributeJSON ...
type AttributeJSON struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// DisplayAttributesJSON ...
type DisplayAttributesJSON struct {
	IsHidden      bool   `json:"isHidden"`
	HideEventDate bool   `json:"hideEventDate"`
	HideEventTime bool   `json:"hideEventTime"`
	PrimaryName   string `json:"primaryName"`
}

// MobileAttributesJSON ...
type MobileAttributesJSON struct {
	EnableApplePassbook     bool `json:"enableApplePassbook"`
	MobileListingNotAllowed bool `json:"mobileListingNotAllowed"`
	StubhubMobileTicket     bool `json:"stubhubMobileTicket"`
}

// EventsJSON ...
type EventsJSON struct {
	ID                   int                    `json:"id"`
	Status               string                 `json:"status"`
	Locale               string                 `json:"locale"`
	Name                 string                 `json:"name"`
	OriginalName         string                 `json:"originalName"`
	Description          string                 `json:"description"`
	EventURL             string                 `json:"eventUrl"`
	WebURI               string                 `json:"webURI"`
	SeoURI               string                 `json:"seoURI"`
	EventDateLocal       StubhubTime            `json:"eventDateLocal"`
	DateOnsale           StubhubTime            `json:"dateOnsale"`
	EventDateUTC         StubhubTime            `json:"eventDateUTC"`
	Venue                VenueJSON              `json:"venue"`
	VenueConfiguration   VenueConfigurationJSON `json:"venueConfiguration"`
	BobID                int                    `json:"bobId"`
	Ancestors            AncestorsJSON          `json:"ancestors"`
	SourceID             string                 `json:"sourceId"`
	Categories           []PropertyExJSON       `json:"categories"`
	Groupings            []PropertyExJSON       `json:"groupings"`
	Performers           []PropertyExJSON       `json:"performers"`
	Geos                 []PropertyExJSON       `json:"geos"`
	CategoriesCollection []PropertyExJSON       `json:"categoriesCollection"`
	GroupingsCollection  []PropertyExJSON       `json:"groupingsCollection"`
	PerformersCollection []PropertyExJSON       `json:"performersCollection"`
	ImageURL             string                 `json:"imageUrl"`
	Images               []ImageJSON            `json:"images"`
	Attributes           []AttributeJSON        `json:"attributes"`
	DisplayAttributes    DisplayAttributesJSON  `json:"displayAttributes"`
	MobileAttributes     MobileAttributesJSON   `json:"mobileAttributes"`
	Score                float32                `json:"score"`
	CreatedDate          StubhubTime            `json:"createdDate"`
	DefaultLocale        string                 `json:"defaultLocale"`
}

// SearchJSON ...
type SearchJSON struct {
	NumFound int          `json:"numFound"`
	Events   []EventsJSON `json:"events"`
}

// SearchErrorJSON ...
type SearchErrorJSON struct {
	Code             string      `json:"code"`
	Description      string      `json:"description"`
	RequestID        string      `json:"requestId"`
	Data             []PointJSON `json:"data"`
	ValidationErrors string      `json:"validationErrors"`
}

// PriceJSON ...
type PriceJSON struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

// ListingJSON ...
type ListingJSON struct {
	ID                    string      `json:"id"`
	Status                string      `json:"status"`
	EventID               string      `json:"eventId"`
	EventDescription      string      `json:"eventDescription"`
	EventDate             StubhubTime `json:"eventDate"`
	InHandDate            StubhubTime `json:"inhandDate"`
	Quantity              int         `json:"quantity"`
	QuantityRemain        int         `json:"quantityRemain"`
	Section               string      `json:"section"`
	Rows                  string      `json:"rows"`
	Seats                 string      `json:"seats"`
	VenueDescription      string      `json:"venueDescription"`
	SplitOption           string      `json:"splitOption"`
	SplitQuantity         int         `json:"splitQuantity"`
	DeliveryOption        string      `json:"deliveryOption"`
	PreDelivered          string      `json:"preDelivered"`
	SaleEndDate           StubhubTime `json:"saleEndDate"`
	SaleMethod            string      `json:"saleMethod"`
	PricePerTicket        PriceJSON   `json:"pricePerTicket"`
	PayoutPerTicket       PriceJSON   `json:"payoutPerTicket"`
	DisplayPricePerTicket PriceJSON   `json:"displayPricePerTicket"`
	PaymentType           string      `json:"paymentType"`
}

// ListingsJSON ...
type ListingsJSON struct {
	NumFound int           `json:"numFound"`
	Listings []ListingJSON `json:"listing"`
}

// ListingsResponseJSON ...
type ListingsResponseJSON struct {
	Listings ListingsJSON `json:"listings"`
}

// ListingErrorJSON ...
type ListingErrorJSON struct {
	Type      string `json:"type"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Parameter string `json:"parameter"`
}

// ListingsErrorsJSON ...
type ListingsErrorsJSON struct {
	Errors   []ListingErrorJSON `json:"errors"`
	NumFound int                `json:"numFound"`
}

// ListingsErrorJSON ...
type ListingsErrorJSON struct {
	Listings ListingsErrorsJSON `json:"listings"`
}
