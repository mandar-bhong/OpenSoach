package models

import "time"

type DBFeedbackInsertRowModel struct {
	CpmIdFk         int64     `db:"cpm_id_fk" json:"cpmidfk"`
	SpIdFk          int64     `db:"sp_id_fk" json:"spidfk"`
	Feedback        int       `db:"feedback" json:"feedback"`
	FeedbackComment *string   `db:"feedback_comment" json:"feedbackcomment"`
	RaisedOn        time.Time `db:"raised_on" json:"raisedon"`
}
