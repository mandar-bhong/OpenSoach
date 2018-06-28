package spl.hkt.opensoach.splapp.viewModels;


import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;

import spl.hkt.opensoach.splapp.Constants;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;

/**
 * Created by samir.s.bukkawar on 3/19/2017.
 */

public class HeaderViewModel extends BaseViewModel {

    private ArrayList<String> locationList;
    private Constants.NETWORK_STATE networkState;

    public ArrayList<String> getLocationList() {
        return locationList;
    }

    public void setLocationList(ArrayList<String> locationList) {
        this.locationList = locationList;
    }

    public Constants.NETWORK_STATE getNetworkState() {
        return networkState;
    }

    public void setNetworkState(Constants.NETWORK_STATE networkState) {
        this.networkState = networkState;
    }

}
