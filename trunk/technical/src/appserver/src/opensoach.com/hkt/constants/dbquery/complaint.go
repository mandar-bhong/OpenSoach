package dbquery

const QUERY_GET_COMPLAINT_TABLE_TOTAL_FILTERED_COUNT = `Select count(*) as count  From spl_hkt_sp_complaint_tbl spcomplaint
Inner Join spl_node_sp_tbl sp on sp.sp_id_fk = spcomplaint.sp_id_fk $WhereCondition$`

const QUERY_COMPLAINT_TABLE_SELECT_BY_FILTER = `Select spcomplaint.id as id,sp_name,complaint_title,description,complaint_by,severity,raised_on,complaint_state,closed_on  From spl_hkt_sp_complaint_tbl spcomplaint
Inner Join spl_node_sp_tbl sp on sp.sp_id_fk = spcomplaint.sp_id_fk $WhereCondition$ ORDER BY $OrderByDirection$ Limit ?,?`

const QUERY_GET_TOP_COMPLAINTS = `select id,complaint_title,raised_on,complaint_state,severity from spl_hkt_sp_complaint_tbl $WhereCondition$
order by severity desc ,raised_on desc
limit ?`

const QUERY_GET_NO_OF_COMPLAINTS_PER_MONTH = `select year(raised_on) as year,
month(raised_on) as month,
count(if(complaint_state=1,1,null)) as open,
count(if(complaint_state=2,1,null)) as closed,
count(if(complaint_state=3,1,null)) as inprogress
from spl_hkt_sp_complaint_tbl
$WhereCondition$
group by month,year`
