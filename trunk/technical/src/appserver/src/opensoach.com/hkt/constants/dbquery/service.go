package dbquery

const QUERY_SERVICE_CONF_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_node_service_conf_tbl $WhereCondition$`

const QUERY_SERVICE_CONF_TABLE_SELECT_BY_FILTER = `Select id,cpm_id_fk,spc_id_fk,conf_type_code,serv_conf_name,short_desc,created_on,updated_on From spl_node_service_conf_tbl $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_SERVICE_INSTANCE_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From  spl_node_service_instance_tbl serv_conf_in
Inner Join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk $WhereCondition$`

const QUERY_SERVICE_INSTANCE_TABLE_SELECT_BY_FILTER = `Select serv_conf_in.id as id,serv_conf_in.sp_id_fk as sp_id_fk,serv_conf.conf_type_code as conf_type_code,serv_conf.serv_conf_name as serv_conf_name From  spl_node_service_instance_tbl serv_conf_in
Inner Join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_SERVICE_INSTANCE_TXN = `select serv_in_txn.id,serv_in_id_fk,fopcode,status,txn_data,txn_date
from spl_node_service_in_txn_tbl  as serv_in_txn inner join spl_node_service_instance_tbl as serv_conf_in on serv_conf_in.id =  serv_in_txn.serv_in_id_fk
where serv_in_txn.cpm_id_fk=? and serv_conf_in.sp_id_fk= ? and txn_date between ? and ?`

const QUERY_GET_SERVICE_CONF_SHORT_LIST = `select id,serv_conf_name from spl_node_service_conf_tbl`

const QUERY_INSERT_SERVICE_CONF_COPY = `insert into spl_node_service_conf_tbl (cpm_id_fk,spc_id_fk,conf_type_code,serv_conf_name,short_desc,serv_conf)
select cpm_id_fk,spc_id_fk,conf_type_code,serv_conf_name,short_desc,serv_conf from spl_node_service_conf_tbl 
where id =:id`
