package com.opensoach.vst.Processor;

import java.util.ArrayList;

import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketUserComplaintDataModel;

public class AckComplaintProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<ArrayList<PacketUserComplaintDataModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        if (request == null) {
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;
        }

        RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketSimpleAckModel ack = (PacketSimpleAckModel) packetDecodeResultModel.Packet.Payload;

        if (!ack.Ack) {
            packetProcessResultModel.IsSuccess = false;
            // TODO: save in local storage
            return packetProcessResultModel;
        }

        // TODO: save in local storage

        packetProcessResultModel.IsSuccess = true;

        return packetProcessResultModel;
    }
}

