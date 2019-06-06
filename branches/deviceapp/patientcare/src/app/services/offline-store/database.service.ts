import { Injectable } from "@angular/core";
import { DatabaseHelper } from "~/app/helpers/database-helper";

@Injectable()
export class DatabaseService {

    public getdbConnection() {
        return DatabaseHelper.getdbConn();
    }

    public closedbConnection() {
        DatabaseHelper.closedbConn();
    }

    public deleteDatabaseInDebugMode() {
        DatabaseHelper.deleteDatabaseInDebugMode();
    }

    public selectAll(key: string): any {
        const response = DatabaseHelper.selectAll(key);
        return response;
    }

    public update(key: string, dataList: Array<any>) {
        const response = DatabaseHelper.update(key, dataList);
        return response;
    }

    public selectByID(key: string, paramList: Array<any>): any {
        const response = DatabaseHelper.selectByID(key, paramList);
        return response;
    }

}