package com.opensoach.hpft.Processor;

import java.util.ArrayList;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.Helper.SyncState;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Manager.SendPacketManager;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.hpft.Model.Communication.PacketSimpleAckModel;

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

