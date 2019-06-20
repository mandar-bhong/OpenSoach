package com.opensoach.vst.Processor;

import java.util.ArrayList;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Helper.SyncState;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;

public class AckDeviceRegProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        AppRepo.getInstance().setIsDeviceSyncInProgress(true);
        CommandRequest<ArrayList<PacketServiceInstanceTxnModel>> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        if (request == null) {
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;
        }

        RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketSimpleAckModel ack = (PacketSimpleAckModel) packetDecodeResultModel.Packet.Payload;

        if (!ack.Ack) {
            packetProcessResultModel.IsSuccess = false;
            return packetProcessResultModel;
        }

        SendPacketManager.Instance().sendOnStateChange(SyncState.DEVICE_REGISTRATION_COMPLETED);

        packetProcessResultModel.IsSuccess = true;
        return packetProcessResultModel;
    }
}

