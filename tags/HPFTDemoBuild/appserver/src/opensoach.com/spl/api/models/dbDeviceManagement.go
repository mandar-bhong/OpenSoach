package models

import "time"

type DBSplMasterDeviceRowModel struct {
	DevId         int       `dbattr:"pri,auto" db:"id" json:"devid"`
	CustId        *int      `db:"cust_id_fk" json:"custid"`
	Serialno      string    `db:"serialno" json:"serialno"`
	DevState      int       `db:"dev_state" json:"devstate"`
	DevStateSince time.Time `db:"dev_state_since" json:"devstatesince"`
}

type DBSplMasterDevDetailsRowModel struct {
	DevId       int64   `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	Make        *string `db:"make" json:"make"`
	Technology  *string `db:"technology" json:"technology"`
	TechVersion *string `db:"tech_version" json:"techversion"`
	ShortDesc   *string `db:"short_desc" json:"shortdesc"`
}

type DBDevStateRowModel struct {
	DevId         int       `dbattr:"pri,auto" db:"id" json:"devid"`
	DevState      int       `db:"dev_state" json:"devstate"`
	DevStateSince time.Time `db:"dev_state_since" json:"devstatesince"`
}

type DBDevCustRowModel struct {
	DevId  int64 `dbattr:"pri,auto" db:"id" json:"devid"`
	CustId int64 `db:"cust_id_fk" json:"custid"`
}

type DBSplCpmDevRowModel struct {
	CpmId int64 `dbattr:"pri" db:"cpm_id_fk" json:"cpmid"`
	DevId int64 `dbattr:"pri" db:"dev_id_fk" json:"devid"`
}

type DBDeviceAssociateProdDataModel struct {
	CustName *string `db:"cust_name" json:"custname"`
	ProdCode *string `db:"prod_code" json:"prodcode"`
}

type DBDeviceShortDataModel struct {
	DevId    int64  `dbattr:"pri,auto" db:"id" json:"devid"`
	Serialno string `db:"serialno" json:"serialno"`
}
