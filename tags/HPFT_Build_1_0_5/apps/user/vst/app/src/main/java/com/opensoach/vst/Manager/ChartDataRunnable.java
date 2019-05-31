package com.opensoach.vst.Manager;

import com.opensoach.vst.Helper.AppHelper;

/**
 * Created by samir.s.bukkawar on 4/9/2017.
 */

public class ChartDataRunnable implements Runnable {

    private Integer mChartId;

    @Override
    public void run() {

        if (mChartId == 0)
            return;

        AppHelper.UpdateChartData(mChartId);
    }

    public ChartDataRunnable(int chartId) {
        mChartId = chartId;
    }
}
