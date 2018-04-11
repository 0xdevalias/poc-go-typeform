package main

import (
	"fmt"
	"log"

	"github.com/0xdevalias/poc-typeform/api"
)

func main() {
	c := api.DefaultClient("TODO")

	r, err := c.RetrieveForm("TODO")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("%#v", r)

	// Display survey structure
	for _, field := range r.Fields {
		fmt.Printf("(%s) '%s'\n", field.Type, field.Title)
		for _, choice := range field.Properties.Choices {
			fmt.Printf("  - %s\n", choice.Label)
		}

		for _, groupField := range field.Properties.Fields {
			fmt.Printf("  (%s) '%s'\n", groupField.Type, groupField.Title)

			for _, groupChoice := range groupField.Properties.Choices {
				fmt.Printf("    - %s\n", groupChoice.Label)
			}
		}
	}
}
