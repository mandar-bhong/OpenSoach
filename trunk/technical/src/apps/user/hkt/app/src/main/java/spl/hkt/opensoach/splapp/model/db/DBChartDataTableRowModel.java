package spl.hkt.opensoach.splapp.model.db;

import java.util.Date;

import spl.hkt.opensoach.splapp.dal.DBConstants;
import spl.hkt.opensoach.splapp.dal.DBTableSchema;

/**
 * Created by samir.s.bukkawar on 2/26/2017.
 */
@DBTableSchema(TableName = DBConstants.TABLE_CHART_DATA)
public class DBChartDataTableRowModel {

    private int chartId;
    private int taskId;
    private int slotId;
    private Date entryTime;
    private Date slotStartTime;
    private Date slotEndTime;
    private int cellState;
    private Date chartDay;
    private boolean isSynced;
    private String authCode;

    public int getChartId() {
        return chartId;
    }

    public void setChartId(int chartId) {
        this.chartId = chartId;
    }

    public int getTaskId() {
        return taskId;
    }

    public void setTaskId(int taskId) {
        this.taskId = taskId;
    }

    public int getSlotId() {
        return slotId;
    }

    public void setSlotId(int slotId) {
        this.slotId = slotId;
    }

    public Date getEntryTime() {
        return entryTime;
    }

    public void setEntryTime(Date entryTime) {
        this.entryTime = entryTime;
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

    public Date getChartDay() {
        return chartDay;
    }

    public void setChartDay(Date chartDay) {
        this.chartDay = chartDay;
    }

    public boolean isSynced() {
        return isSynced;
    }

    public void setSynced(boolean synced) {
        isSynced = synced;
    }

    public String getAuthCode() {
        return authCode;
    }

    public void setAuthCode(String authCode) {
        this.authCode = authCode;
    }
}