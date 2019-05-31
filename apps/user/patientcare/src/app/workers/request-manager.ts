import { CmdModel } from "../models/api/server-cmd-model.js";

export class RequestManager
{
    public static currentSeqNo=0;
    private static requestMap= new Map<number,CmdModel>();

    public static setSequencetNumber(cmd:CmdModel)
    {
        RequestManager.currentSeqNo++;
        cmd.header.seqid=RequestManager.currentSeqNo;
        RequestManager.requestMap.set(cmd.header.seqid, cmd);        
    }

    public static getRequest(seqid:number)
    {
        const cmd= RequestManager.requestMap.get(seqid);
        RequestManager.requestMap.delete(seqid);
        return cmd;
    }
}