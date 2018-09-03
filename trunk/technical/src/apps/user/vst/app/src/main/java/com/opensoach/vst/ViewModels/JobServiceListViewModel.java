package com.opensoach.vst.ViewModels;

import com.opensoach.vst.Views.Adapter.JobServiceDataAdapter;

import java.util.List;

public class JobServiceListViewModel extends BaseViewModel{

    private JobServiceDataAdapter jobServiceDataAdapter;
    private List<JobServiceItemViewModel> data;
    private JobServiceItemViewModel selectedJob;

    public JobServiceListViewModel() {
        this.jobServiceDataAdapter = new JobServiceDataAdapter();
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
}
