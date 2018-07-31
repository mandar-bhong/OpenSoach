package com.opensoach.hpft.AppRepo;

import android.os.Build;

import java.beans.PropertyChangeListener;
import java.beans.PropertyChangeSupport;
import java.util.ArrayList;

import com.opensoach.hpft.BuildConfig;

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

    private String AuthToken;
    private ArrayList<String> authCodeList;
    private boolean isDeviceSyncInProgress;

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
        return  "12345678901234";
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
}