package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;

public class TokenSelectionViewModel extends BaseViewModel{

    private boolean isJobCreate;
    private TokenListViewModel tokenListViewModel;


    public TokenSelectionViewModel(){
        isJobCreate =false;
    }


    public boolean isJobCreate() {
        return isJobCreate;
    }

    @Bindable
    public void setJobCreate(boolean jobCreate) {
        isJobCreate = jobCreate;

    }


    public TokenListViewModel getTokenListViewModel() {
        return tokenListViewModel;
    }

    public void setTokenListViewModel(TokenListViewModel tokenListViewModel) {
        this.tokenListViewModel = tokenListViewModel;
    }

    @Bindable
    public boolean getCreateTokenBtn(){
        if( (!(AppRepo.getInstance().getCurrentRunningMode() == ApplicationConstants.AppRunningMode.JobExecution))) {
            return true;
        }else {
            return false;
        }
    }
}
