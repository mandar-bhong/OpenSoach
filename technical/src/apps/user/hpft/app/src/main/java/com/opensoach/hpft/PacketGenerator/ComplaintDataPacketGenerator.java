package com.opensoach.hpft.PacketGenerator;

import java.util.ArrayList;

import com.opensoach.hpft.Constants.Constants.CommandConstants;
import com.opensoach.hpft.Helper.PacketHelper;
import com.opensoach.hpft.Manager.RequestManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketUserComplaintDataModel;
import com.opensoach.hpft.Processor.AckComplaintProcessor;

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
