package dbquery

const QUERY_SPL_NODE_REPORT_TEMPLATE_TABLE_SELECT_SHORT_DATA_LIST = `select id,report_code,report_desc from spl_node_report_template_tbl`

const QUERY_GET_REPORT_TASK_SUMMARY_PER_MONTH = `select 
status,
count(*) as count,
txn_data ->> '$.taskname' as taskname
from spl_node_service_in_txn_tbl serv_in_txn
inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.id = serv_in_txn.serv_in_id_fk
$WhereCondition$
group by status,taskname`
