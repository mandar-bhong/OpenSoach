package com.opensoach.vst.PacketGenerator;

import android.telephony.gsm.GsmCellLocation;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
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
import com.opensoach.vst.Model.Communication.PacketVehicleDetailsModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;

import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;

public class JobCreatedDataPacketGenerator implements IPacketGenerator<JobServiceViewModel> {

    @Override
    public CommandRequest GenerateRequest(int locationID, JobServiceViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketServiceJobCreatedDataModel> request = new CommandRequest<>();

        PacketServiceJobCreatedDataModel packetServiceJobCreatedDataModel = new PacketServiceJobCreatedDataModel();

        request.Packet = new PacketModel<>();
        request.Packet.Payload = packetServiceJobCreatedDataModel;
        request.Packet.Header = PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_CREATE_JOB, requestID, locationID);

        packetServiceJobCreatedDataModel.ServInId = (Integer) AppRepo.getInstance().getStore().get(ApplicationConstants.APP_STORE_SERVICE_INST_ID);
        packetServiceJobCreatedDataModel.Status = ApplicationConstants.TOKEN_JOB_CREATED;


        packetServiceJobCreatedDataModel.TokenId = data.getTokenItemViewModel().getDbTokenTableRowModel().getId();
        packetServiceJobCreatedDataModel.TokenNo = data.getTokenItemViewModel().getDbTokenTableRowModel().getTokenno();

        Calendar txtDate = Calendar.getInstance();
        packetServiceJobCreatedDataModel.TxnDate = txtDate.getTime();


        PacketVehicleDetailsModel packetVehicleDetailsModel = new PacketVehicleDetailsModel();

        packetVehicleDetailsModel.TokenID = data.getTokenItemViewModel().getDbTokenTableRowModel().getId();

      //  packetVehicleDetailsModel.CustomerDetails = new PacketServiceCustomerDetailsDataModel();
        packetVehicleDetailsModel.VehicleDetails = new PacketServiceVehicleDetailsDataModel();

//        packetVehicleDetailsModel.CustomerDetails.FirstName = data.getJobServiceDetailsViewModel().getFirstName();
//        packetVehicleDetailsModel.CustomerDetails.LastName = data.getJobServiceDetailsViewModel().getLastName();
//        packetVehicleDetailsModel.CustomerDetails.MobileNo = data.getJobServiceDetailsViewModel().getMobileNo();

        packetVehicleDetailsModel.VehicleDetails.KM = data.getJobServiceDetailsViewModel().getKmRuns();
        packetVehicleDetailsModel.VehicleDetails.Petrol = data.getJobServiceDetailsViewModel().getPetrolLevel();

        packetVehicleDetailsModel.Tasks = new ArrayList<>();
        Integer tentetiveCost = 0;

        for(JobServiceItemViewModel model : data.getJobServiceListViewModel().getData()){
            PacketServiceTaskItemDataModel item = new PacketServiceTaskItemDataModel();

            item.taskName = model.getTaskName();
            item.Comment = model.getComment();
            item.Cost = model.getCost();

            tentetiveCost = tentetiveCost + Integer.parseInt( item.Cost);

            packetVehicleDetailsModel.Tasks.add(item);
        }

        packetVehicleDetailsModel.TentetiveCost = tentetiveCost.toString();

        Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();
        packetServiceJobCreatedDataModel.TxnData =  gson.toJson(packetVehicleDetailsModel);

        return  request;
    }

    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

}
