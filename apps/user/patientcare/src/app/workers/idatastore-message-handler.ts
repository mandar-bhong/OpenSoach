import { IDatastoreModel } from "../models/db/idatastore-model.js";

export interface IDatastoreMessageHandler<T extends IDatastoreModel> {
    handleMessage(msg:T): void;
}