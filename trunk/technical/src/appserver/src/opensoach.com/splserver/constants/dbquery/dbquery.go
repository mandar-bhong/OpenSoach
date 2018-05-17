package dbquery

const QUERY_GET_DB_CONN_BY_ID = "select connection_string from spl_master_database_instance_tbl where id = ?"

const QUERY_GET_DB_CONN_BY_CPM_ID = `select dbi.connection_string from spl_master_database_instance_tbl dbi 
inner join spl_master_cust_prod_mapping_tbl cpm on dbi.id = cpm.dbi_id_fk
where cpm.id=?`

const QUERY_SELECT_ALL_PROD_MASTER_SP_CATEGORY_TBL = `select id,spc_name,short_desc,created_on,updated_on from spl_prod_master_sp_category_tbl`
