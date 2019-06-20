package com.opensoach.hpft.Model.Communication;

import java.util.ArrayList;
import java.util.List;

import com.opensoach.hpft.Model.DB.DBChartDataTableRowModel;

/**
 * Created by Mandar on 4/14/2017.
 */

public class DeviceChartDataStartupSyncModel extends DeviceDataBaseModel {

    public List<DBChartDataTableRowModel> unSyncChartData;

    public DeviceChartDataStartupSyncModel(){

        unSyncChartData = new ArrayList<DBChartDataTableRowModel>();
    }

    public List<DBChartDataTableRowModel> getUnSyncChartData() {
        return unSyncChartData;
    }

    public void setUnSyncChartData(List<DBChartDataTableRowModel> unSyncChartData) {
        this.unSyncChartData = unSyncChartData;
    }
}
