package spl.hkt.opensoach.splapp.viewModels;

import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;

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
