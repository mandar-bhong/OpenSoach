package com.opensoach.hpft.ViewModels;

import android.databinding.BaseObservable;
import android.databinding.Bindable;

import com.opensoach.hpft.BR;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.Views.Adapter.TaskDataAdapter;
import com.opensoach.hpft.Views.Adapter.TaskTimeDataAdapter;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 02-08-2018.
 */

public class TaskTimeDataViewModel extends BaseObservable {
    private static final String TAG = "DataViewModel";
    private TaskTimeDataAdapter adapter;
    private List<TaskItemDataModel> data;

    public TaskTimeDataViewModel() {
        data = new ArrayList<>();
        adapter = new TaskTimeDataAdapter();
    }

    public void setUp() {
        // perform set up tasks, such as adding listeners, data population, etc.
        populateData();
    }

    public void tearDown() {
        // perform tear down tasks, such as removing listeners
    }

    @Bindable
    public List<TaskItemDataModel> getData() {
        return this.data;
    }

    @Bindable
    public TaskTimeDataAdapter getAdapter() {
        return this.adapter;
    }

    private void populateData() {
        // populate the data from the source, such as the database.
        for (int i = 0; i < 5; i++) {
            TaskItemDataModel dataModel = new TaskItemDataModel();
            dataModel.setTitle(String.valueOf(i));
            data.add(dataModel);
        }
        notifyPropertyChanged(BR.data);
    }
}
