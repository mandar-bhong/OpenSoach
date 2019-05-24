package spl.hkt.opensoach.splapp.packetGenerator;

import java.util.ArrayList;

import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.RequestManager;
import spl.hkt.opensoach.splapp.model.communication.CommandRequest;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketUserComplaintDataModel;
import spl.hkt.opensoach.splapp.processor.AckComplaintProcessor;
import spl.hkt.opensoach.splapp.processor.AckFeedbackProcessor;

public class FeedbackDataPacketGenerator implements IPacketGenerator<ArrayList<PacketFeedbackDataModel>> {

    @Override
    public CommandRequest GenerateRequest(int locationID, ArrayList<PacketFeedbackDataModel> data) {
        return GetFeedbackDataPacket(locationID, data);
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

    public CommandRequest GetFeedbackDataPacket(int locationID, ArrayList<PacketFeedbackDataModel> complaints) {
        PacketModel<ArrayList<PacketFeedbackDataModel>> packetModel = new PacketModel<>();
        int seqid = RequestManager.Instance().GenerateRequestID();
        packetModel.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_FEEDBACK_DATA, seqid, locationID);

        packetModel.Payload = complaints;

        CommandRequest<ArrayList<PacketFeedbackDataModel>> commandRequest = new CommandRequest<>();
        commandRequest.Packet = packetModel;
        commandRequest.AckProcessor = new AckFeedbackProcessor();

        RequestManager.Instance().AddRequest(seqid, commandRequest);

        return commandRequest;
    }
}
