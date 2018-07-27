package com.opensoach.hospital.Processor;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.AppDBHelper;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Model.Communication.PacketJobCardStatusChangedDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardsStatusChangedDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Model.View.UIJobStateChangedDataModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.List;

/**
 * Created by Mandar on 9/24/2017.
 */

public class JobCardStateChangeProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketJobCardsStatusChangedDataModel> packetLocationDataModel = (PacketModel<PacketJobCardsStatusChangedDataModel>)resultModel.Packet.Payload;
            List<PacketJobCardStatusChangedDataModel> jobCardModels = packetLocationDataModel.Payload.JobCards;

            for (PacketJobCardStatusChangedDataModel jobCardStateModel:jobCardModels) {

                switch (jobCardStateModel.State) {
                    case 1: // Job Started
                    {
                        DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
                        dbJobCardTableRowModel.setJobCardId(jobCardStateModel.JobCardID);
                        dbJobCardTableRowModel.setState(jobCardStateModel.State);
                        dbJobCardTableRowModel.setActualStartDate(jobCardStateModel.StateChangedTime);
                        DatabaseManager.UpdateRow(new DBJobCardTableQueryModel(),dbJobCardTableRowModel,DBJobCardTableQueryModel.UPDATE_STATE_AND_ACTUAL_START_TIME_BY_ID_FILTER);

                        List<JobBriefViewModel> jobVM = AppDBHelper.GetAllJobsViewModels(AppRepo.getInstance().getCurrentLocationId());

                        UIJobStateChangedDataModel uiJobStateChangedDataModel = new UIJobStateChangedDataModel();
                        uiJobStateChangedDataModel.setJobBriefViewModels(jobVM);

                        packetProcessResultModel.CanUpdateUI = true;
                        packetProcessResultModel.UINotifierModel = uiJobStateChangedDataModel;
                        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STATE_UPDATED;

                        packetProcessResultModel.IsSuccess = true;
                    }
                    break;
                    case 2: // Job Completed
                    {
                        DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
                        dbJobCardTableRowModel.setJobCardId(jobCardStateModel.JobCardID);
                        dbJobCardTableRowModel.setState(jobCardStateModel.State);
                        DatabaseManager.DeleteByFilter(new DBJobCardTableQueryModel(),dbJobCardTableRowModel,DBJobCardTableQueryModel.SELECT_ID_FILTER);

                        List<JobBriefViewModel> jobVM = AppDBHelper.GetAllJobsViewModels(AppRepo.getInstance().getCurrentLocationId());

                        UIJobStateChangedDataModel uiJobStateChangedDataModel = new UIJobStateChangedDataModel();
                        uiJobStateChangedDataModel.setJobBriefViewModels(jobVM);

                        packetProcessResultModel.CanUpdateUI = true;
                        packetProcessResultModel.UINotifierModel = uiJobStateChangedDataModel;
                        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STATE_UPDATED;

                        packetProcessResultModel.IsSuccess = true;
                    }
                    break;
                }
            }

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in JobCardProcessor");
        }

        return packetProcessResultModel;
    }
}
