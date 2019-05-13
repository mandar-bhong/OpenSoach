package spl.hkt.opensoach.splapp.packetGenerator;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.CommonHelper;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketPayloadModel;
import spl.hkt.opensoach.splapp.model.communication.PacketServiceInstanceTxnModel;
import spl.hkt.opensoach.splapp.processor.AckChartDataProcessor;

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
