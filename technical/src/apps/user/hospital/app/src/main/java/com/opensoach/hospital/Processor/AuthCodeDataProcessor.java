package com.opensoach.hospital.Processor;

import android.util.Log;

import com.google.gson.JsonElement;
import com.google.gson.JsonParser;
import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.Communication.PacketAuthCodeDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.hospital.Model.DB.DBAuthCodeTableRowModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Model.PacketProcessResultModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 8/27/2017.
 */

public class AuthCodeDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            PacketModel<PacketAuthCodeDataModel> packetAuthCodeDataModel = (PacketModel<PacketAuthCodeDataModel>) packetDecodeResultModel.Packet.Payload;
            ArrayList<String> authCodes = packetAuthCodeDataModel.Payload.AuthCodes;

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(packetDecodeResultModel.JSONPacket);
            String authCodeJSON = root.getAsJsonObject().getAsJsonObject("payload").get("opcodes").toString();

            DBAuthCodeTableRowModel dbAuthCodeTableRowModel = new DBAuthCodeTableRowModel();

            boolean isAuthCodeChanged = false;

            List<DBAuthCodeTableRowModel> dbAuthCodeTableRowModels = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowModel, DBAuthCodeTableQueryModel.SELECT_ALL_FILTER);

            if (authCodeJSON.equals("null")) {
                if (dbAuthCodeTableRowModels.size() > 0) {
                    DBAuthCodeTableRowModel dbAuthCodeTableRowdeleteModel = new DBAuthCodeTableRowModel();

                    int rowDeleted = DatabaseManager.DeleteByFilter(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowdeleteModel, DBAuthCodeTableQueryModel.SELECT_ALL_FILTER);

                    Log.d("", "Deleted rows: " + rowDeleted);
                    isAuthCodeChanged = true;
                }
            } else {
                if (dbAuthCodeTableRowModels.size() > 0) {

                    DBAuthCodeTableRowModel dbAuthModel = dbAuthCodeTableRowModels.get(0);

                    if (!dbAuthModel.getAuthCodeJSON().equals(authCodeJSON)) {
                        DBAuthCodeTableRowModel dbAuthCodeTableRowUpdateModel = new DBAuthCodeTableRowModel();
                        dbAuthCodeTableRowUpdateModel.setAuthCode(authCodeJSON);
                        DatabaseManager.UpdateRow(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowUpdateModel, DBAuthCodeTableQueryModel.UDATE_ALL_FILTER);

                        isAuthCodeChanged = true;
                    }
                } else {
                    DBAuthCodeTableRowModel dbAuthCodeTableRowInsertModel = new DBAuthCodeTableRowModel();
                    dbAuthCodeTableRowInsertModel.setAuthCode(authCodeJSON);
                    DatabaseManager.InsertRow(dbAuthCodeTableRowInsertModel);

                    isAuthCodeChanged = true;
                }

                AppRepo.getInstance().setAuthCodeList(authCodes);
            }

//            if(isAuthCodeChanged){
//                AppNotificationModelBase authCodeNotification = new AppNotificationModelBase();
//                authCodeNotification.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_AUTH_CODE_UPDATE;
//                authCodeNotification.Data = packetAuthCodeDataModel.Payload;
//                packetProcessResultModel.CanUpdateUI = true;
//                packetProcessResultModel.UINotifierModel = authCodeNotification;
//            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception){
            packetProcessResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(exeception,"Error occured in AuthCodeDataProcessor");
        }

        return packetProcessResultModel;
    }
}
