package com.opensoach.hospital.Processor.AckProcessor;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Helper.PacketHelper;
import com.opensoach.hospital.Model.DB.DBEnggPartTableQueryModel;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableQueryModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBLocationTableQueryModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableQueryModel;
import com.opensoach.hospital.Model.DB.DBPartDrawingTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Processor.IProcessor;
import com.opensoach.hospital.Utility.AppLogger;

/**
 * Created by Mandar on 9/5/2017.
 */

public class DeviceRegistrationProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {
            DatabaseManager.DeleteByFilter(new DBLocationTableQueryModel(), new DBLocationTableRowModel(), DBLocationTableQueryModel.SELECT_ALL_FILTER);
            DatabaseManager.DeleteByFilter(new DBEnggPartTableQueryModel(), new DBEnggPartTableRowModel(), DBEnggPartTableQueryModel.SELECT_ALL_FILTER);
            DatabaseManager.DeleteByFilter(new DBJobCardTableQueryModel(), new DBJobCardTableRowModel(), DBJobCardTableQueryModel.SELECT_ALL_FILTER);
            DatabaseManager.DeleteByFilter(new DBPartDrawingTableQueryModel(),new DBPartDrawingTableRowModel(),DBPartDrawingTableQueryModel.SELECT_ALL_FILTER);

            packetProcessResultModel.ServerCommandPacket = PacketHelper.GetDeviceSynCompletedPacket();
            ;
            packetProcessResultModel.CanSendServerCommand = true;
            packetProcessResultModel.IsSuccess = true;

        }catch (Exception ex){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(ex,"Error occured in DeviceRegistrationProcessor");
        }

        return packetProcessResultModel;
    }
}
