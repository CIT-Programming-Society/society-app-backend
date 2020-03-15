package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Don't modify this path
// Add the Viewer and Editor credentials from here: https://console.cloud.google.com/apis/credentials/serviceaccountkey?_ga=2.176794983.336131438.1584302922-1433294965.1584302922
const credentials string = "./credentials/Viewer/Society-App-d55080e09447.json"

func main() {

	ctx := context.Background()
	sa := option.WithCredentialsFile(credentials)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("societies").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	defer client.Close()
}
