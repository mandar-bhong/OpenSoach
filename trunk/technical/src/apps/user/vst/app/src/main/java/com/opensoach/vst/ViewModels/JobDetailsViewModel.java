package com.opensoach.vst.ViewModels;

public class JobDetailsViewModel extends BaseViewModel {

    private TokenSelectionViewModel tokenSelectionViewModel;
    private TokenItemViewModel tokenItemViewModel;

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

    public void setTokenItemViewModel(TokenItemViewModel tokenItemViewModel) {
        this.tokenItemViewModel = tokenItemViewModel;
    }


    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getMobileNo() {
        return mobileNo;
    }

    public void setMobileNo(String mobileNo) {
        this.mobileNo = mobileNo;
    }

    public String getKmRuns() {
        return kmRuns;
    }

    public void setKmRuns(String kmRuns) {
        this.kmRuns = kmRuns;
    }

    public String getPetrolLevel() {
        return petrolLevel;
    }

    public void setPetrolLevel(String petrolLevel) {
        this.petrolLevel = petrolLevel;
    }
}
