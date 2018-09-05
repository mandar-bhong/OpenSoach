package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

public class JobServiceItemViewModel extends BaseViewModel {

    private String taskName;


    public String getTaskName() {
        return taskName;
    }

    @Bindable
    public void setTaskName(String taskName) {
        this.taskName = taskName;
//        notifyPropertyChanged(BR.task);
    }
}
