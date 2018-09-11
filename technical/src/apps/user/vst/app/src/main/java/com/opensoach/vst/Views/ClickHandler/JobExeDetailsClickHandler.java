package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobExeDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Activity.JobExeDetailsActivity;

public class JobExeDetailsClickHandler  {

    public void onClick(View view, JobServiceItemViewModel vm) {
      AppRepo.getInstance().getJobServiceViewModel().setJobServiceItemViewModel(vm);
        Intent i = new Intent(view.getContext(), JobExeDetailsActivity.class);
        view.getContext().startActivity(i);
    }
}
