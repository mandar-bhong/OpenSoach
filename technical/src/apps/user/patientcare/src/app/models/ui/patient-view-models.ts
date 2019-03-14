import { PatientListDBModel } from "../db/patient-list-model";
import { NextActionTimes } from "./next-action-times";

export class PatientListViewModel {
    dbmodel: PatientListDBModel;
    contact: string;
    nextActionTimes: NextActionTimes = new NextActionTimes();
}

