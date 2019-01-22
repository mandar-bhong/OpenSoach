import { IDatastoreModel } from "./idatastore-model";

export class ScheduleDatastoreModel implements IDatastoreModel
{
    uuid:string;
    getModelValues():any[]
    {
        return [this.uuid];
    }
}