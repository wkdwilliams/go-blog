package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func main() {
	a := Address{
		// Street: "123",
		City:   "Unknown",
		State:  "V",
		Zip:    "12345",
	}

	err := a.Validate()

	if err := json.NewEncoder(os.Stdout).Encode(err); err != nil {
		log.Fatal(err)
	}
}
