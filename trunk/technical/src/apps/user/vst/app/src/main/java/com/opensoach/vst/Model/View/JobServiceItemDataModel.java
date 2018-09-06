package com.opensoach.vst.Model.View;

import android.support.annotation.Nullable;

public class JobServiceItemDataModel {

    private boolean isCompleted;
    private String fname ;

    public JobServiceItemDataModel() {
    }


    public boolean isCompleted() {
        return isCompleted;
    }

    public void setIsCompleted(boolean completed) {
        isCompleted = completed;
    }

    @Nullable
    public String getFname() {
        return fname;
    }

    public void setFname(@Nullable String fname) {
        this.fname = fname;
    }


}
