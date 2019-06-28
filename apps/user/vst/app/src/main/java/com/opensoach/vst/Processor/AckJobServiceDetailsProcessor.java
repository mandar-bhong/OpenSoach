package com.opensoach.vst.Processor;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerNVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceInstanceModel;
import com.opensoach.vst.Model.Communication.PacketServiceJobDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceOwnerVehicleDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceTaskItemDataModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.PacketDecodeResultModel;
import com.opensoach.vst.Model.PacketProcessResultModel;
import com.opensoach.vst.Utility.AppLogger;

import java.lang.reflect.Type;
import java.util.ArrayList;

public class AckJobServiceDetailsProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {
        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();

            Type packetType = new TypeToken<PacketModel<PacketSimpleAckModel<ArrayList<PacketServiceJobDetailsDataModel>>>>() {
            }.getType();

            PacketModel<PacketSimpleAckModel<ArrayList<PacketServiceJobDetailsDataModel>>> packet = gson.fromJson(packetDecodeResultModel.JSONPacket, packetType);

            PacketSimpleAckModel<ArrayList<PacketServiceJobDetailsDataModel>> ack = packet.Payload;

            if (!ack.Ack) {
                packetProcessResultModel.IsSuccess = true;
                return packetProcessResultModel;
            }

            if (ack.Data.size() < 1){
                packetProcessResultModel.IsSuccess = true;
                return packetProcessResultModel;
            }

            FillUpdateUIData(packetProcessResultModel,ack.Data.get(0));

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }
        return packetProcessResultModel;
    }


    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, PacketServiceJobDetailsDataModel data) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = data;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_JOB_SERVICE_DETAILS;
    }
}
