
import { ActionsData } from "~/app/models/db/action-datastore.js";
import { ActionItems, AfterXtimeIntervl } from "~/app/models/ui/action-model.js";
import { Schedulardata } from "~/app/models/ui/chart-models.js";
import { ActionHelper } from "./action-helper.js";

export class OutputHelper extends ActionHelper {
    // process variables
    tempdate: Date;
    trueCount: number;
    numberofTimes: number;
    afterXtimeIntervl: AfterXtimeIntervl[]
    arraylenght: number;
    tempActionItems: ActionItems[];
    // end 
    constructor() {
        console.log('output schedular initiated')
        super(); // calling base class
    }
    // code block for creating actions.
    createOutputActions(MedicineSchedularData: Schedulardata) {
        this.tempActionItems = [];
        console.log('Received Data', MedicineSchedularData);
        this.schedulardata = MedicineSchedularData;
        this.numberofTimes = this.schedulardata.conf.numberofTimes;
        return this.generateOutputStickyAction(this.schedulardata);
    }
    // fucntion for generating sticky actions
    generateOutputStickyAction(schedulardata : Schedulardata) {          
        const actions = new ActionsData();      
        actions.actions = this.generateStickyAction(schedulardata);
        actions.enddate = this.generateEndDate(schedulardata);
        return actions;
    }
}// end of class

