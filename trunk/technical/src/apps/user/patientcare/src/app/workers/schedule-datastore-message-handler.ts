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
    handleMessage(msg: ScheduleDatastoreModel): void {
        // const data = <ScheduleDatastoreModel>this.dataModel.data;
        const schedulardata = new Schedulardata();
        schedulardata.data = msg;
        schedulardata.conf = new SchedularConfigData()
        const parsedConf = <SchedularConfigData>JSON.parse(msg.conf);
        schedulardata.conf = parsedConf;
        let actiondata: ActionsData;
        try {
            switch (schedulardata.data.conf_type_code) {
                case ConfigCodeType.MEDICINE:
                    const medicineHelper = new MedicineHelper();
                    console.log('MedicineHelper created');
                    actiondata = <ActionsData>medicineHelper.createMedicineActions(schedulardata);
                    break;
                case ConfigCodeType.INTAKE:
                    console.log('intake invoked');
                    const intakehelper = new IntakeHelper();
                    actiondata = <ActionsData>intakehelper.createIntakeActions(schedulardata);

                    break;
                case ConfigCodeType.MONITOR:
                    console.log('monitor invoked');
                    const monitorhelper = new MonitorHelper()
                    actiondata = <ActionsData>monitorhelper.createMonitorActions(schedulardata);
                    console.log('actions created');
                    break;
                default:
                    break;
            }
            try {
                parsedConf.endDate = actiondata.enddate;
                msg.enddate = actiondata.enddate;
                msg.conf = JSON.stringify(parsedConf);
                actiondata.actions.forEach(element => {
                    const actionsdbdata = new ActionDataStoreModel();
                    Object.assign(actionsdbdata, element);
                    DatabaseHelper.DataStoreInsertUpdate(SYNC_STORE.ACTION, actionsdbdata.getModelValues());
                });
            } catch (e) {
                console.log('action inserting failed....', e.error);
            }
        }
        catch (e) {
            console.error('Action Helper', e);
        }
    }
}