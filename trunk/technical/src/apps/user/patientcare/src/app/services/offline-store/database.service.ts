import { Injectable } from "@angular/core";
var Sqlite = require("nativescript-sqlite");

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
    
    public deleteDatabaseInDebugMode()
    {
        if (Sqlite.exists("patients")) {
            console.log('database exists');
            Sqlite.deleteDatabase("patients");
          }
    }
}