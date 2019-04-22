package spl.hkt.opensoach.splapp.view;


import spl.hkt.opensoach.splapp.viewModels.TaskRowViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface IRowClick {

    TaskRowViewModel getTaskRowViewModel();

    void onRowClick(TaskRowViewModel taskRowViewModel);

}