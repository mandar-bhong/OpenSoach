package models

import (
	"time"
)

type DBSplMasterConfigRowModel struct {
	ConfigKey   string    `dbattr:"pri" db:"config_key" json:"configkey"`
	ConfigValue string    `db:"config_value" json:"configvalue"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCorpTableRowModel struct {
	CorpId         int64     `dbattr:"pri,auto" db:"id" json:"corpid"`
	CorpName       string    `db:"corp_name" json:"corpname"`
	CorpMobileNo   *string   `db:"corp_mobile_no" json:"corpmobileno"`
	CorpEmailId    *string   `db:"corp_email_id" json:"corpemailid"`
	CorpLandlineNo *string   `db:"corp_landline_no" json:"corplandlineno"`
	CreatedOn      time.Time `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCpmDevMappingTableRowModel struct {
	CpmId     int64     `dbattr:"pri" db:"cpm_id_fk" json:"cpmid"`
	DevId     int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCpmSpMappingTableRowModel struct {
	CpmId     int64     `dbattr:"pri" db:"cpm_id_fk" json:"cpmid"`
	SpId      int64     `dbattr:"pri" db:"sp_id_fk" json:"spid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCustDetailsTableRowModel struct {
	CustId         int64     `dbattr:"pri" db:"cust_id_fk" json:"custid"`
	Poc1Name       string    `db:"poc1_name" json:"poc1name"`
	Poc1EmailId    string    `db:"poc1_email_id" json:"poc1emailid"`
	Poc1MobileNo   string    `db:"poc1_mobile_no" json:"poc1mobileno"`
	Poc2Name       *string   `db:"poc2_name" json:"poc2name"`
	Poc2EmailId    *string   `db:"poc2_email_id" json:"poc2emailid"`
	Poc2MobileNo   *string   `db:"poc2_mobile_no" json:"poc2mobileno"`
	Address        *string   `db:"address" json:"address"`
	AddressState   *string   `db:"address_state" json:"addressstate"`
	AddressCity    *string   `db:"address_city" json:"addresscity"`
	AddressPincode *string   `db:"address_pincode" json:"addresspincode"`
	CreatedOn      time.Time `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCustProdCountTableRowModel struct {
	CpcId     int64     `dbattr:"pri,auto" db:"id" json:"cpcid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	DevCnt    int       `db:"dev_cnt" json:"devcnt"`
	SpCnt     int       `db:"sp_cnt" json:"spcnt"`
	UsrCnt    int       `db:"usr_cnt" json:"usrcnt"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCustProdMappingTableRowModel struct {
	CpmId         int64     `dbattr:"pri,auto" db:"id" json:"cpmid"`
	CustId        int64     `db:"cust_id_fk" json:"custid"`
	ProdId        int64     `db:"prod_id_fk" json:"prodid"`
	DbiId         int64     `db:"dbi_id_fk" json:"dbiid"`
	CpmState      int       `db:"cpm_state" json:"cpmstate"`
	CpmStateSince time.Time `db:"cpm_state_since" json:"cpmstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterCustomerTableRowModel struct {
	CustId         int64     `dbattr:"pri,auto" db:"id" json:"custid"`
	CorpId         int64     `db:"corp_id_fk" json:"corpid"`
	CustName       string    `db:"cust_name" json:"custname"`
	CustState      int       `db:"cust_state" json:"custstate"`
	CustStateSince time.Time `db:"cust_state_since" json:"custstatesince"`
	CreatedOn      time.Time `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterDatabaseInstanceTableRowModel struct {
	DbiId            int64     `dbattr:"pri,auto" db:"id" json:"dbiid"`
	DbiName          string    `db:"dbi_name" json:"dbiname"`
	ConnectionString string    `db:"connection_string" json:"connectionstring"`
	ProdId           int       `db:"prod_id_fk" json:"prodid"`
	CreatedOn        time.Time `db:"created_on" json:"createdon"`
	UpdatedOn        time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterDevDetailsTableRowModel struct {
	DevId       int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	DevName     *string   `db:"dev_name" json:"devname"`
	Make        *string   `db:"make" json:"make"`
	Technology  *string   `db:"technology" json:"technology"`
	TechVersion *string   `db:"tech_version" json:"techversion"`
	ShortDesc   *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterDevStatusTableRowModel struct {
	DevId                int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	ConnectionState      int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    time.Time `db:"battery_level_since" json:"batterylevelsince"`
	CreatedOn            time.Time `db:"created_on" json:"createdon"`
	UpdatedOn            time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterDeviceTableRowModel struct {
	DevId         int64     `dbattr:"pri,auto" db:"id" json:"devid"`
	Serialno      string    `db:"serialno" json:"serialno"`
	DevState      int       `db:"dev_state" json:"devstate"`
	DevStateSince time.Time `db:"dev_state_since" json:"devstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterProductTableRowModel struct {
	ProdId    int64       `dbattr:"pri,auto" db:"id" json:"prodid"`
	ProdCode  string    `db:"prod_code" json:"prodcode"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterServicepointTableRowModel struct {
	Id           int       `dbattr:"pri,auto" db:"id" json:"id"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterTotalCountTableRowModel struct {
	Id                int       `dbattr:"pri,auto" db:"id" json:"id"`
	CustCnt           int       `db:"cust_cnt" json:"custcnt"`
	UsrCnt            int       `db:"usr_cnt" json:"usrcnt"`
	DevCnt            int       `db:"dev_cnt" json:"devcnt"`
	SpCnt             int       `db:"sp_cnt" json:"spcnt"`
	DevUnallocatedCnt int       `db:"dev_unallocated_cnt" json:"devunallocatedcnt"`
	DevActiveCnt      int       `db:"dev_active_cnt" json:"devactivecnt"`
	CreatedOn         time.Time `db:"created_on" json:"createdon"`
	UpdatedOn         time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterUserRoleTableRowModel struct {
	UroleId   int       `dbattr:"pri,auto" db:"id" json:"uroleid"`
	ProdId    *int      `db:"prod_id_fk" json:"prodid"`
	UroleCode string    `db:"urole_code" json:"urolecode"`
	UroleName string    `db:"urole_name" json:"urolename"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterUserTableRowModel struct {
	UserId        int64     `dbattr:"pri,auto" db:"id" json:"userid"`
	UsrName       string    `db:"usr_name" json:"usrname"`
	UsrPassword   string    `db:"usr_password" json:"usrpassword"`
	UsrCategory   int       `db:"usr_category" json:"usrcategory"`
	UroleId       *int64    `db:"urole_id_fk" json:"uroleid"`
	UsrState      int       `db:"usr_state" json:"usrstate"`
	UsrStateSince time.Time `db:"usr_state_since" json:"usrstatesince"`
	CreatedOn     time.Time `db:"created_on" json:"createdon"`
	UpdatedOn     time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterUsrCpmTableRowModel struct {
	UcpmId    int       `dbattr:"pri,auto" db:"id" json:"ucpmid"`
	UserId    int       `db:"user_id_fk" json:"userid"`
	CpmId     int       `db:"cpm_id_fk" json:"cpmid"`
	UroleId   int       `db:"urole_id_fk" json:"uroleid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplMasterUsrDetailsTableRowModel struct {
	UsrId              int64     `dbattr:"pri" db:"usr_id_fk" json:"usrid"`
	Fname              *string   `db:"fname" json:"fname"`
	Lname              *string   `db:"lname" json:"lname"`
	MobileNo           *string   `db:"mobile_no" json:"mobileno"`
	AlternateContactNo *string   `db:"alternate_contact_no" json:"alternatecontactno"`
	CreatedOn          time.Time `db:"created_on" json:"createdon"`
	UpdatedOn          time.Time `db:"updated_on" json:"updatedon"`
}
