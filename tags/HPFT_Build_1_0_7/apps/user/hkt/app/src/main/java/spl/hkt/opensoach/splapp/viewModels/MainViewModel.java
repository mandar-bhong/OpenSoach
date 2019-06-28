package spl.hkt.opensoach.splapp.viewModels;

import java.util.HashMap;

import spl.hkt.opensoach.splapp.model.ChartDataModel;

/**
 * Created by samir.s.bukkawar on 3/18/2017.
 */

public class MainViewModel {

    private static MainViewModel singleton;

    private HeaderViewModel headerViewModel;
    private ChartViewModel chartViewModel;
    private FooterViewModel footerViewModel;
    private HashMap<String, ChartDataModel> currenChartDataModelMap;

    /* A private Constructor prevents any other
 * class from instantiating.
 */
    private MainViewModel() {
        currenChartDataModelMap = new HashMap<String, ChartDataModel>();
        headerViewModel = new HeaderViewModel();
    }

    public static MainViewModel getInstance() {
        if (singleton == null)
            singleton = new MainViewModel();
        return singleton;
    }

    public HeaderViewModel getHeaderViewModel() {
        return headerViewModel;
    }

    public ChartViewModel getChartViewModel() {
        return chartViewModel;
    }

    public void setChartViewModel(ChartViewModel chartViewModel) {
        this.chartViewModel = chartViewModel;
    }

    public FooterViewModel getFooterViewModel() {
        return footerViewModel;
    }

    public void setFooterViewModel(FooterViewModel footerViewModel) {
        this.footerViewModel = footerViewModel;
    }

    public void clearCurrenClickeCellModelap() {
        currenChartDataModelMap.clear();
    }

    public void createNewCurrenClickeCellModelMap() {
        currenChartDataModelMap = new HashMap<String, ChartDataModel>();
    }

    public HashMap<String, ChartDataModel> getCurrenChartDataModelMap() {
        return currenChartDataModelMap;
    }

    public void setCurrenChartDataModelMap(HashMap<String, ChartDataModel> currenChartDataModelMap) {
        this.currenChartDataModelMap = currenChartDataModelMap;
    }
}