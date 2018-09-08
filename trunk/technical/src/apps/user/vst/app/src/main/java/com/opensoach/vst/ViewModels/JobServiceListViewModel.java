package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Views.Adapter.JobServiceDataAdapter;

import java.util.List;

public class JobServiceListViewModel extends BaseViewModel{

    private JobServiceDataAdapter jobServiceDataAdapter;
    private List<JobServiceItemViewModel> data;
    private JobServiceItemViewModel selectedJob;
    private  int displayMode;

    public JobServiceListViewModel() {
        this.jobServiceDataAdapter = new JobServiceDataAdapter();
        displayMode = ApplicationConstants.DISPLAY_MODE_JOB_EXECUTION;
    }

    public JobServiceDataAdapter getJobServiceDataAdapter() {
        return jobServiceDataAdapter;
    }

    public List<JobServiceItemViewModel> getData() {
        return data;
    }

    public void setData(List<JobServiceItemViewModel> data) {
        this.data = data;
    }

    public int getDisplayMode() {
        return displayMode;
    }

    public void setDisplayMode(int displayMode) {
        this.displayMode = displayMode;
    }

    @Bindable
    public boolean getShowSummaryButton(){

        if ( getDisplayMode() == ApplicationConstants.DISPLAY_MODE_JOB_CREATION_EDIT){
            return true;
        }else {
            return false;
        }

    }


}
