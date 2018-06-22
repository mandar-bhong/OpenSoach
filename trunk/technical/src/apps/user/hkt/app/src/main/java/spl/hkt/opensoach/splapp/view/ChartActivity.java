package spl.hkt.opensoach.splapp.view;

import android.app.Activity;
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
import android.widget.ImageView;
import android.widget.Spinner;

import java.beans.PropertyChangeEvent;
import java.beans.PropertyChangeListener;
import java.util.ArrayList;

import spl.hkt.opensoach.splapp.Constants;
import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.SPLApplication;
import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.handler.ChartActivityClickHandler;
import spl.hkt.opensoach.splapp.logger.AppLogger;
import spl.hkt.opensoach.splapp.manager.LocationChartRunnable;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.view.ChartConfigModel;
import spl.hkt.opensoach.splapp.model.view.DisplayChartDataModel;
import spl.hkt.opensoach.splapp.viewModels.CellViewModel;
import spl.hkt.opensoach.splapp.viewModels.HeaderViewModel;
import spl.hkt.opensoach.splapp.viewModels.MainViewModel;
import spl.hkt.opensoach.splapp.viewModels.TaskRowViewModel;

public class ChartActivity extends Activity implements ChartTableFragment.OnFragmentInteractionListener, PropertyChangeListener {

    private Handler mScreensaverHandler;
    private Runnable mScreensaverThread;
    private Spinner mLocationSpinner;
    private ImageView mNWStateImageView;
    private ImageView mUploadDataImageView;
    private ImageView mComplaintmageView;
    private Context mContext;
    private Fragment chartTableFragment;
    private  boolean canShowScreenSaver;


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
            mUploadDataImageView.setOnClickListener(new ChartActivityClickHandler());

            mComplaintmageView = (ImageView) findViewById(R.id.imgCommentView);
            mComplaintmageView.setOnClickListener(new ChartActivityClickHandler());


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
        HeaderViewModel headerViewModel = new HeaderViewModel();
        headerViewModel.setLocationList(locationList);
        MainViewModel.getInstance().setHeaderViewModel(headerViewModel);

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
    public void onUserInteraction() {
        super.onUserInteraction();
        stopHandler();
        startHandler();
    }

    @Override
    protected void onResume() {
        super.onResume();
        startHandler();
    }

    @Override
    protected void onPause() {
        super.onPause();
        stopHandler();
    }

    @Override
    protected void onStop() {
        super.onStop();
    }

    @Override
    protected  void onDestroy(){
        AppRepo.getInstance().removePropertyChangeListener(this);
        super.onDestroy();
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

                 case AppRepo.CurrentLocationIdPropName:
                     mUploadDataImageView.setEnabled(true);
                     mComplaintmageView.setEnabled(true);
                     mUploadDataImageView.setAlpha(1f);
                     mComplaintmageView.setAlpha(1f);

                     canShowScreenSaver= true;
                     startHandler();
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
            case AppRepo.CurrentLocationIdPropName:
                //Setting value in bundle is required
                b.putInt("locationid", (int) evt.getNewValue());
                msg.setData(b);
                uiHandler.sendMessage(msg);
                break;
        }
    }

    private void processUserComments(String strComments) {
        //TODO : Send User comments to server
        Log.i("ChartActivity", "User COmments : " + strComments);
    }

}

