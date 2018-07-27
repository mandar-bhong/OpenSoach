package com.opensoach.hospital.ViewModels;

import android.databinding.Bindable;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.BR;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.CommonHelper;
import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableRowModel;
import com.opensoach.hospital.Model.View.JobConfigModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.lang.reflect.Type;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;

/**
 * Created by Mandar on 9/10/2017.
 */

public class JobBoardViewModel extends BaseViewModel implements PropertyChangeListener {


    private int jobCardId;
    private DBJobCardTableRowModel dbJobCardTableRowModel;
    private DBEnggPartTableRowModel dbEnggPartTableRowModel;
    private  List<DBPartDrawingTableRowModel> dbPartDrawingTableRowModels;
    private String jobState;
    private String jobConfig;
    private JobConfigModel jobConfigModel;

    public int getJobCardId() {
        return dbJobCardTableRowModel.getJobCardId();
    }

//    public void setJobCardId(int jobCardId) {
//        this.jobCardId = jobCardId;
//    }

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


        notifyPropertyChanged(BR.jobStatus);
        notifyPropertyChanged(BR.isJobStartButtonVisiable);
        notifyPropertyChanged(BR.jobQuantity);
        notifyPropertyChanged(BR.workType);
    }

    public DBEnggPartTableRowModel getDbEnggPartTableRowModel() {
        return dbEnggPartTableRowModel;
    }

    public void setDbEnggPartTableRowModel(DBEnggPartTableRowModel dbEnggPartTableRowModel) {
        this.dbEnggPartTableRowModel = dbEnggPartTableRowModel;
    }

    public List<DBPartDrawingTableRowModel> getDbPartDrawingTableRowModels() {
        return dbPartDrawingTableRowModels;
    }

    public void setDbPartDrawingTableRowModels(List<DBPartDrawingTableRowModel> dbPartDrawingTableRowModels) {
        this.dbPartDrawingTableRowModels = dbPartDrawingTableRowModels;
    }

    public void setJobState(int state){
        dbJobCardTableRowModel.setState(state);
        notifyPropertyChanged(BR.abortButtonEnabled);
        notifyPropertyChanged(BR.dropButtonEnabled);
        notifyPropertyChanged(BR.startButtonEnabled);
        notifyPropertyChanged(BR.completedButtonEnabled);
    }

    //region Bindable properties

    @Bindable
    public boolean getIsServerConnected() {
        return AppRepo.getInstance().IsServerConnected();
    }

    @Bindable
    public String getJobStatus() {
        if(dbJobCardTableRowModel == null) return "";

        return CommonHelper.GetJobStatusText(dbJobCardTableRowModel);
    }

    @Bindable
    public  int getJobStatusColor() {

        if(dbJobCardTableRowModel == null)
            return super.ContextActivity.getResources().getColor(R.color.colorGray);

        int color = CommonHelper.GetJobStatusTextColor(dbJobCardTableRowModel);

        return  super.ContextActivity.getResources().getColor(color);
    }


    @Bindable()
    public  String getJobDescription() {
        return dbJobCardTableRowModel.getJobCode() + " for " + dbEnggPartTableRowModel.getPartNo() + ", " + dbEnggPartTableRowModel.getPartRevision();
    }

    @Bindable()
    public  String getJobQuantity() {
        return  dbJobCardTableRowModel.getCompletedCount()+"/"+ dbJobCardTableRowModel.getPartCount().toString();
    }

    @Bindable()
    public  boolean getIsJobStartButtonVisiable() {

        if(dbJobCardTableRowModel.getState() == 0){ // Not started job
            return true;
        }else{
            return false;
        }
    }

    @Bindable()
    public  String getWorkID(){
        if (jobConfigModel.workId == null)
            return  "";

        return jobConfigModel.workId.toString();
    }

    @Bindable()
    public  String getStartDTM(){
        if (jobConfigModel.startdtm == null)
            return  "";

        return jobConfigModel.startdtm.toString();
    }

    @Bindable()
    public  String getEndDTM(){
        if (jobConfigModel.enddtm == null)
            return  "";

        return jobConfigModel.enddtm.toString();
    }

    @Bindable()
    public  String getPlannedDates(){

        SimpleDateFormat startDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String startDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getStartDate());

        SimpleDateFormat endDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        String endDateString = endDateFormatter.format(this.dbJobCardTableRowModel.getEndDate());

        return startDateString +" - "+endDateString;
    }

    @Bindable()
    public  String getActualDates(){

        Date minDate = new Date(0);
        SimpleDateFormat startDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        SimpleDateFormat endDateFormatter = new SimpleDateFormat(ApplicationConstants.UI_DATE_FORMAT);
        Date startDate = this.dbJobCardTableRowModel.getActualStartDate();
        if(startDate.getTime()==minDate.getTime())
        {
            startDate=null;
        }
        Date endDate=this.dbJobCardTableRowModel.getActualEndDate();
        if(endDate.getTime()==minDate.getTime())
        {
            endDate=null;
        }
//        if(startDate.getTime()==minDate.getTime()) {
//            return "match";
//        }
//        else
//        {
//            return "no match";
//        }

        if(startDate==null && endDate==null)
        {
            return "NA";
        }

        else if((startDate!=null) && endDate==null)
        {
            String startDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getActualStartDate());
            return startDateString + " - " + "NA";
        }

        else if(startDate==null && endDate!=null)
        {
            String endDateString = endDateFormatter.format(this.dbJobCardTableRowModel.getActualEndDate());
            return "NA" + " - " + endDateString;
        }

        else {

            String startDateString = startDateFormatter.format(this.dbJobCardTableRowModel.getActualStartDate());
            String endDateString = endDateFormatter.format(this.dbJobCardTableRowModel.getActualEndDate());
            return startDateString +" - "+endDateString;
        }
    }

    @Bindable()
    public  String getWorkType(){
        return jobConfigModel.worktype;
    }

    @Bindable()
    public  String getWorkSubType(){
        return jobConfigModel.worksubtype;
    }

    @Bindable()
    public  String getWorkAddress(){
        return jobConfigModel.address;
    }

    @Bindable()
    public boolean getStartButtonEnabled(){
        switch (dbJobCardTableRowModel.getState()){
            case Constants.JOB_STATUS_PENDING:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getCompletedButtonEnabled(){
        switch (dbJobCardTableRowModel.getState()){
            case Constants.JOB_STATUS_INPROGRESS:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getAbortButtonEnabled(){
        switch (dbJobCardTableRowModel.getState()){
            case Constants.JOB_STATUS_INPROGRESS:
                return true;
            default:
                return false;
        }
    }

    @Bindable()
    public boolean getDropButtonEnabled(){
        switch (dbJobCardTableRowModel.getState()){
            case Constants.JOB_STATUS_PENDING:
                return true;
            default:
                return false;
        }
    }



    //endregion Bindable properties

    //region Property Change Handler

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:
                notifyPropertyChanged(BR.isServerConnected);
                break;
        }
    }

    //endregion Property Change Handler
}
