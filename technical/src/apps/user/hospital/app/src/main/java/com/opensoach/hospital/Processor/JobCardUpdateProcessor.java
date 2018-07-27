package com.opensoach.hospital.Processor;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.DataConverter.DataModelConverter;
import com.opensoach.hospital.Model.Communication.PacketJobCardDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardsDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableQueryModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Model.View.UINewJobDataModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 05-12-2017.
 */

public class JobCardUpdateProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {
            PacketModel<PacketJobCardsDataModel> packetLocationDataModel = (PacketModel<PacketJobCardsDataModel>) resultModel.Packet.Payload;
            List<PacketJobCardDataModel> jobCardModels = packetLocationDataModel.Payload.JobCards;

            for (PacketJobCardDataModel jobCard:jobCardModels) {

                DBJobCardTableRowModel dbJobCardTableRowModel = DataModelConverter.ConvertToDBJobCard(jobCard, resultModel.Packet.Header.LocationID);
                DatabaseManager.UpdateRow(new DBJobCardTableQueryModel(), dbJobCardTableRowModel, DBJobCardTableQueryModel.UPDATE_JCODE_PARTCOUNT_STARTTIME_ENDTIME_BY_ID_AND_LOCATION_ID_FILTER);
            }

            //Return is user is at different location
            if(resultModel.Packet.Header.LocationID != AppRepo.getInstance().getCurrentLocationId()) {
                packetProcessResultModel.IsSuccess = true;
                return packetProcessResultModel;
            }

            DBJobCardTableRowModel locationJobsCardTableRowModels = new DBJobCardTableRowModel();
            locationJobsCardTableRowModels.setLocationId(resultModel.Packet.Header.LocationID);

            List<DBJobCardTableRowModel> locationJobCards = DatabaseManager.SelectByFilter(new DBJobCardTableQueryModel(),locationJobsCardTableRowModels, DBJobCardTableQueryModel.SELECT_LOCATION_ID_FILTER);

            List<JobBriefViewModel> jobBriefViewModels = new ArrayList<>() ;

            for (DBJobCardTableRowModel jobCardTableRowModel : locationJobCards){

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

            UINewJobDataModel uiNewJobDataModel = new UINewJobDataModel();
            uiNewJobDataModel.setJobBriefViewModels(jobBriefViewModels);
            packetProcessResultModel.IsSuccess = true;
            packetProcessResultModel.CanUpdateUI = true;
            packetProcessResultModel.UINotifierModel = uiNewJobDataModel;
            packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_NEW_JOB_AVAILABLE;



        } catch (Exception exeception) {
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception, "Error occured in JobCardUpdateProcessor");
        }
        return packetProcessResultModel;
    }
}
