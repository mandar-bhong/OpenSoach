import { SERVER_SYNC_STATE } from "../app-constants";

export class ServerWorkerContext {
    public static authToken: string;
    public static deviceSerialNo: string;
    public static serverUrl: string;

    // SYNC Engine Context
    public static isSyncInprogress: boolean;
    public static syncType: SYNC_TYPE;
    public static syncState: SERVER_SYNC_STATE;
}

export enum SYNC_TYPE {
    NONE = 0,
    FULL = 1,
    DIFFERENTIAL = 2
}

export class CurrentStoreModel {
    currentStoreName: string;
    lastSynched: Date;
}