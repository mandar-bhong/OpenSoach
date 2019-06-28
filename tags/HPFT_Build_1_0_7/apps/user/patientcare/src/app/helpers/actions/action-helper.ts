import { ActionItems, ProcessTime, ActionList } from "~/app/models/ui/action-model.js";
import { Schedulardata } from "~/app/models/ui/chart-models.js";
import { ActionDBModel } from "~/app/models/ui/action-models.js";
import { ActionDataStoreModel } from "~/app/models/db/action-datastore.js";
import { PlatformHelper } from "../platform-helper.js";
import { ActionStatus } from "~/app/app-constants.js";
import { ActionsData } from "~/app/models/db/action-datastore.js";

export class ActionHelper {
    startdate: Date;
    startDateWithoutHours: Date;
    startdatetime: Date;
    actionItems: ActionItems[];
    numberofdays: number;
    schedulardata: Schedulardata;
    actionList: ActionDataStoreModel[];
    enddate: Date;
    foodInstruction: number;
    constructor() {
        console.log('schedular initiated')
    }
    createDateEntries() {
        console.log('in createDateEntries');
        let processTime = ProcessTime;
        this.startdate = new Date(this.schedulardata.data.start_date);
        this.startdatetime = new Date(this.schedulardata.data.start_date);
        this.startDateWithoutHours = new Date(this.schedulardata.data.start_date);
        this.startDateWithoutHours.setHours(0, 0, 0, 0);
        this.numberofdays = this.schedulardata.conf.duration;
        this.enddate = new Date();
        this.enddate.setDate(this.startdate.getDate() + this.numberofdays - 1);

        // for adding  dates in date array
        this.actionItems = [];
        const dt = this.startdate;
        dt.setHours(0, 0, 0, 0);
        for (let i = 0; i < this.numberofdays; i++) {
            let strdate = new Date(dt);
            strdate.setDate(dt.getDate() + i);
            this.actionItems.push({ dateAction: strdate, dayAction: [] });
        }
        console.log(' this.actionItems');
        console.log(this.actionItems);
    }
    // fucntion for determining how much actions are created.
    actionsLength() {
        let actionlen = 0;
        for (let i = 0; i < this.actionItems.length; i++) {
            actionlen += this.actionItems[i].dayAction.length
        }
        return actionlen;
    }
    generateDBActions() {
        this.actionList = [];
        for (let i = 0; i < this.actionItems.length; i++) {
            const dateaction = new Date(this.actionItems[i].dateAction);

            for (let j = 0; j < this.actionItems[i].dayAction.length; j++) {
                const dateval = new Date(dateaction);
                dateval.setMinutes(this.actionItems[i].dayAction[j].time);
                const actionList = new ActionDataStoreModel();
                actionList.scheduled_time = new Date(dateval).toISOString();
                actionList.admission_uuid = this.schedulardata.data.admission_uuid;
                actionList.schedule_uuid = this.schedulardata.data.uuid;
                actionList.conf_type_code = this.schedulardata.data.conf_type_code;
                actionList.is_deleted = ActionStatus.ACTION_ACTIVE;
                actionList.sync_pending = 1;
                actionList.client_updated_at=new Date().toISOString();              
                actionList.updated_by = this.schedulardata.updated_by;
                actionList.uuid = PlatformHelper.API.getRandomUUID();
                // const tempid = Math.random();
                // actionList.uuid = tempid.toString();
                this.actionList.push(actionList);
            }
        }
    }// end of code block  

    // fucntion for getting minutes form date
    getMinutes(): number {
        const hours = this.startdatetime.getHours();
        const minutes = this.startdatetime.getMinutes();
        const hourstominutes = hours * 60;
        const totalminutes = hourstominutes + minutes
        return totalminutes;
    }// end of code block

    // getStartTime(startTime): number {
    //     let t = startTime.toString();
    //     const time = t.split('.');
    //     const minutes = 60 * Number(time[0]);
    //     const totalminutes = minutes + Number(time[1]);
    //     console.log('getStartTime return', totalminutes);
    //     return totalminutes;
    // }
    getScheduleEnddate() {
        if (this.actionList.length > 0) {
            const enddate = this.actionList[this.actionList.length - 1].scheduled_time;
            return enddate;
        }
    }
     // fucntion for generating sticky actions
     generateStickyAction(schedulardata : Schedulardata ) {       
                let actionList:ActionDataStoreModel[] = [];
                
                const actionItem = new ActionDataStoreModel();
                actionItem.scheduled_time =null;       
                actionItem.admission_uuid = schedulardata.data.admission_uuid;
                actionItem.schedule_uuid = schedulardata.data.uuid;
                actionItem.conf_type_code = schedulardata.data.conf_type_code;
                actionItem.is_deleted = ActionStatus.ACTION_ACTIVE;
                actionItem.sync_pending = 1;
                actionItem.client_updated_at=new Date().toISOString();              
                actionItem.updated_by = schedulardata.updated_by;
                actionItem.uuid = PlatformHelper.API.getRandomUUID();  
                actionList.push(actionItem);            
                return actionList;           
    }

    generateEndDate(schedulardata : Schedulardata) {
        let endDate = new Date();
        let startDate = new Date(schedulardata.data.start_date);
        let days = parseInt(schedulardata.conf.duration.toString());
        endDate.setDate(startDate.getDate() + days);
        const date = new Date(endDate).toISOString();
        return date;
    }

}