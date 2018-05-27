package spl.hkt.opensoach.splapp.helper;

import android.content.Context;
import android.os.AsyncTask;
import android.os.Handler;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import spl.hkt.opensoach.splapp.Constants;
import spl.hkt.opensoach.splapp.SPLApplication;
import spl.hkt.opensoach.splapp.SharedPreference.SharedPreferencesHelper;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.dal.DatabaseManager;
import spl.hkt.opensoach.splapp.manager.ChartDataRunnable;
import spl.hkt.opensoach.splapp.manager.ConnectionRetryManager;
import spl.hkt.opensoach.splapp.manager.HttpManager;
import spl.hkt.opensoach.splapp.manager.LocationChangeManager;
import spl.hkt.opensoach.splapp.manager.PacketManager;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.manager.ServerConnectionManager;
import spl.hkt.opensoach.splapp.model.AppNotificationModelBase;
import spl.hkt.opensoach.splapp.model.db.DBAuthCodeTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBAuthCodeTableRowModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartDataTableRowModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBChartTableRowModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableQueryModel;
import spl.hkt.opensoach.splapp.model.db.DBLocationTableRowModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartDataModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartItemDataModel;
import spl.hkt.opensoach.splapp.processor.PacketProcessor;

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
        AppRepo.getInstance().addPropertyChangeListener(new LocationChangeManager());

        HttpManager.ProcessWebSocketURL(AppRepo.getInstance().getServerAPIURL(), AppRepo.getInstance().getDeviceSerial());
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
        // Init all modules
        // PacketManager.getInstance().stop();
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
