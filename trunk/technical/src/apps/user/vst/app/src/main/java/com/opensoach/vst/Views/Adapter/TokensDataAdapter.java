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
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.Views.ClickHandler.TaskTimeClickHandler;
import com.opensoach.vst.databinding.FragmentTaskTimeItemBinding;
import com.opensoach.vst.databinding.FragmentTokenItemBinding;

import java.util.ArrayList;
import java.util.List;

public class TokensDataAdapter extends RecyclerView.Adapter<TokensDataAdapter.DataViewHolder>  {

    private static final String TAG = "DataAdapter";
    private List<TokenItemViewModel> data;

    public TokensDataAdapter() {
        this.data = new ArrayList<>();
    }

    @Override
    public TokensDataAdapter.DataViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View itemView = LayoutInflater.from(parent.getContext()).inflate(R.layout.fragment_token_item,
                new FrameLayout(parent.getContext()), false);
        return new TokensDataAdapter.DataViewHolder(itemView);
    }

    @Override
    public void onBindViewHolder(TokensDataAdapter.DataViewHolder holder, int position) {
        TokenItemViewModel dataModel = data.get(position);
        holder.setViewModel(dataModel);
    }

    @Override
    public int getItemCount() {
        return this.data.size();
    }

    @Override
    public void onViewAttachedToWindow(TokensDataAdapter.DataViewHolder holder) {
        super.onViewAttachedToWindow(holder);
        holder.bind();
    }

    @Override
    public void onViewDetachedFromWindow(TokensDataAdapter.DataViewHolder holder) {
        super.onViewDetachedFromWindow(holder);
        holder.unbind();
    }

    public void updateData(@Nullable List<TokenItemViewModel> data) {

//        if (data == null || data.isEmpty()) {
//            this.data.clear();
//        } else {
//            this.data.addAll(data);
//        }

        this.data = data;

        notifyDataSetChanged();
    }

    static class DataViewHolder extends RecyclerView.ViewHolder {
        /* package */ FragmentTokenItemBinding binding;

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

        /* package */ void setViewModel(TokenItemViewModel viewModel) {
            if (binding != null) {
                binding.setVM(viewModel);
                //binding.setClickHandler(new TaskTimeClickHandler());
            }
        }
    }
}
