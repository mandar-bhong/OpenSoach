package com.opensoach.hpft.Views.ClickHandler;

import android.view.View;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.ViewModels.TaskDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeDataViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeItemViewModel;

import java.util.List;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeClickHandler {

    public void onClick(View view, TaskTimeItemViewModel vm) {

        AppRepo.getInstance().getActiveCard().getTaskDetails().setSelectedItem(vm);

        List<TaskItemDataModel> tasks = ((TaskTimeDataViewModel) vm.Parent).getTaskDataViewModel().getData();


        DBTaskDataTableRowModel dbServiceTaskDataTableRowModel = new DBTaskDataTableRowModel();
        //dbServiceTaskDataTableRowModel.setServConfID(AppRepo.getInstance().getActiveCard().getServConfID());
        dbServiceTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().getSerInID());
        dbServiceTaskDataTableRowModel.setLocationId(AppRepo.getInstance().getActiveCard().getLocationID());
        dbServiceTaskDataTableRowModel.setTaskSlotStartTime(AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime());

        List<DBTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBTaskDataTableQueryModel(), dbServiceTaskDataTableRowModel, DBTaskDataTableQueryModel.SELECT_ID_FILTER);

        for (TaskItemDataModel userModel : tasks) {
            userModel.setIsCompleted(false);
        }

        for(DBTaskDataTableRowModel model : dbRows){
            for (TaskItemDataModel userModel : tasks) {
                if (model.getTitle().equals( userModel.getTitle())) {
                    userModel.setIsCompleted(true);
                    userModel.setServerSyncCompleted(model.isSynced());
                    userModel.setComment(model.getComment());
                    userModel.setObservationValue(model.getValue());
                }
            }
        }

        if (dbRows.size() == 0) {
            for (TaskItemDataModel userModel : tasks) {
                userModel.setIsCompleted(false);
            }
        }

        ((TaskDetailsViewModel) ((TaskTimeDataViewModel) vm.Parent).Parent).getTaskDataAdapter().updateData(tasks);

    }
}
