package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketPayloadModel;

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
