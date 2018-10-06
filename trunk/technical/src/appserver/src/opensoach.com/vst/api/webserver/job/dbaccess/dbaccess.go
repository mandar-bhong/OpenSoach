package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	hkthelper "opensoach.com/vst/api/helper"
	"opensoach.com/vst/constants"
	"opensoach.com/vst/constants/dbquery"
	hktmodels "opensoach.com/vst/models"
)

var SUB_MODULE_NAME = "VST.API.Job.DB"

func GetDBTransaction(dbconn string) (error, *sqlx.Tx) {
	ctx := dbmgr.InsertTxContext{}
	return ctx.GetTransaction(dbconn)
}

func GetJobList(dbConn string, filterModel *hktmodels.DBSearchJobRequestFilterDataModel, listdatareq gmodels.APIDataListRequest, startingRow int) (error, *gmodels.ServerListingResultModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetJobList")

	if isParamValid := hkthelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		hkthelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := hkthelper.GetDBTagFromJSONTag(hktmodels.DBSearchJobResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := hkthelper.GetFilterConditionFormModel(*filterModel)

	if filterModel.StartDate != nil && filterModel.EndDate != nil {

		if whereCondition != "" {
			whereCondition = whereCondition + " and "
		}

		dbStartTime := filterModel.StartDate.Format(pcconst.DB_TIME_FORMAT)
		dbEndTime := filterModel.EndDate.Format(pcconst.DB_TIME_FORMAT)

		whereCondition = whereCondition + " generated_on between '" + dbStartTime + "' and '" + dbEndTime + "'"
	}

	if whereCondition != "" {
		whereCondition = " and " + whereCondition
	}

	countQuery := strings.Replace(dbquery.QUERY_GET_JOB_LIST_BY_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	listQuery := strings.Replace(dbquery.QUERY_JOB_LIST_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	listQuery = strings.Replace(listQuery, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Job Filter Record list filter count query : "+countQuery)
	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Job Filter Record list filter query : "+listQuery)

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
	resdata := &[]hktmodels.DBSearchJobResponseFilterDataModel{}
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

func UpdateJobStatus(tx *sqlx.Tx, updtStruct *hktmodels.DBJobStatusUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing UpdateJobStatus")

	updateCtx := dbmgr.UpdateDeleteTxContext{}
	updateCtx.Tx = tx
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_VST_TOKEN_TBL
	updateErr := updateCtx.UpdateByFilter("TokenId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func GetJobDetailsByTokenId(dbConn string, tokenid int64) (error, *[]hktmodels.DBSplNodeServiceInTxnTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetJobDetailsById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplNodeServiceInTxnTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_JOB_DETAILS_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(tokenid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func InsertJobDeliveredTxn(tx *sqlx.Tx, insrtStruct hktmodels.DBJobDeliveredTxnRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing InsertJobDeliveredTxn.")

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_NODE_SERV_IN_TXN_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}
