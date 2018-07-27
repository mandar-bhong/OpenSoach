package com.opensoach.hospital.Helper;

import android.content.Context;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Model.View.JobToolModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.Utility.JSON.GsonUTCDateAdapter;

import java.lang.reflect.Type;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

/**
 * Created by Mandar on 8/26/2017.
 */

public class CommonHelper {

    public  static String GetPacketJSON (PacketModel model){
        Gson gson = new GsonBuilder().setDateFormat("yyyy-MM-dd'T'HH:mm:ss'Z'").registerTypeAdapter(Date.class,new GsonUTCDateAdapter()).create();
        return   gson.toJson(model);
    }

    public static String ConvertToolJSONTOText(String jsonTools) {

        String toolDisplayData = "";
        try {

            Type collectionType = new TypeToken<List<JobToolModel>>(){}.getType();

            List<JobToolModel> tools= new Gson().fromJson(jsonTools, collectionType);

            for (JobToolModel toolModel:tools ) {
                toolDisplayData = toolDisplayData.concat("&#187;&nbsp;&nbsp;"+toolModel.Name +" - "+ toolModel.Description +"<br/>");
            }

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }

        return toolDisplayData;
    }

    public static Date GetDateOnly(Date date){
            Calendar cal = Calendar.getInstance();
            cal.setTime(date);
            cal.set(Calendar.HOUR_OF_DAY, 0);
            cal.set(Calendar.MINUTE, 0);
            cal.set(Calendar.SECOND, 0);
            cal.set(Calendar.MILLISECOND, 0);
            return cal.getTime();
    }

    public static int GetJobStatus(DBJobCardTableRowModel dbJobCardTableRowModel){

        switch (dbJobCardTableRowModel.getState()) {
            case 0: {//Not Started
                if(dbJobCardTableRowModel.getEndDate().compareTo( new Date()) == -1){
                    return ApplicationConstants.JOB_STATE_NOT_STARTED_END_DELAYED;
                }else if(dbJobCardTableRowModel.getStartDate().compareTo( new Date()) == -1){
                    return ApplicationConstants.JOB_STATE_NOT_STARTED_START_DELAYED;
                }else{
                    return ApplicationConstants.JOB_STATE_NOT_STARTED;
                }
            }
            case 1: {
                if(dbJobCardTableRowModel.getEndDate().compareTo( new Date()) == -1){
                    return ApplicationConstants.JOB_STATE_STARTED_END_DELAYED;
                }else if(dbJobCardTableRowModel.getStartDate().compareTo(dbJobCardTableRowModel.getActualStartDate()) == -1){
                    return ApplicationConstants.JOB_STATE_STARTED_START_DELAYED;
                }else{
                    return ApplicationConstants.JOB_STATE_STARTED;
                }
            }
            case Constants.JOB_STATE_COMPLETED:
                return ApplicationConstants.JOB_STATE_COMPLETED;
            case Constants.JOB_STATE_DROPPED:
                return ApplicationConstants.JOB_STATE_DROPPED;
            case Constants.JOB_STATE_ABORTED:
                return ApplicationConstants.JOB_STATE_ABORTED;
        }

        return 0;
    }

    public static String GetJobStatusText(DBJobCardTableRowModel dbJobCardTableRowModel) {

        String statusText = "";

        int jobStatus = CommonHelper.GetJobStatus(dbJobCardTableRowModel);

        switch (jobStatus){

            case ApplicationConstants.JOB_STATE_NOT_STARTED:
                statusText = "Pending";
                break;
            case ApplicationConstants.JOB_STATE_NOT_STARTED_START_DELAYED:
                //statusText = "Start Delayed";
                //break;
            case ApplicationConstants.JOB_STATE_NOT_STARTED_END_DELAYED:
                //statusText = "End Delayed";
                statusText = "Pending";
                break;


            case ApplicationConstants.JOB_STATE_STARTED_START_DELAYED:
                //statusText = "Start Delayed";
                //break;
            case ApplicationConstants.JOB_STATE_STARTED_END_DELAYED:
                //statusText = "End Delayed";
                //break;
            case ApplicationConstants.JOB_STATE_STARTED:
                statusText = "In Progress";
                break;
            case ApplicationConstants.JOB_STATE_COMPLETED:
                statusText = "Completed";
                break;
            case ApplicationConstants.JOB_STATE_ABORTED:
                statusText = "Aborted";
                break;
            case ApplicationConstants.JOB_STATE_DROPPED:
                statusText = "Dropped";
                break;
        }

        return statusText;
    }

    public static int GetJobStatusTextColor(DBJobCardTableRowModel dbJobCardTableRowModel) {

        int color = R.color.colorJobStatusScheduled;

        int jobStatus = CommonHelper.GetJobStatus(dbJobCardTableRowModel);

        switch (jobStatus){
            case ApplicationConstants.JOB_STATE_NOT_STARTED:
                color = R.color.colorJobStatusScheduled;
                break;
            case ApplicationConstants.JOB_STATE_NOT_STARTED_START_DELAYED:
            case ApplicationConstants.JOB_STATE_STARTED_START_DELAYED:
                color = R.color.color_status_delayed_started;
                break;
            case ApplicationConstants.JOB_STATE_NOT_STARTED_END_DELAYED:
            case ApplicationConstants.JOB_STATE_STARTED_END_DELAYED:
                color = R.color.color_status_delayed_end;
                break;

            case ApplicationConstants.JOB_STATE_STARTED:
                color = R.color.color_status_inprogress;
                break;
//            case ApplicationConstants.JOB_STATE_STARTED_START_DELAYED:
//
//                color = R.color.color_status_delayed_started;
//                break;

            case ApplicationConstants.JOB_STATE_COMPLETED:
                color = R.color.color_status_completed;
                break;
        }

        return color;
    }

    public static void SetCurrentLocationID(Context mContext){
        String locationIdKey = SharedPreferencesHelper.getInstance(mContext).getDataFromSharedPreference(Constants.KEY_LOCATION_ID);

        if(locationIdKey != null && locationIdKey != ""){
            int locationId = Integer.parseInt(locationIdKey);

            boolean isLocationFound = false;

            for(DBLocationTableRowModel locationModel: AppRepo.getInstance().getLocationList()){
                if(locationModel.getLocationId() == locationId ){
                    isLocationFound = true;
                    AppRepo.getInstance().setCurrentLocationId(locationId);
                    break;
                }
            }

            if(!isLocationFound){
                DBLocationTableRowModel dbLocationTableRowModel = AppRepo.getInstance().getLocationList().get(0);
                AppRepo.getInstance().setCurrentLocationId(dbLocationTableRowModel.getLocationId());
                SharedPreferencesHelper.getInstance(mContext).updateSharedPreference(Constants.KEY_LOCATION_ID,String.valueOf(locationId));
            }

        }else{
            DBLocationTableRowModel dbLocationTableRowModel = AppRepo.getInstance().getLocationList().get(0);
            AppRepo.getInstance().setCurrentLocationId(dbLocationTableRowModel.getLocationId());
            SharedPreferencesHelper.getInstance(mContext).updateSharedPreference(Constants.KEY_LOCATION_ID,String.valueOf(dbLocationTableRowModel.getLocationId()));
        }
    }


    public static boolean IsOperatorCodeValid(String opCode){
        return AppRepo.getInstance().getAuthCodeList().contains(opCode);
    }


}
