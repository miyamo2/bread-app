package contentful

import (
	"time"
)

type Metadata struct {
	Tags []interface{} `json:"tags"`
}

type Sys struct {
	Space struct {
		Sys struct {
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
			ID       string `json:"id"`
		} `json:"sys"`
	} `json:"space"`
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Environment struct {
		Sys struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
		} `json:"sys"`
	} `json:"environment"`
	Revision    int `json:"revision"`
	ContentType struct {
		Sys struct {
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
			ID       string `json:"id"`
		} `json:"sys"`
	} `json:"contentType"`
	Locale string `json:"locale"`
}

type Fields struct {
	Seller struct {
		Sys struct {
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
			ID       string `json:"id"`
		} `json:"sys"`
	} `json:"seller"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageList   []struct {
		Sys struct {
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
			ID       string `json:"id"`
		} `json:"sys"`
	} `json:"imageList"`
	Calorie        int    `json:"calorie"`
	Lipid          int    `json:"lipid"`
	Carbohydrate   int    `json:"carbohydrate"`
	Protein        int    `json:"protein"`
	SaltEquivalent int    `json:"saltEquivalent"`
	FoodStuffs     string `json:"foodStuffs"`
	Caution        string `json:"caution"`
	HowToEat       struct {
		Sys struct {
			Type     string `json:"type"`
			LinkType string `json:"linkType"`
			ID       string `json:"id"`
		} `json:"sys"`
	} `json:"howToEat"`
}

type ContentfulEntry struct {
	Metadata Metadata `json:"metadata"`
	Sys      Sys      `json:"sys"`
	Fields   Fields   `json:"fields"`
}
