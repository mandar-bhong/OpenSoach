package com.opensoach.hospital.Helper;

import android.content.Context;
import android.os.AsyncTask;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Communication.CommunicationManager;
import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Manager.ConnectionRetryManager;
import com.opensoach.hospital.Manager.PacketManager;
import com.opensoach.hospital.Manager.SendPacketManager;
import com.opensoach.hospital.Manager.ServerConnectionManager;
import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.Model.DB.DBAuthCodeTableQueryModel;
import com.opensoach.hospital.Model.DB.DBAuthCodeTableRowModel;
import com.opensoach.hospital.Model.DB.DBLocationTableQueryModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.MainViewModel;
import com.opensoach.hospital.Views.PropertyChangeHandler.LocationChangeHandler;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 8/26/2017.
 */

public class AppHelper {

    private static Context mContext;

    public static void Init(Context context) {

        mContext = context;

        Thread.setDefaultUncaughtExceptionHandler(new ExceptionHelper());

        DatabaseManager.getInstance().Init(context);

        ReadAppSettings();

        PacketManager.getInstance().Init();

        CommunicationManager.getInstance().Init(ServerConnectionManager.Instance());

        SendPacketManager.Instance().Init();

        //   ServerConnectionManager.Instance().Connect();

        //PacketProcessor processor = new PacketProcessor();
        //processor.Process();


        //AppRepo.getInstance().addPropertyChangeListener((ChartActivity)SPLApplication.getInstance().getmChartActivity());
        //ScheduleManager scheduleManager = new ScheduleManager();
        //scheduleManager.startScheduler(context,0,0,0);

        //Init Logger
        //com.opensoach.sme.utility.Log4jHelper.Init();


        RegisterHandler();
        PostConnectToServer(context);
    }

    public static void DeInit(){
        AppRepo.getInstance().removePropertyChangeListener(ConnectionRetryManager.Instance());
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

    public static void ExecuteStartUpProcess() {
        Runnable runnable = new Runnable() {
            @Override
            public void run() {

                AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"ExecuteStartUpProcess");
                List<DBLocationTableRowModel> locationModels = DatabaseManager.SelectAll(new DBLocationTableQueryModel(), new DBLocationTableRowModel(),DBLocationTableQueryModel.SORT_BY_NAME_DESC);

                if (locationModels.size() == 0) {
                    AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"No location Exists");
                    AppRepo.getInstance().setIsStartupCompleted();
                    return;
                }
                ;

                AppRepo.getInstance().setLocationList(locationModels);

                DBLocationTableRowModel dbLocationTableRowModel = locationModels.get(0);
                AppRepo.getInstance().setCurrentLocationId(dbLocationTableRowModel.getLocationId());

                AppRepo.getInstance().setIsStartupCompleted();

            }
        };

        AsyncTask.execute(runnable);
    }

    //Delaying server connect till application init e.g. internal view initialization
    private static void PostConnectToServer(Context ctr)
    {
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

    private static ArrayList<String> GetLocationAuthCode(int locationId){
        ArrayList<String> authCodes= new ArrayList<>() ;


        DBAuthCodeTableRowModel dbAuthCodeTableRowModel = new DBAuthCodeTableRowModel();

        List<DBAuthCodeTableRowModel> dbAuthCodeModes = DatabaseManager.SelectByFilter(new DBAuthCodeTableQueryModel(),dbAuthCodeTableRowModel,DBAuthCodeTableQueryModel.SELECT_ALL_FILTER);

        if(dbAuthCodeModes.size() > 0){

            DBAuthCodeTableRowModel dbAuthModel = dbAuthCodeModes.get(0);
            authCodes = DataConvertHelper.ConvertJSONStringArray(dbAuthModel.getAuthCodeJSON());
        }

        return authCodes;
    }


    private static void RegisterHandler(){

        AppRepo.getInstance().addPropertyChangeListener(ConnectionRetryManager.Instance());
        AppRepo.getInstance().addPropertyChangeListener(MainViewModel.getInstance());

        LocationChangeHandler locationChangeHandler =new LocationChangeHandler();
        AppRepo.getInstance().addPropertyChangeListener(locationChangeHandler);
    }


    public static Context GetAppContext(){
        return mContext;
    }


    public static void ProcessUIEvent(final AppNotificationModelBase modelBase) {
        switch (AppRepo.getInstance().getForegroundActivityName()) {
            case ApplicationConstants.FOREGROUND_ACTIVITY_MAIN:
            case ApplicationConstants.FOREGROUND_ACTIVITY_JOB_BOARD:
                if (AppRepo.getInstance().getForegroundActivityHandler() != null) {
                    Handler uiHandler = new Handler(Looper.getMainLooper()) {
                        @Override
                        public void handleMessage(Message message) {
                            AppRepo.getInstance().getForegroundActivityHandler().OnUIUpdateEvent(modelBase);
                        }
                    };

                    Message msg = uiHandler.obtainMessage();
                    uiHandler.sendMessage(msg);
                }
                break;
        }
    }
}
