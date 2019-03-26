package models

import (
	"errors"
	"time"

	prodconst "opensoach.com/prodcore/constants"
)

type ISOTime time.Time

func (s ISOTime) MarshalJSON() ([]byte, error) {
	t := time.Time(s)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(prodconst.ISO_TIME_FORMAT)), nil
}

type PreExecuteFilterHandler func(ctx *DevicePacketProccessExecution, filterModel *SyncConfigModel, request *StoreSyncGetRequestModel) error

type IStoreSync interface {
	GetUuid() string
}

type IStoreCPM interface {
	GetCPMId() int64
}

type StoreSyncModel struct {
	StoreName string `json:"storename"`
}

type StoreEntityModel struct {
	Uuid string `db:"uuid" json:"uuid"`
}

type StoreSyncGetRequestModel struct {
	StoreSyncModel
	UpdatedOn     time.Time              `json:"updatedon"`
	QueryParams   map[string]interface{} `json:"queryparams"`
	FilterHandler PreExecuteFilterHandler
}
type StoreSyncGetResponseModel struct {
	StoreSyncModel
	UpdatedOn *ISOTime    `json:"updatedon"`
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
	DataSourceType int       `db:"data_source" json:"dbsource"`
	QueryData      string    `db:"query_data" json:"querydata"`
}

type StoreConfigModel struct {
}

type SyncConfigTblInfoModel struct {
	Count int `db:"count" json:"count"`
	MaxUpdatedOn *ISOTime `db:"max_updated_on" json:"updatedon"`
}

type CPMIDEntityModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type QueryDataModel struct {
	Select QueryFilterDataModel `json:"select"`
}

type QueryFilterDataModel struct {
	Filters []string `json:"filters"`
}
