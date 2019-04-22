package com.opensoach.hpft.Views.DataBinding;

import android.databinding.BindingAdapter;
import android.support.v7.widget.RecyclerView;

import com.opensoach.hpft.ViewModels.TaskTimeItemViewModel;
import com.opensoach.hpft.Views.Adapter.TaskTimeDataAdapter;

import java.util.List;

/**
 * Created by Mandar on 02-08-2018.
 */


public class TaskTimeRecycleViewDataBinding {
    /**
     * Binds the data to the {@link android.support.v7.widget.RecyclerView.Adapter}, sets the
     * recycler view on the adapter, and performs some other recycler view initialization.
     *
     * @param recyclerView passed in automatically with the data binding
     * @param adapter      must be explicitly passed in
     * @param data         must be explicitly passed in
     */
    @BindingAdapter({"app:adapter", "app:data"})
    public void bind(RecyclerView recyclerView, TaskTimeDataAdapter adapter, List<TaskTimeItemViewModel> data) {
        recyclerView.setAdapter(adapter);
        adapter.updateData(data);

    }
}