
import { action } from "tns-core-modules/ui/dialogs/dialogs";
import { ActionItems, ActionList, SchedularData, ProcessTime, TimeConstants, dayTime, AfterMealTime, BeforeMealTime, BeforeMealTimeInMinutes, AfterMealTimeInMinutes } from "~/app/models/ui/action-model";
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
        this.numberofdays = MedicineSchedularData.conf.duration;
        this.enddate = new Date();
        this.enddate.setDate(this.startdate.getDate() + this.numberofdays - 1);

        // for adding  create date entry
        this.actionItems = [];
        const dt = this.startdate
        for (let i = 0; i < this.numberofdays; i++) {
            let strdate = new Date(dt);
            strdate.setDate(dt.getDate() + i);
            this.actionItems.push({ dateAction: strdate, dayAction: [] });
            console.log('str date', strdate);
            this.frequencyActionasGenerations(strdate, MedicineSchedularData, i);

        }
        console.log(this.actionItems);

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
        // const xx = new Date(this.startdate);
        // xx.setHours(0, 0, 0, 0);
        // console.log('xx', xx);
        // xx.setMinutes(111);
        // console.log('xx after minutes', xx);
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
        const todaysdate = new Date()
        todaysdate.setHours(0, 0, 0, 0);
        receiveddate.setHours(0, 0, 0, 0);
        console.log('received date', receiveddate);
        if (receiveddate.getTime() == todaysdate.getTime()) {
            const hours = this.startdatetime.getHours();
            const minutes = this.startdatetime.getMinutes();
            const hourstominutes = hours * 60;
            const totalminutes = hourstominutes + minutes

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
    beforeAfter(foodInst) {
        let timeConstants = new TimeConstants();
        if (foodInst == ProcessTime.foodInstBeforeMeal) {
            timeConstants.afternoonTime = AfterMealTime.AfternoonAfterMeal;
            timeConstants.morningTime = AfterMealTime.MorningAfterMeal;
            timeConstants.nightTime = AfterMealTime.NightAfteremeal;
        } else {
            timeConstants.afternoonTime = BeforeMealTime.AfternoonbeBeforeMeal;
            timeConstants.morningTime = BeforeMealTime.MorningBeforeMeal;
            timeConstants.nightTime = BeforeMealTime.NightBeforemeal;
        }
        return timeConstants;
    }
    beforeAfterProcessTime(foodInst) {
        let timeConstants = new TimeConstants();
        if (foodInst == ProcessTime.foodInstBeforeMeal) {
            timeConstants.afternoonTime = BeforeMealTimeInMinutes.AfternoonbeBeforeMeal
            timeConstants.morningTime = BeforeMealTimeInMinutes.MorningBeforeMeal;
            timeConstants.nightTime = BeforeMealTimeInMinutes.NightBeforemeal;
        } else {
            timeConstants.afternoonTime = AfterMealTimeInMinutes.AfternoonAfterMeal
            timeConstants.morningTime = AfterMealTimeInMinutes.MorningAfterMeal;
            timeConstants.nightTime = AfterMealTimeInMinutes.NightAfteremeal;
        }
        return timeConstants;
    }
    isNextDateRequired() {
        const lastrec = new Date(this.actionItems[this.actionItems.length - 1].dateAction);
        const newdate = new Date();
        newdate.setDate(lastrec.getDate() + 1);
        this.actionItems.push({ dateAction: newdate, dayAction: [] });
        this.frequencyActionasGenerations(newdate, this.schedulardata, this.actionItems.length - 1);

    }
    actionsLength() {
        let actionlen = 0;
        for (let i = 0; i < this.actionItems.length; i++) {
            actionlen += this.actionItems[i].dayAction.length
        }
        return actionlen;
    }
    frequencyActionasGenerations(strdate, MedicineSchedularData, i) {
        if (MedicineSchedularData.conf.mornFreqInfo.freqMorn) {
            const dayactionTime = this.createDateActions(strdate, dayTime.Morning, MedicineSchedularData.conf.foodInst)
            if (dayactionTime && dayactionTime != null) {
                this.actionItems[i].dayAction.push({ time: dayactionTime });
            }
        }
        if (MedicineSchedularData.conf.aftrnFreqInfo.freqAftrn) {
            const dayactionTime = this.createDateActions(strdate, dayTime.Afternoon, MedicineSchedularData.conf.foodInst)
            if (dayactionTime && dayactionTime != null) {
                this.actionItems[i].dayAction.push({ time: dayactionTime });
            }
        }
        if (MedicineSchedularData.conf.nightFreqInfo.freqNight) {
            const dayactionTime = this.createDateActions(strdate, dayTime.Night, MedicineSchedularData.conf.foodInst)
            if (dayactionTime && dayactionTime != null) {
                this.actionItems[i].dayAction.push({ time: dayactionTime });
            }
        }
    }

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

