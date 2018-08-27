package com.opensoach.vst.Views.Adapter;

import android.databinding.DataBindingUtil;
import android.support.annotation.Nullable;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.TaskTimeItemViewModel;
import com.opensoach.vst.Views.ClickHandler.TaskTimeClickHandler;
import com.opensoach.vst.databinding.FragmentTaskTimeItemBinding;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 02-08-2018.
 */

public class TaskTimeDataAdapter extends RecyclerView.Adapter<TaskTimeDataAdapter.DataViewHolder> {
    private static final String TAG = "TimeDataAdapter";
    private List<TaskTimeItemViewModel> data;

    public TaskTimeDataAdapter() {
        this.data = new ArrayList<>();
    }

    @Override
    public DataViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View itemView = LayoutInflater.from(parent.getContext()).inflate(R.layout.fragment_task_time_item,
                new FrameLayout(parent.getContext()), false);
        return new DataViewHolder(itemView);
    }

    @Override
    public void onBindViewHolder(DataViewHolder holder, int position) {
        TaskTimeItemViewModel dataModel = data.get(position);
        holder.setViewModel(dataModel);
    }

    @Override
    public int getItemCount() {
        return this.data.size();
    }

    @Override
    public void onViewAttachedToWindow(DataViewHolder holder) {
        super.onViewAttachedToWindow(holder);
        holder.bind();
    }

    @Override
    public void onViewDetachedFromWindow(DataViewHolder holder) {
        super.onViewDetachedFromWindow(holder);
        holder.unbind();
    }

    public void updateData(@Nullable List<TaskTimeItemViewModel> data) {

//        if (data == null || data.isEmpty()) {
//            this.data.clear();
//        } else {
//            this.data.addAll(data);
//        }

        this.data = data;

        notifyDataSetChanged();
    }

    static class DataViewHolder extends RecyclerView.ViewHolder {
        /* package */ FragmentTaskTimeItemBinding binding;

        /* package */ DataViewHolder(View itemView) {
            super(itemView);
            bind();
        }

        /* package */ void bind() {
            if (binding == null) {
                binding = DataBindingUtil.bind(itemView);
            }
        }

        /* package */ void unbind() {
            if (binding != null) {
                binding.unbind(); // Don't forget to unbind
            }
        }

        /* package */ void setViewModel(TaskTimeItemViewModel viewModel) {
            if (binding != null) {
                binding.setVM(viewModel);
                binding.setClickHandler(new TaskTimeClickHandler());
            }
        }
    }
}