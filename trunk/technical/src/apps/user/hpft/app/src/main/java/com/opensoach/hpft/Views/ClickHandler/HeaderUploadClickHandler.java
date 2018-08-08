package com.opensoach.hpft.Views.ClickHandler;

import android.content.Intent;
import android.view.View;
import android.widget.Toast;

import com.google.gson.Gson;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Model.Communication.CommandRequest;
import com.opensoach.hpft.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.hpft.Model.DB.DBAuthCodeTableRowModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableQueryModel;
import com.opensoach.hpft.Model.DB.DBServiceTaskDataTableRowModel;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.PacketGenerator.TaskDataPacketGenerator;
import com.opensoach.hpft.R;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.HeaderViewModel;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.Views.Activity.CardDetailsActivity;
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

        ArrayList<TaskItemDataModel> items = new ArrayList<>();
        items.addAll(AppRepo.getInstance().getSelectedTaskDataViewModels());

        //CommandRequest req = new TaskDataPacketGenerator().GenerateRequest(1,items);

        Date startTime = AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem().getTaskTimeDataModel().getStartTime();
        int locationID = AppRepo.getInstance().getActiveCard().getLocationID();

        DBServiceTaskDataTableRowModel dbServiceTaskDataTableRowModel = new DBServiceTaskDataTableRowModel();
        dbServiceTaskDataTableRowModel.setServConfID(AppRepo.getInstance().getActiveCard().ServConfID);
        dbServiceTaskDataTableRowModel.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
        dbServiceTaskDataTableRowModel.setLocationId(locationID);
        dbServiceTaskDataTableRowModel.setTime(startTime);

        List<DBServiceTaskDataTableRowModel> dbRows = DatabaseManager.SelectByFilter(new DBServiceTaskDataTableQueryModel(), dbServiceTaskDataTableRowModel, DBServiceTaskDataTableQueryModel.SELECT_LOCATION_TIME_FILTER);

        if (dbRows.size() > 0) {
            DBServiceTaskDataTableRowModel dbRow = dbRows.get(0);
            dbRow.setData(new Gson().toJson(items));

            DatabaseManager.UpdateRow(new DBServiceTaskDataTableQueryModel(), dbRow, DBServiceTaskDataTableQueryModel.SELECT_LOCATION_TIME_FILTER);

        } else {
            DBServiceTaskDataTableRowModel row = new DBServiceTaskDataTableRowModel();

            row.setServConfID(AppRepo.getInstance().getActiveCard().ServConfID);
            row.setSerInID(AppRepo.getInstance().getActiveCard().SerInID);
            row.setLocationId(locationID);
            row.setTime(startTime);
            row.setData(new Gson().toJson(items));

            DatabaseManager.InsertRow(row);
        }
    }
}
