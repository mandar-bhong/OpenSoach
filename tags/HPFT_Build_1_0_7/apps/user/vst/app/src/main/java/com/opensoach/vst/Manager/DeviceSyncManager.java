package com.opensoach.vst.Manager;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Utility.AppLogger;
import com.opensoach.vst.Model.Communication.PacketBatteryLevelModel;

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
