import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../offline-store/database.service";
import { async } from "rxjs/internal/scheduler/async";


@Injectable()
export class ReportsService {

    constructor(private database: DatabaseService) {

    }
    public getPathlogyReportByUUID(admission_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(admission_uuid);

            this.database.selectByID("getPathlogyReportList", paramList).then(
                (val) => {
                    // console.log('************ getPathlogyReportList Service', val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getPathlogyReportDoc(uuid: string): Promise<any> {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            console.log('uuid', uuid)
            paramList.push(uuid);

            this.database.selectByID("getPathlogyReportDoc", paramList).then(
                (val) => {
                    // console.log('************ get document name', val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }
    public getTreatmentReportByUUID(admission_uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(admission_uuid);

            this.database.selectByID("getTreatmentReportList", paramList).then(
                (val) => {
                    // console.log('************ getTreatmentReportByUUID Service', val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public getTreatmentReportDoc(uuid: string): Promise<any> {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            console.log('uuid', uuid)
            paramList.push(uuid);

            this.database.selectByID("getTreatmentReportDoc", paramList).then(
                (val) => {
                    // console.log('************ get Treatment document name', val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

}