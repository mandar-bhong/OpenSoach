package dbquery

const QUERY_GET_COMPLAINT_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count  From spl_hkt_sp_complaint_tbl spcomplaint
Inner Join spl_node_sp_tbl sp on sp.sp_id_fk = spcomplaint.sp_id_fk $WhereCondition$`

const QUERY_COMPLAINT_TABLE_SELECT_BY_FILTER = `Select spcomplaint.id as id,sp_name,complaint_title,description,complaint_by,severity,raised_on,complaint_state,closed_on  From spl_hkt_sp_complaint_tbl spcomplaint
Inner Join spl_node_sp_tbl sp on sp.sp_id_fk = spcomplaint.sp_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_TOP_FIVE_COMPLAINTS = `select * from spl_hkt_sp_complaint_tbl where complaint_state = 1
order by severity desc ,raised_on desc
limit ?`
