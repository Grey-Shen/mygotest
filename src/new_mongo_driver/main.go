package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	AppName                     string            `bson:"app_name"`
	FullName                    string            `bson:"full_name"`
	ContactInfo                 string            `bson:"contact_information"`
	AppId                       string            `bson:"app_id"`
	EncryptedAppKey             []byte            `bson:"encrypted_app_key"`
	BucketName                  string            `bson:"bucket_name"`
	BucketDomain                string            `bson:"bucket_domain"`
	AccessKey                   string            `bson:"access_key"`
	EncryptedSecretKey          []byte            `bson:"encrypted_secret_key"`
	NamespaceId                 string            `bson:"namespace_id"`
	NamespaceName               string            `bson:"namespace_name"`
	KeyId                       string            `bson:"key_id"`
	BucketId                    string            `bson:"bucket_id"`
	QiniuUid                    int64             `bson:"qiniu_uid"`
	DisabledAt                  *time.Time        `bson:"disabled_at,omitempty"`
	Pipelines                   []string          `bson:"pipelines"`
	IgnoreBizdataKeyWhitelistAt *time.Time        `bson:"ignore_bizdata_key_whitelist_at"`
	SDKVersions                 map[string]string `bson:"sdk_versions"`
	EncryptedEncodedIV          string            `bson:"encrypted_encoded_iv,omitempty"`
}

func main() {
	var (
		err        error
		client     *mongo.Client
		connection *mongo.Collection
		cur        *mongo.Cursor
		app        App
	)

	if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatalln(err)
	}

	connection = client.Database("tassadar_test").Collection("apps")

	cur, err = connection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		err := cur.Decode(&app)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("======app: %#v\n", &app)
	}

	if err := cur.Err(); err != nil {
		return
	}

	return
}
