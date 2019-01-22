import { ActionItems, ProcessTime } from "~/app/models/ui/action-model";
import { Schedulardata } from "~/app/models/ui/chart-models";
import { ActionDBModel } from "~/app/models/ui/action-models";

export class ActionHelper {
    startdate: Date;
    startDateWithoutHours: Date;
    startdatetime: Date;
    actionItems: ActionItems[];
    numberofdays: number;
    schedulardata: Schedulardata;
    actionList: ActionDBModel[];
    enddate: Date;
    foodInstruction: number;
    constructor() {
        console.log('schedular initiated')
    }
    createDateEntries() {
        console.log('in createDateEntries');
        let processTime = ProcessTime;
        this.startdate = new Date(this.schedulardata.conf.startDate);
        this.startdatetime = new Date(this.schedulardata.conf.startDate);
        this.startDateWithoutHours = new Date(this.schedulardata.conf.startDate);
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
            //  dateaction.setHours(0, 0, 0, 0);
            console.log('date actions', dateaction);
            for (let j = 0; j < this.actionItems[i].dayAction.length; j++) {
                const dateval = new Date(dateaction);

                dateval.setMinutes(this.actionItems[i].dayAction[j].time);
                const actionList = new ActionDBModel();
                actionList.exec_time = new Date(dateval);
                actionList.admission_uuid = this.schedulardata.admission_uuid;
                actionList.schedule_uuid = this.schedulardata.uuid;
                actionList.conf_type_code = this.schedulardata.conf_type_code
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


}