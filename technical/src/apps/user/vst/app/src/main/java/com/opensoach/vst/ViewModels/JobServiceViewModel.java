package com.opensoach.vst.ViewModels;

public class JobServiceViewModel extends BaseViewModel {

    private TokenItemViewModel tokenItemViewModel;
    private TokenSelectionViewModel tokenSelectionViewModel;
    //private JobServiceCreationViewModel jobServiceCreationViewModel;
    private JobServiceListViewModel jobServiceListViewModel;
    private JobServiceDetailsViewModel jobServiceDetailsViewModel;





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


    public JobServiceDetailsViewModel getJobServiceDetailsViewModel() {
        return jobServiceDetailsViewModel;
    }

    public void setJobServiceDetailsViewModel(JobServiceDetailsViewModel jobServiceDetailsViewModel) {
        this.jobServiceDetailsViewModel = jobServiceDetailsViewModel;
    }

    public JobServiceListViewModel getJobServiceListViewModel() {
        return jobServiceListViewModel;
    }

    public void setJobServiceListViewModel(JobServiceListViewModel jobServiceListViewModel) {
        this.jobServiceListViewModel = jobServiceListViewModel;
    }


    public TokenItemViewModel getTokenItemViewModel() {
        return tokenItemViewModel;
    }

    public void setTokenItemViewModel(TokenItemViewModel tokenItemViewModel) {
        this.tokenItemViewModel = tokenItemViewModel;
    }


}
