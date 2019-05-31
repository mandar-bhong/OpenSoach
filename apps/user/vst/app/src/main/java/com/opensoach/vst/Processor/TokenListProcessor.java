package com.opensoach.vst.Processor;

import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.Communication.PacketCardListConfigurationModel;
import com.opensoach.vst.Model.Communication.PacketLocationDataModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketTokenListDataModel;
import com.opensoach.vst.Model.DB.DBLocationTableQueryModel;
import com.opensoach.vst.Model.DB.DBLocationTableRowModel;
import com.opensoach.vst.Model.DB.DBTokenTableQueryModel;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.List;

public class TokenListProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel resultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            RequestManager.Instance().CompleteRequest(resultModel.Packet.Header.SeqID);

            Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();

            Type packetType = new TypeToken<PacketModel<PacketSimpleAckModel<ArrayList<PacketTokenListDataModel>>>>() {
            }.getType();

            PacketModel<PacketSimpleAckModel<ArrayList<PacketTokenListDataModel>>> packet =  gson.fromJson(resultModel.JSONPacket,packetType);

            DBTokenTableRowModel dbRowModel = new DBTokenTableRowModel();

            switch (AppRepo.getInstance().getCurrentRunningMode()){
                case Token:
                case JobCreation:
                    break;
                case JobExecution:
                    dbRowModel.setState(2);
                    break;
            }


            DatabaseManager.DeleteByFilter(new DBTokenTableQueryModel(), dbRowModel, DBTokenTableQueryModel.SELECT_ALL);


            ArrayList<DBTokenTableRowModel> list = new ArrayList<>();

            for (PacketTokenListDataModel token : packet.Payload.Data) {

                switch (AppRepo.getInstance().getCurrentRunningMode()){
                    case Token:
                    case JobCreation:
                        break;
                    case JobExecution:
                        if (!(token.State == ApplicationConstants.TOKEN_JOB_CREATED ||
                        token.State == ApplicationConstants.TOKEN_JOB_EXECUTION_INPROGRESS)){
                            continue;
                        }
                        break;
                }

                DBTokenTableRowModel dbTokenTableRowModel = new DBTokenTableRowModel();
                dbTokenTableRowModel.setVehicleno(token.VehicleNo);
                dbTokenTableRowModel.setState(token.State);
                dbTokenTableRowModel.setTokenno(token.TokenID);
                dbTokenTableRowModel.setGeneratedon(token.GeneratedOn);
                dbTokenTableRowModel.setId(token.TokenID);

                DatabaseManager.InsertRow(dbTokenTableRowModel);

                list.add(dbTokenTableRowModel);
            }


            FillUpdateUIData(packetProcessResultModel,list);

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception exeception) {
            Log.d("Exception", exeception.getMessage());
        }
        return packetProcessResultModel;
    }


    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, ArrayList<DBTokenTableRowModel> data) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = data;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_TOKEN_LIST;
    }
}
