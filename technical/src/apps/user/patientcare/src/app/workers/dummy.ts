import { ServerWorkerContext } from "./server-worker-context";

export class Dummy
{
    public static sendToServerCallback: (msg:any)=>void;
    public static DummyMethod()
    {
       console.log('in dummy method');
       ServerWorkerContext.ContextVar1="Dummy packet sent";
       Dummy.sendToServerCallback("dummy packet"); 
    }

    public static Init()
    {
        console.log('in Dummy Init');
    }
}