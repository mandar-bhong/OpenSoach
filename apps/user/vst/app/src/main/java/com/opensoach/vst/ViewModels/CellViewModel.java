package com.opensoach.vst.ViewModels;

import android.view.View;

import com.opensoach.vst.R;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Views.ICellClick;

/**
 * Created by samir.s.bukkawar on 3/6/2017.
 * <p>
 * This model is for each Cell in Chart Table View
 */

public class CellViewModel implements android.view.View.OnClickListener {

    private String TaskName;
    private int slotID;
    private ICellClick iCellClick;
    private View checkBox;
    private boolean isChecked;

   /* private Date CellStartTime;
    private Date CellEndTime;
    private String TaskName;
    private int cellState;
    private Color cellColor;
    private int CellColumnID;
    private int CellRowID;
    private Date taskCompletionTime;
    private boolean isCellSynced;*/


    public String getTaskID() {
        return TaskName;
    }

    public void setTaskName(String taskName) {
        TaskName = taskName;
    }

    public int getSlotID() {
        return slotID;
    }

    public void setSlotID(int slotID) {
        this.slotID = slotID;
    }

    public ICellClick getiCellClick() {
        return iCellClick;
    }

    public void setiCellClick(ICellClick iCellClick) {
        this.iCellClick = iCellClick;
    }


    @Override
    public void onClick(View v) {
        isChecked = !isChecked;
        iCellClick.onCellClick(this);
        if (isChecked)
            checkBox.setBackgroundResource(R.drawable.custom_cell_checked);
        else
            checkBox.setBackgroundResource(R.drawable.custom_cell_available);
    }

    public View getCheckBox() {
        return checkBox;
    }

    public void setCheckBox(View checkBox) {
        this.checkBox = checkBox;
    }

    public void setState(int state) {
        switch (state) {
            case ApplicationConstants.CHART_STATE_ENABLED:
                checkBox.setBackgroundResource(R.drawable.custom_cell_available);
                break;
            case ApplicationConstants.CHART_STATE_BLOCKED:
                checkBox.setBackgroundResource(R.drawable.custom_cell_not_available);
                break;
            case ApplicationConstants.CHART_STATE_ON_TIME:
                checkBox.setBackgroundResource(R.drawable.custom_cell_on_time);
                checkBox.setEnabled(false);
                break;
            case ApplicationConstants.CHART_STATE_DELAYED:
                checkBox.setBackgroundResource(R.drawable.custom_cell_delayed);
                checkBox.setEnabled(false);
                break;
        }
    }

    public boolean isChecked() {
        return isChecked;
    }

}
