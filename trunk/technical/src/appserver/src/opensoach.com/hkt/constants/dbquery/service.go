package dbquery

const QUERY_SERVICE_CONF_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From spl_node_service_conf_tbl $WhereCondition$`

const QUERY_SERVICE_CONF_TABLE_SELECT_BY_FILTER = `Select id,cpm_id_fk,spc_id_fk,conf_type_code,serv_conf_name,short_desc,created_on,updated_on From spl_node_service_conf_tbl $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_SERVICE_INSTANCE_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count From  spl_node_service_instance_tbl serv_conf_in
Inner Join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk $WhereCondition$`

const QUERY_SERVICE_INSTANCE_TABLE_SELECT_BY_FILTER = `Select serv_conf_in.id as id,serv_conf_in.sp_id_fk as sp_id_fk,serv_conf.conf_type_code as conf_type_code,serv_conf.serv_conf_name as serv_conf_name From  spl_node_service_instance_tbl serv_conf_in
Inner Join spl_node_service_conf_tbl serv_conf on serv_conf.id = serv_conf_in.serv_conf_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`
