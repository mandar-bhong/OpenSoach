package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.Views.Adapter.JobServiceDataAdapter;

import java.util.ArrayList;
import java.util.List;

public class JobSummaryViewModel extends  BaseViewModel{

    private String firstName;
    private String lastName;
    private String  mobileNo;
    private String kmRuns;
    private String petrolLevel;
    private JobServiceDataAdapter jobServiceDataAdapter;
    private List<JobServiceItemViewModel> data;
    private JobServiceItemViewModel selectedJob;



    public List<JobServiceItemViewModel> getData() {
        return data;
    }

    public void setData(List<JobServiceItemViewModel> data) {
        this.data = data;
    }

    public JobServiceDataAdapter getJobServiceDataAdapter() {
        return jobServiceDataAdapter;
    }

    public void setJobServiceDataAdapter(JobServiceDataAdapter jobServiceDataAdapter) {
        this.jobServiceDataAdapter = jobServiceDataAdapter;
    }

    public String getFirstName() {
        return firstName;
    }
    @Bindable
    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }
    @Bindable
    public String getLastName() {
        return lastName;
    }

    @Bindable
    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getMobileNo() {
        return mobileNo;
    }

    @Bindable
    public void setMobileNo(String mobileNo) {
        this.mobileNo = mobileNo;
    }


    public String getKmRuns() {
        return kmRuns;
    }

    @Bindable
    public void setKmRuns(String kmRuns) {
        this.kmRuns = kmRuns;
    }

    public String getPetrolLevel() {
        return petrolLevel;
    }

   @Bindable
    public void setPetrolLevel(String petrolLevel) {
        this.petrolLevel = petrolLevel;
    }

    public void setData(ArrayList<JobSummaryViewModel> list) {
//        this.list = list;
    }
}
