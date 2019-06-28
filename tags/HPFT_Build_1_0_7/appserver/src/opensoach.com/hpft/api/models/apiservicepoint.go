package models

import (
	hktmodels "opensoach.com/hpft/models"
)

type APISpCategoryAddRequest struct {
	hktmodels.DBSpCategoryDataModel
}

type APISpAddRequest struct {
	hktmodels.DBSpDataRowModel
}

type APIDevSpAsscociationRequest struct {
	hktmodels.DBDevSpMappingDataModelModel
}

type APIDevSpAsscociationRemoveRequest struct {
	DevId int64 `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	SpId  int64 `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}
