package dbquery

const QUERY_FIELD_OPERATOR_TABLE_SELECT_BY_ID = `Select id,cpm_id_fk,fopcode,fop_name,mobile_no,email_id,short_desc,fop_state,fop_area,created_on,updated_on From spl_hkt_field_operator_tbl Where id = ?`

const QUERY_GET_SPL_MASTER_FOP_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_hkt_field_operator_tbl $WhereCondition$`

const QUERY_SPL_MASTER_FOP_TABLE_SELECT_BY_FILTER = `Select id,fopcode,fop_name,mobile_no,email_id,fop_state,fop_area From spl_hkt_field_operator_tbl $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`
