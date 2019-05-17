package spl.hkt.opensoach.splapp.model.view;

import java.util.ArrayList;

/**
 * Created by Mandar on 4/7/2017.
 */

public class DisplayChartDataModel {

    ArrayList<DisplayChartItemDataModel> chartData;

    public ArrayList<DisplayChartItemDataModel> getChartData() {
        return chartData;
    }

    public  DisplayChartDataModel(){
        chartData = new ArrayList<>();
    }

}
