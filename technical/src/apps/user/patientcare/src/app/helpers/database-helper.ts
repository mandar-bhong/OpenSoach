var Sqlite = require("nativescript-sqlite");

let selectQueries = new Map([
    ["patientlist", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name,patient.uuid as patient_uuid,padmsn.uuid as admission_uuid from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid"],
    ["patientlistbyuuid", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid where padmsn.uuid=?"],
    ["patientlistbymasteruuid", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid where patient.uuid=?"],
    ["chartlist", "select * from schedule_tbl"],
    ["chartInsert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf) values ( ?, ?, ?, ?)"],
    ["monitorConfList", "select uuid,conf_type_code,conf from conf_tbl where conf_type_code = 'Monitor'"],
    ["actionList", "select * from action_tbl"],
    ["actionInsert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,exec_time, status) values ( ?, ?, ?, ?, ?, ?)"],
    ["chartItemByUUID", "select * from schedule_tbl where uuid = ? "],
    ["servicePointList", "select * from service_point_tbl"],
    ["actionTxnInsert", "insert into action_txn_tbl (uuid,admission_uuid,schedule_uuid,txn_data,txn_date,txn_state,conf_type_code,runtime_config_data, status) values ( ?, ?, ?, ?, ?, ?, ?, ?)"],
    ["syncList", "select * from sync_tbl"],
    ["actionTxnList", "select * from action_txn_tbl"],
    ["service_point_tbl_insert", "insert into service_point_tbl (uuid,sp_name,short_desc,sp_state,sp_state_since,updated_on,sync_pending) values ( ?, ?, ?, ?, ?, ?, ?)"],
    ["service_point_tbl_update", "update service_point_tbl set sp_name=?,short_desc=?,sp_state=?,sp_state_since=?,updated_on=?,sync_pending=? where uuid=?"],
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

            // console.log("dataList", dataList);

            var newDatalist = dataList.slice(0);
            newDatalist = newDatalist.concat(newDatalist.splice(0, 1));

            // console.log("newDatalist", newDatalist);

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
                                            console.log("inserting id:",result);
                                            resolve(result);
                                        },
                                        (err) => {
                                            console.log("err",err);
                                            reject(err);
                                        }
                                    );
                            } else {
                                console.log("updating data..");
                                DatabaseHelper.update(tblname.concat("_update"), newDatalist)
                                    .then(
                                        (result) => {
                                            console.log("updated id:",result);
                                            resolve(result);
                                        },
                                        (err) => {
                                            console.log("err",err);
                                            reject(err);
                                        }
                                    );
                            }

                        }

                    });

                });

        });

    }

}