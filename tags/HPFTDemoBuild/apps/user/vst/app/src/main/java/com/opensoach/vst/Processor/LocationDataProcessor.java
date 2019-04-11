package com.opensoach.vst.Processor;

import android.util.Log;

import java.util.List;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.Communication.PacketLocationDataModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.DB.DBLocationTableQueryModel;
import com.opensoach.vst.Model.DB.DBLocationTableRowModel;

/**
 * Created by Mandar on 4/8/2017.
 */

public class LocationDataProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketLocationDataModel[]> packetLocationDataModel = (PacketModel<PacketLocationDataModel[]>) resultModel.Packet;

            List<DBLocationTableRowModel> dbLocationModels = DatabaseManager.SelectAll(new DBLocationTableQueryModel(), new DBLocationTableRowModel());

            for (PacketLocationDataModel location : packetLocationDataModel.Payload) {
                DBLocationTableRowModel existingLocation = null;
                for (DBLocationTableRowModel dbLocationTableRowModel : dbLocationModels) {
                    if (dbLocationTableRowModel.getLocationId() == location.SPID) {
                        existingLocation = dbLocationTableRowModel;
                        break;
                    }
                }

                if (existingLocation != null) {
                    DatabaseManager.DeleteByFilter(new DBLocationTableQueryModel(), existingLocation, DBLocationTableQueryModel.SELECT_ID_FILTER);
                }

                DBLocationTableRowModel dbLocationTableRowModel = new DBLocationTableRowModel();
                dbLocationTableRowModel.setLocationCat(location.CatgoryName);
                dbLocationTableRowModel.setLocationName(location.LocationName);
                dbLocationTableRowModel.setLocationId(location.SPID);
                DatabaseManager.InsertRow(dbLocationTableRowModel);

                if (AppRepo.getInstance().getCurrentLocationId() == 0) {
                    AppRepo.getInstance().setCurrentLocationId(location.SPID);
                }
            }

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception exeception) {
            Log.d("Exception", exeception.getMessage());
        }

        return packetProcessResultModel;
    }
}
