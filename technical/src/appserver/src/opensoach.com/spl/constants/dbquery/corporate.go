package dbquery

const QUERY_GET_SPL_MASTER_CORP_TABLE_TOTAL_FILTERED_COUNT = "SELECT count(*) as count FROM spl_master_corp_tbl $WhereCondition$"

const QUERY_SPL_MASTER_CORP_TABLE_SELECT_BY_FILTER = `SELECT id,corp_name,corp_mobile_no,corp_email_id,corp_landline_no,created_on,updated_on FROM spl_master_corp_tbl $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`
