package dbaccess

import (
	"errors"
	"fmt"
	"strings"

	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants/dbquery"
	lhelper "opensoach.com/spl/helper"
	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.Corporate.DB"

func GetSplMasterCorpTableTotalFilteredRecords(dbConn string, filterModel *lmodels.DBSearchCorpRequestFilterDataModel) (error, *lmodels.DBTotalRecordsModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetSplMasterCorpTableTotalFilteredRecords")

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_GET_SPL_MASTER_CORP_TABLE_TOTAL_FILTERED_COUNT, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter count : "+query)

	selectCtx := dbmgr.SelectContext{}
	data := &lmodels.DBTotalRecordsModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Get()
	if selectErr != nil {
		return selectErr, &lmodels.DBTotalRecordsModel{}
	}
	return nil, data
}

func SplMasterCorpTableSelectByFilter(dbConn string, listdatareq lmodels.DataListRequest, filterModel *lmodels.DBSearchCorpRequestFilterDataModel, startingRow int) (error, *[]lmodels.DBSearchCorpResponseFilterDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing SplMasterCorpTableSelectByFilter")

	if isParamValid := lhelper.DBQueryParamValidate(listdatareq.OrderBy) &&
		lhelper.DBQueryParamValidate(listdatareq.OrderDirection); isParamValid == false {
		return errors.New(fmt.Sprintf("Invalid query paramter %s or %s ", listdatareq.OrderBy, listdatareq.OrderDirection)), nil
	}

	dbMatchedTag := lhelper.GetDBTagFromJSONTag(lmodels.DBSearchCorpResponseFilterDataModel{}, listdatareq.OrderBy)

	whereCondition := lhelper.GetFilterConditionFormModel(*filterModel)

	if whereCondition != "" {
		whereCondition = " where " + whereCondition
	}

	query := strings.Replace(dbquery.QUERY_SPL_MASTER_CORP_TABLE_SELECT_BY_FILTER, "$OrderByDirection$", dbMatchedTag+" "+listdatareq.OrderDirection, 1)
	query = strings.Replace(query, "$WhereCondition$", whereCondition, 1)

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "User Filter Record list filter : "+query)

	limit := listdatareq.Limit
	selectCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSearchCorpResponseFilterDataModel{}
	selectCtx.DBConnection = dbConn
	selectCtx.Dest = data
	selectCtx.Query = query
	selectCtx.QueryType = dbmgr.Query
	selectErr := selectCtx.Select(startingRow, limit)
	if selectErr != nil {
		return selectErr, &[]lmodels.DBSearchCorpResponseFilterDataModel{}
	}
	return nil, data
}
