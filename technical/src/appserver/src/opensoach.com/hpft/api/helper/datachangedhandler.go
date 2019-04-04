package helper

import (
	"reflect"

	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/api/repository"
	constants "opensoach.com/hpft/constants"
	hpftconst "opensoach.com/hpft/constants"
	lmodels "opensoach.com/hpft/models"
	pcmodels "opensoach.com/prodcore/models"
)

var mapTblnameStorename = map[string]string{
	constants.DB_TABLE_SPL_NODE_SP_TBL:                 constants.SYNC_STORE_SERVICE_POINT,
	constants.DB_SPL_HPFT_CONF_TBL:                     constants.SYNC_STORE_CONF,
	constants.DB_SPL_HPFT_PATIENT_MASTER_TBL:           constants.SYNC_STORE_PATIENT_MASTER,
	constants.DB_SPL_HPFT_PATIENT_CONF_TBL:             constants.SYNC_STORE_PATIENT_CONF,
	constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL:        constants.SYNC_STORE_PATIENT_ADMISSION,
	constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL: constants.SYNC_STORE_PERSONAL_DETAILS,
	constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL:  constants.SYNC_STORE_MEDICAL_DETAILS,
	constants.DB_SPL_ACTION_TXN_TBL:                    constants.SYNC_STORE_ACTION_TXN,
	constants.DB_SPL_HPFT_DOCTORS_OPRDERS_TBL:          constants.SYNC_STORE_DOCTORS_ORDERS,
	constants.DB_SPL_HPFT_TREATMENT_TBL:                constants.SYNC_STORE_TREATMENT,
	constants.DB_SPL_HPFT_TREATMENT_DOC_TBL:            constants.SYNC_STORE_TREATMENT_DOC,
	constants.DB_SPL_HPFT_PATHOLOGY_RECORD_TBL:         constants.SYNC_STORE_PATHOLOGY,
	constants.DB_SPL_HPFT_PATHOLOGY_RECORD_DOC_TBL:     constants.SYNC_STORE_PATHOLOGY_DOC,
	constants.DB_SPL_HPFT_ACTION_TBL:                   constants.SYNC_STORE_ACTION,
}

func HandleDatabaseDataChange(tableName string, data interface{}) {

	sendNotifyTask(data, tableName)

}

func sendNotifyTask(taskdata interface{}, tblname string) {

	v := reflect.ValueOf(taskdata)
	iStoreCPM := v.Interface().(pcmodels.IStoreCPM)

	cpmid := iStoreCPM.GetCPMId()

	storename, ok := mapTblnameStorename[tblname]
	if ok == false {
		logger.Context().WithField("tblname", tblname).LogError("", logger.Normal, "Unable to get table storename from map", nil)
	}

	taskDBChangesModel := lmodels.TaskDBChangesModel{}
	taskDBChangesModel.CpmId = cpmid
	taskDBChangesModel.StoreName = storename

	isSendSuccess := repo.Instance().
		SendTaskToServer(hpftconst.TASK_HPFT_API_NOTIFY_DB_CHANGES, "", taskDBChangesModel)

	if isSendSuccess == false {
		logger.Context().WithField("TaskDBChangesModel", taskDBChangesModel).LogError("", logger.Normal, "Unable to submit notify db changes task to server.", nil)
	}

}
