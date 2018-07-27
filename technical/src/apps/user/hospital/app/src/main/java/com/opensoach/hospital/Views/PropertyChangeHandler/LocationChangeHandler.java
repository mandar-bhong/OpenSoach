package com.opensoach.hospital.Views.PropertyChangeHandler;

import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.util.Log;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.DB.DBEnggPartTableQueryModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.MainViewModel;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 9/9/2017.
 */

public class LocationChangeHandler implements PropertyChangeListener {

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.CurrentLocationIdPropName: {


                Handler uiHandler = new Handler(Looper.getMainLooper()) {
                    @Override
                    public void handleMessage(Message message) {

                        MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    }
                };

                Log.d("Connection State: ", evt.getNewValue().toString());

                MainViewModel.getInstance().GridViewModel.getItemsSource().clear();

                DBJobCardTableRowModel dbJobCardTableRowModel = new  DBJobCardTableRowModel();
                dbJobCardTableRowModel.setLocationId((int) evt.getNewValue());

                List<DBJobCardTableRowModel> jobCards = DatabaseManager.SelectByFilter(new DBJobCardTableQueryModel(),dbJobCardTableRowModel, DBJobCardTableQueryModel.SELECT_LOCATION_ID_FILTER);

                List<JobBriefViewModel> jobBriefViewModels = new ArrayList<>() ;

                for (DBJobCardTableRowModel jobCardTableRowModel : jobCards){

                    JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
                    jobBriefViewModels.add(jobBriefViewModel);
                    jobBriefViewModel.ContextActivity = MainViewModel.getInstance().ContextActivity;
                    jobBriefViewModel.setDbJobCardTableRowModel(jobCardTableRowModel);


                    DBEnggPartTableRowModel dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
                    dbEnggPartTableRowModel.setPartId(jobCardTableRowModel.getPartId());

                    List<DBEnggPartTableRowModel> parts =  DatabaseManager.SelectByFilter(new DBEnggPartTableQueryModel(),dbEnggPartTableRowModel,DBEnggPartTableQueryModel.SELECT_ID_FILTER);

                    if(parts.size() >0){
                        jobBriefViewModel.setDbEnggPartTableRowModel(parts.get(0));
                    }

                    MainViewModel.getInstance().GridViewModel.getItemsSource().add(jobBriefViewModel);
                }

                Message msg = uiHandler.obtainMessage();
                Bundle b = new Bundle();
                //b.putBoolean("ConnectionState", (boolean) evt.getNewValue());
                //msg.setData(b);
                uiHandler.sendMessage(msg);

                //SharedPreferencesHelper.getInstance(MainViewModel.getInstance().ContextActivity).updateSharedPreference(Constants.KEY_LOCATION_ID,evt.getNewValue().toString());

            }
            break;
        }


        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:

                break;
        }
    }
}
