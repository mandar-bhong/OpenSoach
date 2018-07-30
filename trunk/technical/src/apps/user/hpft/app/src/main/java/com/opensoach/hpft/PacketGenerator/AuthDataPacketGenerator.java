package com.opensoach.hpft.PacketGenerator;

import com.opensoach.hpft.Constants.CommandConstants;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketAuthenticationModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Processor.AckDeviceRegProcessor;

public class AuthDataPacketGenerator implements IPacketGenerator<String> {

    @Override
    public CommandRequest GenerateRequest(int locationID, String data) {
        PacketModel<PacketAuthenticationModel> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header =PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DEVICE_REG,
                CommandConstants.CMD_DEVICE_REGISTRATION, seqid, 0);

        PacketAuthenticationModel packetAuthenticationModel = new PacketAuthenticationModel();
        packetAuthenticationModel.AuthToken = data;

        packetModel.Payload = packetAuthenticationModel;

        CommandRequest<PacketAuthenticationModel> commandRequest = new CommandRequest<>();
        commandRequest.Packet = packetModel;
        commandRequest.AckProcessor = new AckDeviceRegProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);

        return commandRequest;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
