package com.opensoach.hospital.Processor;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Model.DB.DBEnggPartTableQueryModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBLocationTableQueryModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Model.View.UIServerSyncCompletedModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 9/5/2017.
 */

public class ServerSynCompletedProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            DBJobCardTableRowModel dbJobCardTableRowModel = new  DBJobCardTableRowModel();
            dbJobCardTableRowModel.setLocationId(AppRepo.getInstance().getCurrentLocationId());

            List<DBJobCardTableRowModel> jobCards = DatabaseManager.SelectByFilter(new DBJobCardTableQueryModel(),dbJobCardTableRowModel, DBJobCardTableQueryModel.SELECT_LOCATION_ID_FILTER);

            List<JobBriefViewModel> jobBriefViewModels = new ArrayList<>() ;

            for (DBJobCardTableRowModel jobCardTableRowModel : jobCards){

                JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
                jobBriefViewModels.add(jobBriefViewModel);

                DBEnggPartTableRowModel dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
                dbEnggPartTableRowModel.setPartId(jobCardTableRowModel.getPartId());

                List<DBEnggPartTableRowModel> parts =  DatabaseManager.SelectByFilter(new DBEnggPartTableQueryModel(),dbEnggPartTableRowModel,DBEnggPartTableQueryModel.SELECT_ID_FILTER);

                List<DBEnggPartTableRowModel> parts1 =  DatabaseManager.SelectAll(new DBEnggPartTableQueryModel(),dbEnggPartTableRowModel);

                if(parts.size() >0){
                    jobBriefViewModel.setDbEnggPartTableRowModel(parts.get(0));
                }

                jobBriefViewModel.setDbJobCardTableRowModel(jobCardTableRowModel);
            }

            List<DBLocationTableRowModel> locations = DatabaseManager.SelectAll(new DBLocationTableQueryModel(),new  DBLocationTableRowModel());

            UIServerSyncCompletedModel uiServerSyncCompletedModel = new UIServerSyncCompletedModel();
            uiServerSyncCompletedModel.setJobBriefViewModels(jobBriefViewModels);
            uiServerSyncCompletedModel.setLocations(locations);
            packetProcessResultModel.IsSuccess = true;
            packetProcessResultModel.CanUpdateUI = true;
            packetProcessResultModel.UINotifierModel = uiServerSyncCompletedModel;
            packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_SERVER_DATA_LOAD_COMPLETED;

            AppRepo.getInstance().setIsStartupCompleted();

        }catch (Exception ex){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(ex,"Error occured in ServerSynCompletedProcessor");
        }

        return packetProcessResultModel;
    }
}
