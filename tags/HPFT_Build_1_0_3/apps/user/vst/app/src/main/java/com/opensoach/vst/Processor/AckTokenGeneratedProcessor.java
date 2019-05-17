package com.opensoach.vst.Processor;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Helper.SyncState;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceInstanceTxnModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketTaskCompletedDataModel;
import com.opensoach.vst.Model.Communication.PacketTokenCreateDataModel;
import com.opensoach.vst.Model.Communication.PacketTokenCreateResponseDataModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableQueryModel;
import com.opensoach.vst.Model.DB.DBTaskDataTableRowModel;
import com.opensoach.vst.Model.DB.DBTokenTableQueryModel;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Model.View.ChartConfigModel;
import com.opensoach.vst.Utility.AppLogger;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;

import static com.opensoach.vst.Constants.ApplicationConstants.PACKET_DATE_FORMAT;

public class AckTokenGeneratedProcessor implements IProcessor {


    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        CommandRequest<PacketTokenCreateDataModel> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();
        try {
            RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(packetDecodeResultModel.JSONPacket);
            String ackJSON = root.getAsJsonObject().getAsJsonObject("payload").toString();

            Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();

            TypeToken<PacketSimpleAckModel<PacketTokenCreateResponseDataModel>> typeToken = new TypeToken<PacketSimpleAckModel<PacketTokenCreateResponseDataModel>> () {
            };

            PacketSimpleAckModel<PacketTokenCreateResponseDataModel> ack =  gson.fromJson(ackJSON, typeToken.getType());

            if (!ack.Ack) {
                packetProcessResultModel.IsSuccess = true;
                return packetProcessResultModel;
            }

            DBTokenTableRowModel dbTokenTableRowModel = new DBTokenTableRowModel();
            dbTokenTableRowModel.setVehicleno(request.Packet.Payload.VehicleNumber);
            dbTokenTableRowModel.setGeneratedon(ack.Data.GeneratedOn);
            dbTokenTableRowModel.setTokenno(ack.Data.Token);
            dbTokenTableRowModel.setId(ack.Data.TokenID);

            DatabaseManager.DeleteByFilter(new DBTokenTableQueryModel(), dbTokenTableRowModel, DBTokenTableQueryModel.SELECT_ID_FILTER);

            DatabaseManager.InsertRow(dbTokenTableRowModel);

            FillUpdateUIData(packetProcessResultModel,dbTokenTableRowModel);


            PacketHeaderModel headerModel =  PacketHelper.CreatePacketHeader(CommandConstants.CMD_CAT_DATA, CommandConstants.CMD_DATA_POST_GENERATE_TOKEN,0,request.Packet.Header.LocationID);

            PacketModel<PacketTokenCreateDataModel> packetModel = new PacketModel<>();
            packetModel.Header = headerModel;
            packetModel.Payload=request.Packet.Payload;

            packetProcessResultModel.CanSendServerCommand = true;
            packetProcessResultModel.ServerCommandPacket = CommonHelper.GetPacketJSON(packetModel);


            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }
        return packetProcessResultModel;
    }

    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, DBTokenTableRowModel dbTokenTableRowModel) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = dbTokenTableRowModel;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_TOKEN_CREATED;
    }
}
