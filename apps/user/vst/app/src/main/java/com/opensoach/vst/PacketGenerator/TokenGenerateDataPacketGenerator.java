package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketTokenCreateDataModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.Processor.AckTokenGeneratedProcessor;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;

import java.util.ArrayList;
import java.util.List;

public class TokenGenerateDataPacketGenerator implements IPacketGenerator<CreateTokenViewModel> {

    @Override
    public CommandRequest GenerateRequest(int locationID, CreateTokenViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketTokenCreateDataModel> request = new CommandRequest<>();
        request.Packet = new PacketModel<PacketTokenCreateDataModel>();
        request.Packet.Payload = new PacketTokenCreateDataModel();
        request.Packet.Payload.VehicleNumber = data.getVehicleNumber();

        request.Packet.Header = new PacketHeaderModel();
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_GENERATE_TOKEN, requestID, locationID);

        request.AckProcessor = new AckTokenGeneratedProcessor();

        RequestManager.Instance().AddRequest(requestID,request);

        return request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

}
