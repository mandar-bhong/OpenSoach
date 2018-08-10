package com.opensoach.hpft.Views.ClickHandler;

import android.view.View;
import android.widget.Toast;

import com.google.gson.Gson;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.AppAction;
import com.opensoach.hpft.Manager.SendPacketManager;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
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

            for(TaskItemDataModel model : AppRepo.getInstance().getSelectedTaskDataViewModels()){
                if (model.getIsCompleted() == false){
                    items.add(model);
                }
            }

            Date entryTime = new Date();
            Date slotStartTime = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime();
            Date slotEndTime = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getEndTime();
            int locationID = AppRepo.getInstance().getActiveCard().getLocationID();

            DBServiceTaskDataTableRowModel dbServiceTaskDataTableRowModel = new DBServiceTaskDataTableRowModel();
            dbServiceTaskDataTableRowModel.setServConfID(AppRepo.getInstance().getActiveCard().ServConfID);
            dbServiceTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
            dbServiceTaskDataTableRowModel.setLocationId(locationID);
            dbServiceTaskDataTableRowModel.setEntryTime(entryTime);

            List<DBServiceTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBServiceTaskDataTableQueryModel(), dbServiceTaskDataTableRowModel, DBServiceTaskDataTableQueryModel.SELECT_LOCATION_TIME_FILTER);
            DBServiceTaskDataTableRowModel dbInsertRow = new DBServiceTaskDataTableRowModel();

            if (dbRows.size() > 0) {
                DBServiceTaskDataTableRowModel dbRow = dbRows.get(0);
                dbRow.setData(new Gson().toJson(items));

                DatabaseManager.UpdateRow(new DBServiceTaskDataTableQueryModel(), dbRow, DBServiceTaskDataTableQueryModel.SELECT_LOCATION_TIME_FILTER);

            } else {
                dbInsertRow.setServConfID(AppRepo.getInstance().getActiveCard().ServConfID);
                dbInsertRow.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
                dbInsertRow.setLocationId(locationID);
                dbInsertRow.setEntryTime(entryTime);
                dbInsertRow.setSlotStartTime(slotStartTime);
                dbInsertRow.setSlotEndTime(slotEndTime);
                dbInsertRow.setData(new Gson().toJson(items));

                DatabaseManager.InsertRow(dbInsertRow);
            }

            if (dbRows.size() == 0) {
                dbRows.add(dbInsertRow);
            }

            //SendPacketManager.Instance().send(AppAction.TASK_DATA, dbRows);

        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
            throw  ex;
        }
    }
}
