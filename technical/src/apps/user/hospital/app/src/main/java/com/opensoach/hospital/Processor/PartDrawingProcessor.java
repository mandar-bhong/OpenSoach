package com.opensoach.hospital.Processor;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.Communication.PacketPartDrawingDataModel;
import com.opensoach.hospital.Model.Communication.PacketPartDrawingsDataModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableQueryModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.List;

/**
 * Created by Mandar on 9/10/2017.
 */

public class PartDrawingProcessor implements IProcessor  {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketPartDrawingsDataModel> packetEnggPartsDataModel = (PacketModel<PacketPartDrawingsDataModel>)resultModel.Packet.Payload;
            List<PacketPartDrawingDataModel> partDrawingModels = packetEnggPartsDataModel.Payload.PartDrawings;

            for (PacketPartDrawingDataModel partDrawingModel:partDrawingModels) {


                DBPartDrawingTableRowModel dbPartDrawingTableRowModel = new DBPartDrawingTableRowModel();
                dbPartDrawingTableRowModel.setDrawingId(partDrawingModel.DrawingID);
                dbPartDrawingTableRowModel.setPartId(partDrawingModel.PartID);
                dbPartDrawingTableRowModel.setPath(partDrawingModel.Path);

                List<DBPartDrawingTableRowModel> existingDrawing =  DatabaseManager.SelectByFilter(new DBPartDrawingTableQueryModel(),dbPartDrawingTableRowModel,DBPartDrawingTableQueryModel.SELECT_ID_FILTER);

                if(existingDrawing.size()== 0){
                    DatabaseManager.InsertRow(dbPartDrawingTableRowModel);
                }
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in EnggPartProcessor");
        }

        return packetProcessResultModel;
    }
}
