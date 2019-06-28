package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.TaskItemViewModel;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;

/**
 * Created by Mandar on 02-08-2018.
 */

public class TaskDetailsClickHandler {

    public void onClick(View view, TaskItemViewModel vm) {
        AppRepo.getInstance().setActiveTaskItem(vm);

        Intent i = new Intent(view.getContext(), TaskDetailsActivity.class);
        view.getContext().startActivity(i);
    }
}
