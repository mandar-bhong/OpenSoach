package dbquery

const QUERY_GET_DB_CONN_BY_CPM_ID = `select cpm.id,cpm.prod_id_fk,dbi.connection_string from spl_master_cust_prod_mapping_tbl cpm
inner join spl_master_database_instance_tbl dbi
on cpm.dbi_id_fk = dbi.id
where cpm.id =?`
