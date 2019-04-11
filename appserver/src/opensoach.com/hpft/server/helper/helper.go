package helper

import (
	ghelper "opensoach.com/core/helper"
	hpftmodels "opensoach.com/hpft/models"
	constants "opensoach.com/hpft/server/constants"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func GetStoreTableStruct(packet []byte) (error, pcmodels.StoreSyncApplyRequestModel, *gmodels.DevicePacket) {

	reqModel := pcmodels.StoreSyncApplyRequestModel{}

	devPacket := &gmodels.DevicePacket{}
	devPacket.Payload = &reqModel

	convErr := ghelper.ConvertFromJSONBytes(packet, devPacket)
	if convErr != nil {
		return convErr, pcmodels.StoreSyncApplyRequestModel{}, nil
	}

	switch reqModel.StoreName {
	case constants.DB_PATIENT_MASTER_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientMasterTableRowModel{}
		break
	case constants.DB_PATIENT_ADMISSION_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientAdmissionTableRowModel{}
		break
	case constants.DB_SP_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplNodeSpTableRowModel{}
		break
	case constants.DB_CONF_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftConfTableRowModel{}
		break
	case constants.DB_PATIENT_CONF_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientConfTableRowModel{}
		break
	case constants.DB_PATIENT_PERSONAL_DETAILS_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientPersonalDetailsRowModel{}
		break
	case constants.DB_PATIENT_MEDICAL_DETAILS_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientMedicalDetailsRowModel{}
		break
	case constants.DB_ACTION_TXN_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftActionTxnTableRowModel{}
		break

	}

	convErr = ghelper.ConvertFromJSONBytes(packet, devPacket)
	if convErr != nil {
		return convErr, pcmodels.StoreSyncApplyRequestModel{}, nil
	}

	switch reqModel.StoreName {
	case constants.DB_PATIENT_MASTER_TBL_STORE_NAME:
		ss := *reqModel.Data.(*[]hpftmodels.DBSplHpftPatientMasterTableRowModel)
		reqModel.Data = ss
		break
	case constants.DB_PATIENT_ADMISSION_TBL_STORE_NAME:
		ss := *reqModel.Data.(*[]hpftmodels.DBSplHpftPatientAdmissionTableRowModel)
		reqModel.Data = ss
		break

	case constants.DB_SP_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplNodeSpTableRowModel{}
		break

	case constants.DB_CONF_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftConfTableRowModel{}
		break

	case constants.DB_PATIENT_CONF_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientConfTableRowModel{}
		break

	case constants.DB_PATIENT_PERSONAL_DETAILS_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientPersonalDetailsRowModel{}
		break

	case constants.DB_PATIENT_MEDICAL_DETAILS_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftPatientMedicalDetailsRowModel{}
		break

	case constants.DB_ACTION_TXN_TBL_STORE_NAME:
		reqModel.Data = &[]hpftmodels.DBSplHpftActionTxnTableRowModel{}
		break

	}

	return nil, reqModel, devPacket

}
