package hilal

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertMahasiswa(db string, data Mahasiswa) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("mahasiswa").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID
}

func InsertJurusan(db string, data Jurusan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("jurusan").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID
}

func InsertNilai(db string, data Nilai) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("nilai").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID
}

func InsertAlamat(db string, data Alamat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("alamat").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID
}

func InsertNPM(db string, data NPM) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("npm").InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID
}