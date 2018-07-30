package com.opensoach.hpft;

import android.app.Activity;
import android.app.Application;
import android.content.Context;

import java.util.ArrayList;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.Communication.WebSocketConnector;
import com.opensoach.hpft.Helper.AppHelper;
import com.opensoach.hpft.Constants.Constants.ApplicationConstants;
import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.Model.View.ChartConfigModel;
import com.opensoach.hpft.Model.View.DisplayChartDataModel;
import com.opensoach.hpft.Scheduler.ScheduleManager;
import com.opensoach.hpft.Views.ChartActivity;
import com.opensoach.hpft.Views.TimeChangeListner;
import com.opensoach.hpft.Views.UpdateChartListner;

/**
 * Created by Samir Bukkawar  on 15-Feb-17.
 */

public class SPLApplication extends Application {
    private Context mContext;
    private WebSocketConnector mWSConnector;
    private static SPLApplication mAppInstance;
    private UpdateChartListner mUpdateChartListner;
    private TimeChangeListner mTimeChangeListner;
    private Activity mActivity;
    private Activity mChartActivity;
    private final String TAG = "SPLApplication";

    @Override
    public void onCreate() {
        super.onCreate();
        mContext = this;
        mAppInstance = this;

        //Initialise WebSocket Connection
        // initWSConnection();

        AppHelper.Init(mContext);
        AppHelper.ExecuteStartUpProcess();
        (new ScheduleManager()).startScheduler(this, 1, 30, 1);
    }

    public void registerListner(UpdateChartListner listner, Activity activity) {
        mUpdateChartListner = listner;
        mActivity = activity;
    }

    public void registerTimeChangeListner(TimeChangeListner listner, Activity activity) {
        mTimeChangeListner = listner;
        mChartActivity = activity;
    }

    public static SPLApplication getInstance() {
        return mAppInstance;
    }

    public void updateActivity(final String str) {

        mActivity.runOnUiThread(new Runnable() {
            @Override
            public void run() {
                mUpdateChartListner.callback(str);
            }
        });
    }

    public void updateTimeChange() {
        if (mChartActivity != null && mTimeChangeListner != null) {
            mChartActivity.runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    mTimeChangeListner.notifyTimeChange();
                }
            });
        }
    }

    public void OnUIUpdateEvent(final AppNotificationModelBase model) {
        switch (model.DataProcessStatergyID) {
            case ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA:
                //TODO: convert model.Data to appropriate format
                if (mChartActivity != null) {
                    mChartActivity.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            UpdateChart(model);
                            AppHelper.LoadChartDataAync(AppRepo.getInstance().getCurrentChartId());
                        }
                    });
                }
                break;

            case ApplicationConstants.UI_PROCESSING_STATERGY_CHART_DATA_START_UP_DISPLAY:
                if (mChartActivity != null) {
                    mChartActivity.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            UpdateChartData(model);
                        }
                    });
                }
                break;
            case ApplicationConstants.UI_PROCESSING_STATERGY_AUTH_CODE_UPDATE: {
                ArrayList<String> packetAuthCodeDataModel = (ArrayList) model.Data;

                if(packetAuthCodeDataModel == null) {
                    AppRepo.getInstance().setAuthCodeList(new ArrayList<String>());
                }else{
                    AppRepo.getInstance().setAuthCodeList(packetAuthCodeDataModel);
                }
            }
            break;
        }
    }



    void UpdateChart(AppNotificationModelBase model) {

        ((ChartActivity) mChartActivity).setChartModel((ChartConfigModel) model.Data);
    }

    void UpdateChartData(AppNotificationModelBase model) {
        if(model == null)return;

        ((ChartActivity) mChartActivity).setChartDataModel((DisplayChartDataModel) model.Data);
    }


    public Activity getmChartActivity() {
        return mChartActivity;
    }

    public void setmChartActivity(Activity mChartActivity) {
        this.mChartActivity = mChartActivity;
    }
}