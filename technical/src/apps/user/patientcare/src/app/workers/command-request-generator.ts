import { CmdModel, CmdHeader, AuthTokenModel, GetSyncRequestModel, ApplySyncRequestModel } from "../models/api/server-cmd-model.js";
import { CMD_CATEGORY, CMD_ID } from "../app-constants.js";
import { ServerWorkerContext } from "./server-worker-context.js";
import { RequestManager } from "./request-manager.js";

export class CommandRequestGenerator{

    public static authCmd() {
        // {"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"token":"Dev6AD88A481524BABF"}}

        const authcmd = new CmdModel();
        authcmd.header = new CmdHeader();

        authcmd.header.category = CMD_CATEGORY.CMD_CAT_DEV_REGISTRATION;
        authcmd.header.commandid = CMD_ID.CMD_DEV_REGISTRATION;

        const authTokenModel = new AuthTokenModel();
        authTokenModel.token = ServerWorkerContext.authToken;

        authcmd.payload = authTokenModel;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(authcmd);

        const cmdstring = JSON.stringify(authcmd);

        return cmdstring;
    }

    public static getSyncCmd(strname: string, lastSynched: Date) {
        // {"header":{"crc":"12","category":3,"commandid":50,"seqid":3},"payload":{"storename":"","updatedon":"2018-10-30T00:00:00Z"}}

        const getSyncCmd = new CmdModel();
        getSyncCmd.header = new CmdHeader();

        getSyncCmd.header.category = CMD_CATEGORY.CMD_CAT_SYNC;
        getSyncCmd.header.commandid = CMD_ID.CMD_GET_STORE_SYNC;

        const getrequest = new GetSyncRequestModel();
        getrequest.storename = strname;

        // for first time sync - last synched empty then send default date
        if (lastSynched) {
            getrequest.updatedon = lastSynched;
        } else {
            getrequest.updatedon = new Date("2018-10-30T00:00:00Z");
        }

        getSyncCmd.payload = getrequest;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(getSyncCmd);

        const cmdstring = JSON.stringify(getSyncCmd);

        return cmdstring
    }

    public static applySyncCmd(storename: any, storedata: any[]) {
        // {"header":{"crc":"12","category":3,"commandid":51,"seqid":3},"payload":{"storename":"","storedata":[{"uuid":"PA001","bedno":"A0001"}]}}

        const applySyncCmd = new CmdModel();
        applySyncCmd.header = new CmdHeader();

        applySyncCmd.header.category = CMD_CATEGORY.CMD_CAT_SYNC;
        applySyncCmd.header.commandid = CMD_ID.CMD_APPLY_STORE_SYNC;

        const applyReqModel = new ApplySyncRequestModel();
        applyReqModel.storename = storename;
        applyReqModel.storedata = storedata;

        applySyncCmd.payload = applyReqModel;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(applySyncCmd);

        const cmdstring = JSON.stringify(applySyncCmd);

        return cmdstring;

    }

}