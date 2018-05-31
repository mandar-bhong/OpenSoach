package spl.hkt.opensoach.splapp.manager;

import android.util.Log;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.helper.AppAction;
import spl.hkt.opensoach.splapp.logger.AppLogger;

public class DeviceSyncManager implements PropertyChangeListener {
    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.DeviceSyncCompletedPropName:
                if (AppRepo.getInstance().getIsDeviceSyncInProgress() == false) {
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "Device sync completed");
                    SendPacketManager.Instance().send(AppAction.BATTERY_LEVEL, AppRepo.getInstance().getBatteryLevel());
                }
                break;
        }
    }
}
