package com.opensoach.vst.PacketGenerator;

import java.util.ArrayList;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketFeedbackDataModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Processor.AckFeedbackProcessor;

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
