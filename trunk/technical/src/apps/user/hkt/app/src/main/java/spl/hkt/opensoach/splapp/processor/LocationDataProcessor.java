package spl.hkt.opensoach.splapp.processor;

import android.provider.ContactsContract;
import android.util.Log;

import java.util.List;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.LocationDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketLocationDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableRowModel;

/**
 * Created by Mandar on 4/8/2017.
 */

public class LocationDataProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketLocationDataModel> packetLocationDataModel = (PacketModel<PacketLocationDataModel>) resultModel.Packet.Payload;
            List<LocationDataModel> locations = packetLocationDataModel.Payload.Locations;

            List<DBLocationTableRowModel> dbLocationModels = DatabaseManager.SelectAll(new DBLocationTableQueryModel(), new DBLocationTableRowModel());


            for (LocationDataModel location : locations) {
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
