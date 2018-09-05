package com.opensoach.vst.Views.ClickHandler;

import android.app.AlertDialog;
import android.app.Dialog;
import android.content.DialogInterface;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.ViewModels.JobDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.Views.DialogHelper;

public class JobServiceTaskItemRemoveHandler {

    public void onClick(View view, JobServiceItemViewModel vm) {

        final JobServiceItemViewModel jobServiceItemViewModel = vm;

        AlertDialog.Builder builder = new AlertDialog.Builder(view.getContext());
        builder.setTitle("Confirm");
        builder.setMessage("Are you sure to delete?");

        builder.setPositiveButton("YES", new DialogInterface.OnClickListener() {

            public void onClick(DialogInterface dialog, int which) {
                AppRepo.getInstance().getJobServiceViewModel().
                        getJobServiceListViewModel().
                        getJobServiceDataAdapter().
                        removeItem(jobServiceItemViewModel);

                dialog.dismiss();
            }
        });

        builder.setNegativeButton("NO", new DialogInterface.OnClickListener() {

            @Override
            public void onClick(DialogInterface dialog, int which) {

                // Do nothing
                dialog.dismiss();
            }
        });

        AlertDialog alert = builder.create();
        alert.show();

    }

}
