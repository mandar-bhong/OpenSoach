package spl.hkt.opensoach.splapp.packetGenerator;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketAuthenticationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketPayloadModel;
import spl.hkt.opensoach.splapp.processor.AckDeviceRegProcessor;

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
