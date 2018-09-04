package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.ViewModels.TokenListViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobCreationActivity;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;

public class TokenSelectionHandler {

    public void onClick(View view, TokenSelectionViewModel vm) {

        AppRepo.getInstance().setJobServiceViewModel(new JobServiceViewModel());
        AppRepo.getInstance().getJobServiceViewModel().ContextActivity = vm.ContextActivity;
        AppRepo.getInstance().getJobServiceViewModel().Parent = vm;

        AppRepo.getInstance().getJobServiceViewModel().setTokenSelectionViewModel(vm);

        Intent i = new Intent(view.getContext(), JobCreationActivity.class);
        view.getContext().startActivity(i);

    }
}
