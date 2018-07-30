package com.opensoach.hpft.View;

import android.app.Activity;
import android.app.ActivityManager;
import android.app.Fragment;
import android.app.FragmentTransaction;
import android.content.Context;
import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.util.Log;
import android.view.View;
import android.view.WindowManager;
import android.widget.ArrayAdapter;
import android.widget.FrameLayout;
import android.widget.ImageView;
import android.widget.Spinner;
import android.widget.Toast;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;
import java.util.List;

import com.opensoach.hpft.Constants.Constants;
import com.opensoach.hpft.R;
import com.opensoach.hpft.SPLApplication;
import com.opensoach.hpft.AppRepo.AppRepo;

import com.opensoach.hpft.Handler.ChartActivityClickHandler;
import com.opensoach.hpft.Helper.AppHelper;
import com.opensoach.hpft.Utility.AppLogger;
import com.opensoach.hpft.Manager.BroadCastReceiverManager;
import com.opensoach.hpft.Manager.LocationChartRunnable;
import com.opensoach.hpft.Model.View.ChartConfigModel;
import com.opensoach.hpft.Model.View.DisplayChartDataModel;
import com.opensoach.hpft.ViewModels.MainViewModel;

public class ChartActivity extends Activity implements ChartTableFragment.OnFragmentInteractionListener, PropertyChangeListener {

    private Handler mScreensaverHandler;
    private Runnable mScreensaverThread;
    private Spinner mLocationSpinner;
    private ImageView mNWStateImageView;
    private ImageView mUploadDataImageView;
    private FrameLayout fl_UploadData;
    private FrameLayout fl_comment;
    private ImageView mComplaintmageView;
    private Context mContext;
    private Fragment chartTableFragment;
    private  boolean canShowScreenSaver;
    private boolean doubleBackToExitPressedOnce;
    private Handler mHandler = new Handler();

    private final Runnable mRunnable = new Runnable() {
        @Override
        public void run() {
            doubleBackToExitPressedOnce = false;
        }
    };


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        try {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_chart);
            mContext = this;
            canShowScreenSaver = false;

            AppLogger.getInstance().Log(AppLogger.LogLevel.Debug, "ChartActivity Launched");
            
            AppRepo.getInstance().addPropertyChangeListener(this);

            SPLApplication.getInstance().setmChartActivity(this);

            mNWStateImageView = (ImageView) findViewById(R.id.imgNWState);
            mLocationSpinner = (Spinner) findViewById(R.id.locationSpinner);

            mUploadDataImageView = (ImageView)findViewById(R.id.uploadData);
            mComplaintmageView = (ImageView) findViewById(R.id.imgCommentView);

            fl_UploadData = ((FrameLayout)findViewById(R.id.fl_uploadData));
            fl_comment = ((FrameLayout)findViewById(R.id.fl_comment));

            initMainViewModel();

