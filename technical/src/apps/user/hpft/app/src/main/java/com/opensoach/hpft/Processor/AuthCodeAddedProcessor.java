package com.opensoach.hpft.Processor;

import com.google.gson.Gson;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.List;

import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.Helper.DataConvertHelper;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.hpft.Model.DB.DBAuthCodeTableRowModel;

/**
 * Created by Mandar on 03-07-2018.
 */

public class AuthCodeAddedProcessor implements IProcessor {
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


            Type type = new TypeToken<List<String>>(){}.getType();
            List<String> receivedAuthCodes = new Gson().fromJson(authCodeJSON, type);

            boolean isAuthCodeChanged = false;
            String newAuthCodes ="";

            DBAuthCodeTableRowModel dbAuthCodeTableRowModel =new DBAuthCodeTableRowModel();
            dbAuthCodeTableRowModel.setLocationId(locationId);

            List<DBAuthCodeTableRowModel> dbAuthCodeTableRowModels = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(),dbAuthCodeTableRowModel,DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

            if(dbAuthCodeTableRowModels.size() > 0) {
                String dbAuthCode = dbAuthCodeTableRowModels.get(0).getAuthCodeJSON();

                List<String> dbAuthCodes = DataConvertHelper.ConvertJSONStringArray(dbAuthCode);

                dbAuthCodes.addAll(receivedAuthCodes);

                String newAuthCodeDbItem =  new Gson().toJson(dbAuthCodes);
                newAuthCodes = newAuthCodeDbItem;

                DBAuthCodeTableRowModel dbAuthCodeTableRowUpdateModel = new DBAuthCodeTableRowModel();
                dbAuthCodeTableRowUpdateModel.setLocationId(locationId);
                dbAuthCodeTableRowUpdateModel.setAuthCode(newAuthCodeDbItem);
                DatabaseManager.UpdateRow(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowUpdateModel, DBAuthCodeTableQueryModel.UDATE_AUTHCODE_BY_LOCATIONID_FILTER);

                isAuthCodeChanged = true;

            }else{

                DBAuthCodeTableRowModel dbAuthCodeTableRowInsertModel = new DBAuthCodeTableRowModel();
                dbAuthCodeTableRowInsertModel.setLocationId(locationId);
                dbAuthCodeTableRowInsertModel.setAuthCode(authCodeJSON);
                DatabaseManager.InsertRow(dbAuthCodeTableRowInsertModel);

                newAuthCodes = authCodeJSON;

                isAuthCodeChanged = true;
            }

            if(isAuthCodeChanged){
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
