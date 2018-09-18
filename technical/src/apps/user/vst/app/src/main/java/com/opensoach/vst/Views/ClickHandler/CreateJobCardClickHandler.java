package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobServiceCreationActivity;

public class CreateJobCardClickHandler {

    public void onClick(View view, TokenSelectionViewModel vm) {

        if (vm.getTokenListViewModel().getSelectedToken() == null)
            return;

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceCreationActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);

    }

}
