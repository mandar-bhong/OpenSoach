package dbquery

const QUERY_DELETE_FOP_SP_TABLE_ROW = `Delete From spl_node_fop_sp_tbl Where fop_id_fk = :fop_id_fk`

const QUERY_GET_SP_CATEGORY_SHORT_LIST = `select id,spc_name from spl_node_sp_category_tbl`

const QUERY_DELETE_DEV_SP_MAPPING_TABLE_ROW = `Delete From spl_node_dev_sp_mapping Where dev_id_fk = :dev_id_fk And cpm_id_fk = :cpm_id_fk`
