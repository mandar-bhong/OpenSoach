package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobCreationActivity;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;

import java.util.ArrayList;

public class TokenSelectionHandler {

    public void onClick(View view, TokenSelectionViewModel vm) {

        JobDetailsViewModel jobDetailsViewModel = new JobDetailsViewModel();
        jobDetailsViewModel.Parent = vm;
        jobDetailsViewModel.ContextActivity = vm.ContextActivity;

        JobServiceListViewModel jobServiceListViewModel = new  JobServiceListViewModel();
        jobServiceListViewModel.Parent = jobDetailsViewModel;
        jobServiceListViewModel.ContextActivity = vm.ContextActivity;
        jobServiceListViewModel.setData(new ArrayList<JobServiceItemViewModel>());

        JobServiceViewModel  jobServiceViewModel = new JobServiceViewModel();
        jobServiceViewModel.Parent = vm;
        jobServiceViewModel.ContextActivity = vm.ContextActivity;
        jobServiceViewModel.setJobDetailsViewModel(jobDetailsViewModel);
        jobServiceViewModel.setJobServiceListViewModel(jobServiceListViewModel);
        jobServiceViewModel.setTokenSelectionViewModel(vm);


        AppRepo.getInstance().setJobServiceViewModel(jobServiceViewModel);

        Intent i = new Intent(view.getContext(), JobCreationActivity.class);
        view.getContext().startActivity(i);

    }
}
