package com.opensoach.vst.Helper;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.ArrayList;

import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Model.DB.DBChartDataTableRowModel;
import com.opensoach.vst.Model.View.DisplayChartItemDataModel;


/**
 * Created by Mandar on 4/9/2017.
 */

public class DataConvertHelper {

    public static DisplayChartItemDataModel ConvertDBChartDataToChartDisplayModel(DBChartDataTableRowModel dbModel) {
        DisplayChartItemDataModel displayChartItemDataModel = new DisplayChartItemDataModel();
        displayChartItemDataModel.setChartId(dbModel.getChartId());
        displayChartItemDataModel.setTaskName(dbModel.getTaskName());
        displayChartItemDataModel.setSlotId(dbModel.getSlotId());

        switch (dbModel.getCellState()) {
            case ApplicationConstants.DB_CHART_STATE_ON_TIME:
                displayChartItemDataModel.setState(ApplicationConstants.CHART_STATE_ON_TIME);
                break;
            case ApplicationConstants.DB_CHART_STATE_DELAYED:
                displayChartItemDataModel.setState(ApplicationConstants.CHART_STATE_DELAYED);
                break;
        }

        return displayChartItemDataModel;
    }


    public static ArrayList<String> ConvertJSONStringArray(String JSONData) {

        if (JSONData == null) {
            return new ArrayList<String>();
        }

        TypeToken<ArrayList<String>> typeToken = new TypeToken<ArrayList<String>>() {
        };
        Type packetType = typeToken.getType();
        ArrayList<String> stringArray = new Gson().fromJson(JSONData, packetType);
        return stringArray;
    }
}
