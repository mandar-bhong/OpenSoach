package com.opensoach.hpft.Processor;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Helper.CommonHelper;
import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.Model.Communication.PacketCardListConfigurationModel;
import com.opensoach.hpft.Model.Communication.PacketChartConfigurationModel;
import com.opensoach.hpft.Model.Communication.PacketMedicalDetailsModel;
import com.opensoach.hpft.Model.Communication.PacketModel;
import com.opensoach.hpft.Model.Communication.PacketPatientDetailsModel;
import com.opensoach.hpft.Model.Communication.PacketServiceConfModel;
import com.opensoach.hpft.Model.DB.DBChartTableQueryModel;
import com.opensoach.hpft.Model.DB.DBChartTableRowModel;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Model.View.ChartConfigModel;
import com.opensoach.hpft.Utility.AppLogger;

import java.util.ArrayList;

/**
 * Created by Mandar on 03-08-2018.
 */

public class CardListProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            ArrayList<PacketCardListConfigurationModel> cardList = (ArrayList<PacketCardListConfigurationModel>) packetDecodeResultModel.Packet.Payload;

            if(cardList.size()==0) {
                DBChartTableRowModel deleteModel = new DBChartTableRowModel();
                deleteModel.setLocationId(packetDecodeResultModel.Packet.Header.LocationID);
                DatabaseManager.DeleteByFilter(new DBChartTableQueryModel(), deleteModel, DBChartTableQueryModel.SELECT_LOCATION_ID_FILTER);
            }


            Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();


            for(PacketCardListConfigurationModel model : cardList){
                TypeToken<PacketPatientDetailsModel> patientTypeToken = new TypeToken<PacketPatientDetailsModel>() {};

                model.PatientDetails = gson.fromJson(model.PatientDetailsJSON,patientTypeToken.getType());

                TypeToken<PacketMedicalDetailsModel> medicalTypeToken = new TypeToken<PacketMedicalDetailsModel>() {};
                model.MedicalDetails = gson.fromJson(model.MedicalDetailsJSON,medicalTypeToken.getType());

                TypeToken<PacketServiceConfModel> taskTypeToken = new TypeToken<PacketServiceConfModel>() {};
                model.ServiceConf = gson.fromJson(model.ServConfJSON,taskTypeToken.getType());

                model.LocationID = packetDecodeResultModel.Packet.Header.LocationID;
            }

            FillUpdateUIData(packetProcessResultModel, cardList);
            packetProcessResultModel.IsSuccess = true;
            return packetProcessResultModel;

        }catch (Exception ex){
            AppLogger.getInstance().Log(ex);
        }

        return packetProcessResultModel;
    }

    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, ArrayList<PacketCardListConfigurationModel> data) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = data;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_CARD_LIST_DATA;
    }
}
