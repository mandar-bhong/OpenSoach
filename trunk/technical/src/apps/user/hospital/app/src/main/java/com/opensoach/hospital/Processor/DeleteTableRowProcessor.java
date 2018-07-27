package com.opensoach.hospital.Processor;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.Helper.AppDBHelper;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Model.Communication.PacketDeleteRowDataModel;
import com.opensoach.hospital.Model.Communication.PacketDeleteTableRowDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Model.View.UIDeletedJobDataModel;
import com.opensoach.hospital.Model.View.UIJobStateChangedDataModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.List;

/**
 * Created by Mandar on 9/23/2017.
 */

public class DeleteTableRowProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketDeleteTableRowDataModel> packetEnggPartsDataModel = (PacketModel<PacketDeleteTableRowDataModel>)resultModel.Packet.Payload;
            List<PacketDeleteRowDataModel> deleteRowDataModels = packetEnggPartsDataModel.Payload.DeleteIDs;

            for (PacketDeleteRowDataModel packetDeleteRowDataModel:deleteRowDataModels) {

                switch (packetDeleteRowDataModel.TableID) {
                    case DBConstants.TABLE_JOB_TABLE_ID: {

                        UIDeletedJobDataModel uiDeletedJobDataModel = new UIDeletedJobDataModel();
                        //List<Integer> deletedJobs = new ArrayList<>() ;

                        for (int rowid: packetDeleteRowDataModel.RowIDs){
                            DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
                            dbJobCardTableRowModel.setJobCardId(rowid);

                            //TODO: Verify that row is deleted
                            //DatabaseManager.DeleteByFilter(new DBJobCardTableQueryModel(),dbJobCardTableRowModel,DBJobCardTableQueryModel.SELECT_ID_FILTER);

                            //deletedJobs.add(rowid);
                        }

                        if(resultModel.Packet.Header.LocationID == AppRepo.getInstance().getCurrentLocationId()){
                            //uiDeletedJobDataModel.setDeletedJobs(deletedJobs);

                            List<JobBriefViewModel> jobVM = AppDBHelper.GetAllJobsViewModels(AppRepo.getInstance().getCurrentLocationId());

                            UIJobStateChangedDataModel uiJobStateChangedDataModel = new UIJobStateChangedDataModel();
                            uiJobStateChangedDataModel.setJobBriefViewModels(jobVM);

                            packetProcessResultModel.CanUpdateUI = true;
                            packetProcessResultModel.UINotifierModel = uiJobStateChangedDataModel;
                            packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STATE_UPDATED;

                            packetProcessResultModel.IsSuccess = true;

                        }

                        packetProcessResultModel.IsSuccess = true;
                    }
                    break;
                }
            }
        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }

        return packetProcessResultModel;
    }
}
