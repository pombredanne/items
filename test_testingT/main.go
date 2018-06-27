package main

import (
	"github.com/olivere/elastic"
	"fmt"
	"context"
)


type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main() {

	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	// Create an index
	_, err = client.CreateIndex("ft").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}

	// Add a document to the index
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("ft").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Refresh("true").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}

	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")

	searchResult, err := client.Search().
		Index("ft").   // search in index "ft"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)

	}

	fmt.Println(searchResult)

}
