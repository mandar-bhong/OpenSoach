package com.opensoach.vst.Manager;

import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;

public class BroadCastReceiverManager {

    private static BroadCastReceiverManager singleton;
    private BatteryLevelManager broadcastBatteryLevelReceiver;

    private BroadCastReceiverManager() {
        broadcastBatteryLevelReceiver = new BatteryLevelManager();
    }

    /* Static 'instance' method */
    public static BroadCastReceiverManager Instance() {
        if (singleton == null)
            singleton = new BroadCastReceiverManager();
        return singleton;
    }

    public  void RegisterBatteryLevelReceiver(Context ctx) {
        broadcastBatteryLevelReceiver.setDeregitered(false);
        IntentFilter batteryLevelFilter = new IntentFilter(Intent.ACTION_BATTERY_CHANGED);
        ctx.registerReceiver(broadcastBatteryLevelReceiver, batteryLevelFilter);
    }

    public void DeregisterBatteryLevelReceiver(Context ctx){
        broadcastBatteryLevelReceiver.setDeregitered(true);
        ctx.unregisterReceiver(broadcastBatteryLevelReceiver);
    }
}
