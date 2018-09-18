package dbquery

const QUERY_SPL_NODE_DASHBOARD_DEVICE_SUMMARY = `select connection_state,count(*) as count from
	spl_node_dev_status_tbl as devstate
	inner join spl_node_dev_tbl as dev on dev.dev_id_fk = devstate.dev_id_fk
	where dev.cpm_id_fk = ?
	group by connection_state`

const QUERY_SPL_NODE_DASHBOARD_LOCATION_SUMMARY = `select sp_state,count(*) as count from spl_node_sp_tbl where cpm_id_fk= ? group by sp_state`

const QUERY_SPL_NODE_DASHBOARD_FEEDBACK = `select feedback , count(*) as count from spl_node_feedback_tbl $WhereCondition$ group by feedback`

const QUERY_SPL_NODE_DASHBOARD_TASK = `select status,count(status) as count from spl_node_service_in_txn_tbl serv_in_txn
inner join spl_node_service_instance_tbl serv_conf_in on serv_in_txn.serv_in_id_fk = serv_conf_in.id
$WhereCondition$ group by status`

const QUERY_SPL_NODE_DASHBOARD_COMPLAINT_SUMMARY = `select  complaint_state,count(*) as count from spl_node_sp_complaint_tbl $WhereCondition$ group by complaint_state`

const QUERY_SPL_NODE_DASHBOARD_IN_USE_LOCATION_COUNT = `select count( distinct sp.sp_id_fk) as count  from spl_node_sp_tbl sp
inner join spl_node_dev_sp_mapping devsp on devsp.sp_id_fk = sp.sp_id_fk 
inner join spl_node_service_instance_tbl serv_conf_in on serv_conf_in.sp_id_fk = sp.sp_id_fk 
where sp.cpm_id_fk= ?`

const QUERY_GET_FEEDBACKS_PER_MONTH = `select year(raised_on) as year,
month(raised_on) as month,
count(if(feedback=1,1,null)) as rating1,
count(if(feedback=2,1,null)) as rating2,
count(if(feedback=3,1,null)) as rating3,
count(if(feedback=4,1,null)) as rating4,
count(if(feedback=5,1,null)) as rating5  
from spl_node_feedback_tbl
$WhereCondition$
group by month,year`

const QUERY_GET_NO_OF_COMPLAINTS_PER_MONTH = `select year(raised_on) as year,
month(raised_on) as month,
count(if(complaint_state=1,1,null)) as open,
count(if(complaint_state=2,1,null)) as closed,
count(if(complaint_state=3,1,null)) as inprogress
from spl_node_sp_complaint_tbl
$WhereCondition$
group by month,year`

const QUERY_GET_TOP_COMPLAINTS = `select id,complaint_title,raised_on,complaint_state,severity from spl_node_sp_complaint_tbl $WhereCondition$
order by severity desc ,raised_on desc
limit ?`

const QUERY_GET_TASK_SUMMARY_PER_MONTH = `select year(txn_date) as year,
month(txn_date) as month,
count(if(status=1,1,null)) as ontime,
count(if(status=2,1,null)) as delay 
from spl_node_service_in_txn_tbl
$WhereCondition$
group by month,year`

const QUERY_GET_TOP_FEEDBACKS = `select id,feedback,feedback_comment from spl_node_feedback_tbl $WhereCondition$
order by feedback desc limit ?`

const QUERY_GET_SNAPSHOT_DATA = `select max(txn_date) as txn_date,count(status) as count,status 
from spl_node_service_in_txn_tbl 
where status <> 2 and 
cpm_id_fk = ? and 
txn_date between ? and ? 
group by status`

const QUERY_GET_AVERAGE_TIME = `select(
	select avg(time_to_sec(timediff(b.txn_date,a.txn_date))) as waittime from spl_node_service_in_txn_tbl as a
	inner join spl_node_service_in_txn_tbl  as b
	on a.txn_data->'$.tokenid' = b.txn_data->'$.tokenid'  and a.status= 1 and b.status= 2 
	where a.txn_date $BetweenCondition$) as waittime,
	
	(select avg(time_to_sec(timediff(d.txn_date,c.txn_date))) as jobcreationtime from spl_node_service_in_txn_tbl as c
	inner join spl_node_service_in_txn_tbl  as d
	on c.txn_data->'$.tokenid' = d.txn_data->'$.tokenid'  and c.status = 2 and d.status= 3 
	where c.txn_date $BetweenCondition$) as jobcreationtime,
	
	(select avg(time_to_sec(timediff(f.txn_date,e.txn_date))) as jobexetime from spl_node_service_in_txn_tbl as e
	inner join spl_node_service_in_txn_tbl  as f
	on e.txn_data->'$.tokenid' = f.txn_data->'$.tokenid'  and e.status= 2 and f.status= 4 
	where e.txn_date $BetweenCondition$) as jobexetime,
	
	(select avg(time_to_sec(timediff(h.txn_date,g.txn_date))) as deliverytime from spl_node_service_in_txn_tbl as g
	inner join spl_node_service_in_txn_tbl  as h
	on g.txn_data->'$.tokenid' = h.txn_data->'$.tokenid'  and g.status= 5 and h.status= 6
	where g.txn_date $BetweenCondition$) as deliverytime
	`
