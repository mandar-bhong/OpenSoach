package com.opensoach.hospital.Views.ClickHandler;

import android.content.DialogInterface;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Manager.SendPacketManager;
import com.opensoach.hospital.Model.Communication.Command.DeviceDataDropJobModel;
import com.opensoach.hospital.ViewModels.JobBoardViewModel;

import java.util.Date;

/**
 * Created by Mandar on 18-06-2018.
 */

public class JobDropClickHandler implements DialogInterface.OnClickListener {

    JobBoardViewModel vm;

    public JobDropClickHandler(JobBoardViewModel viewModel){
        vm= viewModel;
    }

    @Override
    public void onClick(DialogInterface dialog, int which) {

        switch (which){
            case DialogInterface.BUTTON_POSITIVE:
                onPositiveButtonClick();
                break;
        }
    }

    void onPositiveButtonClick(){
        DeviceDataDropJobModel deviceDataDropJobModel = DeviceDataDropJobModel.create(
                AppRepo.getInstance().getCurrentLocationId(),
                vm.getJobCardId(),
                "",
                new Date());

        deviceDataDropJobModel.setCommandType(CommandConstants.DEVICE_DATA_COMMAND_DROPED_JOB);
        SendPacketManager.Instance().send(deviceDataDropJobModel);

        vm.setJobState(Constants.JOB_STATUS_DROPED);
    }

}
