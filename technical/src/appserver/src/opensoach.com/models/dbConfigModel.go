package models

type DBMasterConfigRowModel struct {
	ConfigKey   string `db:"config_key" json:"configkey"`
	ConfigValue string `db:"config_value" json:"configvalue"`
}
