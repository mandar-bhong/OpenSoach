import { Injectable } from "@angular/core";
var Sqlite = require("nativescript-sqlite");

let selectQueries = new Map([
    [ "patientlist", "select fname,lname,bed_no,mob_no,status,attended from patient_admission_tbl as padmsn inner join patient_master_tbl as patient on patient.id = padmsn.patient_id" ],
    [ "chartlist", "select id,admission_id,conf_type_code,conf from patient_chart_conf_tbl" ],
    [ "chartInsert", "insert into patient_chart_conf_tbl (admission_id,conf_type_code,conf) values ( ?, ?, ?)" ],
    [ "monitorConfList", "select id,conf_type_code,conf from patient_conf_tbl where conf_type_code = 'Monitor'"],
    [ "actionList", "select id,admission_id,chart_conf_id,exec_time from action_tbl"],
    [ "actionInsert", "insert into action_tbl (admission_id,chart_conf_id,exec_time) values ( ?, ?, ?)" ]
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

    public insert(key: string,dataList:Array<any>) {

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

}