package com.opensoach.hpft.ViewModels;


import java.util.ArrayList;

import com.opensoach.hpft.Constants.Constants;

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
