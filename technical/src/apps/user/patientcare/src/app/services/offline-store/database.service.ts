import { Injectable } from "@angular/core";
var Sqlite = require("nativescript-sqlite");

let selectQueries = new Map([
    [ "patientlist", "select fname,lname,bed_no,mob_no,status,attended,padmsn.sp_uuid,sp_name from patient_admission_tbl as padmsn left join patient_master_tbl as patient on patient.uuid = padmsn.patient_uuid left join service_point_tbl as sp on sp.uuid = padmsn.sp_uuid" ],
    [ "chartlist", "select * from schedule_tbl" ],
    [ "chartInsert", "insert into schedule_tbl (uuid,admission_uuid,conf_type_code,conf) values ( ?, ?, ?, ?)" ],
    [ "monitorConfList", "select uuid,conf_type_code,conf from schedule_tbl where conf_type_code = 'Monitor'"],
    [ "actionList", "select * from action_tbl"],
    [ "actionInsert", "insert into action_tbl (uuid,admission_uuid,conf_type_code,schedule_uuid,exec_time) values ( ?, ?, ?, ?, ?)" ],
    [ "chartItemByUUID", "select * from schedule_tbl where uuid = ? "],
    [ "servicePointList", "select * from service_point_tbl"],
    [ "actionTxnInsert", "insert into action_txn_tbl (uuid,schedule_uuid,txn_data,txn_date,txn_state,conf_type_code,runtime_config_data) values ( ?, ?, ?, ?, ?, ?, ?)" ],
    [ "syncList", "select * from sync_tbl"],
    [ "actionTxnList", "select * from action_txn_tbl"],


]);

@Injectable()
export class DatabaseService {

    public getdbConnection() {

        return new Sqlite('patients');
    }

    public closedbConnection() {
        new Sqlite('patients')
            .then((db) => {
                db.close();
            })
    }

    public deleteDatabaseInDebugMode() {
        if (Sqlite.exists("patients")) {
            console.log('database exists');
            Sqlite.deleteDatabase("patients");
        }
    }

    public selectAll(key: string) :any {

        return new Promise((resolve, reject) => {

        var query: string;

        if (selectQueries.has(key) == true) {
            query = selectQueries.get(key);
        };

        this.getdbConnection()
            .then(db => {   

                db.resultType(Sqlite.RESULTSASOBJECT);

                db.all(query, function (err, resultSet) {

                    if (err){
                        // console.log("select query error:",err);
                        reject(err);
                    }else{
                        // console.log("Result set is:", resultSet);
                        resolve(resultSet);                  
                    }                 
                    
                });
                
            });

        });
      }

    public update(key: string,dataList:Array<any>) {

        return new Promise((resolve, reject) => {

        var query: string;

        if (selectQueries.has(key) == true) {
            query = selectQueries.get(key);
        };

        this.getdbConnection()
            .then(db => {   

                db.execSQL(query,dataList).then(id => {
                    // console.log("INSERT RESULT", id);
                    resolve(id);
                }, error => {
                    // console.log("INSERT ERROR", error);
                    reject(error);
                });
            });

        });

    }

    public selectByID(key: string, paramList:Array<any>) :any {

        return new Promise((resolve, reject) => {

        var query: string;

        if (selectQueries.has(key) == true) {
            query = selectQueries.get(key);
        };

        this.getdbConnection()
            .then(db => {   

                db.resultType(Sqlite.RESULTSASOBJECT);
                
                db.get(query, paramList ,function (err, row) {

                    if (err){
                        reject(err);
                    }else{
                        resolve(row);                  
                    }                 
                    
                });
                
            });

        });
      }

}