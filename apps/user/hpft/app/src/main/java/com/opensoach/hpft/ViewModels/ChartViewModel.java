package com.opensoach.hpft.ViewModels;

import java.util.ArrayList;
import java.util.Date;

import com.opensoach.hpft.Views.IRowClick;
import com.opensoach.hpft.Views.ITableClick;

/**
 * Created by samir.s.bukkawar on 3/6/2017.
 * <p>
 * This \model object is for entire Chart Table View
 */

public class ChartViewModel implements IRowClick {

    private Date taskStartTime;
    private Date taskEndTime;
    private int slotInterval;
    private String locationName;
    private ArrayList<TaskRowViewModel> taskRowViewModelList;
    private ArrayList<String> tableColumnTitleList;
    private ArrayList<String> tableRowTitleList;
    private ITableClick iTableClick;
    private TaskRowViewModel mTaskRowViewModel;
    private int chartId;

    private ChartTaskViewModel chartTaskViewModel;
    private ChartSlotHeaderViewModel chartSlotHeaderViewModel;

    public ChartViewModel(){
        taskRowViewModelList= new ArrayList<TaskRowViewModel>();
    }

    public Date getTaskStartTime() {
        return taskStartTime;
    }

    public void setTaskStartTime(Date taskStartTime) {
        this.taskStartTime = taskStartTime;
    }

    public Date getTaskEndTime() {
        return taskEndTime;
    }

    public void setTaskEndTime(Date taskEndTime) {
        this.taskEndTime = taskEndTime;
    }

    public int getSlotInterval() {
        return slotInterval;
    }

    public void setSlotInterval(int slotInterval) {
        this.slotInterval = slotInterval;
    }

    public String getLocationName() {
        return locationName;
    }

    public void setLocationName(String locationName) {
        this.locationName = locationName;
    }

    public ArrayList<TaskRowViewModel> getTaskRowViewModelList() {
        return taskRowViewModelList;
    }

    public void setTaskRowViewModelList(ArrayList<TaskRowViewModel> taskRowViewModelList) {
        this.taskRowViewModelList = taskRowViewModelList;
    }

    public ArrayList<String> getTableColumnTitleList() {
        return tableColumnTitleList;
    }

    public void setTableColumnTitleList(ArrayList<String> tableColumnTitleList) {
        this.tableColumnTitleList = tableColumnTitleList;
    }

    public ArrayList<String> getTableRowTitleList() {
        return tableRowTitleList;
    }

    public void setTableRowTitleList(ArrayList<String> tableRowTitleList) {
        this.tableRowTitleList = tableRowTitleList;
    }

    public ITableClick getiTableClick() {
        return iTableClick;
    }

    public void setiTableClick(ITableClick iTableClick) {
        this.iTableClick = iTableClick;
    }

    @Override
    public TaskRowViewModel getTaskRowViewModel() {
        return mTaskRowViewModel;
    }

    @Override
    public void onRowClick(TaskRowViewModel taskRowViewModel) {
        mTaskRowViewModel = taskRowViewModel;
        iTableClick.onChartTableClick(this);
    }

    public int getChartId() {
        return chartId;
    }

    public void setChartId(int chartId) {
        this.chartId = chartId;
    }


    public ChartTaskViewModel getChartTaskViewModel() {
        return chartTaskViewModel;
    }

    public void setChartTaskViewModel(ChartTaskViewModel chartTaskViewModel) {
        this.chartTaskViewModel = chartTaskViewModel;
    }

    public ChartSlotHeaderViewModel getChartSlotHeaderViewModel() {
        return chartSlotHeaderViewModel;
    }

    public void setChartSlotHeaderViewModel(ChartSlotHeaderViewModel chartSlotHeaderViewModel) {
        this.chartSlotHeaderViewModel = chartSlotHeaderViewModel;
    }
}