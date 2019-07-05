package com.opensoach.vst.Views.DataBinding;

import android.databinding.BindingAdapter;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.Views.Adapter.JobServiceDataAdapter;
import com.opensoach.vst.Views.Adapter.TokensDataAdapter;

import java.util.List;

public class JobServiceRecycleViewDataBinding {

    @BindingAdapter({"app:adapter", "app:data"})
    public void bind(RecyclerView recyclerView, JobServiceDataAdapter adapter, List<JobServiceItemViewModel> data) {
        recyclerView.setAdapter(adapter);
        adapter.updateData(data);
    }
}
