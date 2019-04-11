package com.opensoach.hpft.Manager;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.Helper.AppAction;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.Model.Communication.PacketBatteryLevelModel;

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
