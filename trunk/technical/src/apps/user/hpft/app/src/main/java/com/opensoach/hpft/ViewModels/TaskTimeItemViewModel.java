package com.opensoach.hpft.ViewModels;

import com.opensoach.hpft.Model.View.TaskTimeItemDataModel;

import java.text.Format;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * Created by Mandar on 07-08-2018.
 */

public class TaskTimeItemViewModel extends  BaseViewModel {

    private TaskTimeItemDataModel taskTimeDataModel;

    public TaskTimeItemDataModel getTaskTimeDataModel() {
        return taskTimeDataModel;
    }

    public void setTaskTimeDataModel(TaskTimeItemDataModel taskTimeDataModel) {
        this.taskTimeDataModel = taskTimeDataModel;
    }

    public String getDisplayText(){
        Format formatter = new SimpleDateFormat("hh:mm a");
        String slotDisplayText = formatter.format(taskTimeDataModel.getStartTime());
        return slotDisplayText;
    }

    public boolean getIsSelectedTime(){

        Date d = new Date();
        if (taskTimeDataModel.getStartTime().getTime() < d.getTime() &&
                taskTimeDataModel.getEndTime().getTime() >    d.getTime() ){

            ((TaskDetailsViewModel) ((TaskTimeDataViewModel)this.Parent).Parent).setSelectedItem(this);

            return  true;
        }else{
            return false;
        }
    }




}
