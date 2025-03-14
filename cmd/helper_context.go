package cmd

import (
	"context"
	"log"

	tfe "github.com/hashicorp/go-tfe"
)

func getClientContext() (*tfe.Client, context.Context) {

	config := &tfe.Config{
		Address: "https://" + *viperString("tfeHostname"),
		Token:   *viperString("tfeToken"),
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create a context
	ctx := context.Background()

	return client, ctx
}
