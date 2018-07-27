package com.opensoach.hospital.Views.ClickHandler;

import android.app.Activity;
import android.view.View;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Manager.SendPacketManager;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataAbortJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataDropJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStartJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStopJobModel;
import com.opensoach.hospital.Model.View.UIViewActionRequestModel;
import com.opensoach.hospital.ViewModels.JobBoardViewModel;

import java.util.Date;

/**
 * Created by Mandar on 29-10-2017.
 */

public class JobBoardClickHandler {

    public void onJobStart(View view, JobBoardViewModel vm) {
        //SendViewActiveCommand(view, CommandConstants.UI_CMD_BACKGROUND_START_JOB, CommandConstants.UI_VIEW_ACTION_CMD_JOB_START, vm.getJobCardId());

        DeviceDataStartJobModel deviceDataStartJobModel = DeviceDataStartJobModel.create(
                AppRepo.getInstance().getCurrentLocationId(),
                vm.getJobCardId(),
                "",
                new Date());

        deviceDataStartJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_START_JOB);
        SendPacketManager.Instance().send(deviceDataStartJobModel);

        vm.setJobState(Constants.JOB_STATUS_INPROGRESS);
    }

    public void onJobStop(View view, JobBoardViewModel vm) {
        //SendViewActiveCommand(view,CommandConstants.UI_CMD_BACKGROUND_STOP_JOB,CommandConstants.UI_VIEW_ACTION_CMD_JOB_STOP,vm.getJobCardId());

        DeviceDataStopJobModel deviceDataStopJobModel = DeviceDataStopJobModel.create(
                AppRepo.getInstance().getCurrentLocationId(),
                vm.getJobCardId(),
                "",
                new Date());

        deviceDataStopJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_STOP_JOB);
        SendPacketManager.Instance().send(deviceDataStopJobModel);

        vm.setJobState(Constants.JOB_STATUS_COMPLETED);
    }

    public void onJobAbort(View view, JobBoardViewModel vm) {
        DeviceDataAbortJobModel deviceDataAbortJobModel = DeviceDataAbortJobModel.create(
                AppRepo.getInstance().getCurrentLocationId(),
                vm.getJobCardId(),
                "",
                new Date());

        deviceDataAbortJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_ABORTED_JOB);
        SendPacketManager.Instance().send(deviceDataAbortJobModel);

        vm.setJobState(Constants.JOB_STATUS_ABORTED);
    }

    public void onJobDrop(View view, JobBoardViewModel vm) {
        DeviceDataDropJobModel deviceDataDropJobModel = DeviceDataDropJobModel.create(
                AppRepo.getInstance().getCurrentLocationId(),
                vm.getJobCardId(),
                "",
                new Date());

        deviceDataDropJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_DROPED_JOB);
        SendPacketManager.Instance().send(deviceDataDropJobModel);

        vm.setJobState(Constants.JOB_STATUS_DROPED);
    }

    public void onQuantityUpdate(View view, JobBoardViewModel vm) {
        SendViewActiveCommand(view,CommandConstants.UI_CMD_BACKGROUND_UPDATE_JOB_UNIT,0,vm.getJobCardId());
    }

    private void SendViewActiveCommand(View view,int command,int requestCommand, Object data){
        UIViewActionRequestModel uiViewActionRequestModel = new UIViewActionRequestModel();
        uiViewActionRequestModel.CommandID = requestCommand;
        uiViewActionRequestModel.Data = data;
        Activity jobBoardActivity = (Activity) view.getContext();
        IViewAction viewAction = (IViewAction)jobBoardActivity;
        viewAction.ExectuteViewTask(command,uiViewActionRequestModel);
    }



}
