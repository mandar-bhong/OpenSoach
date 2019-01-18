
import { action } from "tns-core-modules/ui/dialogs/dialogs";
import { ActionItems, ActionList, SchedularData, ProcessTime, TimeConstants, dayTime, AfterMealTime, BeforeMealTime, BeforeMealTimeInMinutes, AfterMealTimeInMinutes, HospitalTime, Medicinefrequency } from "~/app/models/ui/action-model";
import { ActionHelper } from "./action-helper";
import { ActionDBModel } from "~/app/models/ui/action-models";
import { Schedulardata } from "~/app/models/ui/chart-models";

export class MedicineHelper extends ActionHelper {
    startdate: Date;
    startdatetime: Date;
    enddate: Date;
    tempdate: Date;
    actionItems: ActionItems[];
    numberofdays: number;
    trueCount: number;
    schedulardata: Schedulardata;
    foodInstruction: number;
    actionList: ActionDBModel[];
    startDateWithoutHours: Date;
    constructor() {
        console.log('medicine schedular initiated')
        super();
    }

    medicineActions(MedicineSchedularData: Schedulardata) {
        console.log('received data');
        console.log(MedicineSchedularData);
        this.schedulardata = MedicineSchedularData;
        let processTime = ProcessTime;
        this.startdate = new Date(MedicineSchedularData.conf.startDate);
        this.startdatetime = new Date(MedicineSchedularData.conf.startDate);
        this.startDateWithoutHours = new Date(MedicineSchedularData.conf.startDate);
        this.startDateWithoutHours.setHours(0, 0, 0, 0);
        this.numberofdays = MedicineSchedularData.conf.duration;
        this.enddate = new Date();
        this.enddate.setDate(this.startdate.getDate() + this.numberofdays - 1);

        // for adding  dates in date array
        this.actionItems = [];
        const dt = this.startdate;
        for (let i = 0; i < this.numberofdays; i++) {
            let strdate = new Date(dt);
            strdate.setDate(dt.getDate() + i);
            this.actionItems.push({ dateAction: strdate, dayAction: [] });
            console.log('str date', strdate);
            // creating actions for this Date
            this.frequencyActionasGenerations(strdate, MedicineSchedularData, i);

        }
        console.log(this.actionItems);
        // this code block will executed  when user selected medicine x times in day
        if (MedicineSchedularData.conf.frequency == Medicinefrequency.xTimesInDay) {
            // check wherater we have add extra day or not 
            const totalActions = this.CheckFrequencyCount() * this.numberofdays;
            const createdActions = this.actionsLength();
            if (createdActions < totalActions) {
                this.isNextDateRequired()
            }
            const updateActionsLen = this.actionsLength();
            if (updateActionsLen > totalActions) {
                this.removeActions(updateActionsLen, totalActions);
                console.log('extra actions created');
            }
        }

        this.actionList = [];
        for (let i = 0; i < this.actionItems.length; i++) {
            const dateaction = new Date(this.actionItems[i].dateAction);
            dateaction.setHours(0, 0, 0, 0);
            console.log('date actions', dateaction);
            for (let j = 0; j < this.actionItems[i].dayAction.length; j++) {
                const dateval = new Date(dateaction);

                dateval.setMinutes(this.actionItems[i].dayAction[j].time);
                const actionList = new ActionDBModel();
                actionList.exec_time = new Date(dateval);
                actionList.admission_uuid = MedicineSchedularData.admission_uuid;
                actionList.schedule_uuid = MedicineSchedularData.uuid;
                actionList.conf_type_code = MedicineSchedularData.conf_type_code
                this.actionList.push(actionList);
            }
        }
        console.log('action list');
        console.log(this.actionList);
    } // end of 



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
    createDateActions(date: Date, frq, foodInst) {
        const receiveddate = new Date(date);
        let timeConstants: TimeConstants;
        timeConstants = this.beforeAfter(foodInst);
        let definedTime: TimeConstants;
        definedTime = this.beforeAfterProcessTime(foodInst)
        const startDate = new Date(this.startdatetime);
        startDate.setHours(0, 0, 0, 0);
        receiveddate.setHours(0, 0, 0, 0);
        // if received date from an array &  start date are same then execute following code. 
        if (receiveddate.getTime() == startDate.getTime()) {
            const hours = this.startdatetime.getHours();
            const minutes = this.startdatetime.getMinutes();
            const hourstominutes = hours * 60;
            const totalminutes = hourstominutes + minutes
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
        const newdate = new Date();
        newdate.setDate(lastrec.getDate() + 1);
        this.actionItems.push({ dateAction: newdate, dayAction: [] });
        // if one extra day is required to complete dosage then execute following code.
        this.frequencyActionasGenerations(newdate, this.schedulardata, this.actionItems.length - 1);

    }

    // fucntion for determining how much actions are created.
    actionsLength() {
        let actionlen = 0;
        for (let i = 0; i < this.actionItems.length; i++) {
            actionlen += this.actionItems[i].dayAction.length
        }
        return actionlen;
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
                const dayactionTime = this.createDateActions(strdate, dayTime.Morning, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
            if (MedicineSchedularData.conf.aftrnFreqInfo.freqAftrn) {
                // if user selected afternoon time then created actions of afternoon.
                const dayactionTime = this.createDateActions(strdate, dayTime.Afternoon, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
            if (MedicineSchedularData.conf.nightFreqInfo.freqNight) {
                // if user selected night time then created actions of night.
                const dayactionTime = this.createDateActions(strdate, dayTime.Night, MedicineSchedularData.conf.foodInst)
                if (dayactionTime && dayactionTime != null) {
                    this.actionItems[i].dayAction.push({ time: dayactionTime });
                }
            }
        }
    }

    // fucntion for creating actions after x timeinterval
    createActionAfterXTimeinterval(receivedDate, MedicineSchedularData, i) {
        const receivedActionDate = new Date(receivedDate);
        receivedActionDate.setHours(0, 0, 0, 0);
        const TimeInterval = Math.floor(MedicineSchedularData.conf.intervalHrs * 60);
        let treatmentStartTime = Math.floor(MedicineSchedularData.conf.startTime * 60);
        // cheking received date is same with  startday.
        if (receivedActionDate.getTime() == this.startDateWithoutHours.getTime()) {
             // adding actions from stat time of start date until hospital closetime
            while (treatmentStartTime < HospitalTime.hospitalEndTime) {
                this.actionItems[i].dayAction.push({ time: treatmentStartTime });
                treatmentStartTime += TimeInterval;
            }
        } else {
             // adding actions from hospital start time  until hospital closetime
            let hospitalStartTime = HospitalTime.hospitalStartTime;
            while (hospitalStartTime < HospitalTime.hospitalEndTime) {
                this.actionItems[i].dayAction.push({ time: hospitalStartTime });
                hospitalStartTime += TimeInterval;
            }
        }
    }

   // function for removing deleting unwanted actions.
    removeActions(updateActionsLen, totalActions) {
        const actionDiff = updateActionsLen - totalActions;
        let i = 0;
       while (i < actionDiff) {
            this.actionItems[this.actionItems.length - 1].dayAction.pop();
            i++;
        }
        console.log('this.actionItems', this.actionItems);
    }
}// end of class

