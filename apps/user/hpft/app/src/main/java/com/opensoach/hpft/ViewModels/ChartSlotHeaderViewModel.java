package com.opensoach.hpft.ViewModels;

import com.opensoach.hpft.Model.View.ChartConfigModel;

/**
 * Created by Mandar on 3/28/2017.
 */

public class ChartSlotHeaderViewModel {

    private ChartConfigModel chartDataModel;

    public ChartConfigModel getChartDataModel() {
        return chartDataModel;
    }

    public void setChartDataModel(ChartConfigModel chartDataModel) {
        this.chartDataModel = chartDataModel;
    }
}
