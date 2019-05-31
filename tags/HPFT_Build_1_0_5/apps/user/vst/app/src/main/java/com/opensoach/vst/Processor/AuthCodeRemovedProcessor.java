package com.opensoach.vst.Processor;

import com.google.gson.Gson;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.List;

import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Helper.DataConvertHelper;
import com.opensoach.vst.Utility.AppLogger;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.vst.Model.DB.DBAuthCodeTableRowModel;

/**
 * Created by Mandar on 03-07-2018.
 */

public class AuthCodeRemovedProcessor implements IProcessor {
    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<ArrayList<String>> packetAuthCodeDataModel = (PacketModel)packetDecodeResultModel.Packet.Payload;

            int locationId = packetAuthCodeDataModel.Header.LocationID;

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(packetDecodeResultModel.JSONPacket);
            String authCodeJSON = root.getAsJsonObject().getAsJsonArray("payload").toString();

            Type type = new TypeToken<List<String>>(){}.getType();
            List<String> receivedAuthCodes = new Gson().fromJson(authCodeJSON, type);


            String newAuthCodes ="";


            DBAuthCodeTableRowModel dbAuthCodeTableRowModel =new DBAuthCodeTableRowModel();
            dbAuthCodeTableRowModel.setLocationId(locationId);

            List<DBAuthCodeTableRowModel> dbAuthCodeTableRowModels = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(),dbAuthCodeTableRowModel,DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

            if(dbAuthCodeTableRowModels.size() > 0) {
                String dbAuthCode = dbAuthCodeTableRowModels.get(0).getAuthCodeJSON();

                List<String> dbAuthCodes = DataConvertHelper.ConvertJSONStringArray(dbAuthCode);


                ArrayList<String> authCodeRemoveObj = new ArrayList<>();

                for (int i = 0; i < receivedAuthCodes.size() ; i++) {
                    for (int j = 0; j < dbAuthCodes.size(); j++) {
                        if (receivedAuthCodes.get(i).equals(dbAuthCodes.get(j))) {
                            authCodeRemoveObj.add(dbAuthCodes.get(j));
                        }
                    }
                }

                for (String removeAuthCode : authCodeRemoveObj) {
                    dbAuthCodes.remove(removeAuthCode);
                }

                if (dbAuthCodes.size() == 0 ){
                    DBAuthCodeTableRowModel dbAuthCodeTableRowdeleteModel = new DBAuthCodeTableRowModel();
                    dbAuthCodeTableRowdeleteModel.setLocationId(locationId);

                    DatabaseManager.DeleteByFilter(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowdeleteModel, DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

                    newAuthCodes = "[]";

                }else{
                    String newAuthCodeDbItem = new Gson().toJson(dbAuthCodes);

                    newAuthCodes = newAuthCodeDbItem;

                    DBAuthCodeTableRowModel dbAuthCodeTableRowUpdateModel = new DBAuthCodeTableRowModel();
                    dbAuthCodeTableRowUpdateModel.setLocationId(locationId);
                    dbAuthCodeTableRowUpdateModel.setAuthCode(newAuthCodeDbItem);
                    DatabaseManager.UpdateRow(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowUpdateModel, DBAuthCodeTableQueryModel.UDATE_AUTHCODE_BY_LOCATIONID_FILTER);
                }

                AppNotificationModelBase authCodeNotification = new AppNotificationModelBase();
                authCodeNotification.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_AUTH_CODE_UPDATE;
                authCodeNotification.Data = newAuthCodes;
                packetProcessResultModel.CanUpdateUI = true;
                packetProcessResultModel.UINotifierModel = authCodeNotification;
            }

            packetProcessResultModel.IsSuccess = true;

        }
        catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }

        return packetProcessResultModel;
    }
}
