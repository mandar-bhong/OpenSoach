var Sqlite = require("nativescript-sqlite");

let selectQueries = new Map([
    ["patientlist", "select fname,lname,bed_no,mob_no,status,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.person_accompanying from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid left join patient_personal_details_tbl as pdetails on pdetails.admission_uuid = padmsn.uuid"],
    ["patientlistbyadmissionuuid", "select fname,lname,bed_no,mob_no,status,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.person_accompanying from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid left join patient_personal_details_tbl as pdetails on pdetails.admission_uuid = padmsn.uuid where padmsn.uuid=?"],
    ["patientlistbymasteruuid", "select fname,lname,bed_no,mob_no,status,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid, pdetails.person_accompanying from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid left join patient_personal_details_tbl as pdetails on pdetails.admission_uuid = padmsn.uuid where patient.uuid=?"],
    ["chartlist", "select * from schedule_tbl"],
    ["chartInsert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf) values ( ?, ?, ?, ?)"],
    ["monitorConfList", "select uuid,conf_type_code,conf from conf_tbl where conf_type_code = 'Monitor'"],
    ["actionList", "select * from action_tbl"],
    ["actionInsert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,scheduled_time, status, client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?)"],
    ["chartItemByUUID", "select * from schedule_tbl where uuid = ? "],
    ["getScheduleListActive", "select * from schedule_tbl where status=0 and end_date >=? and admission_uuid=?"],
    ["getScheduleListComplated", "select * from schedule_tbl where end_date <? and admission_uuid=?"],
    ["servicePointList", "select * from service_point_tbl"],
    ["actionTxnInsert", "insert into action_txn_tbl (uuid,admission_uuid,schedule_uuid,txn_data,scheduled_time,txn_state,conf_type_code, updated_on,runtime_config_data) values ( ?, ?, ?, ?, ?, ?, ?, ?,?)"],
    ["syncList", "select * from sync_tbl"],
    ["actionTxnList", "select * from action_txn_tbl"],
    ["service_point_tbl_insert", "insert into service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_by,updated_on,sync_pending,client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["service_point_tbl_update", "update service_point_tbl set sp_name=?,short_desc=?,sp_state=?,sp_state_since=?,updated_by=?,updated_on=?,sync_pending=?, client_updated_at=? where uuid=?"],
    ["schedule_tbl_insert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf,start_date,end_date,status,updated_by,updated_on,sync_pending,client_updated_at) values ( ?, ?, ?, ?, ?,?,?,?,?,?,?)"],
    ["action_tbl_insert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,scheduled_time,is_deleted,updated_by,updated_on, sync_pending,client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["action_tbl_update", "update action_tbl set admission_uuid=?,conf_type_code=?,schedule_uuid=?,scheduled_time=?,is_deleted=?,updated_by=?,updated_on=?, sync_pending=?,client_updated_at=? where uuid = ?"],
    ["action_tbl_delete", "DELETE FROM action_tbl"],
    ["patient_master_tbl_update", "update patient_master_tbl set  patient_reg_no=?, fname=?, lname=?, mob_no=?, age=?, blood_grp=?, gender=?, updated_by=?, updated_on=?, sync_pending=?, client_updated_at=? where uuid=?"],
    ["patient_admission_tbl_update", "update patient_admission_tbl set patient_uuid=?, patient_reg_no=?, bed_no=?, status=?, sp_uuid=?, dr_incharge=?, admitted_on=?, discharged_on=?,updated_by=?, updated_on=?, sync_pending=?, client_updated_at=? where uuid=?"],
    ["action_txn_tbl_insert", "insert into action_txn_tbl (uuid,admission_uuid,schedule_uuid,txn_data,scheduled_time,txn_state,conf_type_code, updated_by,updated_on,runtime_config_data,sync_pending, client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)"],
    ["monitorTxnList", "select schedule.conf,txn.schedule_uuid,txn_data,scheduled_time  from action_txn_tbl as txn left join schedule_tbl as schedule on txn.schedule_uuid = schedule.uuid where schedule.conf_type_code = 'Monitor' order by scheduled_time asc"],
    ["patient_master_tbl_insert", "insert into patient_master_tbl (uuid,patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_by, updated_on, sync_pending,client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["patient_admission_tbl_insert", "insert into patient_admission_tbl (uuid, patient_uuid, patient_reg_no, bed_no, status, sp_uuid, dr_incharge, admitted_on, discharged_on, updated_by, updated_on, sync_pending,client_updated_at) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["userList", "select * from device_access_tbl"],
    ["device_access_tbl_insert", "insert into device_access_tbl (userid, user_fname, user_lname, email, pin ) values ( ?, ?, ?, ?, ?)"],

    ["conf_tbl_insert", "insert into conf_tbl (uuid,conf_type_code,conf, updated_by,updated_on,sync_pending,client_updated_at) values (?, ?, ?, ?, ?, ?, ?)"],
    ["conf_tbl_update", "update conf_tbl set conf_type_code=?,conf=?, updated_by=?,updated_on=?,sync_pending=?,client_updated_at=? where uuid = ?"],
    ["schedule_tbl_update", "update schedule_tbl set admission_uuid=?,conf_type_code=?,conf=?, start_date=?,end_date=?,status=?,updated_by=?,updated_on=?,sync_pending=?,client_updated_at=? where uuid = ?"],
    ["patient_personal_details_tbl_insert", "insert into patient_personal_details_tbl (uuid,patient_uuid,admission_uuid,age,other_details,person_accompanying, updated_by,updated_on,sync_pending,client_updated_at) values (?,?,?,?,?,?,?,?,?,?)"],
    ["patient_personal_details_tbl_update", "update patient_personal_details_tbl set patient_uuid=?,admission_uuid=?,age=?,other_details=?,person_accompanying=?, updated_by=?,updated_on=?,sync_pending=?,client_updated_at=? where uuid = ?"],
    ["patient_medical_details_tbl_insert", "insert into patient_medical_details_tbl (uuid,patient_uuid,admission_uuid,present_complaints,reason_for_admission,history_present_illness,past_history,treatment_before_admission,investigation_before_admission,family_history,allergies,personal_history, updated_by,updated_on,sync_pending,client_updated_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"],
    ["patient_medical_details_tbl_update", "update patient_medical_details_tbl set patient_uuid=?,admission_uuid=?,present_complaints=?,reason_for_admission=?,history_present_illness=?,past_history=?,treatment_before_admission=?,investigation_before_admission=?,family_history=?,allergies=?,personal_history=?, updated_by=?,updated_on=?,sync_pending=?,client_updated_at=? where uuid=?"],
    ["action_txn_tbl_update", "update action_txn_tbl set admission_uuid=?,schedule_uuid=?,txn_data=?,scheduled_time=?,txn_state=?,conf_type_code=?, updated_by=?,updated_on=?,runtime_config_data=?,sync_pending=?, client_updated_at=? where uuid=?"],
    ["doctors_orders_tbl_insert", "insert into doctors_orders_tbl (uuid, admission_uuid, doctor_id, doctors_orders, comment, ack_by, ack_time, status, order_created_time, order_type, document_uuid, document_name, doctype, updated_by, updated_on, sync_pending, client_updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "],
    ["doctors_orders_tbl_update", "update doctors_orders_tbl set admission_uuid=?, doctor_id=?, doctors_orders=?, comment=?, ack_by=?, ack_time=?, status=?, order_created_time=?, order_type=?, document_uuid=?, document_name=?, doctype=?, updated_by=?, updated_on=?, sync_pending=?, client_updated_at=? where uuid=? "],
    ['getdoctororders', 'select * from doctors_orders_tbl where admission_uuid=? '],

    ["getActionListActive", "select * from action_tbl where scheduled_time >=? and admission_uuid=?"],
    ["getActionListComplated", "select * from action_tbl where scheduled_time <? and admission_uuid=?"],
    ["document_tbl_insert", "insert into document_tbl (uuid,doc_path,doc_name,doc_type,datastore,updated_by,updated_on,sync_pending,client_updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["document_tbl_delete", "delete from document_tbl where uuid=?"],
    ["documentget", "select * from document_tbl"],
    ["treatment_tbl_insert", "insert into treatment_tbl (uuid, admission_uuid, treatment_done, treatment_performed_time, details, post_observation, updated_by, updated_on, sync_pending, client_updated_at) values (?,?,?,?,?,?,?,?,?,?)"],
    ["treatment_tbl_update", "update treatment_tbl set admission_uuid=?, treatment_done=?, treatment_performed_time=?, details=?, post_observation=?, updated_by=?, updated_on=?, sync_pending=?, client_updated_at=? where uuid = "],
    ["treatment_doc_tbl_insert", "insert into treatment_doc_tbl (treatment_uuid,document_uuid) values (?,?)"],
    ["pathology_record_tbl_insert", "insert into pathology_record_tbl (uuid, admission_uuid, test_performed, test_performed_time, test_result, comments,updated_by, updated_on, sync_pending, client_updated_at) values (?,?,?,?,?,?,?,?,?,?)"],
    ["pathology_record_tbl_update", "update pathology_record_doc_tbl set admission_uuid=?, test_performed=?, test_performed_time=?, test_result=?, comments=?,updated_by=?, updated_on=?, sync_pending=?, client_updated_at=? where uuid =?"],
    ["pathology_record_doc_tbl_insert", "insert into pathology_record_doc_tbl (pathology_record_uuid,document_uuid) values (?,?)"],
    ["documentlist", "select * from document_tbl"],

    ["patient_admission_details", "select * from patient_admission_tbl where patient_uuid=? "],
    ["patient_personal_details", "select * from patient_master_tbl where uuid=? "],
    ["patient_person_accompanying_details", "select * from patient_personal_details_tbl where admission_uuid=? "],
    ["patient_medical_details", "select * from patient_medical_details_tbl where admission_uuid=? "],
    ["getAllActionList", "select * from action_tbl where admission_uuid=?"],

    ["getActionList", `select act.uuid as action_uuid,act.admission_uuid,act.schedule_uuid,act.scheduled_time,act.conf_type_code,atxn.uuid as action_txn_uuid,atxn.txn_data,
    atxn.client_updated_at,atxn.txn_state,usr.fname,usr.lname, sch.conf
    from action_tbl act
    left join schedule_tbl sch on sch.uuid = act.schedule_uuid
    left join action_txn_tbl atxn on atxn.schedule_uuid = act.schedule_uuid and atxn.scheduled_time = act.scheduled_time
    left join usr_tbl usr on usr.usr_id = act.updated_by
    where act.is_deleted = 0 and act.admission_uuid = ? 
    order by act.scheduled_time asc`],
    ["usr_tbl_insert", "insert into usr_tbl (usr_id,usr_name,urole_name,fname,lname,updated_on,sync_pending,client_updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)"],
    ["usr_tbl_update", "update usr_tbl set usr_name=?,urole_name=?,fname=?,lname=?,updated_on=?,sync_pending=?,client_updated_at=? where usr_id=?"],
    ["action_tbl_next_actions_all",
        `select * from 
    (select action_tbl.admission_uuid,
    action_tbl.schedule_uuid,
    max(action_tbl.scheduled_time) as scheduled_time
    from action_tbl left join action_txn_tbl
    on action_tbl.schedule_uuid = action_txn_tbl.schedule_uuid and action_tbl.scheduled_time = action_txn_tbl.scheduled_time
    where action_tbl.is_deleted = 0 and action_tbl.scheduled_time < ? and action_txn_tbl.uuid IS NULL
    group by action_tbl.schedule_uuid
    UNION
    select action_tbl.admission_uuid,
    action_tbl.schedule_uuid,
    action_tbl.scheduled_time
    from action_tbl  left join action_txn_tbl
    on  action_tbl.schedule_uuid = action_txn_tbl.schedule_uuid and action_tbl.scheduled_time = action_txn_tbl.scheduled_time
    where action_tbl.is_deleted = 0 and action_tbl.scheduled_time >= ? and action_txn_tbl.uuid IS NULL
    ) actions
    order by actions.scheduled_time ASC`],
    ["action_tbl_next_actions_for_admission",
        `select * from 
    (select action_tbl.admission_uuid,
    action_tbl.schedule_uuid,
    max(action_tbl.scheduled_time) as scheduled_time
    from action_tbl left join action_txn_tbl
    on action_tbl.schedule_uuid = action_txn_tbl.schedule_uuid and action_tbl.scheduled_time = action_txn_tbl.scheduled_time
    where action_tbl.is_deleted = 0 and action_tbl.admission_uuid = ? and action_tbl.scheduled_time < ? and action_txn_tbl.uuid IS NULL
    group by action_tbl.schedule_uuid
    UNION
    select action_tbl.admission_uuid,
    action_tbl.schedule_uuid,
    action_tbl.scheduled_time
    from action_tbl  left join action_txn_tbl
    on  action_tbl.schedule_uuid = action_txn_tbl.schedule_uuid and action_tbl.scheduled_time = action_txn_tbl.scheduled_time
    where action_tbl.is_deleted = 0 and action_tbl.admission_uuid = ? and action_tbl.scheduled_time >= ? and action_txn_tbl.uuid IS NULL
    ) actions
    order by actions.scheduled_time ASC`],
    ["getActionForCancel", `select act.uuid,act.admission_uuid,act.conf_type_code,act.schedule_uuid,act.scheduled_time,act.is_deleted,act.updated_by,act.updated_on,act.sync_pending,act.client_updated_at,atxn.uuid as atxn_uuid from action_tbl act
    left join action_txn_tbl atxn on act.scheduled_time = atxn.scheduled_time
    where act.scheduled_time >=? and act.schedule_uuid=? and atxn.uuid IS NULL`],
    ["userList1", "select * from usr_tbl"],
]);

let selectTableName = new Map([
    ["service_point_tbl", "service_point_tbl"],
    ["conf_tbl", "conf_tbl"],
    ["patient_master_tbl", "patient_master_tbl"],
    ["schedule_tbl", "schedule_tbl"],
    ["patient_admission_tbl", "patient_admission_tbl"],
    ["patient_personal_details_tbl", "patient_personal_details_tbl"],
    ["patient_medical_details_tbl", "patient_medical_details_tbl"],
    ["action_txn_tbl", "action_txn_tbl"],
    ["action_tbl", "action_tbl"],
    ["device_access_tbl", "device_access_tbl"],
    ["doctors_orders_tbl", "doctors_orders_tbl"],
    ["document_tbl", "document_tbl"],
    ["treatment_tbl", "treatment_tbl"],
    ["treatment_doc_tbl", "treatment_doc_tbl"],
    ["pathology_record_tbl", "pathology_record_tbl"],
    ["pathology_record_doc_tbl", "pathology_record_doc_tbl"],
    ["mst_user_tbl", "usr_tbl"],
]);


export class DatabaseHelper {

    public static getdbConn() {
        return new Sqlite('patients');
    }

    public static closedbConn() {
        new Sqlite('patients')
            .then((db) => {
                db.close();
            })
    }


    public static deleteDatabaseInDebugMode() {
        if (Sqlite.exists("patients")) {
            console.log('database exists');
            Sqlite.deleteDatabase("patients");
        }
    }

    public static selectAll(key: string): any {

        return new Promise((resolve, reject) => {

            var query: string;

            if (selectQueries.has(key) == true) {
                query = selectQueries.get(key);
            };
            this.getdbConn()
                .then(db => {

                    db.resultType(Sqlite.RESULTSASOBJECT);

                    db.all(query, function (err, resultSet) {

                        if (err) {
                            // console.log("select query error:",err);
                            reject(err);
                        } else {
                            // console.log("Result set is:", resultSet);
                            resolve(resultSet);
                        }

                    });

                });

        });
    }


    public static update(key: string, dataList: Array<any>) {

        return new Promise((resolve, reject) => {

            var query: string;

            if (selectQueries.has(key) == true) {
                query = selectQueries.get(key);
            };

            this.getdbConn()
                .then(db => {
                    db.execSQL(query, dataList).then(id => {
                        // console.log("INSERT RESULT", id);
                        resolve(id);
                    }, error => {
                        // console.log("INSERT ERROR", error);
                        reject(error);
                    });
                });

        });

    }

    public static selectByID(key: string, paramList: Array<any>): any {
        return new Promise((resolve, reject) => {
            var query: string;
            if (selectQueries.has(key) == true) {
                query = selectQueries.get(key);
            };
            this.getdbConn()
                .then(db => {
                    db.resultType(Sqlite.RESULTSASOBJECT);
                    db.all(query, paramList, function (err, result) {
                        if (err) {
                            reject(err);
                        } else {
                            resolve(result);
                        }
                    });
                });
        });
    }

    public static DataStoreInsertUpdate(storename: string, dataList: Array<any>) {

        return new Promise((resolve, reject) => {


            // console.log("dataList in DataStoreInsertUpdate", dataList);
            var updateDatalist = dataList.slice(0);
            updateDatalist = updateDatalist.concat(updateDatalist.splice(0, 1));

            // console.log("updateDatalist", updateDatalist);

            var tblname: string;
            var getQuery: string;
            var getParamList = [];

            if (selectTableName.has(storename) == true) {
                tblname = selectTableName.get(storename);
            };

            switch (storename) {
                case "mst_user_tbl":
                    getQuery = "select * from TABLENAME where usr_id = ?";
                    getParamList = dataList[0];
                    break;
                case "treatment_doc_tbl":
                    getQuery = "select * from TABLENAME where treatment_uuid = ? and document_uuid = ?";
                    getParamList.push(dataList[0])
                    getParamList.push(dataList[1])
                    break;
                case "pathology_record_doc_tbl":
                    getQuery = "select * from TABLENAME where pathology_record_uuid = ? and document_uuid = ?";
                    getParamList.push(dataList[0])
                    getParamList.push(dataList[1])
                    break;
                default:
                    getQuery = "select * from TABLENAME where uuid = ?";
                    getParamList = dataList[0]
                    break;
            }

            // getQuery = "select * from TABLENAME where uuid = ?";

            getQuery = getQuery.replace("TABLENAME", tblname);
            // console.log("getQuery", getQuery);

            this.getdbConn()
                .then(db => {

                    db.resultType(Sqlite.RESULTSASOBJECT);

                    db.get(getQuery, getParamList, function (err, row) {

                        if (err) {
                            console.log("getQuery err", err);
                            reject(err);
                        } else {
                            if (row == null) {
                                console.log("inserting data..");
                                DatabaseHelper.update(tblname.concat("_insert"), dataList)
                                    .then(
                                        (result) => {
                                            console.log("inserting id:", result);
                                            resolve(result);
                                        },
                                        (err) => {
                                            console.log("err", err);
                                            reject(err);
                                        }
                                    );
                            } else {
                                console.log("updating data..");
                                DatabaseHelper.update(tblname.concat("_update"), updateDatalist)
                                    .then(
                                        (result) => {
                                            console.log("updated id:", result);
                                            resolve(result);
                                        },
                                        (err) => {
                                            console.log("err", err);
                                            reject(err);
                                        }
                                    );
                            }

                        }

                    });

                });

        });

    }

    public static getSyncPendingDataStore(storename: string): any {

        return new Promise((resolve, reject) => {

            var tblname: string;
            var getQuery = "select * from TABLENAME where sync_pending = 1";

            if (selectTableName.has(storename) == true) {
                tblname = selectTableName.get(storename);
            };

            getQuery = getQuery.replace("TABLENAME", tblname);
            //  console.log("getQuery", getQuery);

            this.getdbConn()
                .then(db => {

                    db.resultType(Sqlite.RESULTSASOBJECT);

                    db.all(getQuery, function (err, result) {

                        if (err) {
                            reject(err);
                        } else {
                            //  console.log("getSyncPendingDataStore data:", result);
                            resolve(result);
                        }

                    });

                });

        });
    }

    public static updateSyncStoreSyncPending(storename: string, syncType: number, isSyncPending: number): any {

        return new Promise((resolve, reject) => {

            var paramList = [];
            var updateQuery = "";
            const currentTime = new Date().toISOString();
            paramList.push(isSyncPending);

            if (syncType === 1) {
                if (isSyncPending == 1) {
                    paramList.push(currentTime);
                    paramList.push(storename);
                    updateQuery = "update sync_tbl set sync_from_server_pending = ?,sync_from_server_pending_time = ? where store_name = ?";
                } else {
                    paramList.push(storename);
                    updateQuery = "update sync_tbl set sync_from_server_pending = ? where store_name = ?";
                }
            } else if (syncType === 2) {
                if (isSyncPending == 1) {
                    paramList.push(currentTime);
                    paramList.push(storename);
                    updateQuery = "update sync_tbl set sync_to_server_pending = ?,sync_to_server_pending_time = ?  where store_name = ?";
                } else {
                    paramList.push(storename);
                    updateQuery = "update sync_tbl set sync_to_server_pending = ? where store_name = ?";
                }
            }

            this.getdbConn()
                .then(db => {

                    db.execSQL(updateQuery, paramList).then(id => {
                        console.log("updateSyncStoreSyncPending");
                        console.log("affected rows :", id);
                        resolve(id);
                    }, error => {
                        console.log("updateSyncStoreSyncPending");
                        console.log("db error:", error);
                        reject(error);
                    });
                });

        });
    }

    public static updateTableSyncPending(storename: string, syncpendingtime: Date): any {

        return new Promise((resolve, reject) => {

            //  console.log("updateTableSyncPending");
            //   console.log("storename", storename);

            var paramList = [];
            paramList.push(syncpendingtime);
            var tblname: string;
            var updateQuery = "update TABLENAME set sync_pending = 0 where sync_pending = 1 and client_updated_at = ?";

            if (selectTableName.has(storename) == true) {
                tblname = selectTableName.get(storename);
            };

            updateQuery = updateQuery.replace("TABLENAME", tblname);
            //  console.log("updateQuery", updateQuery);

            this.getdbConn()
                .then(db => {

                    db.execSQL(updateQuery, paramList).then(id => {
                        console.log("updateTableSyncPending");
                        console.log("affected rows", id);
                        resolve(id);
                    }, error => {
                        console.log("updateTableSyncPending");
                        console.log("db error", error);
                        reject(error);
                    });
                });

        });
    }

    public static updateSyncStoreLastSynched(storename: string, lastSynched: Date): any {

        return new Promise((resolve, reject) => {

            console.log("lastSynched:", lastSynched);

            const updateQuery = "update sync_tbl set last_synced = ? where store_name = ?"
            var paramList = [];
            paramList.push(lastSynched);
            paramList.push(storename);

            this.getdbConn()
                .then(db => {

                    db.execSQL(updateQuery, paramList).then(id => {
                        console.log("updateSyncStoreLastSynched");
                        console.log("affected rows :", id);
                        resolve(id);
                    }, error => {
                        console.log("updateSyncStoreLastSynched");
                        console.log("db error:", error);
                        reject(error);
                    });
                });

        });
    }
    // fucntion for fetch action for cancel schedule 
    public static getDataByParameters(key: string, paramList: Array<any>): any {
        return new Promise((resolve, reject) => {
            var query: string;
            if (selectQueries.has(key) == true) {
                query = selectQueries.get(key);
            };
            this.getdbConn()
                .then(db => {
                    db.resultType(Sqlite.RESULTSASOBJECT);
                    db.all(query, paramList, function (err, result) {
                        if (err) {
                            console.log('in inner promise error', err);
                            reject(err);
                        } else {
                            resolve(result);
                        }
                    });
                });
        });
    }// end of fucntion
}