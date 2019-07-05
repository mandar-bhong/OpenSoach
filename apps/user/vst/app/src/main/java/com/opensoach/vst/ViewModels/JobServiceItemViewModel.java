package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;

public class JobServiceItemViewModel extends BaseViewModel {

    private String taskName;
    private String comment;
    private String cost;
    private Boolean isTaskCompleted;
    private String note;
    private Boolean isCheckCompleted;

    public Boolean getCheckCompleted() {
       return isCheckCompleted;

    }

    public void setCheckCompleted(Boolean checkCompleted) {
       isCheckCompleted = checkCompleted;
    }

    public JobServiceItemViewModel(){
        isTaskCompleted = false;
        isCheckCompleted = false;
    }

    public Boolean getTaskCompleted() {
        return isTaskCompleted;
    }

    public void setTaskCompleted(Boolean taskCompleted) {
       isTaskCompleted = taskCompleted;
    }

    public String getTaskName() {
        return taskName;
    }


    public String getComment() {
        return comment;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public void setComment(String comment) {
        this.comment = comment;
    }

    public String getCost() {
        return cost;
    }

    public void setCost(String cost) {
        this.cost = cost;
    }

    @Bindable
    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    @Bindable
    public boolean getShowcheckbox(){

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_EXECUTION || AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution && (!(((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_CREATION_SUMMARY ))  ){
            return true;
        }else {
            return false;
        }
    }



    @Bindable
    public boolean getShowMoveRightIcon(){

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_EXECUTION ){
            return true;
        }else {
            return false;
        }
    }


    @Bindable
    public boolean getShowDelete(){

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_CREATION_EDIT && (!(AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution ))){
            return true;
        }else {
            return false;
        }

    }

    @Bindable
    public boolean getTaskCostVisibility(){
        if (AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution){
            return false;
        }else{
            return true;
        }
    }

    @Bindable
    public boolean getTaskNoteVisibility(){
        if (AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution){
            return true;
        }else{
            return false;
        }
    }

    @Bindable
    public boolean getEditble(){
        if(AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution) {
            return false;
        }else {
            return true;
        }
    }


    @Bindable
    public boolean getTaskItemEdit(){
        if(((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_CREATION_SUMMARY){
            return true;
        }else {
            return false;
        }
    }



}
