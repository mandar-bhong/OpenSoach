package com.opensoach.hpft.View;

import com.opensoach.hpft.ViewModels.ChartViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface ITableClick {

    ChartViewModel getChartViewModel();

    void onChartTableClick(ChartViewModel chartViewModel);

}
