package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Model.View.JobServiceItemDataModel;

public class JobServiceItemViewModel extends BaseViewModel {

    private String fname;


    public String getFname() {
        return fname;
    }

    public void setFname(String fname) {
        this.fname = fname;
    }

}
