package com.opensoach.hospital.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableQueryModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableRowModel;
import com.opensoach.hospital.ViewModels.JobBoardViewModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.Views.Activity.JobBoardActivity;

import java.util.List;

/**
 * Created by Mandar on 8/26/2017.
 */

public class JobBriefClickHandler {

    public void onClick(View view, JobBriefViewModel vm) {

        JobBoardViewModel jobBoardViewModel = new JobBoardViewModel();

        DBPartDrawingTableRowModel dbPartDrawingTableRowModel = new DBPartDrawingTableRowModel();
        dbPartDrawingTableRowModel.setPartId(vm.getDbEnggPartTableRowModel().getPartId());
        List<DBPartDrawingTableRowModel> partDrawings = DatabaseManager.SelectByFilter(new DBPartDrawingTableQueryModel(),dbPartDrawingTableRowModel,DBPartDrawingTableQueryModel.SELECT_BY_PART_ID_FILTER);

        jobBoardViewModel.setDbPartDrawingTableRowModels(partDrawings);

        jobBoardViewModel.setDbEnggPartTableRowModel(vm.getDbEnggPartTableRowModel());
        jobBoardViewModel.setDbJobCardTableRowModel(vm.getDbJobCardTableRowModel());

        AppRepo.getInstance().setSelectedJobBoard(jobBoardViewModel);

        Intent i = new Intent(vm.ContextActivity.getBaseContext(), JobBoardActivity.class);
        vm.ContextActivity.startActivity(i);

    }
}
