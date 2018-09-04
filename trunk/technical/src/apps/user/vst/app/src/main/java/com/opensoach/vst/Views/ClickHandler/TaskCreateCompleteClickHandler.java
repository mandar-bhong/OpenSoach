package com.opensoach.vst.Views.ClickHandler;

import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.TaskCreationViewModel;
import com.opensoach.vst.ViewModels.TaskItemViewModel;

public class TaskCreateCompleteClickHandler {

    public void onClick(View view, JobServiceItemViewModel vm) {
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().addItem(vm);
        vm.ContextActivity.finish();
    }
}
