package spl.hkt.opensoach.splapp.processor;

import android.util.Log;

import com.google.gson.JsonElement;
import com.google.gson.JsonParser;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import spl.hkt.opensoach.splapp.Constants;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.helper.AppHelper;
import spl.hkt.opensoach.splapp.helper.ApplicationConstants;
import spl.hkt.opensoach.splapp.helper.CommonHelper;
import spl.hkt.opensoach.splapp.model.AppNotificationModelBase;
import spl.hkt.opensoach.splapp.model.ChartDataModel;
import spl.hkt.opensoach.splapp.model.PacketDecodeResultModel;
import spl.hkt.opensoach.splapp.model.PacketProcessResultModel;
import spl.hkt.opensoach.splapp.model.communication.PacketChartConfigurationModel;
import spl.hkt.opensoach.splapp.model.communication.PacketModel;
import spl.hkt.opensoach.splapp.model.communication.PacketTaskModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableRowModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.util.CommonUtility;

/**
 * Created by Mandar on 2/26/2017. This class can work with db opetation and Memory operation(though layer)
 * also it will set flag to notify UI and send ack/command to server
 */

public class ChartDataProcessor implements IProcessor {

    @Override
    public PacketProcessResultModel Process(PacketDecodeResultModel packetDecodeResultModel) {

        PacketProcessResultModel packetProcessResultModel = new PacketProcessResultModel();

        try {

            JsonParser parser = new JsonParser();
            JsonElement root = parser.parse(packetDecodeResultModel.JSONPacket);
            String chartJSON = root.getAsJsonObject().get("payload").toString();

            ChartConfigModel chartDataModel = CommonHelper.CreateChartModel(packetDecodeResultModel.Packet.Header.LocationID, 0, chartJSON);

            DBChartTableRowModel dbChartTableRowModel = new DBChartTableRowModel();
            dbChartTableRowModel.setLocationId(packetDecodeResultModel.Packet.Header.LocationID);
            dbChartTableRowModel.setServerChartId(chartDataModel.getServerChartId());
            dbChartTableRowModel.setChartId(chartDataModel.getServerChartId());
            dbChartTableRowModel.setChartDispStartDate(new Date());
            //dbChartTableRowModel.setChartDispEndDate(chartDataModel.getSlotEndTime());
            dbChartTableRowModel.setLocationId(packetDecodeResultModel.Packet.Header.LocationID);
            dbChartTableRowModel.setChartPayload(chartJSON);

            int chartCount = DatabaseManager.GetRowCount(new DBChartTableQueryModel(), dbChartTableRowModel, DBChartTableQueryModel.SELECT_CHART_ID_FILTER);

            //TODO: For history management insert every chart as new chart and current chart as history chart
            if (chartCount > 0) {

                List<DBChartTableRowModel> dbCharts = DatabaseManager.SelectByFilter(new DBChartTableQueryModel(), dbChartTableRowModel, DBChartTableQueryModel.SELECT_CHART_ID_FILTER);

                String dbChartJSON = dbCharts.get(0).getChartPayload();

                if(!dbChartJSON.equals(chartJSON)) {
                    dbChartTableRowModel.setChartPayload(chartJSON);
                    DatabaseManager.UpdateRow(new DBChartTableQueryModel(), dbChartTableRowModel, DBChartTableQueryModel.SELECT_SERVER_CHART_ID_FILTER);

                    FillUpdateUIData(packetProcessResultModel,chartDataModel);
                }else{
                    Log.d("Info","No changes found in chart");
                }

            } else {
                DatabaseManager.InsertRow(dbChartTableRowModel);
                FillUpdateUIData(packetProcessResultModel,chartDataModel);
            }

            packetProcessResultModel.IsSuccess = true;

        } catch (Exception ex) {
            //TODO: Log exception error
            Log.d("Exception", ex.getMessage());
        }

        return packetProcessResultModel;
    }


    void FillUpdateUIData(PacketProcessResultModel packetProcessResultModel,ChartConfigModel chartDataModel){
        packetProcessResultModel.CanUpdateUI = true;
        packetProcessResultModel.UINotifierModel = new AppNotificationModelBase();
        packetProcessResultModel.UINotifierModel.Data = chartDataModel;
        packetProcessResultModel.UINotifierModel.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA;
    }

}