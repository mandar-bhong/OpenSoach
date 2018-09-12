package com.opensoach.vst.ViewModels;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.HashMap;

import com.opensoach.vst.Model.ChartDataModel;

/**
 * Created by samir.s.bukkawar on 3/18/2017.
 */

public class MainViewModel extends BaseViewModel implements PropertyChangeListener {

    private static MainViewModel singleton;

    private HeaderViewModel headerViewModel;
    private ChartViewModel chartViewModel;
    private FooterViewModel footerViewModel;
    private CardListViewModel cardListViewModel;
    private HashMap<String, ChartDataModel> currenChartDataModelMap;
    private CreateTokenViewModel createTokenViewModel;
    private TokenListViewModel tokenListViewModel;
    private JobServiceItemViewModel jobServiceItemViewModel;

    /* A private Constructor prevents any other
 * class from instantiating.
 */
    private MainViewModel() {
        currenChartDataModelMap = new HashMap<String, ChartDataModel>();
        headerViewModel = new HeaderViewModel();
        cardListViewModel = new CardListViewModel();
        createTokenViewModel = new CreateTokenViewModel();
        tokenListViewModel = new TokenListViewModel();
        jobServiceItemViewModel = new JobServiceItemViewModel();
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


    public CardListViewModel getCardListViewModel() {
        return cardListViewModel;
    }

    public void setCardListViewModel(CardListViewModel cardListViewModel) {
        this.cardListViewModel = cardListViewModel;
    }

    public CreateTokenViewModel getCreateTokenViewModel() {
        return createTokenViewModel;
    }

    public void setCreateTokenViewModel(CreateTokenViewModel createTokenViewModel) {
        this.createTokenViewModel = createTokenViewModel;
    }


    public TokenListViewModel getTokenListViewModel() {
        return tokenListViewModel;
    }

    public void setTokenListViewModel(TokenListViewModel tokenListViewModel) {
        this.tokenListViewModel = tokenListViewModel;
    }

    @Override
    public void propertyChange(PropertyChangeEvent propertyChangeEvent) {

    }

    public JobServiceItemViewModel getJobServiceItemViewModel() {
        return jobServiceItemViewModel;
    }

    public void setJobServiceItemViewModel(JobServiceItemViewModel jobServiceItemViewModel) {
        this.jobServiceItemViewModel = jobServiceItemViewModel;
    }
}