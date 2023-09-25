package main

import (
	"fmt"
	substack "github.com/mr-destructive/substack-go"
	"log"
)

func main() {
	env, err := substack.loadEnv(".env")
	email := env["EMAIL"]
	password := env["PASSWORD"]
	publicationURL := "https://meetgor.substack.com"

	api, err := substack.NewApi(email, password, publicationURL)
	if err != nil {
		log.Fatalf("Error creating API client: %v", err)
	}

	users, err := api.getPublicationUsers()
	if err != nil {
		log.Fatalf("Error getting publication users: %v", err)
	}
	fmt.Printf("Publication Users: %+v\n", users)

	pub, err := api.getPublication()
	if err != nil {
		log.Fatalf("Error getting publication: %v", err)
	}
	fmt.Printf("Publication: %+v\n", pub)
}
