package com.opensoach.vst.Views;


import com.opensoach.vst.ViewModels.TaskRowViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface IRowClick {

    TaskRowViewModel getTaskRowViewModel();

    void onRowClick(TaskRowViewModel taskRowViewModel);

}