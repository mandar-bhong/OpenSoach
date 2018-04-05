package models

type DBMasterConfigRowModel struct {
	Key      string `db:"param_key" json:"key"`
	Category string `db:"category" json:"category"`
	Value    string `db:"value" json:"value"`
}
