package models

type APITaskHandlerFunc func(msg string, sessionkey string, token string, taskData interface{}) (error, APITaskResultModel)

type APITaskHandlerModel struct {
	Handler     APITaskHandlerFunc
	PayloadType interface{}
}

type APITaskResultModel struct {
	IsSuccess bool                        `json:"isSuceess"`
	Data      interface{}                 `json:"data"`
	ErrorData APITaskResultErrorDataModel `json:"errordata"`
}

type APITaskResultErrorDataModel struct {
	ErrorCode int         `json:"code"`
	Data      interface{} `json:"data"`
}

type APIDBInstanceCpmIdInsertModel struct {
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}

type APIDBInstanceDevInsertRowModel struct {
	DevId int64 `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}
