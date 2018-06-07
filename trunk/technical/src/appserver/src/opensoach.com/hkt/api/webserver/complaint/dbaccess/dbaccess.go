package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	hkthelper "opensoach.com/hkt/api/helper"
	lmodels "opensoach.com/hkt/api/models"
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/constants/dbquery"
	hktmodels "opensoach.com/hkt/models"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

var SUB_MODULE_NAME = "HKT.API.Complaint.DB"

func Insert(dbConn string, insrtStruct *hktmodels.DBComplaintInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing complaint insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_SPL_HKT_SP_COMPLAINT_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func UpdateByFilter(dbConn string, updtStruct *hktmodels.DBComplaintUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing Complaint UpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_TABLE_SPL_HKT_SP_COMPLAINT_TBL
	updateErr := updateCtx.UpdateByFilter("ComplaintId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func ComplaintTableSelectByID(dbConn string, complaintid int64) (error, *[]hktmodels.DBSplHktSpComplaintTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing ComplaintTableSelectByID")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktSpComplaintTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_SPL_HKT_SP_COMPLAINT_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(complaintid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func ComplaintTableSelectAll(dbConn string) (error, *[]hktmodels.DBSplHktSpComplaintTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing ComplaintTableSelectAll")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplHktSpComplaintTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_TABLE_SPL_HKT_SP_COMPLAINT_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectAll()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data

}

func GetComplaintList(dbConn string, filterModel *hktmodels.DBSearchComplaintRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetComplaintList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchComplaintResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_COMPLAINT_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_COMPLAINT_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Complaint Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Complaint Filter Record list filter query : "+listQuery)

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
	resdata := &[]hktmodels.DBSearchComplaintResponseFilterDataModel{}
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

func SelectTopComplaints(dbConn string, filtermodel hktmodels.DBTopComplaintsFilterDataModel, noofcomplaints int) (error, *[]lmodels.APITopActiveComplaintsResponse) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SelectTopFiveComplaints")

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	Query := strings.Replace(dbquery.QUERY_GET_TOP_COMPLAINTS, "$WhereCondition$", whereCondition, 1)

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.APITopActiveComplaintsResponse{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.Query = Query
	selErr := selDBCtx.Select(noofcomplaints)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func GetNoOfComplaintsPerMonth(dbConn string, req lmodels.APIComplaintsByMonthRequest, filtermodel hktmodels.DBNoOfComplaintsPerMonthsFilterDataModel) (error, []hktmodels.DBNoOfComplaintsPerMonthDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetNoOfComplaintsPerMonth")

	data := []hktmodels.DBNoOfComplaintsPerMonthDataModel{}

	whereCondition := hkthelper.GetFilterConditionFormModel(filtermodel)

	if req.StartDate != nil && req.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := req.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := req.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " raised_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_NO_OF_COMPLAINTS_PER_MONTH, "$WhereCondition$", whereCondition, 1)

	selectCtx := dbmgr.SelectContext{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = &data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectCtxErr := selectCtx.Select()
	if selectCtxErr != nil {
		return selectCtxErr, nil
	}

	return nil, data
}
