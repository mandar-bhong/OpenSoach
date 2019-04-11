package com.opensoach.vst.ViewModels;

public class JobExeDetailsViewModel extends  BaseViewModel {

    private TokenItemViewModel tokenItemViewModel;
    private TokenSelectionViewModel tokenSelectionViewModel;

    public TokenItemViewModel getTokenItemViewModel() {
        return tokenItemViewModel;
    }

    public void setTokenItemViewModel(TokenItemViewModel tokenItemViewModel) {
        this.tokenItemViewModel = tokenItemViewModel;
    }

    public TokenSelectionViewModel getTokenSelectionViewModel() {
        return tokenSelectionViewModel;
    }

    public void setTokenSelectionViewModel(TokenSelectionViewModel tokenSelectionViewModel) {
        this.tokenSelectionViewModel = tokenSelectionViewModel;
    }
}
