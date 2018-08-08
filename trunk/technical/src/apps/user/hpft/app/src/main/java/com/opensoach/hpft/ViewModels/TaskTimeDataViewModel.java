package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;

import com.opensoach.hpft.BR;
import com.opensoach.hpft.Model.View.TaskTimeItemDataModel;
import com.opensoach.hpft.Views.Adapter.TaskTimeDataAdapter;

import java.util.ArrayList;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;

/**
 * Created by Mandar on 02-08-2018.
 */

public class TaskTimeDataViewModel extends BaseViewModel {
    private static final String TAG = "DataViewModel";

    private TaskDataViewModel taskDataViewModel;

    private List<TaskTimeItemViewModel> data;

    private TaskTimeItemViewModel selectedItem;

    private List<TaskTimeItemDataModel> timeData;

    public TaskTimeDataViewModel(TaskDataViewModel taskDataViewModel,List<TaskTimeItemDataModel> timeSeries ) {
        data = new ArrayList<>();

        this.taskDataViewModel = taskDataViewModel;
        timeData = timeSeries;
    }

    public void setUp() {
        // perform set up tasks, such as adding listeners, data population, etc.
        populateData();
    }

    public void tearDown() {
        // perform tear down tasks, such as removing listeners
    }

    @Bindable
    public List<TaskTimeItemViewModel> getData() {
        return this.data;
    }

    private void populateData() {
        // populate the data from the source, such as the database.

        data.clear();

        for (int i = 0; i < timeData.size(); i++) {
            TaskTimeItemViewModel dataModel = new TaskTimeItemViewModel();
            dataModel.setTaskTimeDataModel(timeData.get(i));
            dataModel.Parent = this;
            data.add(dataModel);
        }

        notifyPropertyChanged(BR.data);
    }

    @Bindable
    public TaskDataViewModel getTaskDataViewModel() {
        return taskDataViewModel;
    }

    @Bindable
    public void setTaskDataViewModel(TaskDataViewModel taskDataViewModel) {
        this.taskDataViewModel = taskDataViewModel;
    }

    @Bindable
    public List<TaskTimeItemViewModel> getTimeData() {
        return data;
    }

    @Bindable
    public void setTimeData(List<TaskTimeItemDataModel> data) {

    }

    public void setSelectedTimeTaskItem(){
        ((TaskDetailsViewModel)this.Parent).getTaskDataAdapter().updateData(this.getTaskDataViewModel().getData());
        selectedItem = data.get(0);
    }

    public TaskTimeItemViewModel getSelectedItem() {
        return selectedItem;
    }

    public void setSelectedItem(TaskTimeItemViewModel selectItem) {
        this.selectedItem = selectItem;
    }
}
