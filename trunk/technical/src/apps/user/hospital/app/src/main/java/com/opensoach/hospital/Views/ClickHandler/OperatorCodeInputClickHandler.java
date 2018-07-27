package com.opensoach.hospital.Views.ClickHandler;

import android.app.Activity;
import android.view.View;
import android.widget.Toast;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Manager.SendPacketManager;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStartJobModel;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataStopJobModel;
import com.opensoach.hospital.Model.View.UIViewActionRequestModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.ViewModels.OperatorCodeViewModel;

import java.util.Date;

/**
 * Created by Mandar on 26-11-2017.
 */

public class OperatorCodeInputClickHandler {

    public void onOperatorCodeSubmitClick(View view, OperatorCodeViewModel vm,UIViewActionRequestModel requestData) {

        switch (requestData.CommandID) {
            case CommandConstants.UI_VIEW_ACTION_CMD_JOB_START: {
                if (CommonHelper.IsOperatorCodeValid(vm.getOperatorCode())) {
                    DeviceDataStartJobModel deviceDataStartJobModel = DeviceDataStartJobModel.create(
                            AppRepo.getInstance().getCurrentLocationId(),
                            vm.getJobID(),
                            vm.getOperatorCode(),
                            new Date());

                    deviceDataStartJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_START_JOB);

                    SendPacketManager.Instance().send(deviceDataStartJobModel);

                    SendViewActiveCommand(view, CommandConstants.UI_CMD_BACKGROUND_CLOSE_DIALOG, 0, null);
                } else {
                    Toast.makeText(view.getContext(), R.string.msg_operator_code_invalid, Toast.LENGTH_LONG).show();
                }
            }
            break;
            case CommandConstants.UI_VIEW_ACTION_CMD_JOB_STOP: {
                if(CommonHelper.IsOperatorCodeValid(vm.getOperatorCode())) {
                    DeviceDataStopJobModel deviceDataStopJobModel = DeviceDataStopJobModel.create(
                            AppRepo.getInstance().getCurrentLocationId(),
                            vm.getJobID(),
                            vm.getOperatorCode(),
                            new Date());

                    deviceDataStopJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_STOP_JOB);

                    SendPacketManager.Instance().send(deviceDataStopJobModel);

                    SendViewActiveCommand(view,CommandConstants.UI_CMD_BACKGROUND_CLOSE_DIALOG,0,null);
                }else{
                    Toast.makeText(view.getContext(), R.string.msg_operator_code_invalid, Toast.LENGTH_LONG).show();
                }
            }
            break;
        }
    }

    public void onCancel(View view, OperatorCodeViewModel vm,UIViewActionRequestModel requestData){
        SendViewActiveCommand(view,CommandConstants.UI_CMD_BACKGROUND_CLOSE_DIALOG,0,null);
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
