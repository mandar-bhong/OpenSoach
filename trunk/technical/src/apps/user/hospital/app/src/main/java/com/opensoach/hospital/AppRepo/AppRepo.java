package com.opensoach.hospital.AppRepo;


import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Model.View.PropChangeDataModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBoardViewModel;
import com.opensoach.hospital.Views.Interfaces.IUIUpdateEvent;
import com.opensoach.hospital.Views.Notifier.NotifyPropChangeOnUIThread;

import java.beans.PropertyChangeListener;
import java.beans.PropertyChangeSupport;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 8/24/2017.
 */

public class AppRepo  {

    private static AppRepo singleton;

    private final PropertyChangeSupport propertyChangeSupport = new PropertyChangeSupport(this);
    private boolean isServerConnected;
    private int currentLocationId;
    private int currentChartId;
    private String ServerWebSocketURL;
    private boolean isAuthCodeRequired;
    private ArrayList<String> authCodeList;
    private JobBoardViewModel SelectedJobBoard;
    private List<DBLocationTableRowModel> locationList;
    private boolean isStartupConnectionFailedEventRaised;
    private boolean isStartupCompleted;
    private  String foregroundActivityName;
    private IUIUpdateEvent foregroundActivityHandler;

    private  String userName;
    private  String password;


    public static final String IsServerConnectedPropName = "AppRepo.IsServerConnected";
    public static final String CurrentLocationIdPropName = "AppRepo.currentLocationId";
    public static final String IsStartupConnectionFailedEventRaisedPropName = "AppRepo.isStartupConnectionFailedEventRaised";
    public static final String LocationListPropName = "AppRepo.LocationList";
    public static final String IsStartUpCompletedPropName = "AppRepo.IsStartUpCompletedPropName";



    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private AppRepo() {

        currentLocationId = 0;
        currentChartId = 0;
        authCodeList = new ArrayList<String>();
        locationList = new ArrayList<DBLocationTableRowModel>();
        isAuthCodeRequired = false;

    }

    public static AppRepo getInstance() {
        if (singleton == null)
            singleton = new AppRepo();
        return singleton;
    }

    public String getServerWebSocketURL() {
        return ServerWebSocketURL;
    }

    public void setServerWebSocketURL(String serverWebSocketURL) {
        ServerWebSocketURL = serverWebSocketURL;
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

        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"Connection State: Connected: "+newValue);

        new NotifyPropChangeOnUIThread(new PropChangeDataModel(propertyChangeSupport,
                IsServerConnectedPropName,oldValue,newValue)).run();

        if(!newValue && !isStartupConnectionFailedEventRaised){
            isStartupConnectionFailedEventRaised = true;

            new NotifyPropChangeOnUIThread(new PropChangeDataModel(propertyChangeSupport,
                    IsStartupConnectionFailedEventRaisedPropName,!newValue,newValue)).run();

        }else if (newValue){
            isStartupConnectionFailedEventRaised =true;
        }
    }

    public boolean IsServerConnected() {
        return this.isServerConnected;
    }

    public Integer getCurrentLocationId() {
        return currentLocationId;
    }

    public void setCurrentLocationId(int currentLocationId) {
        int oldValue = this.currentLocationId;
        this.currentLocationId = currentLocationId;

        new NotifyPropChangeOnUIThread(new PropChangeDataModel(propertyChangeSupport,
                CurrentLocationIdPropName,oldValue,currentLocationId)).run();
    }

    public Integer getCurrentChartId() {
        return currentChartId;
    }

    public void setCurrentChartId(Integer currentChartId) {
        this.currentChartId = currentChartId;
    }

    public boolean isAuthCodeRequired() {
        return isAuthCodeRequired;
    }

    public void setAuthCodeRequired(boolean authCodeRequired) {
        isAuthCodeRequired = authCodeRequired;
    }

    public ArrayList<String> getAuthCodeList() {
        return authCodeList;
    }

    public void setAuthCodeList(ArrayList<String> authCodeList) {
        this.authCodeList = authCodeList;
    }

    public JobBoardViewModel getSelectedJobBoard() {
        return SelectedJobBoard;
    }

    public void setSelectedJobBoard(JobBoardViewModel selectedJobBoard) {
        SelectedJobBoard = selectedJobBoard;
    }

    public List<DBLocationTableRowModel> getLocationList() {
        return locationList;
    }

    public void setLocationList(List<DBLocationTableRowModel> locationList) {
        List<DBLocationTableRowModel> oldValue = this.locationList;
        this.locationList = locationList;

        new NotifyPropChangeOnUIThread(new PropChangeDataModel(propertyChangeSupport,
                LocationListPropName,oldValue,locationList)).run();
    }


    public void setIsStartupCompleted(){

        isStartupCompleted =true;
        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"setting setIsStartupCompleted");

        new NotifyPropChangeOnUIThread(new PropChangeDataModel(propertyChangeSupport,
                IsStartUpCompletedPropName,false,true)).run();
    }

    public boolean getIsStartUpCompleted(){
        return isStartupCompleted;
    }


    public String getForegroundActivityName() {
        return foregroundActivityName;
    }

    public void setForegroundActivityName(String foregroundActivityName) {
        this.foregroundActivityName = foregroundActivityName;
    }

    public IUIUpdateEvent getForegroundActivityHandler() {
        return foregroundActivityHandler;
    }

    public void setForegroundActivityHandler(IUIUpdateEvent foregroundActivityHandler) {
        this.foregroundActivityHandler = foregroundActivityHandler;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

}
