package spl.hkt.opensoach.splapp.processor;

import android.util.Log;

import com.google.gson.JsonElement;
import com.google.gson.JsonParser;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.ApplicationConstants;
import spl.hkt.opensoach.splapp.helper.CommonHelper;
import spl.hkt.opensoach.splapp.model.AppNotificationModelBase;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartConfigurationModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableRowModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;

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