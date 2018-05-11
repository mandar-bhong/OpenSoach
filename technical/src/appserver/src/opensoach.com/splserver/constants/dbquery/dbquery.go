package dbquery

const QUERY_GET_DB_CONN_BY_ID = "select connection_string from spl_master_database_instance_tbl where id = ?"

const QUERY_GET_DB_CONN_BY_CPM_ID = `select dbi.connection_string from spl_master_database_instance_tbl dbi 
inner join spl_master_cust_prod_mapping_tbl cpm on dbi.id = cpm.dbi_id_fk
where cpm.id=?`
