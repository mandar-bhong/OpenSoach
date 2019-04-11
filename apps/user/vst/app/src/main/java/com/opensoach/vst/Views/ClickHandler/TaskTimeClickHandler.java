package com.opensoach.vst.Views.ClickHandler;

import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.Model.View.TaskItemDataModel;
import com.opensoach.vst.ViewModels.TaskDetailsViewModel;
import com.opensoach.vst.ViewModels.TaskTimeDataViewModel;
import com.opensoach.vst.ViewModels.TaskTimeItemViewModel;

import java.util.List;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeClickHandler {

    public void onClick(View view, TaskTimeItemViewModel vm) {

        TaskTimeItemViewModel previouslySelected = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem();

        AppRepo.getInstance().getActiveCard().getTaskDetails().setSelectedItem(vm);

        vm.setActiveSelected();

        if (previouslySelected != null){
            previouslySelected.setActiveSelected();
        }

        AppRepo.getInstance().getSelectedTaskDataViewModels().clear();

        List<TaskItemDataModel> tasks = ((TaskTimeDataViewModel) vm.Parent).getTaskDataViewModel().getData();


        DBTaskDataTableRowModel dbServiceTaskDataTableRowModel = new DBTaskDataTableRowModel();
        //dbServiceTaskDataTableRowModel.setServConfID(AppRepo.getInstance().getActiveCard().getServConfID());
        dbServiceTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().getSerInID());
        dbServiceTaskDataTableRowModel.setLocationId(AppRepo.getInstance().getActiveCard().getLocationID());
        dbServiceTaskDataTableRowModel.setTaskSlotStartTime(AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime());

        List<DBTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBTaskDataTableQueryModel(), dbServiceTaskDataTableRowModel, DBTaskDataTableQueryModel.SELECT_ID_FILTER);

        for (TaskItemDataModel userModel : tasks) {
            userModel.setIsCompleted(false);
            userModel.setServerSyncCompleted(false);
            userModel.setComment(null);
            userModel.setObservationValue(null);
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

        ((TaskDetailsViewModel) ((TaskTimeDataViewModel) vm.Parent).Parent).getTaskDataAdapter().updateData(tasks);

    }
}
