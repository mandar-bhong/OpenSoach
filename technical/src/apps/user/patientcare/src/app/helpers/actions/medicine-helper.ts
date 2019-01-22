
import { ProcessTime, TimeConstants, dayTime, AfterMealTime, BeforeMealTime, BeforeMealTimeInMinutes, AfterMealTimeInMinutes, Medicinefrequency, AfterXtimeIntervl, DayTimes, ActionItems } from "~/app/models/ui/action-model";
import { ActionHelper } from "./action-helper";
import { Schedulardata } from "~/app/models/ui/chart-models";

export class MedicineHelper extends ActionHelper {
    // process variables
    tempdate: Date;
    trueCount: number;
    numberofTimes: number;
    afterXtimeIntervl: AfterXtimeIntervl[]
    arraylenght: number;
    tempActionItems: ActionItems[];
    // end 
    constructor() {
        console.log('medicine schedular initiated')
        super(); // calling base class
    }
    // code block for creating actions.
    createMedicineActions(MedicineSchedularData: Schedulardata) {
        this.tempActionItems = [];
        console.log('received data', MedicineSchedularData);
        this.schedulardata = MedicineSchedularData;
        this.numberofTimes = this.schedulardata.conf.numberofTimes;
        this.createDateEntries(); // calling base class function for creating date entries.
        // creating new array without memory ref.
        const tempdata = this.actionItems.slice();
        this.tempActionItems = tempdata;
        // ittrating date array for creating action entries
        for (let i = 0; i <= this.actionItems.length - 1; i++) {
            // creating actions for this Date
            this.frequencyActionasGenerations(this.actionItems[i].dateAction, MedicineSchedularData, i);
        }
        // this code block will executed  when user selected medicine x times in day
        if (MedicineSchedularData.conf.frequency == Medicinefrequency.xTimesInDay) {
            // check wherater we have add extra day or not 
            const totalActions = this.CheckFrequencyCount() * this.numberofdays;
            const createdActions = this.actionsLength();
            if (createdActions < totalActions) {
                this.isNextDateRequired(); // checking if medicine dosage remening.if yes then adding one extra daty in array.
            }
            const updateActionsLen = this.actionsLength();
        // removing unwanted day entries
            if (updateActionsLen > totalActions) {
                this.removeActions(updateActionsLen, totalActions);
                console.log('extra actions created');
            }
        } else {
            this.actionItems = this.tempActionItems;
            console.log('processed action list', this.actionItems);
        }
        // calling base function for  create final  action entries.
        this.generateDBActions();
        console.log('action list');
        console.log(this.actionList);
    } // end of code block
// code  block for checking frequency count. 
    CheckFrequencyCount() {
        let trueCount = 0;
        if (this.schedulardata.conf.mornFreqInfo.freqMorn) {
            trueCount++;
        }
        if (this.schedulardata.conf.aftrnFreqInfo.freqAftrn) {
            trueCount++;
        }
        if (this.schedulardata.conf.nightFreqInfo.freqNight) {
            trueCount++;
        }
        return trueCount;
    }
    // function for create date action entries.
    createDateActionsEntries(date: Date, frq, foodInst) {
        const receiveddate = new Date(date);
        let timeConstants: TimeConstants;
        // getiing time constants based on food  instructions in string format.
        timeConstants = this.beforeAfter(foodInst);
        let definedTime: TimeConstants;
         // getiing time constants based on food  instructions in minutes
        definedTime = this.beforeAfterProcessTime(foodInst)
        const startDate = new Date(this.startdatetime);
        startDate.setHours(0, 0, 0, 0);
        //  receiveddate.setHours(0, 0, 0, 0);
        // if received date from an array &  start date are same then execute following code. 
        if (receiveddate.getTime() == startDate.getTime()) {
            const totalminutes = this.getMinutes();
            console.log('total start time in minutes', totalminutes);
            // checking wheater current time elapse scheduled start time.
            if (frq == dayTime.Morning) {
                if (totalminutes <= definedTime.morningTime) {
                    return definedTime.morningTime;
                }
            }
            if (frq == dayTime.Afternoon) {
                if (totalminutes <= definedTime.afternoonTime) {
                    return definedTime.afternoonTime;
                }
            }
            if (frq == dayTime.Night) {
                if (totalminutes <= definedTime.nightTime) {
                    return definedTime.nightTime;
                }
            }
        } else {
            // else then return respective  timings 
            if (frq == dayTime.Morning) {
                return definedTime.morningTime;
            }
            if (frq == dayTime.Afternoon) {
                return definedTime.afternoonTime;
            }
            if (frq == dayTime.Night) {
                return definedTime.nightTime;
            }
        }
    }
    // this is unused code but if wants time in string format  then this code block will executed
    beforeAfter(foodInst) {
        // function will  return respective time values based on food instruction.before or after.
        let timeConstants = new TimeConstants();
        if (foodInst == ProcessTime.foodInstBeforeMeal) {
            // return  after meal time  in string format 
            timeConstants.afternoonTime = AfterMealTime.AfternoonAfterMeal;
            timeConstants.morningTime = AfterMealTime.MorningAfterMeal;
            timeConstants.nightTime = AfterMealTime.NightAfteremeal;
        } else {
            // return before meal time in string format 
            timeConstants.afternoonTime = BeforeMealTime.AfternoonbeBeforeMeal;
            timeConstants.morningTime = BeforeMealTime.MorningBeforeMeal;
            timeConstants.nightTime = BeforeMealTime.NightBeforemeal;
        }
        return timeConstants;
    }
    // function for returning time in minutes based on food instructions.
    beforeAfterProcessTime(foodInst) {
        let timeConstants = new TimeConstants();
        if (foodInst == ProcessTime.foodInstBeforeMeal) {
            // return  before meal time  in minutes
            timeConstants.afternoonTime = BeforeMealTimeInMinutes.AfternoonbeBeforeMeal
            timeConstants.morningTime = BeforeMealTimeInMinutes.MorningBeforeMeal;
            timeConstants.nightTime = BeforeMealTimeInMinutes.NightBeforemeal;
        } else {
            // return  after meal time  in string minutes
            timeConstants.afternoonTime = AfterMealTimeInMinutes.AfternoonAfterMeal
            timeConstants.morningTime = AfterMealTimeInMinutes.MorningAfterMeal;
            timeConstants.nightTime = AfterMealTimeInMinutes.NightAfteremeal;
        }
        return timeConstants;
    }
    // fucntion for decide is next is required for complete dosage completions.
    isNextDateRequired() {
        const lastrec = new Date(this.actionItems[this.actionItems.length - 1].dateAction);
        const newdate = new Date(lastrec);
        newdate.setDate(lastrec.getDate() + 1);
        this.actionItems.push({ dateAction: newdate, dayAction: [] });
        // if one extra day is required to complete dosage then execute following code.
        this.frequencyActionasGenerations(newdate, this.schedulardata, this.actionItems.length - 1);
    }


