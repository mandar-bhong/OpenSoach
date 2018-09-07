package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.Views.Activity.JobServiceListActivity;
import java.util.ArrayList;

public class JobDetailsCompleteClickHandler {

        public void onClick(View view, JobServiceDetailsViewModel vm) {

        Intent i = new Intent(MainViewModel.getInstance().ContextActivity, JobServiceListActivity.class);
        MainViewModel.getInstance().ContextActivity.startActivity(i);
    }

}
