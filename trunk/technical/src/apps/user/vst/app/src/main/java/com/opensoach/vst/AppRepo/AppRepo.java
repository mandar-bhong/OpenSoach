package com.opensoach.vst.AppRepo;

import java.beans.PropertyChangeListener;
import java.beans.PropertyChangeSupport;
import java.util.ArrayList;
import java.util.List;

import com.opensoach.vst.BuildConfig;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.Model.View.TaskItemDataModel;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;
import com.opensoach.vst.ViewModels.JobExeDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.JobSummaryViewModel;
import com.opensoach.vst.ViewModels.TaskItemViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;

/**
 * Created by Mandar on 2/25/2017. This class will have all application level data
 */

public final class AppRepo {

    private static AppRepo singleton;

    private final PropertyChangeSupport propertyChangeSupport = new PropertyChangeSupport(this);
    private boolean isServerConnected;
    private int currentLocationId;
    private int currentChartId;
    private String ServerWebSocketURL;

    private String ServerAPIHOST;
    private String ServerAPIURL;


    private boolean isDeviceAuthorized;

    private  int BatteryLevel;
    private  boolean isChartRendered;
    private ApplicationConstants.AppRunningMode currentRunningMode;

    private String AuthToken;
    private ArrayList<String> authCodeList;
    private boolean isDeviceSyncInProgress;
    private CardBriefViewModel activeCard;
    private TaskItemViewModel activeTaskItem;
    private TokenSelectionViewModel selectedToken;
    private JobServiceViewModel jobServiceViewModel;
    private JobSummaryViewModel jobSummaryViewModel;
    private JobServiceDetailsViewModel jobServiceDetailsViewModel;
    private JobExeDetailsViewModel jobExeDetailsViewModel;

    private List<TaskItemDataModel> selectedTaskDataViewModels;

    public static final String IsServerConnectedPropName = "AppRepo.IsServerConnected";
    public static final String CurrentLocationIdPropName = "AppRepo.currentLocationId";
    public static final String DeviceSyncCompletedPropName = "AppRepo.isDeviceSyncInProgress";
    public static final String DeviceAuthorizedPropName = "AppRepo.isDeviceAuthorized";
    public static final String IsChartRenderedPropName = "AppRepo.isChartRendered";

    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private AppRepo() {

        currentLocationId = 0;
        currentChartId = 0;
        authCodeList = new ArrayList<>();

        selectedTaskDataViewModels = new ArrayList<>();

        isDeviceAuthorized = true;

        ServerAPIHOST = BuildConfig.ServiceAPIHost;
        ServerAPIURL = "http://" + ServerAPIHOST + "/api/v1/endpoint/deviceauthorization";

    }

    public static AppRepo getInstance() {
        if (singleton == null)
            singleton = new AppRepo();
        return singleton;
    }

    public String getDeviceSerial() {
        return  "1345494544733456";
        //devid- 1 , S.no - 1234567890123456
        //devid- 2 , S.no - 1345494544733456
        //devid- 3 , S.no - 1155623421323222
        //return Build.SERIAL;
    }

    public String getServerAPIURL() {
        return ServerAPIURL;
    }

    public void setServerAPIURL(String serverAPIURL) {
        ServerAPIURL = serverAPIURL;
    }

    public String getServerWebSocketURL() {
        return ServerWebSocketURL;
    }

    public void setServerWebSocketURL(String serverWebSocketURL) {
        ServerWebSocketURL = serverWebSocketURL;
    }

    public String getAuthToken() {
        return AuthToken;
    }

    public void setAuthToken(String authToken) {
        AuthToken = authToken;
    }

    public void addPropertyChangeListener(PropertyChangeListener listener) {
        this.propertyChangeSupport.addPropertyChangeListener(listener);
    }

    public void removePropertyChangeListener(PropertyChangeListener listener) {
        this.propertyChangeSupport.removePropertyChangeListener(listener);
    }

    public void IsServerConnected(boolean newValue) {
        boolean oldValue = this.isServerConnected;
        this.isServerConnected = newValue;
        this.propertyChangeSupport.firePropertyChange(IsServerConnectedPropName, oldValue, newValue);
    }

    public boolean IsServerConnected() {
        return this.isServerConnected;
    }

    public Integer getCurrentLocationId() {
        return currentLocationId;
    }

