package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Helper.AppHelper;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.Views.Activity.JobExeDetailsActivity;
import com.opensoach.vst.Views.Activity.LandingPageActivity;

public class JobServiceModeClickHandler {

    public void onTokenCreate(View view) {

        AppRepo.getInstance().setDeviceSerialNo("1234567890123456");

        Start(view);
    }

    public void onJobCreate(View view) {
        AppRepo.getInstance().setDeviceSerialNo("1345494544733456");

        Start(view);
    }

    public void onJobExe(View view) {
        AppRepo.getInstance().setDeviceSerialNo("1155623421323222");

        Start(view);
    }

    private void Start(View view){
        AppHelper.Init(view.getContext());

        AppHelper.ExecuteStartUpProcess();

        Intent myIntent = new Intent(view.getContext(),
                LandingPageActivity.class);
        view.getContext().startActivity(myIntent);
    }

}
