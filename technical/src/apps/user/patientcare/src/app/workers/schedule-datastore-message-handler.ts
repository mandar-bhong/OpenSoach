import { IDatastoreMessageHandler } from "./idatastore-message-handler.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";

export class ScheduleDatastoreMessageHandler implements IDatastoreMessageHandler<ScheduleDatastoreModel>
{
    handleMessage(msg: ScheduleDatastoreModel): void {                
    }
}