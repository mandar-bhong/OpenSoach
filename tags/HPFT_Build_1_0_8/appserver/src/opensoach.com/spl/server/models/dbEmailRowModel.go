package models

type DBEmailTemplateRowModel struct {
	ID       int64   `dbattr:"pri" db:"id" json:"emiltmlid"`
	Code     string  `db:"code" json:"code"`
	Subject  string  `db:"subject" json:"subject"`
	Body     string  `db:"body" json:"body"`
	Bcc      *string `db:"bcc" json:"bcc"`
	MaxRetry int     `db:"maxretry" json:"maxretry"`
}

type DBEmailRowModel struct {
	ID         int64   `dbattr:"pri,auto" db:"id" json:"emilid"`
	TemplateID int64   `db:"email_tml_id_fk" json:"emailtmlid"`
	Subject    string  `db:"subject" json:"subject"`
	Body       string  `db:"body" json:"body"`
	Bcc        *string `db:"bcc" json:"bcc"`
	Status     int     `db:"status" json:"status"`
	Comment    *string `db:"comment" json:"comment"`
}
