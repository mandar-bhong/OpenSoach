export class ServerWorkerContext {
    public static ContextVar1: string;
    public static authToken: string;
    public static deviceSerialNo: string;
    public static serverUrl: string;

    // SYNC Engine Context
    public static isSyncInprogress: boolean;
    public static syncType: SYNC_TYPE;
    public static syncDirection: SYNC_DIRECTION;
}

export enum SYNC_TYPE {
    NONE = 0,
    FULL = 1,
    DIFFERENTIAL = 2
}

export enum SYNC_DIRECTION {
    NONE = 0,
    TO_SERVER = 1,
    FROM_SERVER = 2
}