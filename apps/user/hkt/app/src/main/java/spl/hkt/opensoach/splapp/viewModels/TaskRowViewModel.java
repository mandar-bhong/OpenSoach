package spl.hkt.opensoach.splapp.viewModels;

import java.util.ArrayList;
import java.util.Date;
import spl.hkt.opensoach.splapp.view.ICellClick;
import spl.hkt.opensoach.splapp.view.IRowClick;
import spl.hkt.opensoach.splapp.viewModels.CellViewModel;


/**
 * Created by samir.s.bukkawar on 3/6/2017.
 * <p>
 * This model is for each Task (Row in Chart Table)
 */

public class TaskRowViewModel implements ICellClick {

    private String taskName;
    private Date taskStartTime;
    private Date taskEndTime;
    private ArrayList<CellViewModel> cellViewModelList;
    private IRowClick iRowClick;
    private CellViewModel mCellViewModel;

    public TaskRowViewModel(){
        cellViewModelList =new ArrayList<CellViewModel>();
    }

    public String getTaskName() {
        return taskName;
    }

    public void setTaskName(String taskName) {
        this.taskName = taskName;
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

    public ArrayList<CellViewModel> getCellViewModelList() {
        return cellViewModelList;
    }

    public void setCellViewModelList(ArrayList<CellViewModel> cellViewModelList) {
        this.cellViewModelList = cellViewModelList;
    }

    public IRowClick getiRowClick() {
        return iRowClick;
    }

    public void setiRowClick(IRowClick iRowClick) {
        this.iRowClick = iRowClick;
    }

    public CellViewModel getmCellViewModel() {
        return mCellViewModel;
    }

    public void setmCellViewModel(CellViewModel mCellViewModel) {
        this.mCellViewModel = mCellViewModel;
    }

    @Override
    public CellViewModel getCellViewModel() {
        return mCellViewModel;
    }

    @Override
    public void onCellClick(CellViewModel cellViewModel) {
        mCellViewModel = cellViewModel;
       iRowClick.onRowClick(this);
    }
}
