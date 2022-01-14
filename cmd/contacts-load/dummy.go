package main

import (
	"context"
	_ "embed"
	"encoding/json"

	contactapi "github.com/bingoohuang/contacts-api"
	"github.com/sirupsen/logrus"
)

//go:embed contacts.json
var contactsJson []byte

func main() {
	var contacts []contactapi.Contact

	//unmarshall data
	if err := json.Unmarshal(contactsJson, &contacts); err != nil {
		logrus.Error("unmarshall an error occurred", err)
	}

	logrus.Info("Data\n", len(contacts))

	//import mongo client
	client := contactapi.ConnectMongoDb("mongodb://localhost:27017")
	logrus.Info(client)

	defer client.Disconnect(context.TODO())

	collection := client.Database("ContactDb").Collection("contacts")

	logrus.Warn("Total data count:", len(contacts))

	for _, item := range contacts {
		collection.InsertOne(context.TODO(), item)
	}

	logrus.Info("Data import finished...")
}
