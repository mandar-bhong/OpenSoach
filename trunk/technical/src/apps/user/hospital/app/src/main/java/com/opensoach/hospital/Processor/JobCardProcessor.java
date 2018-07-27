package com.opensoach.hospital.Processor;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.DataConverter.DataModelConverter;
import com.opensoach.hospital.Model.Communication.PacketJobCardDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardsDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.List;

/**
 * Created by Mandar on 9/4/2017.
 */

public class JobCardProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketJobCardsDataModel> packetLocationDataModel = (PacketModel<PacketJobCardsDataModel>)resultModel.Packet.Payload;
            List<PacketJobCardDataModel> jobCardModels = packetLocationDataModel.Payload.JobCards;

            for (PacketJobCardDataModel jobCard:jobCardModels) {

                DBJobCardTableRowModel dbJobCardTableRowModel = DataModelConverter.ConvertToDBJobCard(jobCard,resultModel.Packet.Header.LocationID);

                DatabaseManager.InsertRow(dbJobCardTableRowModel);
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in JobCardProcessor");
        }

        return packetProcessResultModel;
    }
}
