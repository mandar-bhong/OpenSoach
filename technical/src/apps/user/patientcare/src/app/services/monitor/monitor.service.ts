import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ActionDBModel, ActionTxnDBModel } from "~/app/models/ui/action-models";
import { ActionDataStoreModel } from "~/app/models/db/action-datastore";

@Injectable()
export class MonitorService {

    constructor(private database: DatabaseService) {

    }

    public getTempActionTxn(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorTxnList").then(
                (val) => {
                    // console.log("temp data", val);

                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Temperature');
                    // console.log("filterdata",filterdata);

                    resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
    public getBloodPreActionTxn(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorTxnList").then(
                (val) => {
                    // console.log("blood pressure data", val);

                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Blood pressure');
                    // console.log("filterdata",filterdata);

                    resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getRespirationActionTxn(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorTxnList").then(
                (val) => {
                    // console.log("Respiration data", val);

                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Respiration Rate');
                    // console.log("filterdata",filterdata);

                    resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getPulseActionTxn(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorTxnList").then(
                (val) => {
                    // console.log("Pulse data", val);

                    const filterdata  = val.filter(a => JSON.parse(a.conf).name ==='Pulse Rate');
                    // console.log("filterdata",filterdata);

                    resolve(filterdata);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public getUserAccountList(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("userList").then(
                (val) => {
                    // console.log("User Account List", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
}