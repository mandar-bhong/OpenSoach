package spl.hkt.opensoach.splapp.manager;

import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.helper.AppAction;

public class BroadCastReceiverManager {

    public static void RegisterBatteryLevelReceiver(Context ctx) {
        BroadcastReceiver batteryLevelReceiver = new BroadcastReceiver() {
            public void onReceive(Context context, Intent intent) {
                context.unregisterReceiver(this);
                int rawlevel = intent.getIntExtra("level", -1);
                int scale = intent.getIntExtra("scale", -1);
                int level = -1;
                if (rawlevel >= 0 && scale > 0) {
                    level = (rawlevel * 100) / scale;
                }

                AppRepo.getInstance().setBatteryLevel(level);

                //TODO Raise event once 5% change is occured

                if (AppRepo.getInstance().getIsDeviceSyncInProgress() == false){
                    SendPacketManager.Instance().send(AppAction.BATTERY_LEVEL, level);
                }
            }
        };
        IntentFilter batteryLevelFilter = new IntentFilter(Intent.ACTION_BATTERY_CHANGED);
        //registerReceiver(batteryLevelReceiver, batteryLevelFilter);
        ctx.registerReceiver(batteryLevelReceiver, batteryLevelFilter);
    }
}
