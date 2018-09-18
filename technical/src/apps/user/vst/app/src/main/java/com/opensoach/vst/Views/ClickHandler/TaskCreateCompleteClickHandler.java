package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;

public class TaskCreateCompleteClickHandler {

    public void onClick(View view, JobServiceItemViewModel vm) {
      AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().addItem(vm);

       ((Activity) view.getContext()).finish();
       

    }
}
