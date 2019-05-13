package spl.hkt.opensoach.splapp.manager;

import android.util.Log;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.helper.AppAction;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.model.communication.PacketBatteryLevelModel;

public class DeviceSyncManager implements PropertyChangeListener {

    private static DeviceSyncManager singleton;

    private DeviceSyncManager() {

    }

    /* Static 'instance' method */
    public static DeviceSyncManager Instance() {
        if (singleton == null)
            singleton = new DeviceSyncManager();
        return singleton;
    }

    @Override
    public void propertyChange(PropertyChangeEvent evt) {
        switch (evt.getPropertyName()) {
            case AppRepo.DeviceSyncCompletedPropName:
                if (AppRepo.getInstance().getIsDeviceSyncInProgress() == false) {
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "Device sync completed");
                    PacketBatteryLevelModel packetBatteryLevelModel = new PacketBatteryLevelModel();
                    packetBatteryLevelModel.BatteryLevel= AppRepo.getInstance().getBatteryLevel();
                    SendPacketManager.Instance().send(AppAction.BATTERY_LEVEL, packetBatteryLevelModel);
                }
                break;
        }
    }
}
