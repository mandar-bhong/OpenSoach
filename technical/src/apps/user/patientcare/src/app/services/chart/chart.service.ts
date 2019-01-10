import { Injectable, Version } from "@angular/core";
import { DatabaseService } from "../../services/offline-store/database.service";
import { ChartDBModel } from "~/app/models/ui/chart-models";

@Injectable()
export class ChartService {

    constructor(private database: DatabaseService) {

    }

    public getChartList(): any {

        return new Promise((resolve, reject) => {

            this.database.selectAll("chartlist").then(
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
            listData.push(data.admissionid);
            listData.push(data.conf_type_code);
            listData.push(data.conf);

            this.database.insert("chartInsert", listData).then(
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

    public getChartByUUID(uuid: string) {
        return new Promise((resolve, reject) => {

            const paramList = new Array<any>();

            paramList.push(uuid);

            this.database.selectByID("chartItemByUUID",paramList).then(
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
}