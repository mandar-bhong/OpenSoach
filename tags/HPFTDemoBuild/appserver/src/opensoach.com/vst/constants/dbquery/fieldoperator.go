package dbquery

const QUERY_FIELD_OPERATOR_TABLE_SELECT_BY_ID = `Select id,cpm_id_fk,fopcode,fop_name,mobile_no,email_id,short_desc,fop_state,fop_area,created_on,updated_on From spl_node_field_operator_tbl Where id = ?`

const QUERY_GET_SPL_MASTER_FOP_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_node_field_operator_tbl $WhereCondition$`

const QUERY_SPL_MASTER_FOP_TABLE_SELECT_BY_FILTER = `Select id,fopcode,fop_name,mobile_no,email_id,fop_state,fop_area From spl_node_field_operator_tbl $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_FIELD_OPERATOR_SHORT_LIST = `select id,fopcode,fop_name from spl_node_field_operator_tbl where cpm_id_fk=?`

const QUERY_DELETE_FOP_SP_TABLE_ROW = `Delete From spl_node_fop_sp_tbl Where fop_id_fk = :fop_id_fk and sp_id_fk = :sp_id_fk`

const QUERY_GET_FOP_SP_ASSOCIATIONS = `select fop_id_fk,fopsp.sp_id_fk,sp_name from spl_node_fop_sp_tbl fopsp
inner join spl_node_sp_tbl sp on sp.sp_id_fk = fopsp.sp_id_fk
where fop_id_fk = ? and fopsp.cpm_id_fk = ?`
