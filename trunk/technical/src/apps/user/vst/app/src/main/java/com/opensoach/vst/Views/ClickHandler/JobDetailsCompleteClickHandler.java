package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobServiceListActivity;

import java.util.ArrayList;

public class JobDetailsCompleteClickHandler {

    public void onClick(View view, JobDetailsViewModel vm) {

        AppRepo.getInstance().getJobServiceViewModel().setJobDetailsViewModel(vm);
        AppRepo.getInstance().getJobServiceViewModel().setJobServiceListViewModel(new JobServiceListViewModel());
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().ContextActivity = vm.ContextActivity;
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().Parent = vm;
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().setData( new ArrayList<JobServiceItemViewModel>());


        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceListActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }

}