    // fucntion for geberatating actions for x times in day.
    frequencyActionasGenerations(strdate, MedicineSchedularData, i) {
        // if user has selectes  medicine after x time interval then following code blcok will executed
        if (MedicineSchedularData.conf.frequency == Medicinefrequency.AfterXTimeInterval) {
            this.createActionAfterXTimeinterval(strdate, MedicineSchedularData, i);
        } else {
            // if user selcted x times in day then following code block.
            if (MedicineSchedularData.conf.mornFreqInfo.freqMorn) {
                // if user selected moringi time then created actions of moring.
                const dayactionTime = this.createDateActionsEntries(strdate, dayTime.Morning, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
            if (MedicineSchedularData.conf.aftrnFreqInfo.freqAftrn) {
                // if user selected afternoon time then created actions of afternoon.
                const dayactionTime = this.createDateActionsEntries(strdate, dayTime.Afternoon, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
            if (MedicineSchedularData.conf.nightFreqInfo.freqNight) {
                // if user selected night time then created actions of night.
                const dayactionTime = this.createDateActionsEntries(strdate, dayTime.Night, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
        }
    }

    // fucntion for creating actions after x timeinterval
    createActionAfterXTimeinterval(receivedDate, MonitorSchedularData, i) {
        console.log(' createActionAfterXTimeinterval function called');
        const receivedActionDate = new Date(receivedDate);
        // const TimeInterval = Math.floor(MedicineSchedularData.conf.intervalHrs * 60);
        const TimeInterval = MonitorSchedularData.conf.intervalHrs;
        let treatmentStartTime = Math.floor(MonitorSchedularData.conf.startTime * 60);
        let xIntervalStartTime = treatmentStartTime;
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
            // getting minutes in from schedule start date.
            const totalminutes = this.getMinutes();
            if (totalminutes > xIntervalStartTime) {
                this.generateXTimesActions(totalminutes, receivedActionDate, TimeInterval, i)
            } else {
                this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
            }
        } else {
            this.generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i)
        }
    } // end of code block.

    // function for removing  unwanted actions.
    removeActions(updateActionsLen, totalActions) {
        const actionDiff = updateActionsLen - totalActions;
        let i = 0;
        while (i < actionDiff) {
            this.actionItems[this.actionItems.length - 1].dayAction.pop();
            i++;
        }
        console.log('this.actionItems', this.actionItems);
    }
// code block will generate actions based on x time interval
    generateXTimesActions(xIntervalStartTime, receivedActionDate, TimeInterval, i) {
        let index = i;
        for (let x = 0; x < this.numberofTimes; x++) {
            if (xIntervalStartTime > DayTimes.dayEndTime) {
                const nextDate = new Date(receivedActionDate);
                nextDate.setDate(receivedActionDate.getDate() + 1);
                xIntervalStartTime -= DayTimes.dayEndTime;
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
    }
}// end of class

