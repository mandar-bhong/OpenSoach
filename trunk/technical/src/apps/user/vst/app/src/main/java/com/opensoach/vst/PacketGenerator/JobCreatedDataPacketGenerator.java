package com.opensoach.vst.PacketGenerator;

import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceJobCreatedDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceTaskItemDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketTokenClaimDataModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;

import java.util.ArrayList;

public class JobCreatedDataPacketGenerator implements IPacketGenerator<JobServiceViewModel> {

    @Override
    public CommandRequest GenerateRequest(int locationID, JobServiceViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketServiceJobCreatedDataModel> request = new CommandRequest<>();

        PacketServiceJobCreatedDataModel packetServiceJobCreatedDataModel = new PacketServiceJobCreatedDataModel();

        request.Packet = new PacketModel<>();
        request.Packet.Payload = packetServiceJobCreatedDataModel;
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CHART_DATA, requestID, locationID);


        packetServiceJobCreatedDataModel.CustomerDetails = new PacketServiceCustomerDetailsDataModel();
        packetServiceJobCreatedDataModel.VehicleDetails = new PacketServiceVehicleDetailsDataModel();

        packetServiceJobCreatedDataModel.TokenId = data.getTokenItemViewModel().getDbTokenTableRowModel().getId();
        packetServiceJobCreatedDataModel.TokenNo = data.getTokenItemViewModel().getDbTokenTableRowModel().getTokenno();


        packetServiceJobCreatedDataModel.CustomerDetails.FirstName = data.getJobServiceDetailsViewModel().getFirstName();
        packetServiceJobCreatedDataModel.CustomerDetails.LastName = data.getJobServiceDetailsViewModel().getLastName();
        packetServiceJobCreatedDataModel.CustomerDetails.MobileNo = data.getJobServiceDetailsViewModel().getMobileNo();

        packetServiceJobCreatedDataModel.VehicleDetails.KM = data.getJobServiceDetailsViewModel().getKmRuns();
        packetServiceJobCreatedDataModel.VehicleDetails.Petrol = data.getJobServiceDetailsViewModel().getPetrolLevel();

        packetServiceJobCreatedDataModel.Tasks = new ArrayList<>();


        for(JobServiceItemViewModel model : data.getJobServiceListViewModel().getData()){
            PacketServiceTaskItemDataModel item = new PacketServiceTaskItemDataModel();

            item.taskName = model.getTaskName();
            item.Comment = model.getComment();
            item.Cost = model.getCost();

            packetServiceJobCreatedDataModel.Tasks.add(item);
        }



        return  request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

}
