package com.opensoach.hpft.Processor;

import android.util.Log;

import java.util.List;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Model.Communication.PacketLocationDataModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.DB.DBLocationTableQueryModel;
import com.opensoach.hpft.Model.DB.DBLocationTableRowModel;

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
