package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobSummaryViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Activity.JobServiceSummaryActivity;

public class JobServiceSummaryClickHandler {

    public void onClick(View view) {

//        JobServiceItemViewModel jobServiceItemViewModel = new JobServiceItemViewModel();
//        jobServiceItemViewModel.ContextActivity = vm.ContextActivity;
//        jobServiceItemViewModel.Parent = vm;


        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().setDisplayMode(ApplicationConstants.DISPLAY_MODE_JOB_CREATION_SUMMARY);

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceSummaryActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);

    }
}
