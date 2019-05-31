package models

import (
	"mime/multipart"

	gmodels "opensoach.com/models"
)

type DocumentStoreInsertRowModel struct {
	Uuid string `db:"uuid"`
	CPMIDEntityModel
	Name        string `db:"name"`
	DocType     string `db:"doctype"`
	URI         string `db:"location"`
	StorageType int    `db:"location_type"`
	Persisted   int    `db:"persisted" json:"persisted"`
	Updated_by  int64  `db:"updated_by" json:"updatedby"`
}

type DocumentStoreRowModel struct {
	Uuid         string `db:"uuid"`
	DocId        int64  `db:"id" dbattr:"pri,auto"`
	Name         string `db:"name"`
	DocType      string `db:"doctype"`
	Location     string `db:"location"`
	LocationType int    `db:"location_type"`
	Persisted    int    `db:"persisted" json:"persisted"`
}

type DocumentStoreUpdateRowModel struct {
	Uuid string `db:"uuid" json:"uuid"`
	CPMIDEntityModel
	Name         string `db:"name" json:"name"`
	DocType      string `db:"doctype" json:"doctype"`
	Location     string `db:"location" json:"location"`
	LocationType int    `db:"location_type"`
	Persisted    int    `db:"persisted" json:"persisted"`
}

type DocumentStoreDataModel struct {
	DocumentID  string
	StorageType int
	NestedPath  []string
	FileData    *multipart.Form
	DBContext   gmodels.IDBConnection
}

type IDocumentStore interface {
	Get() (error, DocumentData)
	Save() error
	SaveDBRecord() error
	SaveDocument() error
	Update() error
}

type DocumentData struct {
	ByteData    []byte
	ContentType string
}
