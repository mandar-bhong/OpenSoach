package spl.hkt.opensoach.splapp.packetGenerator;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketPayloadModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.processor.AckComplaintProcessor;

public class ComplaintDataPacketGenerator implements IPacketGenerator<ArrayList<PacketUserComplaintDataModel>> {

    @Override
    public CommandRequest GenerateRequest(int locationID, ArrayList<PacketUserComplaintDataModel> data) {
        return GetComplaintDataPacket(locationID, data);
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

    public CommandRequest GetComplaintDataPacket(int locationID, ArrayList<PacketUserComplaintDataModel> complaints) {
        PacketModel<ArrayList<PacketUserComplaintDataModel>> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_COMPLAINT_DATA, seqid, locationID);

        packetModel.Payload = complaints;

        CommandRequest<ArrayList<PacketUserComplaintDataModel>> commandRequest = new CommandRequest<>();
        commandRequest.Packet = packetModel;
        commandRequest.AckProcessor = new AckComplaintProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);

        return commandRequest;
    }
}
