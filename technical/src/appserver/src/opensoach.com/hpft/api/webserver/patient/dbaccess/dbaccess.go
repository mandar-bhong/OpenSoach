package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hpft/api/helper"
	"opensoach.com/hpft/constants"
	"opensoach.com/hpft/constants/dbquery"
	hktmodels "opensoach.com/hpft/models"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.Patient.DB"

func GetDBTransaction(dbconn string) (error, *sqlx.Tx) {
	ctx := dbmgr.InsertTxContext{}
	return ctx.GetTransaction(dbconn)
}

func Insert(dbConn string, insrtStruct *hktmodels.DBPatientMasterInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MASTER_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func ServiceInstanceInsert(tx *sqlx.Tx, insrtStruct *hktmodels.DBServiceInstanceInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing ServiceInstance insert.")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SERVICE_INSTANCE
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateByFilter(dbConn string, updtStruct *hktmodels.DBPatientUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient UpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MASTER_TBL
	updateErr := updateCtx.UpdateByFilter("PatientId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func UpdatePatientStatus(dbConn string, updtStruct *hktmodels.DBPatientUpdateStatusRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Device UpdatePatientStatus")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL
	updateErr := updateCtx.UpdateByFilter("AdmissionId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetPatientById(dbConn string, patientId int64) (error, *[]hktmodels.DBSplHpftPatientMasterTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientMasterTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_PATIENT_MASTER_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(patientId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetPatientByFilter(dbConn string, filterStruct *hktmodels.DBPatientFilterModel) (error, *[]hktmodels.DBSplHpftPatientMasterTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientMasterTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MASTER_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectByFilter(*filterStruct)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func AdmissionTblInsert(dbConn string, insrtStruct *hktmodels.DBAdmissionTblInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing admission tbl insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func AdmissionTblUpdateByFilter(dbConn string, updtStruct *hktmodels.DBAdmissionTblUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient AdmissionTblUpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_ADMISSION_TBL
	updateErr := updateCtx.UpdateByFilter("AdmissionId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetAdmissionById(dbConn string, admissionId int64) (error, *[]hktmodels.DBSplHpftPatientAdmissionTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetAdmissionById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientAdmissionTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_PATIENT_ADMISSION_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(admissionId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func PersonalDetailsInsert(dbConn string, insrtStruct *hktmodels.DBPersonalDetailsInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing PersonalDetailsInsert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func PersonalDetailsUpdateByFilter(dbConn string, updtStruct *hktmodels.DBPersonalDetailsUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient PersonalDetailsUpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("PersonalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func PersonalDetailsUpdatePersonAccompanying(dbConn string, updtStruct *hktmodels.DBPersonalDetailsUpdatePersonAccompanyingRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient PersonalDetailsUpdatePersonAccompanying")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_PERSONAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("PersonalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetPersonalDetailsById(dbConn string, personalDetailsId int64) (error, *[]hktmodels.DBSplHpftPatientPersonalDetailsRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPersonalDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientPersonalDetailsRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_PATIENT_PERSONAL_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(personalDetailsId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func MedicalDetailsInsert(dbConn string, insrtStruct *hktmodels.DBMedicalDetailsInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing MedicalDetailsInsert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func MedicalDetailsUpdateByFilter(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdatePresentComplaints(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdatePresentComplaintsRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdatePresentComplaints")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateReasonForAdmission(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateReasonForAdmissionRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateReasonForAdmission")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateHistoryPresentIllness(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateHistoryPresentIllnessRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateHistoryPresentIllness")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdatePastHistory(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdatePastHistoryRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdatePastHistory")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateTreatmentBeforeAdmission(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateTreatmentBeforeAdmissionRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateTreatmentBeforeAdmission")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateInvestigationBeforeAdmission(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateInvestigationBeforeAdmissionRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateInvestigationBeforeAdmission")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateFamilyHistory(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateFamilyHistoryRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateFamilyHistory")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdateAllergies(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdateAllergiesRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdateAllergies")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func MedicalDetailsUpdatePersonalHistory(dbConn string, updtStruct *hktmodels.DBMedicalDetailsUpdatePersonalHistoryRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient MedicalDetailsUpdatePersonalHistory")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_MEDICAL_DETAILS_TBL
	updateErr := updateCtx.UpdateByFilter("MedicalDetailsId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetMedicalDetailsById(dbConn string, medicalDetailsId int64) (error, *[]hktmodels.DBSplHpftPatientMedicalDetailsRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetMedicalDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientMedicalDetailsRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_PATIENT_MEDICAL_DETAILS_TABLE_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(medicalDetailsId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetPatientList(dbConn string, filterModel *hktmodels.DBSearchPatientRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchPatientResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_SPL_PATIENT_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_SPL_PATIENT_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient Filter Record list filter query : "+listQuery)

	data := &gmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &hktmodels.DBTotalRecordsModel{}
	selectCtxCount.DBConnection = dbConn
	selectCtxCount.Dest = dataCount
	selectCtxCount.Query = countQuery
	selectCtxCount.QueryType = dbmgr.Query
	selectCtxCountErr := selectCtxCount.Get()
	if selectCtxCountErr != nil {
		return selectCtxCountErr, nil
	}

	data.RecordCount = dataCount.TotalRecords

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	resdata := &[]hktmodels.DBSearchPatientResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = resdata
	selectCtx.Query = listQuery
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, nil
	}

	data.RecordList = resdata

	return nil, data
}

func GetPatientConfList(dbConn string, cpmid int64) (error, *[]hktmodels.DBSplHpftPatientConfTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientConfList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientConfTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_PATIENT_CONF_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetPatientConfById(dbConn string, confId int64) (error, *[]hktmodels.DBSplHpftPatientConfTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetMedicalDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHpftPatientConfTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_SPL_HPFT_PATIENT_CONF_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(confId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func PatientConfUpdateByFilter(dbConn string, updtStruct *hktmodels.DBPatientConfUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing patient PatientConfUpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_HPFT_PATIENT_CONF_TBL
	updateErr := updateCtx.UpdateByFilter("PatientConfId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetPatientMasterList(dbConn string, filterModel *hktmodels.DBSearchPatientMasterRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientMasterList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchPatientMasterResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_PATIENT_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_SPL_MASTER_PATIENT_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient master Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Patient master Filter Record list filter query : "+listQuery)

	data := &gmodels.ServerListingResultModel{}

	selectCtxCount := dbmgr.SelectContext{}
	dataCount := &hktmodels.DBTotalRecordsModel{}
	selectCtxCount.DBConnection = dbConn
	selectCtxCount.Dest = dataCount
	selectCtxCount.Query = countQuery
	selectCtxCount.QueryType = dbmgr.Query
	selectCtxCountErr := selectCtxCount.Get()
	if selectCtxCountErr != nil {
		return selectCtxCountErr, nil
	}

	data.RecordCount = dataCount.TotalRecords

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	resdata := &[]hktmodels.DBSearchPatientMasterResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = resdata
	selectCtx.Query = listQuery
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, nil
	}

	data.RecordList = resdata

	return nil, data
}

func GetAdmissionStatusById(dbConn string, admissionId int64) (error, *[]hktmodels.DBPatientAdmissionStatusInfoModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetAdmissionStatusById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBPatientAdmissionStatusInfoModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_PATIENT_ADMISSION_TABLE_STATUS_SELECT_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(admissionId)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
