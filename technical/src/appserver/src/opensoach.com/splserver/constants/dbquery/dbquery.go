package dbquery

const QUERY_GET_DB_CONN_BY_ID = "select connection_string from spl_master_database_instance_tbl where id = ?"

const QUERY_GET_DB_CONN_BY_CPM_ID = `select dbi.connection_string from spl_master_database_instance_tbl dbi 
inner join spl_master_cust_prod_mapping_tbl cpm on dbi.id = cpm.dbi_id_fk
where cpm.id=?`

const QUERY_SELECT_ALL_PROD_MASTER_SP_CATEGORY_TBL = `select id,spc_name,short_desc,created_on,updated_on from spl_prod_master_sp_category_tbl`

const QUERY_SELECT_ALL_HKT_MASTER_TASK_LIB_TBL = `select * from spl_hkt_master_task_lib_tbl`

const QUERY_SELECT_EMAIL_TML_BY_CODE = `select id,code,subject,body,bcc,maxretry from spl_master_email_template_tbl where code = ?`

const QUERY_UPDATE_EMAIL_EMAIL_STATUS = `update spl_master_email_tbl set status = :status, comment=:comment where id = :id`
