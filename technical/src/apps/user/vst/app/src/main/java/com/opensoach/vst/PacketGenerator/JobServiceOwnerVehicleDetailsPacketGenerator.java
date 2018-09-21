package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceJobCreatedDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceOwnerVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceVehicleDetailsDataModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;

public class JobServiceOwnerVehicleDetailsPacketGenerator  implements IPacketGenerator<JobServiceViewModel> {
    @Override
    public CommandRequest GenerateRequest(int locationID, JobServiceViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketServiceOwnerVehicleDetailsDataModel> request = new CommandRequest<>();


        PacketServiceOwnerVehicleDetailsDataModel custVechDetails = new PacketServiceOwnerVehicleDetailsDataModel();
        custVechDetails.CustomerDetails = new PacketServiceCustomerDetailsDataModel();
        custVechDetails.CustomerDetails.FirstName = data.getJobServiceDetailsViewModel().getFirstName();
        custVechDetails.CustomerDetails.LastName = data.getJobServiceDetailsViewModel().getLastName();
        custVechDetails.CustomerDetails.MobileNo = data.getJobServiceDetailsViewModel().getMobileNo();


        custVechDetails.VehicleNo = data.getTokenItemViewModel().getVehicleNo();


        request.Packet = new PacketModel<>();
        request.Packet.Payload = custVechDetails;
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_OWNER_VEHICLE_DETAILS, requestID, locationID);


        return request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }
}
