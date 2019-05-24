package com.opensoach.vst.Manager;

import com.opensoach.vst.Helper.AppHelper;

/**
 * Created by Mandar on 4/8/2017.
 */

public class  LocationChartRunnable implements Runnable {

    private  Integer locId;

    public LocationChartRunnable(int locationId){
        locId = locationId;
    }

    @Override
    public void run() {

        if(locId == 0)
            return;

        AppHelper.LoadLocationChart(locId);
    }
}