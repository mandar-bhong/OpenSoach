package models

type DBFeedbackInsertRowModel struct {
	CpmIdFk  int64 `db:"cpm_id_fk" json:"cpmidfk"`
	SpIdFk   int64 `db:"sp_id_fk" json:"spidfk"`
	Feedback int   `db:"feedback" json:"feedback"`
}
