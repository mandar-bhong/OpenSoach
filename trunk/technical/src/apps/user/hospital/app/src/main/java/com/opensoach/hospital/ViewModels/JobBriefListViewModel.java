package com.opensoach.hospital.ViewModels;

import android.app.FragmentManager;

import com.opensoach.hospital.Views.Miscellaneous.JobBriefViewAdaptor;

import java.util.List;

/**
 * Created by Mandar on 8/25/2017.
 */

public class JobBriefListViewModel extends BaseViewModel {

    public FragmentManager fragmentManager;
    public List<JobBriefViewModel> JobsBrief;
    public  Object DataContext;

    public JobBriefListViewModel JobsBriefList;

    public JobBriefViewAdaptor GridAdaptor;

    public JobBriefViewAdaptor getGridAdaptor() {
        return GridAdaptor;
    }

    public void setGridAdaptor(JobBriefViewAdaptor gridAdaptor) {
        GridAdaptor = gridAdaptor;
    }

    public JobBriefListViewModel(){
        GridAdaptor = new JobBriefViewAdaptor();
    }

    public List<JobBriefViewModel> getJobsBrief() {
        return JobsBrief;
    }

    public void setJobsBrief(List<JobBriefViewModel> jobsBrief) {
        this.JobsBrief = jobsBrief;
    }

    public JobBriefListViewModel getJobsBriefList() {
        return JobsBriefList;
    }

    public void setJobsBriefList(JobBriefListViewModel jobsBriefList) {
        JobsBriefList = jobsBriefList;
    }

    public Object getDataContext() {
        return DataContext;
    }

    public void setDataContext(Object dataContext) {
        DataContext = dataContext;
    }
}
