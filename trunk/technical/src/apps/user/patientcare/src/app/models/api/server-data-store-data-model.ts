import { SYNC_STORE } from "~/app/app-constants";
import { IDatastoreModel } from "../db/idatastore-model";

export class ServerDataStoreDataModel
{
	public datastore: SYNC_STORE;
	public data:IDatastoreModel;
}