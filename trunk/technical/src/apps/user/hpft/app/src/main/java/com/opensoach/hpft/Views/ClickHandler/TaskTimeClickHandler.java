package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.ViewModels.TaskDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeDataViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeItemViewModel;
import com.opensoach.hpft.Views.Activity.TaskDetailsActivity;

import java.util.List;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeClickHandler {

    public void onClick(View view, TaskTimeItemViewModel vm) {

        AppRepo.getInstance().getActiveCard().getTaskDetails().setSelectedItem(vm);

        List<TaskItemDataModel> tasks = ((TaskTimeDataViewModel) vm.Parent).getTaskDataViewModel().getData();


        DBServiceTaskDataTableRowModel dbServiceTaskDataTableRowModel = new DBServiceTaskDataTableRowModel();
        dbServiceTaskDataTableRowModel.setServConfID(AppRepo.getInstance().getActiveCard().getServConfID());
        dbServiceTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().getSerInID());
        dbServiceTaskDataTableRowModel.setLocationId(AppRepo.getInstance().getActiveCard().getLocationID());
        dbServiceTaskDataTableRowModel.setTime(AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime());

        List<DBServiceTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBServiceTaskDataTableQueryModel(), dbServiceTaskDataTableRowModel, DBServiceTaskDataTableQueryModel.SELECT_LOCATION_TIME_FILTER);

        if (dbRows.size() > 0) {
            DBServiceTaskDataTableRowModel dbRow = dbRows.get(0);

            TypeToken<List<TaskItemDataModel>> taskItemTypeToken = new TypeToken<List<TaskItemDataModel>>() {
            };

            List<TaskItemDataModel> dbCompletedTasks = new Gson().fromJson(dbRow.getData(), taskItemTypeToken.getType());

            for (TaskItemDataModel model : dbCompletedTasks) {

                for (TaskItemDataModel userModel : tasks) {
                    if (model.getTitle().equals( userModel.getTitle())) {
                        userModel.setIsCompleted(model.getIsCompleted());
                    }
                }
            }
        }else {
            for (TaskItemDataModel userModel : tasks) {
                userModel.setIsCompleted(false);
            }
        }

        ((TaskDetailsViewModel) ((TaskTimeDataViewModel) vm.Parent).Parent).getTaskDataAdapter().updateData(tasks);

    }
}
