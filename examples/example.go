package main

import (
	"fmt"
	substack "github.com/mr-destructive/substack-go"
	"log"
)

func main() {
	publicationURL := "https://meetgor.substack.com"
	client, err := substack.NewApi("email@example.com", "supersecret", publicationURL)
	if err != nil {
		log.Fatalf("Error creating API client: %v", err)
	}

	users, err := client.PublicationUsers()
	if err != nil {
		log.Fatalf("Error getting publication users: %v", err)
	}
	fmt.Printf("Publication Users: %+v\n", users)

	pub, err := client.Publication()
	if err != nil {
		log.Fatalf("Error getting publication: %v", err)
	}
	fmt.Printf("Publication: %+v\n", pub)
}
