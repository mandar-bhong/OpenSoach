package com.opensoach.hospital.Processor;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.util.Log;

import com.opensoach.hospital.Communication.CommunicationManager;
import com.opensoach.hospital.Helper.AppHelper;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 8/26/2017.
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

        try {
            // msg="{\"header\":{\"crc\":\"\",\"category\":2,\"commandid\":3,\"seqid\":1,\"locationid\":1},\"payload\":{\"chartid\":1,\"locationid\":0,\"customerid\":0,\"locationcategoryid\":0,\"chartname\":\"Chart1\",\"starttime\":300,\"endtime\":900,\"slotinterval\":60,\"tasks\":[{\"taskid\":1,\"taskname\":\"Task1\",\"taskorder\":0},{\"taskid\":2,\"taskname\":\"Task2\",\"taskorder\":0},{\"taskid\":3,\"taskname\":\"Task3\",\"taskorder\":0},{\"taskid\":4,\"taskname\":\"Task4\",\"taskorder\":0},{\"taskid\":5,\"taskname\":\"Task5\",\"taskorder\":0}]}}";

            //     PacketModel packetModel = new Gson().fromJson(msg, PacketModel.class);
            AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"Packet Processor: Received Packet: "+msg);

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

            PacketProcessResultModel processResult = result.Processor.Process(result);

            if (!processResult.IsSuccess) {
                //TODO: Log error and discard this packet
                return;
            }

        if (processResult.CanUpdateUI) {
            AppHelper.ProcessUIEvent(processResult.UINotifierModel);
        }

            if (processResult.CanSendServerCommand) {
                //TODO: Handle Server Online state
                CommunicationManager.getInstance().SendPacket(processResult.ServerCommandPacket);
            }
        }catch (Exception ex){
            AppLogger.getInstance().Log(ex,"Error occured in PacketProcessor");
        }

    }
}
