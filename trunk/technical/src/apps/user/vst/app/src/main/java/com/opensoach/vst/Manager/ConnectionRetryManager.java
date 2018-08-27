package com.opensoach.vst.Manager;

import android.util.Log;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.Timer;
import java.util.TimerTask;


import com.opensoach.vst.AppRepo.AppRepo;

/**
 * Created by Mandar on 3/4/2017.
 */

public class ConnectionRetryManager implements PropertyChangeListener {

    private static ConnectionRetryManager singleton;

    private Timer _retryTimer;

    public Boolean IsRetryInProgress;


    private ConnectionRetryManager() {

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

    public boolean DeInit(){
        StopRetryConnection();
        return  true;
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
        _retryTimer = new Timer("ConnectionRetryManager", true);
        _retryTimer.schedule(new ConnectionRetryTask(), 1000, 1000 * 60 * 1);//1000-Milisecond,60-sec,1-min
    }

    public void StopRetryConnection() {

        _retryTimer.cancel();
    }


}

 class ConnectionRetryTask extends TimerTask  {
    @Override
    public void run() {
        if (ConnectionRetryManager.Instance().IsRetryInProgress) return;

        ConnectionRetryManager.Instance().IsRetryInProgress = true;

        ServerConnectionManager.Instance().Connect();

        ConnectionRetryManager.Instance().IsRetryInProgress = false;
    }
}
