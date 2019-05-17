package com.opensoach.hpft.Views;


import com.opensoach.hpft.ViewModels.TaskRowViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface IRowClick {

    TaskRowViewModel getTaskRowViewModel();

    void onRowClick(TaskRowViewModel taskRowViewModel);

}