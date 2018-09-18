package com.opensoach.vst.Views.Adapter;

import android.databinding.DataBindingUtil;
import android.support.annotation.Nullable;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.Views.ClickHandler.JobExeDetailsClickHandler;
import com.opensoach.vst.Views.ClickHandler.JobServiceCreationHandler;
import com.opensoach.vst.Views.ClickHandler.JobServiceExeHandler;
import com.opensoach.vst.databinding.FragmentJobServiceItemBinding;

import java.util.ArrayList;
import java.util.List;

public class JobServiceDataAdapter extends RecyclerView.Adapter<JobServiceDataAdapter.DataViewHolder> {

    private List<JobServiceItemViewModel> data;

    public JobServiceDataAdapter() {
        this.data = new ArrayList<>();
    }

    @Override
    public JobServiceDataAdapter.DataViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View itemView = LayoutInflater.from(parent.getContext()).inflate(R.layout.fragment_job_service_item,
                new FrameLayout(parent.getContext()), false);
        return new JobServiceDataAdapter.DataViewHolder(itemView);
    }

    @Override
    public void onBindViewHolder(JobServiceDataAdapter.DataViewHolder holder, int position) {
        JobServiceItemViewModel dataModel = data.get(position);
        holder.setViewModel(dataModel);
    }

    @Override
    public int getItemCount() {
        return this.data.size();
    }


    @Override
    public void onViewAttachedToWindow(JobServiceDataAdapter.DataViewHolder holder) {
        super.onViewAttachedToWindow(holder);
        holder.bind();
    }

    @Override
    public void onViewDetachedFromWindow(JobServiceDataAdapter.DataViewHolder holder) {
        super.onViewDetachedFromWindow(holder);
        holder.unbind();
    }

    public  void NotifyDataSetChanged(){
        notifyDataSetChanged();
    }

    public void updateData(@Nullable List<JobServiceItemViewModel> data) {
        this.data = data;
        notifyDataSetChanged();
    }

    public void addItem(@Nullable JobServiceItemViewModel item) {
        if (this.data != null){

            this.data.add(item);

            notifyDataSetChanged();
        }
    }


    public void removeItem(@Nullable JobServiceItemViewModel item) {
        if (this.data != null){

            this.data.remove(item);

            notifyDataSetChanged();
        }
    }



    static class DataViewHolder extends RecyclerView.ViewHolder {
        /* package */ FragmentJobServiceItemBinding binding;

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

        /* package */ void setViewModel(JobServiceItemViewModel viewModel) {
            if (binding != null) {
                binding.setVM(viewModel);
				binding.setClickHandler(new JobServiceCreationHandler());
				binding.setExeHandler(new JobServiceExeHandler());
            }
        }


    }
}


