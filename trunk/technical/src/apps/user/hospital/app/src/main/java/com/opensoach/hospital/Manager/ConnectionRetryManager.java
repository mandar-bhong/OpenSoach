package com.opensoach.hospital.Manager;

import android.util.Log;

import com.opensoach.hospital.AppRepo.AppRepo;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.Timer;
import java.util.TimerTask;

/**
 * Created by Mandar on 8/26/2017.
 */

public class ConnectionRetryManager implements PropertyChangeListener {

    private static ConnectionRetryManager singleton;

    private Timer _retryTimer;

    public Boolean IsRetryInProgress;
    private Boolean isFirstTime;


    private ConnectionRetryManager() {
        isFirstTime = true;
    }

    /* Static 'instance' method */
    public static ConnectionRetryManager Instance() {
        if (singleton == null)
            singleton = new ConnectionRetryManager();
        return singleton;
    }

    public boolean Init() {
        IsRetryInProgress = false;
        StartRetryConnection();
        return true;
    }

    @Override
    public void propertyChange(PropertyChangeEvent evt) {

        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:
                Log.d("Connection State: ", evt.getNewValue().toString());
                if ((boolean) evt.getNewValue()) {
                    StopRetryConnection();
                } else {
                    StartRetryConnection();
                }
                break;
        }
    }

    public void StartRetryConnection() {

        int retryTimeInSec = 1000 * 15 * 1;

        if(isFirstTime){
            isFirstTime = false;
            _retryTimer = new Timer("ConnectionRetryManager", true);
            _retryTimer.schedule(new ConnectionRetryTask(), 0 , retryTimeInSec);//1000-Milisecond,60-sec,1-min
        }else{
            _retryTimer = new Timer("ConnectionRetryManager", true);
            _retryTimer.schedule(new ConnectionRetryTask(), retryTimeInSec, retryTimeInSec);//1000-Milisecond,60-sec,1-min
        }


    }

    public void StopRetryConnection() {

        _retryTimer.cancel();
        _retryTimer = null;
    }
}


class ConnectionRetryTask extends TimerTask {
    @Override
    public void run() {
        Thread.currentThread().setName("ConnectionRetyManager");

        if (ConnectionRetryManager.Instance().IsRetryInProgress) return;

        ConnectionRetryManager.Instance().IsRetryInProgress = true;

        ServerConnectionManager.Instance().Connect();

        ConnectionRetryManager.Instance().IsRetryInProgress = false;

    }
}