package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Constants.ApplicationConstants;

public class JobServiceItemViewModel extends BaseViewModel {

    private String taskName;
    private String comment;
    private String cost;


    public String getTaskName() {
        return taskName;
    }


    public String getComment() {
        return comment;
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

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_EXECUTION){
            return true;
        }else {
            return false;
        }
    }



    @Bindable
    public boolean getShowMoveRightIcon(){

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_EXECUTION){
            return true;
        }else {
            return false;
        }
    }


    @Bindable
    public boolean getShowDelete(){

        if ( ((JobServiceListViewModel)Parent).getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_CREATION_EDIT){
            return true;
        }else {
            return false;
        }

    }

}
