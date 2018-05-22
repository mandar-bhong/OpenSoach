package spl.hkt.opensoach.splapp.processor;

import android.provider.ContactsContract;
import android.util.Log;

import java.util.List;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.PacketLocationDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableRowModel;

/**
 * Created by Mandar on 4/8/2017.
 */

public class LocationDataProcessor implements IProcessor  {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketLocationDataModel> packetLocationDataModel = (PacketModel<PacketLocationDataModel>)resultModel.Packet.Payload;
            List<Integer> locationids = packetLocationDataModel.Payload.LocationIds;

            List<DBLocationTableRowModel> dbLocationModels = DatabaseManager.SelectAll(new DBLocationTableQueryModel(),new DBLocationTableRowModel());


            for (Integer locationid:locationids) {
                boolean isLocationExists=false;
                for (DBLocationTableRowModel dbLocationTableRowModel:dbLocationModels) {
                    if(dbLocationTableRowModel.getLocationId()  == locationid) {
                        isLocationExists = true;
                    }
                }

                if(!isLocationExists){
                    DBLocationTableRowModel dbLocationTableRowModel = new DBLocationTableRowModel();
                    dbLocationTableRowModel.setLocationCat(0);
                    dbLocationTableRowModel.setLocationName("");
                    dbLocationTableRowModel.setLocationId(locationid);
                    DatabaseManager.InsertRow(dbLocationTableRowModel);
                }

                if(AppRepo.getInstance().getCurrentLocationId() == 0){
                    AppRepo.getInstance().setCurrentLocationId(locationid);
                }
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            Log.d("Exception", exeception.getMessage());
        }

        return packetProcessResultModel;
    }
}
