package models

import (
	"time"
)

type IStoreSync interface {
	GetUuid() string
}

type StoreSyncModel struct {
	StoreName string `json:"storename"`
}

type StoreEntityModel struct {
	Uuid string `db:"uuid" json:"uuid"`
}

type StoreSyncGetRequestModel struct {
	StoreSyncModel
	UpdatedOn   time.Time              `json:"updatedon"`
	QueryParams map[string]interface{} `json:"queryparams"`
}
type StoreSyncGetResponseModel struct {
	StoreSyncModel
	UpdatedOn time.Time   `json:"updatedon"`
	Count     int         `json:"count"`
	Data      interface{} `json:"data"`
}

type StoreSyncApplyRequestModel struct {
	StoreSyncModel
	Data interface{} `json:"storedata"`
}

type StoreSyncApplyResponseModel struct {
}

type SyncConfigModel struct {
	Id             int       `db:"id" json:"storeid"`
	StoreName      string    `db:"store_name" json:"storename"`
	HasQuery       string    `db:"has_qry" json:"hasquery"`
	SelectCountQry string    `db:"select_count_qry" json:"selectcountqry"`
	SelectQry      string    `db:"select_qry" json:"selectqry"`
	InsertQry      string    `db:"insert_qry" json:"insertqry"`
	UpdateQry      string    `db:"update_qry" json:"updateqry"`
	UpdatedOn      time.Time `db:"updated_on" json:"updatedon"`
}

type StoreConfigModel struct {
}

type SyncConfigTblInfoModel struct {
	Count        int       `db:"count" json:"count"`
	MaxUpdatedOn time.Time `db:"max_updated_on" json:"updatedon"`
}
