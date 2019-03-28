import { ActionHelper } from "./action-helper.js";
import { Schedulardata } from "~/app/models/ui/chart-models.js";
import { Medicinefrequency, DayTimes, ActionItems, Monitorfrequency } from "~/app/models/ui/action-model.js";
import { ActionsData } from "~/app/models/db/action-datastore.js";

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
        if (this.schedulardata.conf.frequency == 0 || this.schedulardata.conf.frequency == 1) {
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
        const actios = new ActionsData();
        actios.actions = this.actionList;
        actios.enddate = this.getScheduleEnddate();
       // console.log('final actions')
        return actios;
    } else {
        console.log('in as required');
        const actios = new ActionsData();
        actios.actions = [];
        actios.enddate = null;
        return;
    }
    }// end of fucntions.

    // fucntion for geberatating actions for x times in day.
    frequencyActionasGenerations(strdate, MonitorSchedularData, i) {
        // if user has selectes  medicine after x time interval then following code blcok will executed
        console.log('frequencyActionasGenerations called ');
        if (MonitorSchedularData.conf.frequency == Monitorfrequency.AfterXTimeInterval) {
            console.log('createActionAfterXTimeinterval called ');
            this.createActionAfterXTimeinterval(strdate, MonitorSchedularData, i);
        } else {
            // if user selected specific times in a day then following code block will executed.
            this.generateActionsOnSpecificTime(strdate, MonitorSchedularData, i)
        }
    }// end of class

    // fucntion for creating actions after x timeinterval
    createActionAfterXTimeinterval(receivedDate, SchedularData, i) {
        let index = i;
        const receivedActionDate = new Date(receivedDate);
        const TimeInterval = SchedularData.conf.interval;
        let scheduleTimeOnStartDate = this.getStartTime(SchedularData.conf.startTime);
        let scheduleTime = this.getStartTime(SchedularData.conf.startTime);
        console.log('scheduled time', scheduleTime);
        const scheduleCreationTime = this.getMinutes();
        let position = 0;
        for (let x = 0; x < this.numberofTimes; x++) {
            if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
                if (scheduleTimeOnStartDate >= scheduleCreationTime) {
                    if (scheduleTime > DayTimes.dayEndTime) {
                        position++;
                        const nextDate = new Date(receivedActionDate);
                        nextDate.setDate(receivedActionDate.getDate() + 1);
                        scheduleTime -= DayTimes.dayEndTime;
                        const arraylen = this.actionItems.length - 1;
                        if (i >= arraylen) {
                            this.tempActionItems.push({ dateAction: nextDate, dayAction: [] });
                            this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                        } else {
                            // this.tempActionItems[i].dayAction.push({ time: xIntervalStartTime });
                            this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                        }
                        index++;
                    } else {
                        this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                    }
                }
            } else {
                if (scheduleTime > DayTimes.dayEndTime) {
                    position++;
                    const nextDate = new Date(receivedActionDate);
                    nextDate.setDate(receivedActionDate.getDate() + 1);
                    scheduleTime -= DayTimes.dayEndTime;
                    const arraylen = this.actionItems.length - 1;
                    if (i >= arraylen) {
                        this.tempActionItems.push({ dateAction: nextDate, dayAction: [] });
                        this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                    } else {
                        // this.tempActionItems[i].dayAction.push({ time: xIntervalStartTime });
                        this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                    }
                    index++;
                } else {
                    this.tempActionItems[index + position].dayAction.push({ time: scheduleTime });
                }
            }
            scheduleTime = Number(scheduleTime) + Number(TimeInterval);
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
        console.log('generateActionsOnSpecificTime executed');
        const receivedActionDate = new Date(receivedDate);
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
            const totalminutes = this.getMinutes();
            // checking schedule start date time period.
            for (let h = 0; h < MonitorSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime = this.getStartTime(MonitorSchedularData.conf.specificTimes[h]);
                // cheking schedule start time ellpased  of not 
                if (receivedSpecificTime >= totalminutes) {
                    this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
                }
            }
        } else {
            for (let h = 0; h < MonitorSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime = this.getStartTime(MonitorSchedularData.conf.specificTimes[h]);
                this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
            }
        }
    }// end of code block.
}// end of class. 