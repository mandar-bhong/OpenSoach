var Sqlite = require("nativescript-sqlite");

let selectQueries = new Map([
    ["patientlist", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid"],
    ["patientlistbyadmissionuuid", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid where padmsn.uuid=?"],
    ["patientlistbymasteruuid", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid where patient.uuid=?"],
    ["chartlist", "select * from schedule_tbl"],
    ["chartInsert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf) values ( ?, ?, ?, ?)"],
    ["monitorConfList", "select uuid,conf_type_code,conf from conf_tbl where conf_type_code = 'Monitor'"],
    ["actionList", "select * from action_tbl"],
    ["actionInsert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,exec_time, status) values ( ?, ?, ?, ?, ?, ?)"],
    ["chartItemByUUID", "select * from schedule_tbl where uuid = ? "],
    ["getScheduleListActive", "select * from schedule_tbl where end_date >=?"],
    ["getScheduleListComplated", "select * from schedule_tbl where end_date <?"],
    ["servicePointList", "select * from service_point_tbl"],
    ["actionTxnInsert", "insert into action_txn_tbl (uuid,admission_uuid,schedule_uuid,txn_data,txn_date,txn_state,conf_type_code, updated_on,runtime_config_data) values ( ?, ?, ?, ?, ?, ?, ?, ?,?)"],
    ["syncList", "select * from sync_tbl"],
    ["actionTxnList", "select * from action_txn_tbl"],
    ["service_point_tbl_insert", "insert into service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_on,sync_pending) values ( ?, ?, ?, ?, ?, ?, ?)"],
    ["service_point_tbl_update", "update service_point_tbl set sp_name=?,short_desc=?,sp_state=?,sp_state_since=?,updated_on=?,sync_pending=? where uuid=?"],
    ["schedule_tbl_insert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf,end_date,sync_pending,sync_pending_time) values ( ?, ?, ?, ?,?,?,?)"],
    ["action_tbl_insert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,exec_time, sync_pending,sync_pending_time) values ( ?, ?, ?, ?, ?, ?, ?)"],
    ["action_tbl_delete", "DELETE FROM action_tbl"],
    ["patient_master_tbl_update", "update patient_master_tbl set  patient_reg_no=?, fname=?, lname=?, mob_no=?, age=?, blood_grp=?, gender=?, updated_on=?, sync_pending=?, sync_pending_time=? where uuid=?"],
    ["patient_admission_tbl_update", "update patient_admission_tbl set patient_uuid=?, patient_reg_no=?, bed_no=?, status=?, sp_uuid=?, dr_incharge=?, admitted_on=?, discharged_on=?, updated_on=?, sync_pending=?, sync_pending_time=? where uuid=?"],
    ["action_txn_tbl_insert", "insert into action_txn_tbl (uuid,admission_uuid,schedule_uuid,txn_data,txn_date,txn_state,conf_type_code,updated_on,runtime_config_data,sync_pending, sync_pending_time) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)"],
    ["monitorTxnList", "select schedule.conf,txn.schedule_uuid,txn_data,txn_date  from action_txn_tbl as txn left join schedule_tbl as schedule on txn.schedule_uuid = schedule.uuid where schedule.conf_type_code = 'Monitor'"],
    ["patient_master_tbl_insert", "insert into patient_master_tbl (uuid,patient_reg_no, fname, lname, mob_no, age, blood_grp, gender, updated_on, sync_pending,sync_pending_time) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["patient_admission_tbl_insert", "insert into patient_admission_tbl (uuid, patient_uuid, patient_reg_no, bed_no, status, sp_uuid, dr_incharge, admitted_on, discharged_on, updated_on, sync_pending,sync_pending_time) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["userList", "select * from user_account_tbl"],

    ["user_account_tbl_insert", "insert into user_account_tbl (id, user_fname, user_lname, email, pin ) values ( ?, ?, ?, ?, ?)"],
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
    ["user_account_tbl", "user_account_tbl"],
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
            var getQuery = "select * from TABLENAME where uuid = ?";

            if (selectTableName.has(storename) == true) {
                tblname = selectTableName.get(storename);
            };

            getQuery = getQuery.replace("TABLENAME", tblname);
            // console.log("getQuery", getQuery);

            this.getdbConn()
                .then(db => {

                    db.resultType(Sqlite.RESULTSASOBJECT);

                    db.get(getQuery, dataList[0], function (err, row) {

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


    public static filterWithParam(key: string, paramList: Array<any>): any {
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

    public static updateSyncStoreSyncPending(storename: string, syncType: number): any {

        return new Promise((resolve, reject) => {

            var paramList = [];
            var updateQuery = "";
            const currentTime = Date.now();
            paramList.push(currentTime);
            paramList.push(storename);

            if (syncType === 1) {
                updateQuery = "update sync_tbl set sync_from_server_pending = 1,sync_from_server_pending_time = ? where store_name = ?";
            } else if (syncType === 2) {
                updateQuery = "update sync_tbl set sync_to_server_pending = 1,sync_to_server_pending_time = ?  where store_name = ?";
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

    public static updateTableSyncPending(storename: string): any {

        return new Promise((resolve, reject) => {

            //  console.log("updateTableSyncPending");
            //   console.log("storename", storename);

            var tblname: string;
            var updateQuery = "update TABLENAME set sync_pending = 0 where sync_pending = 1";

            if (selectTableName.has(storename) == true) {
                tblname = selectTableName.get(storename);
            };

            updateQuery = updateQuery.replace("TABLENAME", tblname);
            //  console.log("updateQuery", updateQuery);

            this.getdbConn()
                .then(db => {

                    db.execSQL(updateQuery, []).then(id => {
                        console.log("updateTableSyncPending");
                        //  console.log("db result", id);
                        resolve(id);
                    }, error => {
                        console.log("updateTableSyncPending");
                        console.log("db error", error);
                        reject(error);
                    });
                });

        });
    }

}