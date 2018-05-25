package spl.hkt.opensoach.splapp.apprepo;

import java.beans.PropertyChangeListener;
import java.beans.PropertyChangeSupport;
import java.util.ArrayList;

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

    private String AuthToken;
    private ArrayList<String> authCodeList;

    public static final String IsServerConnectedPropName = "AppRepo.IsServerConnected";
    public static final String CurrentLocationIdPropName = "AppRepo.currentLocationId";

    /* A private Constructor prevents any other
     * class from instantiating.
     */
    private AppRepo() {

        currentLocationId = 0;
        currentChartId = 0;
        authCodeList = new ArrayList<>();

        ServerAPIHOST = "172.105.232.148";
        ServerAPIURL = "http://"+ServerAPIHOST +"/api/v1/endpoint/deviceauthorization";

    }

    public static AppRepo getInstance() {
        if (singleton == null)
            singleton = new AppRepo();
        return singleton;
    }

    public String getDeviceSerial() {
        return "1234567890123456";
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
}