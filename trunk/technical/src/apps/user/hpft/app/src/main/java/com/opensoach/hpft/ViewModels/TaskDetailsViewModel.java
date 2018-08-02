package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

/**
 * Created by Mandar on 01-08-2018.
 */

public class TaskDetailsViewModel extends BaseViewModel {

    private TaskDataViewModel taskDataViewModel;
    private TaskTimeDataViewModel taskTimeDataViewModel;
    private String title;


@Bindable
    public TaskDataViewModel getTaskDataViewModel() {
        return taskDataViewModel;
    }

    @Bindable
    public void setTaskDataViewModel(TaskDataViewModel taskDataViewModel) {
        this.taskDataViewModel = taskDataViewModel;
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

    public void setTaskTimeDataViewModel(TaskTimeDataViewModel taskTimeDataViewModel) {
        this.taskTimeDataViewModel = taskTimeDataViewModel;
    }
}
