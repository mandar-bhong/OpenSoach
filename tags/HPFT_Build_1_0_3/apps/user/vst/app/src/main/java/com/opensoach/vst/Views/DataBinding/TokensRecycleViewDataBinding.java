package com.opensoach.vst.Views.DataBinding;

import android.databinding.BindingAdapter;
import android.support.v7.widget.RecyclerView;

import com.opensoach.vst.ViewModels.TaskTimeItemViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.Views.Adapter.TaskTimeDataAdapter;
import com.opensoach.vst.Views.Adapter.TokensDataAdapter;

import java.util.List;

public class TokensRecycleViewDataBinding {

    @BindingAdapter({"app:adapter", "app:data"})
    public void bind(RecyclerView recyclerView, TokensDataAdapter adapter, List<TokenItemViewModel> data) {
        recyclerView.setAdapter(adapter);
        adapter.updateData(data);
    }
}
