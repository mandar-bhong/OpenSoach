import { IDatastoreMessageHandler } from "./idatastore-message-handler.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { Schedulardata, SchedularConfigData } from "../models/ui/chart-models.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ActionDataStoreModel, ActionsData } from "../models/db/action-datastore.js";
import { SYNC_STORE, ConfigCodeType } from "../app-constants.js";
import { MonitorHelper } from "../helpers/actions/monitor-helper.js";
import { IntakeHelper } from "../helpers/actions/intake-helper.js";
import { MedicineHelper } from "../helpers/actions/medicine-helper.js";

export class ScheduleDatastoreMessageHandler implements IDatastoreMessageHandler<ScheduleDatastoreModel>
{
    handleMessage(msg: ScheduleDatastoreModel): ActionDataStoreModel[] {
        // const data = <ScheduleDatastoreModel>this.dataModel.data;
        const schedulardata = new Schedulardata();
        schedulardata.data = msg;
        schedulardata.updated_by = msg.updated_by;
        schedulardata.conf = new SchedularConfigData()
        const parsedConf = <SchedularConfigData>JSON.parse(msg.conf);
        schedulardata.conf = parsedConf;
        let actiondata: ActionsData;
        let days: number;
        try {
            switch (schedulardata.data.conf_type_code) {
                case ConfigCodeType.MEDICINE:
                    const medicineHelper = new MedicineHelper();                   
                    actiondata = <ActionsData>medicineHelper.createMedicineActions(schedulardata);
                    break;
                case ConfigCodeType.INTAKE:                 
                    const intakehelper = new IntakeHelper();
                    actiondata = <ActionsData>intakehelper.createIntakeActions(schedulardata);
                    break;
                case ConfigCodeType.MONITOR:               
                    const monitorhelper = new MonitorHelper()
                    actiondata = <ActionsData>monitorhelper.createMonitorActions(schedulardata);                  
                    break;
                case ConfigCodeType.OUTPUT:
                    //   actiondata = <ActionsData>monitorhelper.createMonitorActions(schedulardata);
                    actiondata = new ActionsData();
                    let startDate = new Date(schedulardata.data.start_date);
                
                    let endDate = new Date();
                    days = parseInt(schedulardata.conf.duration.toString());                
                    endDate.setDate(startDate.getDate() + days);                   
                    actiondata.actions = [];
                    actiondata.enddate = new Date(endDate).toISOString();
                    break;
                default:
                    break;
            }
            try {                
                msg.end_date = actiondata.enddate;  
                return actiondata.actions;
            } catch (e) {
                console.log('action inserting failed....', e.error);
            }
        }
        catch (e) {
            console.error('Action Helper', e);
        }
    }
}