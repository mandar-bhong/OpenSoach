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
package com.opensoach.hpft.Views.Adapter;

import android.databinding.DataBindingUtil;
import android.support.annotation.Nullable;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.FrameLayout;


import com.opensoach.hpft.Model.View.TaskItemDataModel;
import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.TaskItemViewModel;
import com.opensoach.hpft.Views.ClickHandler.TaskDetailsClickHandler;
import com.opensoach.hpft.databinding.FragmentTaskItemBinding;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Gregory Rasmussen on 7/26/17.
 */
public class TaskDataAdapter extends RecyclerView.Adapter<TaskDataAdapter.DataViewHolder> {
    private static final String TAG = "DataAdapter";
    private List<TaskItemDataModel> data;

    public TaskDataAdapter() {
        this.data = new ArrayList<>();
    }

    @Override
    public DataViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View itemView = LayoutInflater.from(parent.getContext()).inflate(R.layout.fragment_task_item,
                new FrameLayout(parent.getContext()), false);
        return new DataViewHolder(itemView);
    }

    @Override
    public void onBindViewHolder(DataViewHolder holder, int position) {
        TaskItemDataModel dataModel = data.get(position);
        holder.setViewModel(new TaskItemViewModel(dataModel));
    }

    @Override
    public int getItemCount() {
        return this.data == null ? 0 : this.data.size();
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

    public void updateData(@Nullable List<TaskItemDataModel> data) {

//        if (data == null || data.isEmpty()) {
//            this.data.clear();
//        } else {
//            this.data.addAll(data);
//        }

        this.data = data;

        notifyDataSetChanged();
    }

    static class DataViewHolder extends RecyclerView.ViewHolder {
        /* package */ FragmentTaskItemBinding binding;

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

        /* package */ void setViewModel(TaskItemViewModel viewModel) {
            if (binding != null) {
                binding.setVM(viewModel);
                binding.setClickHandler(new TaskDetailsClickHandler());
            }
        }
    }
}