    public void setCurrentLocationId(int currentLocationId) {
        //Location Change might required chart change
        int oldValue = this.currentLocationId;
        this.currentLocationId = currentLocationId;
        this.propertyChangeSupport.firePropertyChange(CurrentLocationIdPropName, oldValue, currentLocationId);
    }

    public Integer getCurrentChartId() {
        return currentChartId;
    }

    public void setCurrentChartId(Integer currentChartId) {
        this.currentChartId = currentChartId;
    }

    public ArrayList<String> getAuthCodeList() {
        return authCodeList;
    }

    public void setAuthCodeList(ArrayList<String> authCodeList) {
        this.authCodeList = authCodeList;
    }

    public boolean getIsDeviceSyncInProgress() {
        return this.isDeviceSyncInProgress;
    }

    public void setIsDeviceSyncInProgress(boolean isDeviceSyncInProgress) {
        this.isDeviceSyncInProgress = isDeviceSyncInProgress;
        this.propertyChangeSupport.firePropertyChange(DeviceSyncCompletedPropName, !this.isDeviceSyncInProgress, this.isDeviceSyncInProgress);
    }

    public boolean IsDeviceAuthorized() {
        return isDeviceAuthorized;
    }

    public void setIsDeviceAuthorized(boolean deviceAuthorized) {
        boolean oldValue = isDeviceAuthorized;
        isDeviceAuthorized = deviceAuthorized;
        this.propertyChangeSupport.firePropertyChange(DeviceAuthorizedPropName, oldValue, this.isDeviceAuthorized);
    }


    public int getBatteryLevel() {
        return BatteryLevel;
    }

    public void setBatteryLevel(int batteryLevel) {
        BatteryLevel = batteryLevel;
    }

    public boolean isChartRendered() {
        return isChartRendered;
    }

    public void setChartRendered(boolean chartRendered) {
        boolean oldValue = isChartRendered;
        isChartRendered = chartRendered;
        this.propertyChangeSupport.firePropertyChange(IsChartRenderedPropName, oldValue, isChartRendered);
    }


    public CardBriefViewModel getActiveCard() {
        return activeCard;
    }

    public void setActiveCard(CardBriefViewModel activeCard) {
        this.activeCard = activeCard;
    }

    public TaskItemViewModel getActiveTaskItem() {
        return activeTaskItem;
    }

    public void setActiveTaskItem(TaskItemViewModel activeTaskItem) {
        this.activeTaskItem = activeTaskItem;
    }

    public List<TaskItemDataModel> getSelectedTaskDataViewModels() {
        return selectedTaskDataViewModels;
    }


    public TokenSelectionViewModel getSelectedToken() {
        return selectedToken;
    }

    public void setSelectedToken(TokenSelectionViewModel selectedToken) {
        this.selectedToken = selectedToken;
    }


    public JobServiceViewModel getJobServiceViewModel() {
        return jobServiceViewModel;
    }

    public void setJobServiceViewModel(JobServiceViewModel jobServiceViewModel) {
        this.jobServiceViewModel = jobServiceViewModel;
    }

    public JobSummaryViewModel getJobSummaryViewModel() {
        return jobSummaryViewModel;
    }

    public void setJobSummaryViewModel(JobSummaryViewModel jobSummaryViewModel) {
        this.jobSummaryViewModel = jobSummaryViewModel;
    }


    public ApplicationConstants.AppRunningMode getCurrentRunningMode() {
        return currentRunningMode;
    }

    public void setCurrentRunningMode(ApplicationConstants.AppRunningMode currentRunningMode) {
        this.currentRunningMode = currentRunningMode;
    }



    public JobServiceDetailsViewModel getJobServiceDetailsViewModel() {
        return jobServiceDetailsViewModel;
    }

    public void setJobServiceDetailsViewModel(JobServiceDetailsViewModel jobServiceDetailsViewModel) {
        this.jobServiceDetailsViewModel = jobServiceDetailsViewModel;
    }

    public JobExeDetailsViewModel getJobExeDetailsViewModel() {
        return jobExeDetailsViewModel;
    }

    public void setJobExeDetailsViewModel(JobExeDetailsViewModel jobExeDetailsViewModel) {
        this.jobExeDetailsViewModel = jobExeDetailsViewModel;
    }
}