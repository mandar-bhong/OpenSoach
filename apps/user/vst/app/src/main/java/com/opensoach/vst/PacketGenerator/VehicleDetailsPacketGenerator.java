package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketVehicleDetailsRequestDataModel;
import com.opensoach.vst.Processor.AckVehicleDetailsProcessor;

public class VehicleDetailsPacketGenerator implements IPacketGenerator<String> {

    @Override
    public CommandRequest GenerateRequest(int locationID, String vehicleno) {
        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketVehicleDetailsRequestDataModel> request = new CommandRequest<>();
        request.Packet = new PacketModel<PacketVehicleDetailsRequestDataModel>();
        request.Packet.Payload = new PacketVehicleDetailsRequestDataModel();
        request.Packet.Payload.VehicleNumber = vehicleno;

        request.Packet.Header = new PacketHeaderModel();
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_CONFIG,
                CommandConstants.CMD_CONFIG_GET_VEHICLE_DETAILS,requestID,locationID);


        request.AckProcessor = new AckVehicleDetailsProcessor();

        RequestManager.Instance().AddRequest(requestID,request);

        return request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
