package com.opensoach.vst.ViewModels;


import android.app.Activity;
import android.databinding.Bindable;
import android.databinding.BindingAdapter;
import android.graphics.drawable.Drawable;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.view.View;
import android.widget.ImageView;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.BR;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.R;

/**
 * Created by samir.s.bukkawar on 3/19/2017.
 */

public class HeaderViewModel extends BaseViewModel implements PropertyChangeListener {

    private ArrayList<String> locationList;
    private Constants.NETWORK_STATE networkState;
    private boolean backButtonVisiable;
    private boolean uploadEnabled;
    private boolean uploadVisiable;


    public HeaderViewModel() {
        this.networkState = Constants.NETWORK_STATE.WEB_SOCKET_UNAUTHORIZED;
        uploadEnabled= true;
        uploadVisiable = false;
    }

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

    public void setNetworkState(Constants.NETWORK_STATE networkState) {
        this.networkState = networkState;
        notifyPropertyChanged(BR.nwState);
    }

    @Bindable
    public boolean isBackButtonVisiable() {
        return backButtonVisiable;
    }

    public void setBackButtonVisiable(boolean backButtonVisiable) {
        this.backButtonVisiable = backButtonVisiable;
        notifyPropertyChanged(BR.backButtonVisiable);
    }

    public void onBackClick(View view) {
        Activity currentActivity =(Activity) view.getContext();
        currentActivity.finish();
    }

    @Bindable
    public Drawable getNwState() {
        switch (networkState) {
            case WEB_SOCKET_CONNECTED: {
                return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.online);
            }
            case WEB_SOCKET_DISSCONNECTED: {
                return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.offline);
            }
            case WEB_SOCKET_UNAUTHORIZED: {
                return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.unauthorized);
            }

            case NW_NOT_AVAILABLE: {
                return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.offline);
            }
            default: {
                return MainViewModel.getInstance().ContextActivity.getResources().getDrawable(R.drawable.offline);
            }
        }
    }


    public void setUploadEnabled(boolean isUploadEnabled){
        uploadEnabled = isUploadEnabled;
        notifyPropertyChanged(BR.uploadEnabled);
        notifyPropertyChanged(BR.uploadAlpha);
    }

    @Bindable
    public boolean getUploadEnabled(){
        return  uploadEnabled;
    }

    @Bindable
    public float getUploadAlpha(){
        if (uploadEnabled == false){
            return  0.2f;
        }else{
            return  1.0f;
        }
    }


    @Bindable
    public boolean getUploadVisiable() {
        return uploadVisiable;
    }

    public void setUploadVisiable(boolean uploadVisiable) {
        this.uploadVisiable = uploadVisiable;
        notifyPropertyChanged(BR.uploadVisiable);
    }

    @Override
    public void propertyChange(PropertyChangeEvent evt) {

        Handler uiHandler = new Handler(Looper.getMainLooper()) {
            @Override
            public void handleMessage(Message message) {

                switch (  message.getData().getString("PropertyName")){
                    case AppRepo.IsServerConnectedPropName:
                        boolean isConnected = message.getData().getBoolean("ConnectionState");
                        if (isConnected) {
                            setNetworkState(Constants.NETWORK_STATE.WEB_SOCKET_CONNECTED);
                        } else {
                            setNetworkState(Constants.NETWORK_STATE.WEB_SOCKET_DISSCONNECTED);
                        }
                        break;
                    case AppRepo.DeviceAuthorizedPropName:

                        boolean isAuthorized = message.getData().getBoolean("IsAuthorized");

                        if (isAuthorized == false){
                            setNetworkState(Constants.NETWORK_STATE.WEB_SOCKET_UNAUTHORIZED);
                        }

                        break;
                }
            }
        };

        Message msg = uiHandler.obtainMessage();
        Bundle b = new Bundle();
        b.putString("PropertyName", evt.getPropertyName());

        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:
                b.putBoolean("ConnectionState", (boolean) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;

            case AppRepo.DeviceAuthorizedPropName:
                b.putBoolean("IsAuthorized", (boolean) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;
        }
    }
}
