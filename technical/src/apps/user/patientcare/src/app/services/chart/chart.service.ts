import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ChartDBModel } from "~/app/models/ui/chart-models";
import { Subscription, Subject } from "rxjs";
import { ServerDataStoreDataModel } from "~/app/models/api/server-data-store-data-model";
import { ScheduleDatastoreModel } from "~/app/models/db/schedule-model";

@Injectable()
export class ChartService {
    scheduleDataContext = new Subject<ServerDataStoreDataModel<ScheduleDatastoreModel>[]>();
    constructor(private database: DatabaseService) {

    }

    public getChartList(): any {
        return new Promise((resolve, reject) => {
            this.database.selectAll("chartlist").then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
    public getScheduleList(key: string): any {
        return new Promise((resolve, reject) => {
            const paramList = new Array<any>();
            const dt = new Date().toISOString();
            console.log('dt', dt);
            paramList.push(dt);
            console.log('param list', paramList);
            this.database.selectByID(key, paramList).then(
                (val) => {
                    console.log("chart data", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }



    public insertChartItem(data: ChartDBModel) {

        return new Promise((resolve, reject) => {

            const listData = new Array<any>();

            listData.push(data.uuid);
            listData.push(data.admission_uuid);
            listData.push(data.conf_type_code);
            listData.push(data.conf);

            this.database.update("chartInsert", listData).then(
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

    public getChartByUUID(uuid: string): any {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(uuid);

            this.database.selectByID("chartItemByUUID", paramList).then(
                (val) => {
                    console.log("chart item", val);
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });
    }

    public getMonitorConf(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("monitorConfList").then(
                (val) => {
                    resolve(val);
                },
                (error) => {
                    reject(error);
                }
            );

        });

    }
    setScheduleContext(value) {
        this.scheduleDataContext.next(value);
    }
}// end of init 