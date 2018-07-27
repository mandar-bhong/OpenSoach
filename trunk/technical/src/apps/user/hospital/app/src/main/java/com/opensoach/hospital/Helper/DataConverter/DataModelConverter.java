package com.opensoach.hospital.Helper.DataConverter;

import com.opensoach.hospital.Model.Communication.PacketEnggPartToolDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardDataModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

/**
 * Created by Mandar on 9/8/2017.
 */

public class DataModelConverter {

    public  static JobBriefViewModel JobCard(DBJobCardTableRowModel dbModel){
        JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
//        jobBriefViewModel.setJob(dbModel.getJobCode());
//        //jobBriefViewModel.setPart(dbModel.getPartNumber());
//        //TODO: Set customer info
//        //jobBriefViewModel.setCustomer();
//        jobBriefViewModel.setTargatedDate(dbModel.getStartDate().toString());
//        jobBriefViewModel.setQuantity(dbModel.getPartCount().toString());

        return jobBriefViewModel;
    }


    public  static DBJobCardTableRowModel ConvertToDBJobCard(PacketJobCardDataModel jobCard,int locationID){

        DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
        dbJobCardTableRowModel.setJobCardId(jobCard.JobID);
        dbJobCardTableRowModel.setCustomer(jobCard.Customer);
        dbJobCardTableRowModel.setLocationId(locationID);
        dbJobCardTableRowModel.setPartId(jobCard.PartID);
        dbJobCardTableRowModel.setPartCount(jobCard.PartCount);
        dbJobCardTableRowModel.setJobCode(jobCard.JobCode);
        dbJobCardTableRowModel.setStartDate(jobCard.StartTime);
        dbJobCardTableRowModel.setEndDate(jobCard.EndTime);
        dbJobCardTableRowModel.setActualStartDate(jobCard.ActualStartTime);
        dbJobCardTableRowModel.setActualEndDate(jobCard.ActualEndTime);
        dbJobCardTableRowModel.setState(jobCard.State);
        dbJobCardTableRowModel.setComments(jobCard.Comments);
        dbJobCardTableRowModel.setCompletedCount(jobCard.QuantityCompleted);
        dbJobCardTableRowModel.setJobConfig(jobCard.JobConfig);

        return dbJobCardTableRowModel;
    }


    public  static DBEnggPartTableRowModel ConvertToDBEnggPart(PacketEnggPartToolDataModel packetEnggPartToolDataModel,String enggPartTools){

        DBEnggPartTableRowModel dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
        dbEnggPartTableRowModel.setPartId(packetEnggPartToolDataModel.EnggPart.PartID);
        dbEnggPartTableRowModel.setPartNo(packetEnggPartToolDataModel.EnggPart.PartNo);
        dbEnggPartTableRowModel.setPartRevision(packetEnggPartToolDataModel.EnggPart.PartRevision);
        dbEnggPartTableRowModel.setInternalPartNo(packetEnggPartToolDataModel.EnggPart.InternalPartNo);
        dbEnggPartTableRowModel.setProcess(packetEnggPartToolDataModel.EnggPart.Process);
        dbEnggPartTableRowModel.setToolJSON(enggPartTools);

        return dbEnggPartTableRowModel;
    }
}
