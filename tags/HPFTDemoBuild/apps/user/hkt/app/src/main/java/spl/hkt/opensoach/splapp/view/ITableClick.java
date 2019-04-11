package spl.hkt.opensoach.splapp.view;

import spl.hkt.opensoach.splapp.viewModels.ChartViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface ITableClick {

    ChartViewModel getChartViewModel();

    void onChartTableClick(ChartViewModel chartViewModel);

}
