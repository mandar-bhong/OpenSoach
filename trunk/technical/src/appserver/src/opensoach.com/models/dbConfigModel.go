package models

type DBMasterConfigRowModel struct {
	Key      string `db:"key" json:"key"`
	Category string `db:"category" json:"category"`
	Value    string `db:"value" json:"value"`
}
