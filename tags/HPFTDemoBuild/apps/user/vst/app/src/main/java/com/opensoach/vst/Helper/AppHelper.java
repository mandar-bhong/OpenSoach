package com.opensoach.vst.Helper;

import android.content.Context;
import android.os.AsyncTask;
import android.os.Handler;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.SPLApplication;
import com.opensoach.vst.SharedPreference.SharedPreferencesHelper;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Utility.AppLogger;
import com.opensoach.vst.Communication.CommunicationManager;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Manager.BroadCastReceiverManager;
import com.opensoach.vst.Manager.ChartDataRunnable;
import com.opensoach.vst.Manager.ConnectionRetryManager;
import com.opensoach.vst.Manager.DeviceSyncManager;
import com.opensoach.vst.Manager.HttpManager;
import com.opensoach.vst.Manager.LocationChangeManager;
import com.opensoach.vst.Manager.PacketManager;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Manager.ServerConnectionManager;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.vst.Model.DB.DBAuthCodeTableRowModel;
import com.opensoach.vst.Model.DB.DBChartDataTableQueryModel;
import com.opensoach.vst.Model.DB.DBChartDataTableRowModel;
import com.opensoach.vst.Model.DB.DBChartTableQueryModel;
import com.opensoach.vst.Model.DB.DBChartTableRowModel;
import com.opensoach.vst.Model.DB.DBLocationTableQueryModel;
import com.opensoach.vst.Model.DB.DBLocationTableRowModel;
import com.opensoach.vst.Model.View.ChartConfigModel;
import com.opensoach.vst.Model.View.DisplayChartDataModel;
import com.opensoach.vst.Model.View.DisplayChartItemDataModel;
import com.opensoach.vst.Processor.PacketProcessor;
import com.opensoach.vst.ViewModels.MainViewModel;

/**
 * Created by Mandar on 2/25/2017.
 */

public class AppHelper {
    private static Context mContext;

    public static void Init(Context context) {

        mContext = context;
        // Init all modules

        // init Database
        DatabaseManager dbManager = new DatabaseManager();
        dbManager.Init(context);

        ReadAppSettings();

        PacketManager.getInstance().Init();
        CommunicationManager.getInstance().Init(ServerConnectionManager.Instance());

        SendPacketManager.Instance().Init();

        //   ServerConnectionManager.Instance().Connect();

        PacketProcessor processor = new PacketProcessor();
        //processor.Process();


        //AppRepo.getInstance().addPropertyChangeListener((ChartActivity)SPLApplication.getInstance().getmChartActivity());
        //ScheduleManager scheduleManager = new ScheduleManager();
        //scheduleManager.startScheduler(context,0,0,0);

        //Init Logger
        //Log4jHelper.Init();

        //ConnectionRetryManager.Instance().Init();
        AppRepo.getInstance().addPropertyChangeListener(ConnectionRetryManager.Instance());
        AppRepo.getInstance().addPropertyChangeListener(LocationChangeManager.Instance());
        AppRepo.getInstance().addPropertyChangeListener(DeviceSyncManager.Instance());
        AppRepo.getInstance().addPropertyChangeListener(MainViewModel.getInstance().getHeaderViewModel());

        BroadCastReceiverManager.Instance().RegisterBatteryLevelReceiver(mContext);

        ExecuteStartUpProcess();

        HttpManager.ProcessWebSocketURL();
    }

    public static void ReadAppSettings() {
        SharedPreferencesHelper preferencesHelper = SharedPreferencesHelper.getInstance(mContext);
        String locationName = preferencesHelper.getDataFromSharedPreference(Constants.KEY_LOCATION_NAME);
        String locationId = preferencesHelper.getDataFromSharedPreference(Constants.KEY_LOCATION_ID);

        if (locationName == null) {
            //TODO Retrieve the location and set to Shared Preferences
            preferencesHelper.updateSharedPreference(Constants.KEY_LOCATION_NAME, "");
            preferencesHelper.updateSharedPreference(Constants.KEY_LOCATION_ID, "");
        }
    }

    public static void DeInit() {
        AppRepo.getInstance().removePropertyChangeListener(ConnectionRetryManager.Instance());
        AppRepo.getInstance().removePropertyChangeListener(LocationChangeManager.Instance());
        AppRepo.getInstance().removePropertyChangeListener(DeviceSyncManager.Instance());

        CommunicationManager.getInstance().DeInit();
        SendPacketManager.Instance().DeInit();
        ConnectionRetryManager.Instance().DeInit();


        if (mContext == null)return;
        BroadCastReceiverManager.Instance().DeregisterBatteryLevelReceiver(mContext);
    }

