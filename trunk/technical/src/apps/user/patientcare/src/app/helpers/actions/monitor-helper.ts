import { ActionHelper } from "./action-helper";
import { Schedulardata } from "~/app/models/ui/chart-models";
import { Medicinefrequency, DayTimes, ActionItems, Monitorfrequency } from "~/app/models/ui/action-model";

export class MonitorHelper extends ActionHelper {
    // process variables
    numberofTimes: number;
    arraylenght: number;
    tempActionItems: ActionItems[];
    // end 
    constructor() {
        console.log('monitor schedular initiated')
        super();
    }
    //  code block for creating monitor actions.
    createMonitorActions(MonitorSchedularData: Schedulardata) {
        this.tempActionItems = [];
        console.log('received data', MonitorSchedularData);
        this.schedulardata = MonitorSchedularData;
        this.numberofTimes = this.schedulardata.conf.numberofTimes;
        // creating date entries  using base class fucntion.
        this.createDateEntries();
        // creating array by value (without memory ref)
        const tempdata = this.actionItems.slice();
        this.tempActionItems = tempdata;
        // generating actions based on date entryes.
        for (let i = 0; i <= this.actionItems.length - 1; i++) {
            // creating actions for this Date
            this.frequencyActionasGenerations(this.actionItems[i].dateAction, MonitorSchedularData, i);
        }
        // if user selected after x time interval then
        if (MonitorSchedularData.conf.frequency == Monitorfrequency.AfterXTimeInterval) {
            this.actionItems = this.tempActionItems;
        }
        // creating final  actions.
        this.generateDBActions();
        console.log('action list');
        console.log(this.actionList);
    }// end of fucntions.

    // fucntion for geberatating actions for x times in day.
    frequencyActionasGenerations(strdate, MonitorSchedularData, i) {
        // if user has selectes  medicine after x time interval then following code blcok will executed
        console.log('frequencyActionasGenerations called ');
        if (MonitorSchedularData.conf.frequency == Monitorfrequency.AfterXTimeInterval) {
            this.createActionAfterXTimeinterval(strdate, MonitorSchedularData, i);
        } else {
            // if user selected specific times in a day then following code block will executed.
            this.generateActionsOnSpecificTime(strdate, MonitorSchedularData, i)
        }
    }// end of class

    // fucntion for creating actions after x timeinterval
    createActionAfterXTimeinterval(receivedDate, MonitorSchedularData, i) {
        console.log(' createActionAfterXTimeinterval function called');
        const receivedActionDate = new Date(receivedDate);
        // const TimeInterval = Math.floor(MedicineSchedularData.conf.intervalHrs * 60);
        const TimeInterval = MonitorSchedularData.conf.intervalHrs;
        let treatmentStartTime = Math.floor(MonitorSchedularData.conf.startTime * 60);
        let xIntervalStartTime = treatmentStartTime;
        //cheking schedule start date periods
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
            // geting minutes from date- base class function
            const totalminutes = this.getMinutes();
            // cheking schedule create time is greater than schedule start time if
            if (totalminutes > xIntervalStartTime) {
                //if greater then pass schedule create time as start time.
                this.generateXTimesActions(totalminutes, receivedActionDate, TimeInterval, i)
            } else {
                this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
            }
        } else {
            this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
        }
    } // end of code block.
    //code block for generatring actions after x time interval
    generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i) {
        let index = i;
        for (let x = 0; x < this.numberofTimes; x++) {
            // if time exceed 24 hours
            if (xIntervalStartTime > DayTimes.dayEndTime) {
                const nextDate = new Date(receivedActionDate);
                nextDate.setDate(receivedActionDate.getDate() + 1);
                // if time exceeds 24 time then reduce 24*60 time from intervaltime
                xIntervalStartTime -= DayTimes.dayEndTime;
                const arraylen = this.actionItems.length - 1;
                // is item have last index in array if it is then add new element in array as next date.
                if (i >= arraylen) {
                    this.tempActionItems.push({ dateAction: nextDate, dayAction: [] });
                    this.tempActionItems[index + 1].dayAction.push({ time: xIntervalStartTime });
                } else {
                    // this.tempActionItems[i].dayAction.push({ time: xIntervalStartTime });
                    this.tempActionItems[index + 1].dayAction.push({ time: xIntervalStartTime });
                }
                index++;
            } else {
                this.tempActionItems[index].dayAction.push({ time: xIntervalStartTime });
            }
            // increatement in schedule start time
            xIntervalStartTime += TimeInterval;
        }
    }// end of class

    // code block for generate actions on specific times
    generateActionsOnSpecificTime(receivedDate, MonitorSchedularData, i) {
        const receivedActionDate = new Date(receivedDate);
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
            const totalminutes = this.getMinutes();
            // checking schedule start date time period.
            for (let h = 0; h < MonitorSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime = Math.floor(MonitorSchedularData.conf.specificTimes[h] * 60);
                // cheking schedule start time ellpased  of not 
                if (receivedSpecificTime >= totalminutes) {
                    this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
                }
            }
        } else {
            for (let h = 0; h < MonitorSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime = Math.floor(MonitorSchedularData.conf.specificTimes[h] * 60);
                this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
            }
        }
    }// end of code block.
}// end of class. 