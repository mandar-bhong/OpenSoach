package com.opensoach.hpft.PacketGenerator;

import com.opensoach.hpft.Constants.Constants.CommandConstants;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketPayloadModel;

public class DeviceSyncCompletedDataPacketGenerator implements IPacketGenerator {

    @Override
    public CommandRequest GenerateRequest(int locationID, Object data) {
        PacketModel<PacketPayloadModel> packetModel = new PacketModel<>();
         int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,
                CommandConstants.CMD_CONFIG_DEVICE_SYNC_COMPLETED, seqid, 0);

        packetModel.Payload = new PacketPayloadModel();;

        CommandRequest<PacketPayloadModel> commandRequest = new CommandRequest<>();
        commandRequest.Packet= packetModel;

        RequestManager.Instance().AddRequest(seqid, commandRequest);

        return commandRequest;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
