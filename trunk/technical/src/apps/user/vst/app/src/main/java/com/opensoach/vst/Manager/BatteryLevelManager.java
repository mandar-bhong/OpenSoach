package com.opensoach.vst.Manager;

import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Model.Communication.PacketBatteryLevelModel;

/**
 * Created by Mandar on 25-06-2018.
 */

public class BatteryLevelManager extends BroadcastReceiver {

    boolean isDeregitered;

    public boolean isDeregitered() {
        return isDeregitered;
    }

    public void setDeregitered(boolean deregitered) {
        isDeregitered = deregitered;
    }



    @Override
    public void onReceive(Context context, Intent intent) {

        if (isDeregitered){
            context.unregisterReceiver(this);
            return;
        }

        //context.unregisterReceiver(this);
        int rawlevel = intent.getIntExtra("level", -1);
        int scale = intent.getIntExtra("scale", -1);
        int level = -1;
        if (rawlevel >= 0 && scale > 0) {
            level = (rawlevel * 100) / scale;
        }

        if (AppRepo.getInstance().getBatteryLevel() == level)
            return;

        AppRepo.getInstance().setBatteryLevel(level);

        //TODO Raise event once 5% change is occured

        if (AppRepo.getInstance().getIsDeviceSyncInProgress() == false) {
            PacketBatteryLevelModel packetBatteryLevelModel = new PacketBatteryLevelModel();
            packetBatteryLevelModel.BatteryLevel = level;
            SendPacketManager.Instance().send(AppAction.BATTERY_LEVEL, packetBatteryLevelModel);
        }
    }
}
