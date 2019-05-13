package com.opensoach.hpft.Model;

import java.util.Date;

/**
 * Created by samir.s.bukkawar on 3/25/2017.
 */

public class ChartDataModel {
    private int chartId;
    private String taskName;
    private int slotId;
    private Date slotStartTime;
    private Date slotEndTime;
    private int cellState;
    private Date entryDate;
    private String authCode;


    public int getChartId() {
        return chartId;
    }

    public void setChartId(int chartId) {
        this.chartId = chartId;
    }

    public String getTaskName() {
        return taskName;
    }

    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    public int getSlotId() {
        return slotId;
    }

    public void setSlotId(int slotId) {
        this.slotId = slotId;
    }

    public Date getSlotStartTime() {
        return slotStartTime;
    }

    public void setSlotStartTime(Date slotStartTime) {
        this.slotStartTime = slotStartTime;
    }

    public Date getSlotEndTime() {
        return slotEndTime;
    }

    public void setSlotEndTime(Date slotEndTime) {
        this.slotEndTime = slotEndTime;
    }

    public int getCellState() {
        return cellState;
    }

    public void setCellState(int cellState) {
        this.cellState = cellState;
    }

    public Date getEntryDate() {
        return entryDate;
    }

    public void setEntryDate(Date entryDate) {
        this.entryDate = entryDate;
    }

    public String getAuthCode() {
        return authCode;
    }

    public void setAuthCode(String authCode) {
        this.authCode = authCode;
    }


}
