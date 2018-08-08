package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.hpft.ViewModels.TaskDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeDataViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeItemViewModel;
import com.opensoach.hpft.Views.Activity.TaskDetailsActivity;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeClickHandler {

    public void onClick(View view, TaskTimeItemViewModel vm) {
        ((TaskDetailsViewModel)((TaskTimeDataViewModel)vm.Parent).Parent).getTaskDataAdapter().updateData(((TaskTimeDataViewModel)vm.Parent).getTaskDataViewModel().getData());

        //TODO:Need to bind data from local database
    }
}
