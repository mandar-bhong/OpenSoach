package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.Views.Activity.JobServiceTaskCreationActivity;

public class JobServiceTaskListCreateHandler {

    public void onClick(View view, JobServiceListViewModel vm) {

        Intent i = new Intent(view.getContext(), JobServiceTaskCreationActivity.class);
        view.getContext().startActivity(i);
    }


}
