package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.Views.Activity.CardDetailsActivity;
import com.opensoach.hpft.Views.Activity.TaskDetailsActivity;

/**
 * Created by Mandar on 02-08-2018.
 */

public class TaskDetailsClickHandler {

    public void onClick(View view, TaskItemViewModel vm) {
        //AppRepo.getInstance().setActiveCard(vm);

        Intent i = new Intent(view.getContext(), TaskDetailsActivity.class);
        view.getContext().startActivity(i);
    }
}
