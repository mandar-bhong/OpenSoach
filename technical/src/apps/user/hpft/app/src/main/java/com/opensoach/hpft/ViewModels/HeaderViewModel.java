package com.opensoach.hpft.ViewModels;


import android.databinding.BindingAdapter;
import android.graphics.drawable.Drawable;
import android.widget.ImageView;

import java.util.ArrayList;

import com.opensoach.hpft.Constants.Constants;
import com.opensoach.hpft.R;

/**
 * Created by samir.s.bukkawar on 3/19/2017.
 */

public class HeaderViewModel extends BaseViewModel {

    private ArrayList<String> locationList;
    private Constants.NETWORK_STATE networkState;

    @BindingAdapter("android:background")
    public static void setImageDrawable(ImageView view, Drawable drawable) {
        view.setImageDrawable(drawable);
    }

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

    public Drawable getNwState() {
        return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.online);
    }

}
