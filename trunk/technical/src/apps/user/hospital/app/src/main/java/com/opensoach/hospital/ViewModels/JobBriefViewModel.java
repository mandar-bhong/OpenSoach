package com.opensoach.hospital.ViewModels;

import android.databinding.Bindable;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hospital.BR;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.View.JobConfigModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;

import java.lang.reflect.Type;
import java.text.SimpleDateFormat;

/**
 * Created by Mandar on 8/25/2017.
 */

public class JobBriefViewModel extends BaseViewModel {

    private JobBriefViewModel dataContext;
    private String job = "abc";
    private String part;
    private String customer;
    private String targatedDate;
    private String actualStartDate;
    private String quantity;
    private String status;
    private int statusColor;

    private DBJobCardTableRowModel dbJobCardTableRowModel;
    private DBEnggPartTableRowModel dbEnggPartTableRowModel;
    private JobConfigModel jobConfigModel;


    public DBJobCardTableRowModel getDbJobCardTableRowModel() {
        return dbJobCardTableRowModel;
    }

    public void setDbJobCardTableRowModel(DBJobCardTableRowModel dbJobCardTableRowModel) {
        this.dbJobCardTableRowModel = dbJobCardTableRowModel;

        jobConfigModel = new JobConfigModel();

        try {
            if (this.dbJobCardTableRowModel != null) {
                Gson gson = new Gson();
                Type jobConfigType = new TypeToken<JobConfigModel>() {
                }.getType();
                jobConfigModel = gson.fromJson(this.dbJobCardTableRowModel.getJobConfig(), jobConfigType);
            }
        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }

        notifyPropertyChanged(BR._all);
    }

    public DBEnggPartTableRowModel getDbEnggPartTableRowModel() {
        return dbEnggPartTableRowModel;
    }

    public void setDbEnggPartTableRowModel(DBEnggPartTableRowModel dbEnggPartTableRowModel) {
        this.dbEnggPartTableRowModel = dbEnggPartTableRowModel;
        notifyPropertyChanged(BR._all);
    }

    public String getJob() {
        return dbJobCardTableRowModel.getJobCode();
    }

    public String getTargatedDate() {
        SimpleDateFormat startDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String startDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getStartDate());

        SimpleDateFormat endDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String endDateString = endDateFormatter.format(this.dbJobCardTableRowModel.getEndDate());

        return startDateString +" - "+endDateString;
    }

    public String getTargatedStartDate() {
        SimpleDateFormat startDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String startDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getStartDate());
        return startDateString;
    }

    public String getTargatedEndDate() {
        SimpleDateFormat endDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String endDateString = endDateFormatter.format(this.dbJobCardTableRowModel.getEndDate());
        return endDateString;
    }

    public String getCustomer() {
        return dbJobCardTableRowModel.getCustomer();
    }

    public String getPart() {
        return dbEnggPartTableRowModel.getPartNo() +", "+dbEnggPartTableRowModel.getPartRevision();
    }

    public String getActualStartDate() {
        if (dbJobCardTableRowModel.getState() == 0)
            return "----";

        SimpleDateFormat startDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String actualStartDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getActualStartDate());
        return actualStartDateString;
    }

    @Bindable
    public String getQuantity() {
        return  dbJobCardTableRowModel.getCompletedCount()+"/"+ dbJobCardTableRowModel.getPartCount().toString();
    }

    @Bindable
    public int getQuantityProgress() {
        return (int) ((dbJobCardTableRowModel.getCompletedCount() * 100) / dbJobCardTableRowModel.getPartCount());
    }


    @Bindable()
    public  String getWorkID(){
        if (jobConfigModel.workId == null)
            return  "";

        return jobConfigModel.workId.toString() ;
    }

    @Bindable()
    public  String getWorkType(){
        if (jobConfigModel.worktype == null)
            return  "";

        return jobConfigModel.worktype;
    }

    @Bindable()
    public  String getAddress(){
        if (jobConfigModel.address == null)
            return  "";

        return jobConfigModel.address;
    }


    @Bindable()
    public boolean getInprogressVisibility() {

        switch (dbJobCardTableRowModel.getState()) {
            case Constants.JOB_STATUS_INPROGRESS:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getAbortedVisibility() {

        switch (dbJobCardTableRowModel.getState()) {
            case Constants.JOB_STATUS_ABORTED:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getPendingVisibility() {

        switch (dbJobCardTableRowModel.getState()) {
            case Constants.JOB_STATUS_PENDING://pending
                return true;
            default:
                return false;
        }
    }


    @Bindable()
    public boolean getDropedVisibility() {
        switch (dbJobCardTableRowModel.getState()) {
            case Constants.JOB_STATUS_DROPED:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getCompletedVisibility() {
        switch (dbJobCardTableRowModel.getState()) {
            case Constants.JOB_STATUS_COMPLETED:
                return true;
            default:
                return false;
        }
    }


    public String getStatus() {
        return CommonHelper.GetJobStatusText(dbJobCardTableRowModel);
    }

    public int getStatusTextColor() {

        int color = CommonHelper.GetJobStatusTextColor(dbJobCardTableRowModel);

        return ContextActivity.getResources().getColor(color);
    }

    public int getStatusColor() {

        int color = R.color.colorGray;

        //int jobStatus = CommonHelper.GetJobStatus(dbJobCardTableRowModel);

        color =  CommonHelper.GetJobStatusTextColor(dbJobCardTableRowModel);
//
//        switch (jobStatus){
//
//            case ApplicationConstants.JOB_STATE_NOT_STARTED:
//            case ApplicationConstants.JOB_STATE_NOT_STARTED_START_DELAYED:
//            case ApplicationConstants.JOB_STATE_NOT_STARTED_END_DELAYED:
//                color = R.color.colorGray;
//                break;
//
//            case ApplicationConstants.JOB_STATE_STARTED:
//            case ApplicationConstants.JOB_STATE_STARTED_START_DELAYED:
//            case ApplicationConstants.JOB_STATE_STARTED_END_DELAYED:
//                color = R.color.color_status_inprogress;
//                color = R.color.color_status_inprogress;
//                break;
//
//            case ApplicationConstants.JOB_STATE_COMPLETED:
//                color = R.color.color_status_inprogress;
//                break;
//        }

        return ContextActivity.getResources().getColor(color);
    }





}
