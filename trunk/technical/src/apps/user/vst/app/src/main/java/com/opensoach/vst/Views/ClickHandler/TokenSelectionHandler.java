package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
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

        JobServiceDetailsViewModel jobDetailsViewModel = new JobServiceDetailsViewModel();
        jobDetailsViewModel.Parent = vm;
        jobDetailsViewModel.ContextActivity = vm.ContextActivity;

        JobServiceListViewModel jobServiceListViewModel = new  JobServiceListViewModel();
        jobServiceListViewModel.Parent = jobDetailsViewModel;
        jobServiceListViewModel.ContextActivity = vm.ContextActivity;
        jobServiceListViewModel.setData(new ArrayList<JobServiceItemViewModel>());

        JobServiceViewModel  jobServiceViewModel = new JobServiceViewModel();
        jobServiceViewModel.Parent = vm;
        jobServiceViewModel.ContextActivity = vm.ContextActivity;

        jobServiceViewModel.setJobServiceDetailsViewModel(jobDetailsViewModel);
        jobServiceViewModel.setJobServiceListViewModel(jobServiceListViewModel);
        jobServiceViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
        AppRepo.getInstance().setJobServiceViewModel(jobServiceViewModel);

        JobDetailsViewModel jobDetailsViewModel1 = new JobDetailsViewModel();
        jobDetailsViewModel1.Parent = vm;
        jobDetailsViewModel1.ContextActivity = vm.ContextActivity;

        jobDetailsViewModel1.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
        AppRepo.getInstance().setJobDetailsViewModel(jobDetailsViewModel1);

        Intent i = new Intent(view.getContext(), JobCreationActivity.class);
        view.getContext().startActivity(i);

    }
}
