package com.opensoach.vst.Processor;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.util.Log;

import com.opensoach.vst.SPLApplication;
import com.opensoach.vst.Communication.CommunicationManager;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;

/**
 * Created by Mandar on 2/22/2017.
 * This class will handle all received packet from communication layer
 */

    public class PacketProcessor extends Handler {

    @Override
    public void handleMessage(Message msg) {
        super.handleMessage(msg);

        Bundle b = msg.getData();
        String strResponse = b.getString("Server_Received_Packet");
        Log.d("Server_Received_Packet",strResponse);
        Process(strResponse);
    }

    public void Process(String msg) {
       // msg="{\"header\":{\"crc\":\"\",\"category\":2,\"commandid\":3,\"seqid\":1,\"locationid\":1},\"payload\":{\"chartid\":1,\"locationid\":0,\"customerid\":0,\"locationcategoryid\":0,\"chartname\":\"Chart1\",\"starttime\":300,\"endtime\":900,\"slotinterval\":60,\"tasks\":[{\"taskid\":1,\"taskname\":\"Task1\",\"taskorder\":0},{\"taskid\":2,\"taskname\":\"Task2\",\"taskorder\":0},{\"taskid\":3,\"taskname\":\"Task3\",\"taskorder\":0},{\"taskid\":4,\"taskname\":\"Task4\",\"taskorder\":0},{\"taskid\":5,\"taskname\":\"Task5\",\"taskorder\":0}]}}";

   //     PacketModel packetModel = new Gson().fromJson(msg, PacketModel.class);

        PacketDecodeResultModel result = HeaderDecoder.Decode(msg);

        if (!result.IsSuccess) {
            //TODO: Log error and discard this packet
            return;
        }

        PayloadDecoder.Decode(result, msg);

        if (!result.IsSuccess) {
            //TODO: Log error and discard this packet
            return;
        }

        if(result.Processor!=null) {
            PacketProcessResultModel processResult = result.Processor.Process(result);

            if (!processResult.IsSuccess) {
                //TODO: Log error and discard this packet
                return;
            }

            if (processResult.CanUpdateUI) {
                SPLApplication.getInstance().OnUIUpdateEvent(processResult.UINotifierModel);
            }

            if (processResult.CanSendServerCommand) {
                //TODDO: Handle Server Online state
                CommunicationManager.getInstance().SendPacket(processResult.ServerCommandPacket);
            }
        }

    }
}
