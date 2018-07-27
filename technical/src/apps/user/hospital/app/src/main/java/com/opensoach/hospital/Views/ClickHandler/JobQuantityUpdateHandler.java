package com.opensoach.hospital.Views.ClickHandler;

import android.app.Activity;
import android.view.View;
import android.widget.Toast;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Manager.SendPacketManager;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataJobQuantityUpdateModel;
import com.opensoach.hospital.Model.View.UIViewActionRequestModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.ViewModels.JobQuantityViewModel;

import java.util.Date;

/**
 * Created by Mandar on 25-11-2017.
 */

public class JobQuantityUpdateHandler {

    public void onQuantityUpdate(View view, JobQuantityViewModel vm) {

        if (
                (vm.getOperatorCode() == "" || vm.getOperatorCode() == "") &&
                (vm.getFinishedQuantity() == null || vm.getFinishedQuantity() == "") &&
                (vm.getComment() == null || vm.getComment() == "")
                ) {
            Toast.makeText(view.getContext(), R.string.msg_all_fields_empty, Toast.LENGTH_SHORT).show();
            return;
        }

        if ((vm.getFinishedQuantity() == null || vm.getFinishedQuantity() == "" )&& (vm.getComment() == null || vm.getComment() == "") ){
            Toast.makeText(view.getContext(), R.string.msg_quantity_and_comment_empty, Toast.LENGTH_SHORT).show();
            return;
        }

        if (CommonHelper.IsOperatorCodeValid(vm.getOperatorCode())) {
            DeviceDataJobQuantityUpdateModel deviceDataJobQuantityUpdateModel = DeviceDataJobQuantityUpdateModel.create(
                    AppRepo.getInstance().getCurrentLocationId(),
                    vm.getJobID(),
                    vm.getOperatorCode(),
                    vm.getQuantity(),
                    new Date(),
                    vm.getComment());

            deviceDataJobQuantityUpdateModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_UPDATE_JOB_UNIT);

            SendPacketManager.Instance().send(deviceDataJobQuantityUpdateModel);

            SendViewActiveCommand(view, CommandConstants.UI_CMD_BACKGROUND_CLOSE_DIALOG, 0, null);
        } else {
            Toast.makeText(view.getContext(), R.string.msg_operator_code_invalid, Toast.LENGTH_SHORT).show();
        }
    }


    public void onCancel(View view, JobQuantityViewModel vm){
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
