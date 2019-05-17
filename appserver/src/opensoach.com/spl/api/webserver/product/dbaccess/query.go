package dbaccess

const QUERY_SELECT_CPM_DETAILS = `select cpm.id, prod.prod_code,cpm.cust_id_fk,cust.cust_name,dbinst.connection_string from spl_master_cust_prod_mapping_tbl cpm
INNER JOIN spl_master_product_tbl prod on cpm.prod_id_fk = prod.id
INNER JOIN spl_master_database_instance_tbl dbinst on cpm.dbi_id_fk = dbinst.id 
INNER JOIN spl_master_customer_tbl cust on cpm.cust_id_fk = cust.id
where cpm.id = :1`
