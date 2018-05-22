package spl.hkt.opensoach.splapp.viewModels;

import android.view.View;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.helper.ApplicationConstants;
import spl.hkt.opensoach.splapp.view.ICellClick;

/**
 * Created by samir.s.bukkawar on 3/6/2017.
 * <p>
 * This model is for each Cell in Chart Table View
 */

public class CellViewModel implements android.view.View.OnClickListener {

    private int TaskID;
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


    public int getTaskID() {
        return TaskID;
    }

    public void setTaskID(int taskID) {
        TaskID = taskID;
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
            checkBox.setBackgroundResource(R.drawable.checkboxselected);
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
                checkBox.setBackgroundResource(R.drawable.checkbox_selector);
                break;
            case ApplicationConstants.CHART_STATE_BLOCKED:
                checkBox.setBackgroundResource(R.drawable.checkbox_selector);
                break;
            case ApplicationConstants.CHART_STATE_ON_TIME:
                //  checkBox.setButtonDrawable(R.drawable.checkbox_selected_ontime);
                checkBox.setBackgroundResource(R.drawable.checkbox_selected_ontime);
                //   checkBox.setChecked(true);
                checkBox.setEnabled(false);
                break;
            case ApplicationConstants.CHART_STATE_DELAYED:
                //  checkBox.setButtonDrawable(R.drawable.checkbox_selected_delayed);
                checkBox.setBackgroundResource(R.drawable.checkbox_selected_delayed);
                //   checkBox.setChecked(true);
                checkBox.setEnabled(false);
                break;
        }
    }

    public boolean isChecked() {
        return isChecked;
    }

}
