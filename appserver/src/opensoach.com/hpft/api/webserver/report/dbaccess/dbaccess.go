package dbaccess

import (
	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	patientdbaccess "opensoach.com/hpft/api/webserver/patient/dbaccess"
	spdbaccess "opensoach.com/hpft/api/webserver/servicepoint/dbaccess"
	"opensoach.com/hpft/constants"
	"opensoach.com/hpft/constants/dbquery"
	hktmodels "opensoach.com/hpft/models"
)

var SUB_MODULE_NAME = "HPFT.API.Report.DB"

func GetReportInfo(dbConn string, reportid int64) (error, *[]hktmodels.DBSplNodeReportTemplateTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportInfo")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeReportTemplateTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_SPL_NODE_REPORT_TEMPLATE_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(reportid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetReportQueryData(dbConn string, query string, args ...interface{}) (error, []string, [][]string) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportQueryData")

	engine, err := sqlx.Connect("mysql", dbConn)

	if err != nil {
		return err, nil, nil
	}

	rows, err := engine.Queryx(query, args...)

	if err != nil {
		return err, nil, nil
	}

	cols, err := rows.Columns()
	if err != nil {
		//fmt.Println("Failed to get columns", err)
		return err, nil, nil
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))

	var resultRows [][]string

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {

		result := make([]string, len(cols))

		err = rows.Scan(dest...)
		if err != nil {
			//fmt.Println("Failed to scan row", err)
			return err, nil, nil
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}

		resultRows = append(resultRows, result)
	}

	return nil, cols, resultRows
}

func GetReportShortDataList(dbConn string) (error, *[]hktmodels.DBReportTemplateShortDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportShortDataList")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBReportTemplateShortDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SPL_NODE_REPORT_TEMPLATE_TABLE_SELECT_SHORT_DATA_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetReportInfoByCode(dbConn string, reportcode string) (error, *[]hktmodels.DBSplNodeReportTemplateTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetReportInfoByCode")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeReportTemplateTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Query = dbquery.QUERY_SELECT_REPORT_TEMPLATE_BY_REPORT_CODE
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(reportcode)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetPatientAdmissionReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftPatientAdmissionTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientAdmissionReportData")

	err, data := patientdbaccess.GetAdmissionById(dbConn, admissionid)

	return err, data

}

func GetPatientMasterReportData(dbConn string, patientid int64) (error, *[]hktmodels.DBSplHpftPatientMasterTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientMasterReportData")

	err, data := patientdbaccess.GetPatientById(dbConn, patientid)

	return err, data

}

func GetPatientPersonalDetailsReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftPatientPersonalDetailsRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientPersonalDetailsReportData")

	err, data := patientdbaccess.GetPersonalDetailsByAdmissionId(dbConn, admissionid)

	return err, data

}

func GetPatientMedicalDetailsReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftPatientMedicalDetailsRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientMedicalDetailsReportData")

	err, data := patientdbaccess.GetMedicalDetailsDetailsByAdmissionId(dbConn, admissionid)

	return err, data

}

func GetPatientDoctorsOrdersReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftDoctorsOrdersTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientDoctorsOrdersReportData")

	err, data := patientdbaccess.GetPatientDoctorsOrdersByAdmissionId(dbConn, admissionid)

	return err, data

}

func GetPatientTreatmentReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftTreatmentTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientTreatmentReportData")

	err, data := patientdbaccess.GetPatientTreatmentByAdmissionId(dbConn, admissionid)

	return err, data

}

func GetPatientPathologicalRecordReportData(dbConn string, admissionid int64) (error, *[]hktmodels.DBSplHpftPathologyRecordTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetPatientPathologicalRecordReportData")

	err, data := patientdbaccess.GetPatientPathologyRecordsByAdmissionId(dbConn, admissionid)

	return err, data

}

func GetUserData(dbConn string, usrid int64) (error, *[]hktmodels.PatientUserInfo) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetUserData")

	err, data := patientdbaccess.GetUserInfoById(dbConn, usrid)

	return err, data

}

func GetServicePointData(dbConn string, spid int64) (error, *[]hktmodels.DBSplNodeSpTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetServiPointData")

	err, data := spdbaccess.ServicePointSelectByID(dbConn, spid)

	return err, data

}
