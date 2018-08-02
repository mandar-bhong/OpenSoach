package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 01-08-2018.
 */

public class TaskDetailsViewModel extends BaseViewModel {

    private TaskDataViewModel dataViewModel;
    private String title;


@Bindable
    public TaskDataViewModel getDataViewModel() {
        return dataViewModel;
    }

    @Bindable
    public void setDataViewModel(TaskDataViewModel dataViewModel) {
        this.dataViewModel = dataViewModel;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }
}
