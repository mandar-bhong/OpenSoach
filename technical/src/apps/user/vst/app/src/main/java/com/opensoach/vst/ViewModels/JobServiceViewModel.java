package com.opensoach.vst.ViewModels;

public class JobServiceViewModel extends BaseViewModel {

    private TokenSelectionViewModel tokenSelectionViewModel;
    //private JobServiceCreationViewModel jobServiceCreationViewModel;
    private JobDetailsViewModel jobDetailsViewModel;
    private JobServiceListViewModel jobServiceListViewModel;


    public TokenSelectionViewModel getTokenSelectionViewModel() {
        return tokenSelectionViewModel;
    }

    public void setTokenSelectionViewModel(TokenSelectionViewModel tokenSelectionViewModel) {
        this.tokenSelectionViewModel = tokenSelectionViewModel;
    }


//    public JobServiceCreationViewModel getJobServiceCreationViewModel() {
//        return jobServiceCreationViewModel;
//    }
//
//    public void setJobServiceCreationViewModel(JobServiceCreationViewModel jobServiceCreationViewModel) {
//        this.jobServiceCreationViewModel = jobServiceCreationViewModel;
//    }


    public JobDetailsViewModel getJobDetailsViewModel() {
        return jobDetailsViewModel;
    }

    public void setJobDetailsViewModel(JobDetailsViewModel jobDetailsViewModel) {
        this.jobDetailsViewModel = jobDetailsViewModel;
    }


    public JobServiceListViewModel getJobServiceListViewModel() {
        return jobServiceListViewModel;
    }

    public void setJobServiceListViewModel(JobServiceListViewModel jobServiceListViewModel) {
        this.jobServiceListViewModel = jobServiceListViewModel;
    }
}