    public static void ExecuteStartUpProcess() {

        Runnable runnable = new Runnable() {
            @Override
            public void run() {

                List<DBLocationTableRowModel> locationModels = DatabaseManager.SelectAll(new DBLocationTableQueryModel(), new DBLocationTableRowModel());

                if (locationModels.size() == 0) {
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "No location Exists");
                    return;
                }
                ;

                DBLocationTableRowModel dbLocationTableRowModel = locationModels.get(0);

                AppRepo.getInstance().setCurrentLocationId(dbLocationTableRowModel.getLocationId());

                ArrayList<String> authCodes = GetLocationAuthCode(dbLocationTableRowModel.getLocationId());

                AppRepo.getInstance().setAuthCodeList(authCodes);
            }
        };

        AsyncTask.execute(runnable);
    }

    public static void LoadChartDataAync(int chartId) {
        Runnable runnable = new ChartDataRunnable(chartId);
        AsyncTask.execute(runnable);
    }

    public static void LoadLocationChart(Integer locationId) {

        try {
            DBChartTableRowModel dbChartTableRowModel = new DBChartTableRowModel();
            dbChartTableRowModel.setLocationId(locationId);
            List<DBChartTableRowModel> charts = DatabaseManager.SelectByFilter(new DBChartTableQueryModel(), dbChartTableRowModel, DBChartTableQueryModel.SELECT_LOCATION_ID_FILTER);

            if (charts.size() <= 0) {
                //Log: Location not found for "locationId"
                AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "Chart not found for locationid ");
                return;
            }

            DBChartTableRowModel chartModel = charts.get(0);

            AppRepo.getInstance().setCurrentChartId(chartModel.getChartId());

            ChartConfigModel chartConfigModel = CommonHelper.CreateChartModel(chartModel);

            AppNotificationModelBase notificationModelBase = new AppNotificationModelBase();
            notificationModelBase.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA;
            notificationModelBase.Data = chartConfigModel;

            SPLApplication.getInstance().OnUIUpdateEvent(notificationModelBase);

            //Start Chart Data processing
            // UpdateChartData(chartModel.getChartId());

        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }
    }

    public static void UpdateChartData(Integer chartId) {

        try {
            //Start Chart Data processing
            SimpleDateFormat chartStartDate = new SimpleDateFormat("yyyy-MM-dd");
            Date dateWithoutTime = chartStartDate.parse(chartStartDate.format(new Date()));

            DBChartDataTableRowModel dbChartDataTableRowModel = new DBChartDataTableRowModel();
            dbChartDataTableRowModel.setChartId(chartId);
            dbChartDataTableRowModel.setChartDay(dateWithoutTime);

            List<DBChartDataTableRowModel> chartDataItems = DatabaseManager.SelectByFilter(new DBChartDataTableQueryModel(), dbChartDataTableRowModel, DBChartDataTableQueryModel.FILTER_BY_CHARTID_TODAY);

            if (chartDataItems.size() <= 0) {
                return;
            }

            NotifyChartDataStatusUpdate(chartDataItems);
        } catch (Exception ex) {
            AppLogger.getInstance().Log(ex);
        }
    }

    public static void NotifyChartDataStatusUpdate(List<DBChartDataTableRowModel> chartDataItems) {

        DisplayChartDataModel displayChartDataModel = new DisplayChartDataModel();

        for (DBChartDataTableRowModel dbModel : chartDataItems) {
            DisplayChartItemDataModel displayChartItemDataModel = DataConvertHelper.ConvertDBChartDataToChartDisplayModel(dbModel);
            displayChartDataModel.getChartData().add(displayChartItemDataModel);
        }

        AppNotificationModelBase notificationChartDataModelBase = new AppNotificationModelBase();
        notificationChartDataModelBase.DataProcessStatergyID = ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA_START_UP_DISPLAY;
        notificationChartDataModelBase.Data = displayChartDataModel;

        SPLApplication.getInstance().OnUIUpdateEvent(notificationChartDataModelBase);

    }


    private static ArrayList<String> GetLocationAuthCode(int locationId) {
        ArrayList<String> authCodes = new ArrayList<>();


        DBAuthCodeTableRowModel dbAuthCodeTableRowModel = new DBAuthCodeTableRowModel();
        dbAuthCodeTableRowModel.setLocationId(locationId);

        List<DBAuthCodeTableRowModel> dbAuthCodeModes = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(), dbAuthCodeTableRowModel, DBAuthCodeTableQueryModel.SELECT_BY_LOCATION_FILTER);

        if (dbAuthCodeModes.size() > 0) {

            DBAuthCodeTableRowModel dbAuthModel = dbAuthCodeModes.get(0);
            authCodes = DataConvertHelper.ConvertJSONStringArray(dbAuthModel.getAuthCodeJSON());
        }

        return authCodes;
    }

    //Delaying server connect till application init e.g. internal view initialization
    private static void PostConnectToServer(Context ctr) {
        Handler hdl = new Handler(ctr.getMainLooper());
        hdl.post(new Runnable() {

            public Runnable init() {
                return this;
            }

            @Override
            public void run() {
                ConnectionRetryManager.Instance().Init();
            }
        }.init());
    }

    public static void OnWebSocketURLReceived() {
        PostConnectToServer(mContext);
    }

}
