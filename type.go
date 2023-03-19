package hilal

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// struct untuk tabel Mahasiswa
type Mahasiswa struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name         string             `bson:"name,omitempty" json:"name,omitempty"`
    NPM          string             `bson:"npm,omitempty" json:"npm,omitempty"`
    Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
    EmailAddress string             `bson:"email_address,omitempty" json:"email_address,omitempty"`
}

// struct untuk tabel Jurusan
type Jurusan struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name     string             `bson:"name,omitempty" json:"name,omitempty"`
    NPM      string             `bson:"npm,omitempty" json:"npm,omitempty"`
    Jurusan  string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

// struct untuk tabel Nilai
type Nilai struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name      string             `bson:"name,omitempty" json:"name,omitempty"`
    NPM       string             `bson:"npm,omitempty" json:"npm,omitempty"`
    Sejarah   float32            `bson:"sejarah,omitempty" json:"sejarah,omitempty"`
    Matematika float32           `bson:"matematika,omitempty" json:"matematika,omitempty"`
    Inggris   float32            `bson:"inggris,omitempty" json:"inggris,omitempty"`
    Pkn       float32            `bson:"pkn,omitempty" json:"pkn,omitempty"`
}

// struct untuk tabel Alamat
type Alamat struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name         string             `bson:"name,omitempty" json:"name,omitempty"`
    NPM          string             `bson:"npm,omitempty" json:"npm,omitempty"`
    EmailAddress string             `bson:"email_address,omitempty" json:"email_address,omitempty"`
}

// struct untuk tabel NPM
type NPM struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name         string             `bson:"name,omitempty" json:"name,omitempty"`
    EmailAddress string             `bson:"email_address,omitempty" json:"email_address,omitempty"`
    Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}