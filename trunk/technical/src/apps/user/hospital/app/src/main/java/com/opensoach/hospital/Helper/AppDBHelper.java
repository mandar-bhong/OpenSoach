package com.opensoach.hospital.Helper;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.DB.DBEnggPartTableQueryModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 9/24/2017.
 */

public class AppDBHelper {

    public static List<JobBriefViewModel> GetAllJobsViewModels(int locationID){

        DBJobCardTableRowModel dbJobCardTableRowModel = new  DBJobCardTableRowModel();
        dbJobCardTableRowModel.setLocationId(locationID);

        List<DBJobCardTableRowModel> jobCards = DatabaseManager.SelectByFilter(new DBJobCardTableQueryModel(),dbJobCardTableRowModel, DBJobCardTableQueryModel.SELECT_LOCATION_ID_FILTER);


        List<JobBriefViewModel> jobBriefViewModels = new ArrayList<>() ;

        for (DBJobCardTableRowModel jobCardTableRowModel : jobCards){

            JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
            jobBriefViewModels.add(jobBriefViewModel);

            jobBriefViewModel.setDbJobCardTableRowModel(jobCardTableRowModel);

            DBEnggPartTableRowModel dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
            dbEnggPartTableRowModel.setPartId(jobCardTableRowModel.getPartId());

            List<DBEnggPartTableRowModel> parts =  DatabaseManager.SelectByFilter(new DBEnggPartTableQueryModel(),dbEnggPartTableRowModel,DBEnggPartTableQueryModel.SELECT_ID_FILTER);

            if(parts.size() >0){
                jobBriefViewModel.setDbEnggPartTableRowModel(parts.get(0));
            }
        }

        return jobBriefViewModels;
    }
}
