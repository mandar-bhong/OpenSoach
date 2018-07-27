package com.opensoach.hospital.ViewModels;

import android.databinding.Bindable;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.BR;
import com.opensoach.hospital.Helper.AppHelper;
import com.opensoach.hospital.Views.ClickHandler.TestButtonClickHander;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;

/**
 * Created by Mandar on 8/25/2017.
 */

public class MainViewModel extends BaseViewModel implements PropertyChangeListener {

    private static MainViewModel singleton;
    HeaderViewModel headerViewModel;

    private ArrayList<String> locationList;

    private MainViewModel() {
    }

    public static MainViewModel getInstance() {
        if (singleton == null)
            singleton = new MainViewModel();
        return singleton;
    }

    public JobGridViewModel GridViewModel;

    public JobBriefGridViewModel JobBriefGridViewVM;


    public JobGridViewModel getGridViewModel() {
        return GridViewModel;
    }

    public void setGridViewModel(JobGridViewModel gridViewModel) {
        this.GridViewModel = gridViewModel;
    }

    public HeaderViewModel getHeaderViewModel() {
        return headerViewModel;
    }

    public void setHeaderViewModel(HeaderViewModel headerViewModel) {
        this.headerViewModel = headerViewModel;
    }



    //region Bindable properties

    public void setJobStatusTextChanged(){
        notifyPropertyChanged(BR.jobStatusText);
    }

    @Bindable
    public String getJobStatusText(){
        if (this.GridViewModel != null && this.GridViewModel.getItemsSource()!= null &&  this.GridViewModel.getItemsSource().size() > 0){
            return "You have been assigned the following jobs";
        }else{
            return "No jobs are currently assigned to you";
        }
    }

    @Bindable
    public boolean getIsServerConnected() {
        return AppRepo.getInstance().IsServerConnected();
    }

    //endregion Bindable properties

    //region Property Change Handler

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:
                notifyPropertyChanged(BR.isServerConnected);
                break;
            case AppRepo.IsStartupConnectionFailedEventRaisedPropName: {
                AppHelper.ExecuteStartUpProcess();
            }
            break;
            case AppRepo.CurrentLocationIdPropName:{
                headerViewModel.notifyChange();
            }
            break;
            case AppRepo.LocationListPropName: {
                headerViewModel.notifyChange();
            }
            break;
        }
    }

    //endregion Property Change Handler


    //region Test Data
    TestButtonClickHander Handler = new TestButtonClickHander();

    public TestButtonClickHander getHandler() {
        return Handler;
    }

    public void setHandler(TestButtonClickHander handler) {
        Handler = handler;
    }

//endregion Test Data
}
