package dbquery

const QUERY_GET_SP_CATEGORY_SHORT_LIST = `select id,spc_name from spl_node_sp_category_tbl where cpm_id_fk=?`

const QUERY_DELETE_DEV_SP_MAPPING_TABLE_ROW = `Delete From spl_node_dev_sp_mapping Where dev_id_fk = :dev_id_fk And cpm_id_fk = :cpm_id_fk`

const QUERY_GET_NODE_SP_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count from spl_node_sp_tbl sp 
left join spl_node_dev_sp_mapping devsp  on devsp.sp_id_fk = sp.sp_id_fk
inner join spl_node_sp_category_tbl spc on spc.id = sp.spc_id_fk
left join spl_node_dev_tbl dev on dev.dev_id_fk = devsp.dev_id_fk
left join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.sp_id_fk = sp.sp_id_fk $WhereCondition$`

const QUERY_NODE_SP_TABLE_SELECT_BY_FILTER = `select sp.sp_id_fk,sp.sp_name,sp.spc_id_fk,spc.spc_name,sp.sp_state,sp.sp_state_since 
from spl_node_sp_tbl sp 
inner join spl_node_sp_category_tbl spc on spc.id = sp.spc_id_fk
left join spl_node_dev_sp_mapping devsp  on devsp.sp_id_fk = sp.sp_id_fk
left join spl_node_dev_tbl dev on dev.dev_id_fk = devsp.dev_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_SERVICEPOINT_SHORT_LIST = `select sp_id_fk,sp_name from spl_node_sp_tbl where cpm_id_fk = ?`

const QUERY_GET_SERVICE_POINT_BY_ID = `select * from spl_node_sp_tbl where sp_id_fk = ?`
