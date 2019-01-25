import { ActionHelper } from "./action-helper.js";
import { Schedulardata } from "~/app/models/ui/chart-models.js";
import { Medicinefrequency, DayTimes, ActionItems, Monitorfrequency, Intakefrequency } from "~/app/models/ui/action-model.js";
import { ActionsData } from "~/app/models/db/action-datastore";

export class IntakeHelper extends ActionHelper {
    // process variables 
    numberofTimes: number;
    arraylenght: number;
    tempActionItems: ActionItems[];
    //end
    constructor() {
        console.log('Intake schedular initiated')
        super();
    }
    createIntakeActions(IntakeSchedularData: Schedulardata) {
        this.tempActionItems = [];
        console.log('received data', IntakeSchedularData);
        this.schedulardata = IntakeSchedularData;
        this.numberofTimes = this.schedulardata.conf.numberofTimes;
        // creating date entries
        this.createDateEntries();
        // creating array without memory ref.
        const tempdata = this.actionItems.slice();
        this.tempActionItems = tempdata;
        // generating actions based on date entryes.
        for (let i = 0; i <= this.actionItems.length - 1; i++) {
            // creating actions for this Date
            this.frequencyActionasGenerations(this.actionItems[i].dateAction, IntakeSchedularData, i);
        }
        // if user selected after x time intervals.
        if (IntakeSchedularData.conf.frequency == Intakefrequency.AfterXTimeInterval) {
            this.actionItems = this.tempActionItems;
        }
        // creating final actions.
        this.generateDBActions();
        const actios = new ActionsData();
        actios.actions = this.actionList;
        actios.enddate = this.getScheduleEnddate();
        return actios;
    }// end of fucntions.

    // fucntion for geberatating actions for x times in day.
    frequencyActionasGenerations(strdate, IntakeSchedularData, i) {
        // if user has selectes  medicine after x time interval then following code blcok will executed
        console.log('frequencyActionasGenerations called ');
        if (IntakeSchedularData.conf.frequency == Intakefrequency.AfterXTimeInterval) {
            this.createActionAfterXTimeinterval(strdate, IntakeSchedularData, i);
        } else {
            // if user selected specific times in a day then following code block will executed.
            this.generateActionsOnSpecificTime(strdate, IntakeSchedularData, i)
        }
    }// end of class

    // fucntion for creating actions after x timeinterval
    createActionAfterXTimeinterval(receivedDate, SchedularData, i) {
        console.log(' createActionAfterXTimeinterval function called');
        let index = i;
        const receivedActionDate = new Date(receivedDate);
        const TimeInterval = SchedularData.conf.intervalHrs;
        let scheduleTime = this.getStartTime(SchedularData.conf.startTime);
        console.log('scheduled time', scheduleTime);
        const scheduleCreationTime = this.getMinutes();
        for (let x = 0; x < this.numberofTimes; x++) {
            if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
                if (scheduleTime >= scheduleCreationTime) {
                    if (scheduleTime > DayTimes.dayEndTime) {
                        const nextDate = new Date(receivedActionDate);
                        nextDate.setDate(receivedActionDate.getDate() + 1);
                        scheduleTime -= DayTimes.dayEndTime;
                        const arraylen = this.actionItems.length - 1;
                        if (i >= arraylen) {
                            this.tempActionItems.push({ dateAction: nextDate, dayAction: [] });
                            this.tempActionItems[index + 1].dayAction.push({ time: scheduleTime });
                        } else {
                            // this.tempActionItems[i].dayAction.push({ time: xIntervalStartTime });
                            this.tempActionItems[index + 1].dayAction.push({ time: scheduleTime });
                        }
                        index++;
                    } else {
                        this.tempActionItems[index].dayAction.push({ time: scheduleTime });
                    }
                }
            } else {
                if (scheduleTime > DayTimes.dayEndTime) {
                    const nextDate = new Date(receivedActionDate);
                    nextDate.setDate(receivedActionDate.getDate() + 1);
                    scheduleTime -= DayTimes.dayEndTime;
                    const arraylen = this.actionItems.length - 1;
                    if (i >= arraylen) {
                        this.tempActionItems.push({ dateAction: nextDate, dayAction: [] });
                        this.tempActionItems[index + 1].dayAction.push({ time: scheduleTime });
                    } else {
                        // this.tempActionItems[i].dayAction.push({ time: xIntervalStartTime });
                        this.tempActionItems[index + 1].dayAction.push({ time: scheduleTime });
                    }
                    index++;
                } else {
                    this.tempActionItems[index].dayAction.push({ time: scheduleTime });
                }
            }
            scheduleTime += TimeInterval;
        }
        // const receivedActionDate = new Date(receivedDate);
        // const TimeInterval = IntakeSchedularData.conf.intervalHrs;
        // // checking schedule start and its time periods
        // let treatmentStartTime = Math.floor(IntakeSchedularData.conf.startTime * 60);
        // let xIntervalStartTime = treatmentStartTime;
        // if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
        //     // getting totoal minutes from start date
        //     const totalminutes = this.getMinutes();
        //     if (totalminutes > xIntervalStartTime) {
        //         this.generateXTimesActions(totalminutes, receivedActionDate, TimeInterval, i)
        //     } else {
        //         this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
        //     }
        // } else {
        //     this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
        // }
    } // end of code block.
    // code block for generate actions after x time intervalF
    generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i) {
        let index = i;
        for (let x = 0; x < this.numberofTimes; x++) {
            if (xIntervalStartTime > DayTimes.dayEndTime) {
                const nextDate = new Date(receivedActionDate);
                nextDate.setDate(receivedActionDate.getDate() + 1);
                xIntervalStartTime -= DayTimes.dayEndTime;
                // checking item haveing last index in array.if it is then add one extra element at last possition.
                const arraylen = this.actionItems.length - 1;
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
            xIntervalStartTime += TimeInterval;
        }
    }// end of class

    // code block for generate actions on specific times
    generateActionsOnSpecificTime(receivedDate, IntakeSchedularData, i) {
        const receivedActionDate = new Date(receivedDate);
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
            const totalminutes = this.getMinutes();
            for (let h = 0; h < IntakeSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime = this.getStartTime(IntakeSchedularData.conf.specificTimes[h]);
                if (receivedSpecificTime >= totalminutes) {
                    this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
                }
            }
        } else {
            for (let h = 0; h < IntakeSchedularData.conf.specificTimes.length; h++) {
                const receivedSpecificTime =this.getStartTime(IntakeSchedularData.conf.specificTimes[h]);
                this.actionItems[i].dayAction.push({ time: receivedSpecificTime });
            }
        }
    }// end of code block.
}// end of class. 