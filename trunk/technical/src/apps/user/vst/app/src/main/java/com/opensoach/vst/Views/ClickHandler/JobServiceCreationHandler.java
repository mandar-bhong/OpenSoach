package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.Views.Activity.JobServiceSummaryActivity;
import com.opensoach.vst.Views.Activity.JobServiceTaskCreationActivity;

public class JobServiceCreationHandler {

    public void onCreateTask(View view, JobServiceListViewModel vm) {

        Intent i = new Intent(view.getContext(), JobServiceTaskCreationActivity.class);
        view.getContext().startActivity(i);
    }

    public void onTaskCreateCompleted(View view, JobServiceItemViewModel vm) {
        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().addItem(vm);

        ((Activity) view.getContext()).finish();


    }

    public void onShowSummary(View view) {

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceSummaryActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }


    public void onSummaryConfirmClick(View view) {
        JobServiceViewModel jobServiceViewModel  = AppRepo.getInstance().getJobServiceViewModel();

        TokenItemViewModel tokenItemViewModel = jobServiceViewModel.getTokenItemViewModel();
        JobServiceDetailsViewModel jobServiceDetailsViewModel =  jobServiceViewModel.getJobServiceDetailsViewModel();
        JobServiceListViewModel jobServiceListViewModel  = jobServiceViewModel.getJobServiceListViewModel();


        SendPacketManager.Instance().send(AppAction.CREATE_JOB_COMFIRM,jobServiceViewModel);

        AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_JOB_SUBMITTED,true);
        ((Activity)view.getContext()).finish();

    }
}
