package com.opensoach.vst.Model.DB;

import java.util.Date;

import com.opensoach.vst.DAL.DBConstants;
import com.opensoach.vst.DAL.DBTableSchema;

/**
 * Created by samir.s.bukkawar on 2/26/2017.
 */
@DBTableSchema(TableName = DBConstants.TABLE_CHART)
public class DBChartTableRowModel {
    private int chartId;
    private int serverChartId;
    private int locationId;
    private String chartPayload;
    private Date chartDispStartDate;
    private Date chartDispEndDate;
    private String chartName;


    public int getChartId() {
        return chartId;
    }

    public void setChartId(int chartId) {
        this.chartId = chartId;
    }

    public int getServerChartId() {
        return serverChartId;
    }

    public void setServerChartId(int serverChartId) {
        this.serverChartId = serverChartId;
    }

    public String getChartPayload() {
        return chartPayload;
    }

    public void setChartPayload(String chartPayload) {
        this.chartPayload = chartPayload;
    }

    public int getLocationId() {
        return locationId;
    }

    public void setLocationId(int locationId) {
        this.locationId = locationId;
    }

    public Date getChartDispStartDate() {
        return chartDispStartDate;
    }

    public void setChartDispStartDate(Date chartDispStartDate) {
        this.chartDispStartDate = chartDispStartDate;
    }

    public Date getChartDispEndDate() {
        return chartDispEndDate;
    }

    public void setChartDispEndDate(Date chartDispEndDate) {
        this.chartDispEndDate = chartDispEndDate;
    }

    public String getChartName() {
        return chartName;
    }

    public void setChartName(String chartName) {
        this.chartName = chartName;
    }
}