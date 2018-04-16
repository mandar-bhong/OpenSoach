package models

type DBSearchCustomerDataModel struct {
	Name  *string `db:"cust_name" json:"name"`
	State *int    `db:"cust_state" json:"state"`
}
