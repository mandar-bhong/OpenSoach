package dbquery

const QUERY_GET_FEEDBACK_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count  from spl_node_feedback_tbl $WhereCondition$`

const QUERY_FEEDBACK_TABLE_SELECT_BY_FILTER = `select id,feedback,feedback_comment from spl_node_feedback_tbl
$WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`
