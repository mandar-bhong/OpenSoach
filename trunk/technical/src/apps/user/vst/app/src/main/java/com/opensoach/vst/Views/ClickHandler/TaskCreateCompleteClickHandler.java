package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;

import com.opensoach.vst.ViewModels.TaskItemViewModel;
import com.opensoach.vst.Views.Activity.JobCreationActivity;
import com.opensoach.vst.Views.Activity.JobServiceCreationActivity;
import com.opensoach.vst.Views.Activity.JobServiceListActivity;

public class TaskCreateCompleteClickHandler {

    public void onClick(View view, JobServiceItemViewModel vm) {
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().addItem(vm);

       ((Activity) view.getContext()).finish();

    }
}
