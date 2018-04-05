package models

type DBMasterConfigRowModel struct {
	Config_key   string `db:"config_key" json:"config_key"`
	Config_value string `db:"config_value" json:"config_value"`
}
