package db

import "cloud.google.com/go/civil"

//License : data strucutre of License table in google cloud storage
type License struct {
	Name               string      `json:"name"`               // Name of license definition
	Type               string      `json:"type"`               // The type of the license.
	Exclusive          Price       `json:"exclusive"`          // Defines whether licensor allows selling the content for exclusive access by the publisher; if this field is present, exclusive field value is monetary value that the licensor is asking for exclusivity.
	Payment            Payment     `json:"payment"`            // Defines the terms of paymen for the article views. There are different possibilities for different licence types.
	Owner              LegalEntity `json:"owner"`              // Defines mitem owner
	Distribution       string      `json:"distribution"`       // Defines rules for mitem distribution
	AdsPolicy          AdsPolicy   `json:"adsPolicy"`          // Defines set of rules for ads insertion into article
	AllowModifications bool        `json:"allowModifications"` // Indicated if mitem can be modified by distributor, changes to the license must always be approved by the owner
	CustomConditions   string      `json:"customConditions"`   // Any custom conditions owner and distributor agree to
}

// Price defines price structure containing currency and value in "cents"
type Price struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

// Payment defines the terms of paymen for the article views. There are different possibilities for different licence types.
type Payment struct {
	Model string `json:"model"` // Defines the payment model for the article.
	Price Price  `json:"price"` // Field meaning depends on the paymend model chosen.
}

// LegalEntity defines mosaiq legal enity object with type and ID
type LegalEntity struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// AdsPolicy defines set of rules for including ads in the article
type AdsPolicy struct {
	Allowed bool `json:"allowed"` // if ads are allowed in the mitem
	// TODO: define restrictions, either as nested structure or relation to other entity
	//Restrictions interface{} `json:"restrictions"`
}

type TestEntity struct {
	TestDateTimeField civil.DateTime
	TestFloatField    float32
	TestIntField      int
	TestStringField   string
}
