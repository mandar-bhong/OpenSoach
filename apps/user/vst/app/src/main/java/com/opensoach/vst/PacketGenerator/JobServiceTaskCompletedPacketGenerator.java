package com.opensoach.vst.PacketGenerator;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceGetJobDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceJobCreatedDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceOwnerVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceTaskItemDataModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;

import java.util.ArrayList;
import java.util.Calendar;

public class JobServiceTaskCompletedPacketGenerator implements IPacketGenerator<JobServiceItemViewModel> {

    @Override
    public CommandRequest GenerateRequest(int locationID, JobServiceItemViewModel data) {

        Integer requestID  = RequestManager.Instance().GenerateRequestID();

        CommandRequest<PacketServiceJobCreatedDataModel> request = new CommandRequest<>();
        RequestManager.Instance().AddRequest(requestID,request);

        PacketServiceJobCreatedDataModel packetServiceJobCreatedDataModel = new PacketServiceJobCreatedDataModel();

        packetServiceJobCreatedDataModel.ServInId = (Integer) AppRepo.getInstance().getStore().get(ApplicationConstants.APP_STORE_SERVICE_INST_ID);
        packetServiceJobCreatedDataModel.Status = ApplicationConstants.TOKEN_JOB_EXECUTION_INPROGRESS;

        packetServiceJobCreatedDataModel.TokenId = AppRepo.getInstance().getJobServiceViewModel().getTokenItemViewModel().getDbTokenTableRowModel().getId();
        packetServiceJobCreatedDataModel.TokenNo = AppRepo.getInstance().getJobServiceViewModel().getTokenItemViewModel().getDbTokenTableRowModel().getTokenno();

        Calendar txtDate = Calendar.getInstance();
        packetServiceJobCreatedDataModel.TxnDate = txtDate.getTime();


        PacketServiceTaskItemDataModel packetServiceTaskItemDataModel = new PacketServiceTaskItemDataModel();

        packetServiceTaskItemDataModel.taskName = data.getTaskName();
        packetServiceTaskItemDataModel.Note = data.getNote();


        Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();
        packetServiceJobCreatedDataModel.TxnData =  gson.toJson(packetServiceTaskItemDataModel);

        PacketModel<PacketServiceJobCreatedDataModel> packet = new PacketModel<>();
        packet.Header =   PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA,
                CommandConstants.CMD_DATA_JOB_TASK_COMPLETED,requestID,locationID);
        packet.Payload = packetServiceJobCreatedDataModel;

        request.Packet = packet;


        return  request;
    }


    @Override
    public CommandRequest GenerateUnsyncRequest(int locationID) {
        return null;
    }

}
