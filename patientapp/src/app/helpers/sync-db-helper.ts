import { DatabaseHelper } from "./database-helper.js";


export class SyncDb {

    public static getSyncList(): any {

        return new Promise((resolve, reject) => {


            DatabaseHelper.selectAll("syncList").then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }

}