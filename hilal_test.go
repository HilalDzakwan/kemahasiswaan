package hilal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertMahasiswa(t *testing.T) {
	db := MongoConnect("kemahasiswaan")
	defer db.Drop(context.Background())

	data := Mahasiswa{
		Name:         "John Doe",
		NPM:          "12345",
		Alamat:       "Jalan Raya",
		EmailAddress: "johndoe@example.com",
	}

	result := InsertMahasiswa("kemahasiswaan", data)

	assert.NotNil(t, result, "insertion should succeed")
}

func TestInsertJurusan(t *testing.T) {
	db := MongoConnect("kemahasiswaan")
	defer db.Drop(context.Background())

	data := Jurusan{
		Name:    "Jurusan A",
		NPM:     "12345",
		Jurusan: "Teknik Informatika",
	}

	result := InsertJurusan("kemahasiswaan", data)

	assert.NotNil(t, result, "insertion should succeed")
}

func TestInsertNilai(t *testing.T) {
	db := MongoConnect("kemahasiswaan")
	defer db.Drop(context.Background())

	data := Nilai{
		Name:       "John Doe",
		NPM:        "12345",
		Sejarah:    90,
		Matematika: 80,
		Inggris:    85,
		Pkn:        75,
	}

	result := InsertNilai("kemahasiswaan", data)

	assert.NotNil(t, result, "insertion should succeed")
}

func TestInsertAlamat(t *testing.T) {
	db := MongoConnect("kemahasiswaan")
	defer db.Drop(context.Background())

	data := Alamat{
		Name:         "John Doe",
		NPM:          "12345",
		EmailAddress: "johndoe@example.com",
	}

	result := InsertAlamat("kemahasiswaan", data)

	assert.NotNil(t, result, "insertion should succeed")
}

func TestInsertNPM(t *testing.T) {
	db := MongoConnect("kemahasiswaan")
	defer db.Drop(context.Background())

	data := NPM{
		Name:         "John Doe",
		EmailAddress: "johndoe@example.com",
		Alamat:       "Jalan Raya",
	}

	result := InsertNPM("kemahasiswaan", data)

	assert.NotNil(t, result, "insertion should succeed")
}
