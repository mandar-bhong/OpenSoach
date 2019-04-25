import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ActionDBModel, ActionTxnDBModel } from "~/app/models/ui/action-models";
import { ActionDataStoreModel } from "~/app/models/db/action-datastore";
import { UserDetailDBModel } from "~/app/models/ui/user-auth-models";

@Injectable()
export class UserAuthService {

    constructor(private database: DatabaseService) {

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
    public insertDeviceAccessItem(data: UserDetailDBModel) {

        return new Promise((resolve, reject) => {

            const listData = new Array<any>();

            listData.push(data.userid);
            listData.push(data.first_name);
            listData.push(data.last_name);
            listData.push(data.email);
            listData.push(data.pin);

            this.database.update("device_access_tbl_insert", listData).then(
                (val) => {
                    console.log("device access data", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getUserAccountList1(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("userList1").then(
                (val) => {
                    // console.log("User Account List 1", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
}