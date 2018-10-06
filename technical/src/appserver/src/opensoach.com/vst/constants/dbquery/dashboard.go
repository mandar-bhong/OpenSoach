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

const QUERY_GET_AVERAGE_TIME = `select
avg(t.waittime) as waittime,
avg(t.jobcreationtime) as jobcreationtime,
avg(t.jobexetime) as jobexetime,
avg(t.deliverytime) as deliverytime
from
(
select
token.id as tokenid,
(select time_to_sec(timediff(b.txn_date,a.txn_date)) as waittime 
from spl_node_service_in_txn_tbl as a,spl_node_service_in_txn_tbl as b
where a.status = 1 and b.status = 2 and a.txn_data->'$.tokenid'= token.id and b.txn_data->'$.tokenid'= token.id) as waittime,
(select time_to_sec(timediff(d.txn_date,c.txn_date)) as jobcreationtime 
from spl_node_service_in_txn_tbl as c,spl_node_service_in_txn_tbl as d
where c.status = 2 and d.status = 3 and c.txn_data->'$.tokenid'= token.id and d.txn_data->'$.tokenid'= token.id) as jobcreationtime,
(select time_to_sec(timediff(f.txn_date,e.txn_date)) as jobexetime 
from spl_node_service_in_txn_tbl as e,spl_node_service_in_txn_tbl as f
where e.status = 2 and f.status = 4 and e.txn_data->'$.tokenid'= token.id and f.txn_data->'$.tokenid'= token.id limit 1) as jobexetime,
(select time_to_sec(timediff(h.txn_date,g.txn_date)) as deliverytime 
from spl_node_service_in_txn_tbl as g,spl_node_service_in_txn_tbl as h
where g.status = 5 and h.status = 6 and g.txn_data->'$.tokenid'= token.id and h.txn_data->'$.tokenid'= token.id) as deliverytime
from spl_vst_token token
inner join spl_node_service_in_txn_tbl serv_in_txn on token.id = serv_in_txn.txn_data->'$.tokenid'
$WhereCondition$
group by tokenid
) as t`

const QUERY_GET_VEHICLE_SUMMARY_PER_MONTH = `select year(txn_date) as year,
	month(txn_date) as month,
	count(if(status=6,1,null)) as vehicleserviced
	from spl_node_service_in_txn_tbl
	$WhereCondition$
	group by month,year`

const QUERY_GET_VEHICLE_SUMMARY_PER_WEEK = `select
	date(txn_date) as servicedate,
	count(if(status=6,1,null)) as vehicleserviced
	from spl_node_service_in_txn_tbl
	$WhereCondition$
	group by servicedate`

const QUERY_GET_AVG_TIME_SUMMARY_PER_MONTH = `select
year(generated_on) as year,
month(generated_on) as month,
avg(t1.waittime) as waittime,
avg(t1.waittime) as jobcreationtime,
avg(t1.waittime) as jobexetime,
avg(t1.waittime) as deliverytime
from
(
select
token.id as tokenid,
generated_on,
(select time_to_sec(timediff(b.txn_date,a.txn_date)) as waittime 
from spl_node_service_in_txn_tbl as a,spl_node_service_in_txn_tbl as b
where a.status = 1 and b.status = 2 and a.txn_data->'$.tokenid'= token.id and b.txn_data->'$.tokenid'= token.id) as waittime,
(select time_to_sec(timediff(d.txn_date,c.txn_date)) as jobcreationtime 
from spl_node_service_in_txn_tbl as c,spl_node_service_in_txn_tbl as d
where c.status = 2 and d.status = 3 and c.txn_data->'$.tokenid'= token.id and d.txn_data->'$.tokenid'= token.id) as jobcreationtime,
(select time_to_sec(timediff(f.txn_date,e.txn_date)) as jobexetime 
from spl_node_service_in_txn_tbl as e,spl_node_service_in_txn_tbl as f
where e.status = 2 and f.status = 4 and e.txn_data->'$.tokenid'= token.id and f.txn_data->'$.tokenid'= token.id limit 1) as jobexetime,
(select time_to_sec(timediff(h.txn_date,g.txn_date)) as deliverytime 
from spl_node_service_in_txn_tbl as g,spl_node_service_in_txn_tbl as h
where g.status = 5 and h.status = 6 and g.txn_data->'$.tokenid'= token.id and h.txn_data->'$.tokenid'= token.id) as deliverytime
from spl_vst_token token
inner join spl_node_service_in_txn_tbl serv_in_txn on token.id = serv_in_txn.txn_data->'$.tokenid'
$WhereCondition$
group by tokenid) as t1
group by month,year`

const QUERY_GET_AVG_TIME_SUMMARY_PER_WEEK = `select
date(generated_on) as servicedate,
avg(t1.waittime) as waittime,
avg(t1.waittime) as jobcreationtime,
avg(t1.waittime) as jobexetime,
avg(t1.waittime) as deliverytime
from
(
select
token.id as tokenid,
generated_on,
(select time_to_sec(timediff(b.txn_date,a.txn_date)) as waittime 
from spl_node_service_in_txn_tbl as a,spl_node_service_in_txn_tbl as b
where a.status = 1 and b.status = 2 and a.txn_data->'$.tokenid'= token.id and b.txn_data->'$.tokenid'= token.id) as waittime,
(select time_to_sec(timediff(d.txn_date,c.txn_date)) as jobcreationtime 
from spl_node_service_in_txn_tbl as c,spl_node_service_in_txn_tbl as d
where c.status = 2 and d.status = 3 and c.txn_data->'$.tokenid'= token.id and d.txn_data->'$.tokenid'= token.id) as jobcreationtime,
(select time_to_sec(timediff(f.txn_date,e.txn_date)) as jobexetime 
from spl_node_service_in_txn_tbl as e,spl_node_service_in_txn_tbl as f
where e.status = 2 and f.status = 4 and e.txn_data->'$.tokenid'= token.id and f.txn_data->'$.tokenid'= token.id limit 1) as jobexetime,
(select time_to_sec(timediff(h.txn_date,g.txn_date)) as deliverytime 
from spl_node_service_in_txn_tbl as g,spl_node_service_in_txn_tbl as h
where g.status = 5 and h.status = 6 and g.txn_data->'$.tokenid'= token.id and h.txn_data->'$.tokenid'= token.id) as deliverytime
from spl_vst_token token
inner join spl_node_service_in_txn_tbl serv_in_txn on token.id = serv_in_txn.txn_data->'$.tokenid'
group by tokenid) as t1
$WhereCondition$
group by servicedate`
