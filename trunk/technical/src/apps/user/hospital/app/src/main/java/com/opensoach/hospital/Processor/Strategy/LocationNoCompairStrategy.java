package com.opensoach.hospital.Processor.Strategy;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.Communication.PacketLocationDataModel;
import com.opensoach.hospital.Model.Communication.PacketLocationsDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Processor.IProcessor;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.List;

/**
 * Created by Mandar on 9/4/2017.
 */

public class LocationNoCompairStrategy implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketLocationsDataModel> packetLocationDataModel = (PacketModel<PacketLocationsDataModel>)resultModel.Packet.Payload;
            List<PacketLocationDataModel> locationModels = packetLocationDataModel.Payload.Locations;


            for (PacketLocationDataModel locationModel:locationModels) {

                DBLocationTableRowModel dbLocationTableRowModel = new DBLocationTableRowModel();
                dbLocationTableRowModel.setLocationName(locationModel.Name);
                dbLocationTableRowModel.setLocationId(locationModel.LocationId);
                DatabaseManager.InsertRow(dbLocationTableRowModel);

                if (AppRepo.getInstance().getCurrentLocationId() == 0) {
                    AppRepo.getInstance().setCurrentLocationId(locationModel.LocationId);
                }
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in LocationNoCompairStrategy");
        }

        return packetProcessResultModel;
    }
}
