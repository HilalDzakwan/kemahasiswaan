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
// insert data ke koleksi Mahasiswa
func insertMahasiswa(db *mongo.Database, mhs Mahasiswa) (insertedID interface{}, err error) {
	result, err := db.Collection("Mahasiswa").InsertOne(context.Background(), mhs)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Mahasiswa: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Jurusan
func insertJurusan(db *mongo.Database, jur Jurusan) (insertedID interface{}, err error) {
	result, err := db.Collection("Jurusan").InsertOne(context.Background(), jur)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Jurusan: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Nilai
func insertNilai(db *mongo.Database, n Nilai) (insertedID interface{}, err error) {
	result, err := db.Collection("Nilai").InsertOne(context.Background(), n)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Nilai: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi Alamat
func insertAlamat(db *mongo.Database, alamat Alamat) (insertedID interface{}, err error) {
	result, err := db.Collection("Alamat").InsertOne(context.Background(), alamat)
	if err != nil {
		return nil, fmt.Errorf("failed to insert Alamat: %v", err)
	}
	return result.InsertedID, nil
}

// insert data ke koleksi NPM
func insertNPM(db *mongo.Database, npm NPM) (insertedID interface{}, err error) {
	result, err := db.Collection("NPM").InsertOne(context.Background(), npm)
	if err != nil {
		return nil, fmt.Errorf("failed to insert NPM: %v", err)
	}
	return result.InsertedID, nil
}