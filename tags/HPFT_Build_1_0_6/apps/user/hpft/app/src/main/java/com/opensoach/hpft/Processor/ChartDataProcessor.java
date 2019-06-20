package com.opensoach.hpft.Processor;

import android.util.Log;

import java.util.ArrayList;
import java.util.Date;

import com.opensoach.hpft.DAL.DatabaseManager;
import com.opensoach.hpft.Constants.ApplicationConstants;
import com.opensoach.hpft.Helper.CommonHelper;
import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.Model.PacketDecodeResultModel;
import com.opensoach.hpft.Model.PacketProcessResultModel;
import com.opensoach.hpft.Model.Communication.PacketChartConfigurationModel;
import com.opensoach.hpft.Model.DB.DBChartTableQueryModel;
import com.opensoach.hpft.Model.DB.DBChartTableRowModel;
import com.opensoach.hpft.Model.View.ChartConfigModel;

/**
 * Created by Mandar on 2/26/2017. This class can work with db opetation and Memory operation(though layer)
 * also it will set flag to notify UI and send ack/command to server
 */

public class ChartDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            ArrayList<PacketChartConfigurationModel> charts = (ArrayList<PacketChartConfigurationModel>) packetDecodeResultModel.Packet.Payload;

            if(charts.size()==0)
            {
                // TODO: Here only one chart per servicepoint is considered, however in future there can be multiple charts
                DBChartTableRowModel deleteModel = new DBChartTableRowModel();
                deleteModel.setLocationId(packetDecodeResultModel.Packet.Header.LocationID);

                DatabaseManager.DeleteByFilter(new DBChartTableQueryModel(), deleteModel, DBChartTableQueryModel.SELECT_LOCATION_ID_FILTER);

                packetProcessResultModel.IsSuccess = true;

                ChartConfigModel chartDataModel = new ChartConfigModel();
                FillUpdateUIData(packetProcessResultModel, chartDataModel);
                return packetProcessResultModel;
            }

            // TODO: Here only one chart per servicepoint is considered, however in future there can be multiple charts
            PacketChartConfigurationModel packetChartConfigurationModel = charts.get(0);

            DBChartTableRowModel dbChartTableRowModel = new DBChartTableRowModel();
            dbChartTableRowModel.setLocationId(packetDecodeResultModel.Packet.Header.LocationID);
            dbChartTableRowModel.setServerChartId(packetChartConfigurationModel.ChartID);
            dbChartTableRowModel.setChartId(packetChartConfigurationModel.ChartID);
            dbChartTableRowModel.setChartName(packetChartConfigurationModel.ChartName);
            dbChartTableRowModel.setChartDispStartDate(new Date());
            //dbChartTableRowModel.setChartDispEndDate(chartDataModel.getSlotEndTime());
            dbChartTableRowModel.setChartPayload(packetChartConfigurationModel.ServConf);

            ChartConfigModel chartDataModel = CommonHelper.CreateChartModel(dbChartTableRowModel);

            // TODO: Here only one chart per servicepoint is considered, however in future there can be multiple charts
            DatabaseManager.DeleteByFilter(new DBChartTableQueryModel(), dbChartTableRowModel, DBChartTableQueryModel.SELECT_LOCATION_ID_FILTER);

            DatabaseManager.InsertRow(dbChartTableRowModel);
            FillUpdateUIData(packetProcessResultModel, chartDataModel);

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {
            //TODO: Log exception error
            Log.d("Exception", ex.getMessage());
        }

        return packetProcessResultModel;
    }

    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel, ChartConfigModel chartDataModel) {
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = chartDataModel;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA;
    }

}