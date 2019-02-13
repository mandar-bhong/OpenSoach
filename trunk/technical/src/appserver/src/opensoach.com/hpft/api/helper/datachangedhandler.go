package helper

import (
	"reflect"

	repo "opensoach.com/hpft/api/repository"
	constants "opensoach.com/hpft/constants"
	hpftconst "opensoach.com/hpft/constants"
	lmodels "opensoach.com/hpft/models"
	pcmodels "opensoach.com/prodcore/models"
)

var mapTblnameStorename = map[string]string{
	constants.DB_TABLE_SPL_NODE_SP_TBL:                 constants.SYNC_STORE_NAME_SERVICE_POINT,
	constants.DB_SPL_HPFT_CONF_TBL:                     constants.SYNC_STORE_NAME_CONF,
	constants.DB_SPL_HPFT_PATIENT_MASTER_TBL:           constants.SYNC_STORE_NAME_PATIENT_MASTER,
	constants.DB_SPL_HPFT_PATIENT_CONF_TBL:             constants.SYNC_STORE_NAME_PATIENT_CONF,
	constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL:        constants.SYNC_STORE_NAME_PATIENT_ADMISSION,
	constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL: constants.SYNC_STORE_NAME_PERSONAL_DETAILS,
	constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL:  constants.SYNC_STORE_NAME_MEDICAL_DETAILS,
	constants.DB_SPL_ACTION_TXN_TBL:                    constants.SYNC_STORE_NAME_ACTION_TXN,
	constants.DB_SPL_HPFT_DOCTORS_OPRDERS_TBL:          constants.SYNC_STORE_NAME_DOCTORS_ORDERS,
}

func HandleDatabaseDataChange(tableName string, data interface{}) {

	switch tableName {

	case constants.DB_SPL_HPFT_PATIENT_MASTER_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_PATIENT_CONF_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_ACTION_TXN_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_TABLE_SPL_NODE_SP_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_CONF_TBL:

		sendNotifyTask(data, tableName)

		break

	case constants.DB_SPL_HPFT_DOCTORS_OPRDERS_TBL:

		sendNotifyTask(data, tableName)

		break

	}

}

func sendNotifyTask(taskdata interface{}, tblname string) {

	v := reflect.ValueOf(taskdata)
	iStoreCPM := v.Interface().(pcmodels.IStoreCPM)

	cpmid := iStoreCPM.GetCPMId()

	storename := mapTblnameStorename[tblname]

	taskDBChangesModel := lmodels.TaskDBChangesModel{}
	taskDBChangesModel.CpmId = cpmid
	taskDBChangesModel.StoreName = storename

	isSendSuccess := repo.Instance().
		SendTaskToServer(hpftconst.TASK_HPFT_API_NOTIFY_DB_CHANGES, "", taskDBChangesModel)

	if isSendSuccess == false {
		// logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to submit task to server.", nil)
	}

}
