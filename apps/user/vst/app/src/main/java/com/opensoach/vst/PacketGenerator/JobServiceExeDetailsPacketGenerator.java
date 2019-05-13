package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceGetJobDetailsDataModel;
import com.opensoach.vst.Processor.AckJobServiceDetailsProcessor;

import java.util.ArrayList;

public class JobServiceExeDetailsPacketGenerator implements IPacketGenerator<Integer> {

    @Override
    public CommandRequest GenerateRequest(int locationID, Integer tokenID) {
        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketServiceGetJobDetailsDataModel> request = new CommandRequest<>();
        RequestManager.Instance().AddRequest(requestID,request);


        PacketServiceGetJobDetailsDataModel packetServiceGetExeDetailsDataModel = new PacketServiceGetJobDetailsDataModel();
        packetServiceGetExeDetailsDataModel.TokenIDs = new ArrayList<>();
        packetServiceGetExeDetailsDataModel.TokenIDs.add(tokenID);


        PacketModel<PacketServiceGetJobDetailsDataModel> packet = new PacketModel<>();
        packet.Header =   PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,
                CommandConstants.CMD_CONFIG_GET_TOKEN_JOB_DETAILS,requestID,locationID);
        packet.Payload = packetServiceGetExeDetailsDataModel;

        request.Packet = packet;
        request.AckProcessor = new AckJobServiceDetailsProcessor();

        return request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
