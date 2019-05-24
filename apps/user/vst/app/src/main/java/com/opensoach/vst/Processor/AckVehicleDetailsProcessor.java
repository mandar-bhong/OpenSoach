package com.opensoach.vst.Processor;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.SyncState;
import com.opensoach.vst.Manager.RequestManager;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.Communication.CommandRequest;
import com.opensoach.vst.Model.Communication.PacketLocationDataModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerNVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceOwnerVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketTokenListDataModel;
import com.opensoach.vst.Model.Communication.PacketVehicleDetailsRequestDataModel;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Utility.AppLogger;

import java.lang.reflect.Type;
import java.util.ArrayList;

public class AckVehicleDetailsProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        CommandRequest<PacketVehicleDetailsRequestDataModel> request = (CommandRequest) RequestManager.Instance().GetRequest(packetDecodeResultModel.Packet.Header.SeqID);

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {
            RequestManager.Instance().CompleteRequest(packetDecodeResultModel.Packet.Header.SeqID);


            Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();

            Type packetType = new TypeToken<PacketModel<PacketSimpleAckModel<PacketServiceOwnerVehicleDetailsDataModel>>>() {
            }.getType();

            PacketModel<PacketSimpleAckModel<PacketServiceOwnerVehicleDetailsDataModel>> packet = gson.fromJson(packetDecodeResultModel.JSONPacket, packetType);

            PacketSimpleAckModel<PacketServiceOwnerVehicleDetailsDataModel> ack = packet.Payload;

            if (!ack.Ack) {
                packetProcessResultModel.IsSuccess = true;
                return packetProcessResultModel;
            }

            packetType = new TypeToken<PacketServiceCustomerNVehicleDetailsDataModel>(){}.getType();
            PacketServiceCustomerNVehicleDetailsDataModel customerDetails = gson.fromJson(ack.Data.CustomerDetails, packetType);

            if (customerDetails.CustomerDetails != null){
                AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_VEHICLE_DETAILS_PREFIX + ack.Data.VehicleNo, customerDetails);

                FillUpdateUIData(packetProcessResultModel,customerDetails.CustomerDetails);
            }

            packetProcessResultModel.IsSuccess = true;

        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }
        return packetProcessResultModel;
    }

    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, PacketServiceCustomerDetailsDataModel data) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = data;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_VEHICLE_DETAILS;
    }
}
