package com.opensoach.vst.Processor;

import android.util.Log;

import com.google.gson.JsonElement;
import com.google.gson.JsonParser;

import java.util.ArrayList;
import java.util.List;

import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.vst.Model.DB.DBAuthCodeTableRowModel;

/**
 * Created by Mandar on 4/23/2017.
 */

public class AuthCodeDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<ArrayList<String>> packetAuthCodeDataModel = (PacketModel)packetDecodeResultModel.Packet.Payload;
            int locationId = packetAuthCodeDataModel.Header.LocationID;
            //List<String> authCodes = packetAuthCodeDataModel.Payload.AuthCodes;

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(packetDecodeResultModel.JSONPacket);
            String authCodeJSON = root.getAsJsonObject().getAsJsonArray("payload").toString();

            DBAuthCodeTableRowModel dbAuthCodeTableRowModel =new DBAuthCodeTableRowModel();
            dbAuthCodeTableRowModel.setLocationId(locationId);

            boolean isAuthCodeChanged = false;

            List<DBAuthCodeTableRowModel> dbAuthCodeTableRowModels = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(),dbAuthCodeTableRowModel,DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

            if(authCodeJSON.equals("null")) {
                if (dbAuthCodeTableRowModels.size() > 0) {
                    DBAuthCodeTableRowModel dbAuthCodeTableRowdeleteModel = new DBAuthCodeTableRowModel();
                    dbAuthCodeTableRowdeleteModel.setLocationId(locationId);

                    int rowDeleted = DatabaseManager.DeleteByFilter(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowdeleteModel, DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

                    Log.d("","Deleted rows: " + rowDeleted);
                    isAuthCodeChanged = true;
                }
            }else {
                if (dbAuthCodeTableRowModels.size() > 0) {

                    DBAuthCodeTableRowModel dbAuthModel = dbAuthCodeTableRowModels.get(0);

                    if (!dbAuthModel.getAuthCodeJSON().equals(authCodeJSON)) {
                        DBAuthCodeTableRowModel dbAuthCodeTableRowUpdateModel = new DBAuthCodeTableRowModel();
                        dbAuthCodeTableRowUpdateModel.setLocationId(locationId);
                        dbAuthCodeTableRowUpdateModel.setAuthCode(authCodeJSON);
                        DatabaseManager.UpdateRow(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowUpdateModel, DBAuthCodeTableQueryModel.UDATE_AUTHCODE_BY_LOCATIONID_FILTER);

                        isAuthCodeChanged = true;
                    }
                } else {
                    DBAuthCodeTableRowModel dbAuthCodeTableRowInsertModel = new DBAuthCodeTableRowModel();
                    dbAuthCodeTableRowInsertModel.setLocationId(locationId);
                    dbAuthCodeTableRowInsertModel.setAuthCode(authCodeJSON);
                    DatabaseManager.InsertRow(dbAuthCodeTableRowInsertModel);

                    isAuthCodeChanged = true;
                }
            }

            if(isAuthCodeChanged){
                AppNotificationModelBase authCodeNotification = new AppNotificationModelBase();
                authCodeNotification.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_AUTH_CODE_UPDATE;
                authCodeNotification.Data = packetAuthCodeDataModel.Payload;
                packetProcessResultModel.CanUpdateUI = true;
                packetProcessResultModel.UINotifierModel = authCodeNotification;
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            Log.d("Exception", exeception.getMessage());
        }

        return packetProcessResultModel;
    }
}
