package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

public class JobServiceDetailsViewModel extends BaseViewModel {


    private TokenItemViewModel tokenItemViewModel;
    private TokenSelectionViewModel tokenSelectionViewModel;

    private JobCustomerDetailsViewModel jobCustomerDetailsViewModel;

    private String firstName;
    private String lastName;
    private String mobileNo;
    private String kmRuns;
    private String petrolLevel;



    public TokenSelectionViewModel getTokenSelectionViewModel() {
        return tokenSelectionViewModel;
    }

    public void setTokenSelectionViewModel(TokenSelectionViewModel tokenSelectionViewModel) {
        this.tokenSelectionViewModel = tokenSelectionViewModel;
    }

    public TokenItemViewModel getTokenItemViewModel() {
        return tokenItemViewModel;
    }

    @Bindable
    public void setTokenItemViewModel(TokenItemViewModel tokenItemViewModel) {
        this.tokenItemViewModel = tokenItemViewModel;
    }

    public JobCustomerDetailsViewModel getJobCustomerDetailsViewModel() {
        return jobCustomerDetailsViewModel;
    }

    public void setJobCustomerDetailsViewModel(JobCustomerDetailsViewModel jobCustomerDetailsViewModel) {
        this.jobCustomerDetailsViewModel = jobCustomerDetailsViewModel;
    }

    public String getFirstName() {
        return firstName;
    }

    @Bindable
    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    @Bindable
    public String getMobileNo() {
        return mobileNo;
    }

    public void setMobileNo(String mobileNo) {
        this.mobileNo = mobileNo;
    }

    @Bindable
    public void setLastName(String lastName) {
        this.lastName = lastName;
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

}
