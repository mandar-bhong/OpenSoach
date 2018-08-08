package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

import com.opensoach.hpft.Model.Communication.PacketServiceConfModel;
import com.opensoach.hpft.Model.View.TaskTimeItemDataModel;
import com.opensoach.hpft.Views.Adapter.TaskDataAdapter;
import com.opensoach.hpft.Views.Adapter.TaskTimeDataAdapter;

import java.util.HashMap;
import java.util.List;

/**
 * Created by Mandar on 01-08-2018.
 */

public class TaskDetailsViewModel extends BaseViewModel {

    private TaskTimeDataViewModel taskTimeDataViewModel;
    private String title;

    private TaskTimeDataAdapter taskTimeDataAdapter;
    private TaskDataAdapter taskDataAdapter;

    private PacketServiceConfModel packetServiceConf;

    public TaskDetailsViewModel() {
        taskTimeDataAdapter = new TaskTimeDataAdapter();
        taskDataAdapter = new TaskDataAdapter();
    }

    public TaskDetailsViewModel(PacketServiceConfModel packetServiceConfModel,
                                List<TaskTimeItemDataModel> timeSeries ) {

        this();

        this.packetServiceConf = packetServiceConfModel;

        TaskDataViewModel taskDataViewModel = new TaskDataViewModel();
        taskDataViewModel.setUp(packetServiceConfModel.TaskConfig);
        taskTimeDataViewModel = new TaskTimeDataViewModel(taskDataViewModel,timeSeries );
        taskTimeDataViewModel.Parent = this;
    }

    @Bindable
    public TaskTimeDataAdapter getTaskTimeDataAdapter() {
        return this.taskTimeDataAdapter;
    }

    @Bindable
    public TaskDataAdapter getTaskDataAdapter() {
        return this.taskDataAdapter;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public TaskTimeDataViewModel getTaskTimeDataViewModel() {
        return taskTimeDataViewModel;
    }

}
