package bahaya

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertOneDoc(t *testing.T) {
	// buat koneksi ke MongoDB
	dbname := "testdb"
	db := MongoConnect(dbname)

	// persiapkan data untuk diinsert
	doc := map[string]interface{}{
		"_id":           primitive.NewObjectID(),
		"name":          "John Doe",
		"email_address": "johndoe@example.com",
	}

	// insert data ke koleksi
	collection := "testcollection"
	insertedID := InsertOneDoc(dbname, collection, doc)

	// verifikasi hasil insert
	result := db.Collection(collection).FindOne(context.Background(), map[string]interface{}{"_id": insertedID})
	if result.Err() != nil {
		t.Errorf("failed to find inserted document: %v", result.Err())
	}
	var insertedDoc map[string]interface{}
	err := result.Decode(&insertedDoc)
	if err != nil {
		t.Errorf("failed to decode inserted document: %v", err)
	}
	for k, v := range doc {
		if insertedDoc[k] != v {
			t.Errorf("expected %v, got %v", v, insertedDoc[k])
		}
	}

	// hapus data yang sudah diinsert
	_, err = db.Collection(collection).DeleteOne(context.Background(), map[string]interface{}{"_id": insertedID})
	if err != nil {
		t.Errorf("failed to delete inserted document: %v", err)
	}

	// tutup koneksi ke MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		t.Errorf("failed to disconnect from MongoDB: %v", err)
	}
}
