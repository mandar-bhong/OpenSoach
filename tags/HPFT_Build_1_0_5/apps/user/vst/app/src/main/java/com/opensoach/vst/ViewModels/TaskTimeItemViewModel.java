package com.opensoach.vst.ViewModels;

import android.databinding.Bindable;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.BR;
import com.opensoach.vst.Model.View.TaskTimeItemDataModel;

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

    @Bindable
    public boolean isActiveSelected() {
        if (AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem() == this ){
            return true;
        }else{
            return false;
        }
    }


    public void setActiveSelected() {
        notifyPropertyChanged(BR.activeSelected);
    }


    public boolean getIsDisabled(){

        Date d = new Date();
        if (taskTimeDataModel.getStartTime().getTime() < d.getTime() &&
                taskTimeDataModel.getEndTime().getTime() >    d.getTime() ){
            return  false;
        }

        if (taskTimeDataModel.getStartTime().getTime() < d.getTime() )
            return  false;
        else
            return  true;
    }




}