            ArrayAdapter locArrAdapter = new ArrayAdapter(this, android.R.layout.simple_spinner_item, MainViewModel.getInstance().getHeaderViewModel().getLocationList());
            locArrAdapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);
            //Setting the ArrayAdapter data on the Spinner
            mLocationSpinner.setAdapter(locArrAdapter);

            // Begin the transaction
            FragmentTransaction ft = getFragmentManager().beginTransaction();
            chartTableFragment = new ChartTableFragment();
            ft.replace(R.id.fragmentPlaceHolder, chartTableFragment);
            ft.commit();


            //init ScreenSaver Handler
            mScreensaverHandler = new Handler();
            mScreensaverThread = new Runnable() {
                @Override
                public void run() {
                    Intent intent = new Intent(getApplicationContext(), ScreenSaverActivity.class);
                    startActivity(intent);
                }
            };

            new LocationChartRunnable(AppRepo.getInstance().getCurrentLocationId()).run();
        } catch (Exception ex) {
            Log.d("ChartActivityOnCreErr", ex.getMessage());
        }
    }

    @Override
    public void onWindowFocusChanged(boolean hasFocus) {
        super.onWindowFocusChanged(hasFocus);

        if (hasFocus) {
            getWindow().getDecorView().setSystemUiVisibility(
                    View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                            | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY);

            getWindow().addFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN);
        }
    }


    @Override
    public void onFragmentInteraction(Uri uri) {
        Log.i("ChartActivity", "onFragmentInteraction");
    }

    private void initMainViewModel() {
        //TODO init MainViewModel
        //Set Chart Data
        // MainViewModel.getInstance().setChartViewModel(CommonUtility.getChartViewModel());

        //Set Location List
        ArrayList<String> locationList = new ArrayList<String>();
        //Temp Add location
        locationList.add("ServicePoint1");

        //TODO How to initi HeaderViewModel

        MainViewModel.getInstance().getHeaderViewModel().setLocationList(locationList);

        //Temp set NW STATE
        MainViewModel.getInstance().getHeaderViewModel().setNetworkState(Constants.NETWORK_STATE.WEB_SOCKET_CONNECTED);

        setNWStateIcon(Constants.NETWORK_STATE.WEB_SOCKET_DISSCONNECTED);
    }

    private void processClickedCellList(String mAuthCode) {
        /*ArrayList<CellViewModel> clickedCellList = MainViewModel.getInstance().getCurrenClickeCellModelList();

        if (mAuthCode == null || mAuthCode.isEmpty() || mAuthCode.length() == 0 || mAuthCode.equals("")) {
            Toast.makeText(this, getResources().getString(R.string.enter_auth_code), Toast.LENGTH_LONG).show();
        } else {
            //TODO Update this to NW
            sendToPacketManager(clickedCellList);

            for (CellViewModel cellViewModel : clickedCellList)
                updateChartModel(cellViewModel);

            //TODO Need to confirm before removing CurrenClickeCellModelList
            //Remove CurrenClickeCellModelList
            MainViewModel.getInstance().setCurrenClickeCellModelList(new ArrayList<CellViewModel>());
        }*/
    }

    public void stopHandler() {
        mScreensaverHandler.removeCallbacks(mScreensaverThread);
    }

    public void startHandler() {
        if (!canShowScreenSaver)return;
        mScreensaverHandler.postDelayed(mScreensaverThread, Constants.SCREEN_IDLE_TIMEOUT);
    }

    @Override
    public void onBackPressed() {

        if (doubleBackToExitPressedOnce) {
            super.onBackPressed();
            finish();
            return;
        }

        this.doubleBackToExitPressedOnce = true;
        Toast.makeText(this, "Please click BACK again to exit", Toast.LENGTH_SHORT).show();

        mHandler.postDelayed(mRunnable, 2000);
    }

    @Override
    public void onUserInteraction() {
        super.onUserInteraction();
        stopHandler();
        startHandler();
    }

    @Override
    protected void onResume() {
        super.onResume();
        startHandler();
        BroadCastReceiverManager.Instance().RegisterBatteryLevelReceiver(this);
    }

    @Override
    protected void onPause() {
        super.onPause();
        stopHandler();
        BroadCastReceiverManager.Instance().DeregisterBatteryLevelReceiver(this);
    }

    @Override
    protected void onStop() {
        super.onStop();
    }

    @Override
    protected  void onDestroy(){
        AppLogger.getInstance().Log(AppLogger.LogLevel.Error, "Activity getting distroyed");
        AppRepo.getInstance().removePropertyChangeListener(this);

        removeFromRecent(this);

        AppHelper.DeInit();

        super.onDestroy();

        finishAffinity();
        System.exit(0);
    }

    private void setNWStateIcon(Constants.NETWORK_STATE state) {
        switch (state) {
            case WEB_SOCKET_CONNECTED: {
                mNWStateImageView.setBackground(getResources().getDrawable(R.drawable.online));
                break;
            }
            case WEB_SOCKET_DISSCONNECTED: {
                mNWStateImageView.setBackground(getResources().getDrawable(R.drawable.offline));
                break;
            }
            case WEB_SOCKET_UNAUTHORIZED: {
                mNWStateImageView.setBackground(getResources().getDrawable(R.drawable.unauthorized));
                break;
            }

            case NW_NOT_AVAILABLE: {
                mNWStateImageView.setBackground(getResources().getDrawable(R.drawable.offline));
                break;
            }
            default: {
                mNWStateImageView.setBackground(getResources().getDrawable(R.drawable.offline));
                break;
            }
        }
    }


    public void setChartModel(ChartConfigModel model) {
        Handler hdl = new Handler(this.getMainLooper());
        hdl.post(new Runnable() {
            Activity executionContext;
            ChartConfigModel chartConfigModel;

            public Runnable init(Activity exeContext, ChartConfigModel model) {
                executionContext = exeContext;
                chartConfigModel = model;
                return this;
            }

            @Override
            public void run() {
                ((ChartTableFragment) chartTableFragment).setChart(executionContext, chartConfigModel);
            }
        }.init(this, model));
    }

    public void setChartDataModel(DisplayChartDataModel model) {
        Handler hdl = new Handler(this.getMainLooper());
        hdl.post(new Runnable() {
            Activity executionContext;
            DisplayChartDataModel chartDataModel;

            public Runnable init(Activity exeContext, DisplayChartDataModel model) {
                executionContext = exeContext;
                chartDataModel = model;
                return this;
            }

            @Override
            public void run() {
                ((ChartTableFragment) chartTableFragment).setChartData(executionContext, chartDataModel);
            }
        }.init(this, model));
    }


    @Override
    public void propertyChange(PropertyChangeEvent evt) {

        Handler uiHandler = new Handler(Looper.getMainLooper()) {
            @Override
            public void handleMessage(Message message) {

             switch (  message.getData().getString("PropertyName")){
                 case AppRepo.IsServerConnectedPropName:
                     boolean isConnected = message.getData().getBoolean("ConnectionState");
                     if (isConnected) {
                         setNWStateIcon(Constants.NETWORK_STATE.WEB_SOCKET_CONNECTED);
                     } else {
                         setNWStateIcon(Constants.NETWORK_STATE.WEB_SOCKET_DISSCONNECTED);
                     }
                     break;
                 case AppRepo.DeviceAuthorizedPropName:

                     boolean isAuthorized = message.getData().getBoolean("IsAuthorized");

                     if (isAuthorized == false){
                         setNWStateIcon(Constants.NETWORK_STATE.WEB_SOCKET_UNAUTHORIZED);
                     }

                     break;

                 case AppRepo.IsChartRenderedPropName:
                     boolean ischartrendered = message.getData().getBoolean("ischartrendered");

                     if (ischartrendered){
                         fl_UploadData.setClickable(true);
                         fl_UploadData.setOnClickListener(new ChartActivityClickHandler());

                         fl_comment.setClickable(true);
                         fl_comment.setOnClickListener(new ChartActivityClickHandler());

                         mUploadDataImageView.setEnabled(true);
                         mComplaintmageView.setEnabled(true);
                         mUploadDataImageView.setAlpha(1f);
                         mComplaintmageView.setAlpha(1f);

                         canShowScreenSaver= true;
                         startHandler();
                     }else{
                         fl_UploadData.setClickable(false);
                         fl_UploadData.setOnClickListener(null);
                         fl_comment.setClickable(false);
                         fl_comment.setOnClickListener(null);

                         mUploadDataImageView.setEnabled(false);
                         mComplaintmageView.setEnabled(false);
                         mUploadDataImageView.setAlpha(.2f);
                         mComplaintmageView.setAlpha(0.2f);

                         canShowScreenSaver= false;
                         stopHandler();
                     }


                     break;
             }
            }
        };

        Message msg = uiHandler.obtainMessage();
        Bundle b = new Bundle();
        b.putString("PropertyName", evt.getPropertyName());

        switch (evt.getPropertyName()) {
            case AppRepo.IsServerConnectedPropName:
                b.putBoolean("ConnectionState", (boolean) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;

            case AppRepo.DeviceAuthorizedPropName:
                b.putBoolean("IsAuthorized", (boolean) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;
            case AppRepo.IsChartRenderedPropName:
                //Setting value in bundle is required
                b.putBoolean("ischartrendered", (boolean) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;
        }
    }

    private void processUserComments(String strComments) {
        //TODO : Send User comments to server
        Log.i("ChartActivity", "User COmments : " + strComments);
    }


    public void removeFromRecent(Context context)
    {
        ActivityManager am = (ActivityManager)getSystemService(Context.ACTIVITY_SERVICE);
        if(am != null) {
            List<ActivityManager.AppTask> tasks = am.getAppTasks();
            if (tasks != null && tasks.size() > 0) {
                tasks.get(0).setExcludeFromRecents(true);
            }
        }
    }

}

