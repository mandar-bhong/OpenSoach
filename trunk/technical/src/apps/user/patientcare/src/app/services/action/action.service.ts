import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ActionDBModel, ActionTxnDBModel } from "~/app/models/ui/action-models";
import { ActionDataStoreModel } from "~/app/models/db/action-datastore";

@Injectable()
export class ActionService {

    constructor(private database: DatabaseService) {

    }

    public getActionList(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("actionList").then(
                (val) => {
                    // console.log("action data",val);                  
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }

    public getActionTxnList(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("actionTxnList").then(
                (val) => {
                    console.log("action data service", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );
        });

    }

    public getActionActiveList(key: string, admission_uuid: string, getAllFlag: boolean): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            if (getAllFlag === true) {
                const dt = new Date().toISOString();
                console.log('dt', dt);
                paramList.push(dt);
                console.log('filter data console');
            }
            paramList.push(admission_uuid);
            // console.log('param list', paramList);
            this.database.selectByID(key, paramList).then(
                (val) => {
                    // console.log("Action data", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }

    public getallActionActiveList(admission_uuid: string): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            paramList.push(admission_uuid);
            // console.log('param list', paramList);
            this.database.selectByID("getActionList", paramList).then(
                (val) => {
                    console.log("Action All data new ", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }


    public insertActionItem(data: ActionDataStoreModel) {

        return new Promise((resolve, reject) => {
            const listData = new Array<any>();

            listData.push(data.uuid);
            listData.push(data.admission_uuid);
            listData.push(data.conf_type_code);
            listData.push(data.schedule_uuid);
            listData.push(data.scheduled_time);
            listData.push(data.sync_pending);

            this.database.update("actionInsert", listData).then(
                (val) => {
                    // console.log("chart data",val);                  
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public insertActionTxnItem(data: ActionTxnDBModel) {

        return new Promise((resolve, reject) => {

            const listData = new Array<any>();

            listData.push(data.uuid);
            listData.push(data.admission_uuid);
            listData.push(data.schedule_uuid);
            listData.push(data.txn_data);
            listData.push(data.scheduled_time);
            listData.push(data.txn_state);
            listData.push(data.conf_type_code);
            listData.push(data.runtime_config_data);
            listData.push(data.status);

            this.database.update("actionTxnInsert", listData).then(
                (val) => {
                    // console.log("chart data",val);                  
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    // end of code block.
    public getDoctorsList(key: string, admission_uuid: string): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            paramList.push(admission_uuid);
            this.database.selectByID(key, paramList).then(
                (val) => {
                    resolve(val);
                }, (error) => {
                    reject(error);
                }
            );
        });
    }

}