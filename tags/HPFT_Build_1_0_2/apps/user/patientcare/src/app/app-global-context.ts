import { APP_MODE, SERVER_SYNC_STATE } from "./app-constants";

export class AppGlobalContext
{
    public static AppMode: APP_MODE;
    public static SerialNumber: string;
    public static Token: string;
    public static SyncState:SERVER_SYNC_STATE;
    public static WebsocketUrl:string;
}