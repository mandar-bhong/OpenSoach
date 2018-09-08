package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.content.pm.ApplicationInfo;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobCreationActivity;
import com.opensoach.vst.Views.Activity.JobServiceListActivity;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;

import java.util.ArrayList;

public class TokenSelectionHandler {

    public void onClick(View view, TokenSelectionViewModel vm) {



        if (  AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.Token)
        {


        }
        else if( AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution)
        {
            JobServiceDetailsViewModel jobDetailsViewModel = new JobServiceDetailsViewModel();
            jobDetailsViewModel.Parent = vm;
            jobDetailsViewModel.ContextActivity = vm.ContextActivity;

            jobDetailsViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
            AppRepo.getInstance().setJobServiceDetailsViewModel(jobDetailsViewModel);

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

            Intent i = new Intent(view.getContext(), JobServiceListActivity.class);
            view.getContext().startActivity(i);

        }
        else if( AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobCreation)
        {

            JobServiceDetailsViewModel jobDetailsViewModel = new JobServiceDetailsViewModel();
            jobDetailsViewModel.Parent = vm;
            jobDetailsViewModel.ContextActivity = vm.ContextActivity;

            jobDetailsViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
            AppRepo.getInstance().setJobServiceDetailsViewModel(jobDetailsViewModel);

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


            Intent i = new Intent(view.getContext(), JobCreationActivity.class);
            view.getContext().startActivity(i);


        } else {}


    }
}
