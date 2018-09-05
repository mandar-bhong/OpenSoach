package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.ViewModels.CreateTokenViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Activity.JobCreationActivity;
import com.opensoach.vst.Views.Activity.JobServiceCreationActivity;
import com.opensoach.vst.Views.Activity.JobServiceListActivity;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;

public class JobServiceTaskListCreateHandler {

    public void onClick(View view, JobServiceListViewModel vm) {

        Intent i = new Intent(view.getContext(), JobServiceCreationActivity.class);
        view.getContext().startActivity(i);
    }


}
