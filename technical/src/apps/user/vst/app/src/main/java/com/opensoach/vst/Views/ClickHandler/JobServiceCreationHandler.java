package com.opensoach.vst.Views.ClickHandler;

import android.app.Activity;
import android.app.AlertDialog;
import android.content.DialogInterface;
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
import com.opensoach.vst.Views.Activity.JobServiceTaskListActivity;

public class JobServiceCreationHandler {

    public void onShowJobTaskList(View view, JobServiceDetailsViewModel vm) {

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceTaskListActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }

    public void onCreateTask(View view, JobServiceListViewModel vm) {

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceTaskCreationActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }

    public void onTaskCreateCompleted(View view, JobServiceItemViewModel vm) {

        if(AppRepo.getInstance().getJobServiceViewModel().getJobServiceItemViewModel() == vm){
            ((Activity) view.getContext()).finish();
            return;
        }


        AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel().getJobServiceDataAdapter().addItem(vm);
        Intent i = new Intent(view.getContext(), JobServiceTaskListActivity.class);
        view.getContext().startActivity(i);
    }

    public void onTaskMarkCompleted(View view, JobServiceItemViewModel vm) {

        if ( vm.getTaskCompleted() ) {
            ((Activity) view.getContext()).finish();
            return;
        }

        vm.setTaskCompleted(true);

        SendPacketManager.Instance().send(AppAction.UPDATE_JOB_TASK_COMPLETED,vm);
        ((Activity) view.getContext()).finish();
    }

    public void onTaskCancle(View view, JobServiceItemViewModel vm) {
        ((Activity) view.getContext()).finish();
    }

    public void onShowSummary(View view) {

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceSummaryActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }

    public void onEditTask(View view ,JobServiceItemViewModel vm ) {
        AppRepo.getInstance().getJobServiceViewModel().setJobServiceItemViewModel(vm);
        
        Intent i = new Intent(view.getContext(), JobServiceTaskCreationActivity.class);
        view.getContext().startActivity(i);
    }

    public void onSummaryConfirmClick(View view) {
        JobServiceViewModel jobServiceViewModel  = AppRepo.getInstance().getJobServiceViewModel();

        TokenItemViewModel tokenItemViewModel = jobServiceViewModel.getTokenItemViewModel();
        JobServiceDetailsViewModel jobServiceDetailsViewModel =  jobServiceViewModel.getJobServiceDetailsViewModel();
        JobServiceListViewModel jobServiceListViewModel  = jobServiceViewModel.getJobServiceListViewModel();

        SendPacketManager.Instance().send(AppAction.UPADATE_VEHICLE_OWNER_DETAILS,jobServiceViewModel);

        SendPacketManager.Instance().send(AppAction.CREATE_JOB_COMFIRM,jobServiceViewModel);

        AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_JOB_SUBMITTED,true);
        ((Activity)view.getContext()).finish();

    }

    public void onServiceTaskRemove(final View view, JobServiceItemViewModel vm) {

        final JobServiceItemViewModel jobServiceItemViewModel = vm;

        AlertDialog.Builder builder = new AlertDialog.Builder(view.getContext());
        builder.setTitle("Confirm");
        builder.setMessage("Are you sure to delete?");

        builder.setPositiveButton("YES", new DialogInterface.OnClickListener() {

            public void onClick(DialogInterface dialog, int which) {
                AppRepo.getInstance().getJobServiceViewModel().
                        getJobServiceListViewModel().
                        getJobServiceDataAdapter().
                        removeItem(jobServiceItemViewModel);

                dialog.dismiss();
                Intent i = new Intent(view.getContext(), JobServiceTaskListActivity.class);
                view.getContext().startActivity(i);
            }
        });

        builder.setNegativeButton("NO", new DialogInterface.OnClickListener() {

            @Override
            public void onClick(DialogInterface dialog, int which) {

                // Do nothing
                dialog.dismiss();
            }
        });

        AlertDialog alert = builder.create();
        alert.show();

    }

    public void onShowTaskDetails(View view, JobServiceItemViewModel vm) {

        AppRepo.getInstance().getJobServiceViewModel().setJobServiceItemViewModel(vm);
        Intent i = new Intent(view.getContext(), JobServiceTaskCreationActivity.class);
        view.getContext().startActivity(i);
    }
}
