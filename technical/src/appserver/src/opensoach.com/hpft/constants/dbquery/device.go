package dbquery

const QUERY_SPL_NODE_DEVICE_TABLE_SELECT_SHORT_DATA_LIST = `select dev_id_fk,dev_name,serialno from spl_node_dev_tbl where cpm_id_fk = ?`

const QUERY_GET_DEVICE_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count  from spl_node_dev_tbl dev
left join spl_node_dev_status_tbl devstate on devstate.dev_id_fk = dev.dev_id_fk $WhereCondition$`

const QUERY_DEVICE_TABLE_SELECT_BY_FILTER = `select dev.dev_id_fk,dev.dev_name,dev.serialno,dev.created_on,dev.updated_on,devstate.connection_state,devstate.connection_state_since,devstate.sync_state,devstate.sync_state_since,devstate.battery_level,devstate.battery_level_since 
from spl_node_dev_tbl dev
left join spl_node_dev_status_tbl devstate on devstate.dev_id_fk = dev.dev_id_fk
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_SELECT_DEVICE_WITH_NO_SP_ASSOCIATION_SHORT_DATA_LIST = `select dev_id_fk,dev_name,serialno from spl_node_dev_tbl 
where dev_id_fk NOT IN (select dev_id_fk from  spl_node_dev_sp_mapping) and cpm_id_fk = ?`

const QUERY_SPL_NODE_DEVICE_TABLE_SELECT_BY_ID = `select * from spl_node_dev_tbl where dev_id_fk = ?`
