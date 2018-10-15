package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Model.Communication.PacketServiceTaskItemDataModel;
import com.opensoach.vst.PacketGenerator.JobCreatedDataPacketGenerator;

public class JobServiceViewModel extends BaseViewModel {

    private TokenItemViewModel tokenItemViewModel;
    private TokenSelectionViewModel tokenSelectionViewModel;
    private JobServiceListViewModel jobServiceListViewModel;
    private JobServiceDetailsViewModel jobServiceDetailsViewModel;
    private JobServiceItemViewModel jobServiceItemViewModel;
    private  int clickMode;


    public JobServiceViewModel() {
        clickMode = ApplicationConstants.CLIK_MODE_ITEM_EDITABLE;
    }

//    public int getClickMode() {
//        return clickMode;
//    }
//
//    public void setClickMode(int clickMode) {
//        this.clickMode = clickMode;
//    }



    public TokenSelectionViewModel getTokenSelectionViewModel() {
        return tokenSelectionViewModel;
    }

    public void setTokenSelectionViewModel(TokenSelectionViewModel tokenSelectionViewModel) {
        this.tokenSelectionViewModel = tokenSelectionViewModel;
    }


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

    public JobServiceItemViewModel getJobServiceItemViewModel() {
        return jobServiceItemViewModel;
    }

    public void setJobServiceItemViewModel(JobServiceItemViewModel jobServiceItemViewModel) {
        this.jobServiceItemViewModel = jobServiceItemViewModel;
    }

    public String getCostSum() {
        int totalTentCost = 0;

        for(JobServiceItemViewModel jobServiceItemViewModel : getJobServiceListViewModel().getData()){
            PacketServiceTaskItemDataModel item = new PacketServiceTaskItemDataModel();
            item.Cost = jobServiceItemViewModel.getCost();
            if (!(item.Cost  == null || item.Cost  == "")){
                totalTentCost = totalTentCost + Integer.parseInt(item.Cost );
            }
        }
        return String.valueOf(totalTentCost);
    }

}
