package com.opensoach.vst.PacketGenerator;

import java.util.ArrayList;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketUserComplaintDataModel;
import com.opensoach.vst.Processor.AckComplaintProcessor;

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
