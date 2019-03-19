import { IDatastoreMessageHandler } from "./idatastore-message-handler.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { Schedulardata, SchedularConfigData } from "../models/ui/chart-models.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ActionDataStoreModel, ActionsData } from "../models/db/action-datastore.js";
import { SYNC_STORE, ConfigCodeType, ActionStatus } from "../app-constants.js";
import { MonitorHelper } from "../helpers/actions/monitor-helper.js";
import { IntakeHelper } from "../helpers/actions/intake-helper.js";
import { MedicineHelper } from "../helpers/actions/medicine-helper.js";

export class CancelScheduleDatastoreMessageHandler implements IDatastoreMessageHandler<ScheduleDatastoreModel>{
    handleMessage(msg: ScheduleDatastoreModel): Promise<ActionDataStoreModel[]> {
        return new Promise(async (resolve, reject) => {         
            const schedulardata = new Schedulardata();
            schedulardata.data = msg;
            schedulardata.conf = new SchedularConfigData()
            const parsedConf = <SchedularConfigData>JSON.parse(msg.conf);
            schedulardata.conf = parsedConf;
            let actiondata: ActionsData;
            try {             
                actiondata = new ActionsData();
                actiondata.actions = [];            
                try {
                    const success = await this.getChartList(schedulardata.data.uuid, msg.updated_by);
                    success.forEach((item) => {
                        item.is_deleted = ActionStatus.ACTION_DELETED;
                        item.updated_by = msg.updated_by;
                        item.sync_pending = 1;
                        item.client_updated_at = new Date().toISOString();
                        actiondata.actions.push(item);
                        resolve(actiondata.actions);
                    });
                    msg.end_date = new Date().toISOString();
                }
                catch (e) {
                    console.error('Action not received', e);
                }
            }
            catch (e) {
                console.error('Cancel Action Helper', e);
            }
        });
    }
    async getChartList(uuid, updated_by): Promise<any> {
        return new Promise((resolve, reject) => {
            console.log('getChartList calling');
            const currentDate = new Date().toISOString();
            let paramData = [currentDate, uuid];
            console.log('parameter', paramData);
            DatabaseHelper.getDataByParameters("getActionForCancel", paramData).then(
                (success) => {               
                   resolve(success);
                },
                (error) => {
                    console.log('getChartList response Failed', error);
                    reject(error);
                });
        });
    }

}
