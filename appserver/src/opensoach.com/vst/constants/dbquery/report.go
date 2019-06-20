package dbquery

const QUERY_SPL_NODE_REPORT_TEMPLATE_TABLE_SELECT_SHORT_DATA_LIST = `select id,report_code,report_desc from spl_node_report_template_tbl`

const QUERY_SELECT_REPORT_TEMPLATE_BY_REPORT_CODE = `select * from spl_node_report_template_tbl where report_code = ?`
