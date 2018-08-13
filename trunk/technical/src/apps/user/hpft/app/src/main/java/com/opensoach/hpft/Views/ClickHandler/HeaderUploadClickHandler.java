package com.opensoach.hpft.Views.ClickHandler;

import android.view.View;
import android.widget.Toast;

import com.google.gson.Gson;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.AppAction;
import com.opensoach.hpft.Manager.SendPacketManager;
import com.opensoach.hpft.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.R;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.ViewModels.HeaderViewModel;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.Views.DialogHelper;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * Created by Mandar on 06-08-2018.
 */

public class HeaderUploadClickHandler {
    public void onClick(View view, HeaderViewModel vm) {

        DialogHelper.showSingleLineEditTextAlert(
                view.getContext(),
                view.getContext().getResources().getString(R.string.dialog_enter_auth_code),
                new DialogHelper.DialogCallBack() {

                    @Override
                    public boolean onSucess(String authText) {

                        if (AppRepo.getInstance().getAuthCodeList().contains(authText)) {
                            SendData(authText);
                            return true;
                        } else {
                            Toast.makeText(
                                    MainViewModel.getInstance().ContextActivity,
                                    MainViewModel.getInstance().ContextActivity.getResources().getString(R.string.invalid_auth_code),
                                    Toast.LENGTH_LONG).show();

                            return false;
                        }
                    }

                    @Override
                    public void onSucess(String strData1, String strData2) {

                    }

                    @Override
                    public void onSucess(String strData1, String strData2, String strData3) {

                    }
                });

    }


    private void SendData(String authText) {

        try {
            ArrayList<TaskItemDataModel> items = new ArrayList<>();

            for (TaskItemDataModel model : AppRepo.getInstance().getSelectedTaskDataViewModels()) {
                if (model.isServerSyncCompleted() == false) {
                    items.add(model);
                }
            }

            Date entryTime = new Date();
            Date slotStartTime = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime();
            Date slotEndTime = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getEndTime();
            int locationID = AppRepo.getInstance().getActiveCard().getLocationID();

            DBTaskDataTableRowModel dbTaskDataTableRowModel = new DBTaskDataTableRowModel();
            dbTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
            dbTaskDataTableRowModel.setLocationId(locationID);
            dbTaskDataTableRowModel.setTaskSlotStartTime(slotStartTime);

            List<DBTaskDataTableRowModel> updatedRows = new ArrayList<>();

            for (TaskItemDataModel model : items) {

                dbTaskDataTableRowModel.setTitle(model.getTitle());

                List<DBTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBTaskDataTableQueryModel(), dbTaskDataTableRowModel, DBTaskDataTableQueryModel.SELECT_TITLE_FILTER);

                if (dbRows.size() == 0) {
                    DBTaskDataTableRowModel dbInsertRow = new DBTaskDataTableRowModel();
                    dbInsertRow.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
                    dbInsertRow.setLocationId(locationID);
                    dbInsertRow.setTaskTime(entryTime);
                    dbInsertRow.setTaskSlotStartTime(slotStartTime);
                    dbInsertRow.setTaskSlotEndTime(slotEndTime);
                    dbInsertRow.setTitle(model.getTitle());
                    DatabaseManager.InsertRow(dbInsertRow);
                    updatedRows.add(dbInsertRow);
                }
            }


            SendPacketManager.Instance().send(AppAction.TASK_DATA, updatedRows);

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
            throw ex;
        }
    }
}
