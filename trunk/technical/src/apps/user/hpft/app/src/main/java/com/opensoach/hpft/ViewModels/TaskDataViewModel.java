/*
 * Copyright (c) 2018 Phunware Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:

 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.

 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package com.opensoach.hpft.ViewModels;

import android.databinding.BaseObservable;
import android.databinding.Bindable;


import com.opensoach.hpft.BR;
import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.Views.Adapter.TaskDataAdapter;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Gregory Rasmussen on 7/26/17.
 */
public class TaskDataViewModel extends BaseObservable {
    private static final String TAG = "DataViewModel";
    private TaskDataAdapter adapter;
    private List<TaskItemDataModel> data;

    public TaskDataViewModel() {
        data = new ArrayList<>();
        adapter = new TaskDataAdapter();
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
    public TaskDataAdapter getAdapter() {
        return this.adapter;
    }

    private void populateData() {
        // populate the data from the source, such as the database.
        for (int i = 0; i < 15; i++) {
            TaskItemDataModel dataModel = new TaskItemDataModel();
            dataModel.setTitle(String.valueOf(i));
            data.add(dataModel);
        }
        notifyPropertyChanged(BR.data);
    }
}
