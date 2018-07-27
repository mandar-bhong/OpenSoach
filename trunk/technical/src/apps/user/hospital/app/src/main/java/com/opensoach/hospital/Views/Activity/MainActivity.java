package com.opensoach.hospital.Views.Activity;

import android.content.Context;
import android.databinding.DataBindingUtil;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.os.Process;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.view.WindowManager;
import android.widget.Toast;

import com.opensoach.hospital.AppRepo.AppRepo;
import com.opensoach.hospital.Helper.AppHelper;
import com.opensoach.hospital.Helper.ApplicationConstants;
import com.opensoach.hospital.Helper.ObjectCompararHelper;
import com.opensoach.hospital.Model.AppNotificationModelBase;
import com.opensoach.hospital.Model.View.UIDeletedJobDataModel;
import com.opensoach.hospital.Model.View.UIJobStateChangedDataModel;
import com.opensoach.hospital.Model.View.UINewJobDataModel;
import com.opensoach.hospital.Model.View.UIServerSyncCompletedModel;
import com.opensoach.hospital.R;
import com.opensoach.hospital.Utility.AppLogger;
import com.opensoach.hospital.ViewModels.HeaderViewModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.JobGridViewModel;
import com.opensoach.hospital.ViewModels.MainViewModel;
import com.opensoach.hospital.Views.Interfaces.IFragment;
import com.opensoach.hospital.Views.Interfaces.IUIUpdateEvent;
import com.opensoach.hospital.databinding.ActivityMainBinding;

import java.util.ArrayList;
import java.util.List;

public class MainActivity extends AppCompatActivity implements IFragment<MainViewModel>,IUIUpdateEvent {

    private Thread.UncaughtExceptionHandler defaultUEH;
    private Context mContext;
    private static MainActivity instance;
    private boolean doubleBackToExitPressedOnce;
    private Handler mHandler = new Handler();

    private final Runnable mRunnable = new Runnable() {
        @Override
        public void run() {
            doubleBackToExitPressedOnce = false;
        }
    };

    public MainActivity()  {

    }

    public static MainActivity getInstance() {
        return instance;
    }


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        instance = this;
        mContext = this;

        AppLogger.getInstance().Log(AppLogger.LogLevel.Debug,"Starting App");

        AppHelper.Init(mContext);
        //AppHelper.ExecuteStartUpProcess();

        ActivityMainBinding activityMainBinding = DataBindingUtil.setContentView(this, R.layout.activity_main);

        setDataContext(MainViewModel.getInstance());

        //method to hide the Keyboard onCreate of Activity
        hideSoftKeyboard();


//        new Handler().postDelayed(new Runnable() {
//            @Override
//            public void run() {
//                Intent i = new Intent(MainActivity.this, SplashScreenActivity.class);
//                startActivity(i);
//            }
//        },0);

    }

    public void hideSoftKeyboard() {
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_STATE_HIDDEN);
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
    public MainViewModel getDataContext() {
        return null;
    }

    @Override
    public void setDataContext(MainViewModel viewModel) {
        ActivityMainBinding activityMainBinding = DataBindingUtil.setContentView(this, R.layout.activity_main);

        //viewModel.AppContext = this.getBaseContext();
        viewModel.ContextActivity = this;

        JobGridViewModel jobGridViewModel = new JobGridViewModel();
        jobGridViewModel.ContextActivity = this;
        jobGridViewModel.setItemsSource(new ArrayList<JobBriefViewModel>());
		//jobGridViewModel.setItemsSource(GenerateData());
        viewModel.setGridViewModel(jobGridViewModel);

        viewModel.setHeaderViewModel(new HeaderViewModel());

        activityMainBinding.setMainViewModel(viewModel);
    }


    @Override
    public void onBackPressed() {

        if (doubleBackToExitPressedOnce) {
            super.onBackPressed();
            return;
        }

        this.doubleBackToExitPressedOnce = true;
        Toast.makeText(this, "Please click BACK again to exit", Toast.LENGTH_SHORT).show();

        mHandler.postDelayed(mRunnable, 2000);
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();

        AppHelper.DeInit();

        if (mHandler != null) {
            mHandler.removeCallbacks(mRunnable);
        }

        android.os.Process.killProcess(Process.myPid());
    }

    @Override
    protected void onResume() {
        super.onResume();
        AppRepo.getInstance().setForegroundActivityName(ApplicationConstants.FOREGROUND_ACTIVITY_MAIN);
        AppRepo.getInstance().setForegroundActivityHandler(this);
    }


@Override
    public void OnUIUpdateEvent(final AppNotificationModelBase model) {
    switch (model.DataProcessStatergyID) {
        case ApplicationConstants.UI_PROCESSING_STATERGY_SERVER_DATA_LOAD_COMPLETED: {

            Handler uiHandler = new Handler(Looper.getMainLooper()) {
                @Override
                public void handleMessage(Message message) {

                    UIServerSyncCompletedModel uiServerSyncCompletedModel = (UIServerSyncCompletedModel) model;

                    MainViewModel.getInstance().GridViewModel.getItemsSource().clear();

                    for (JobBriefViewModel jobBriefViewModel : uiServerSyncCompletedModel.getJobBriefViewModels()) {

                        jobBriefViewModel.ContextActivity = MainViewModel.getInstance().ContextActivity;
                        MainViewModel.getInstance().GridViewModel.getItemsSource().add(jobBriefViewModel);
                    }

                    AppRepo.getInstance().setLocationList(((UIServerSyncCompletedModel) model).getLocations());
                    MainViewModel.getInstance().getHeaderViewModel().setLocations(((UIServerSyncCompletedModel) model).getLocations());
                    MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    MainViewModel.getInstance().setJobStatusTextChanged();
                }
            };

            Message msg = uiHandler.obtainMessage();
            Bundle b = new Bundle();

            //msg.setData();
            uiHandler.sendMessage(msg);
        }
        break;

        case ApplicationConstants.UI_PROCESSING_STATERGY_NEW_JOB_AVAILABLE: {
            Handler uiHandler = new Handler(Looper.getMainLooper()) {
                @Override
                public void handleMessage(Message message) {

                    UINewJobDataModel uiNewJobDataModel = (UINewJobDataModel) model;

                    MainViewModel.getInstance().GridViewModel.getItemsSource().clear();

                    for (JobBriefViewModel jobBriefViewModel : uiNewJobDataModel.getJobBriefViewModels()) {

                        jobBriefViewModel.ContextActivity = MainViewModel.getInstance().ContextActivity;
                        MainViewModel.getInstance().GridViewModel.getItemsSource().add(jobBriefViewModel);
                    }

                    MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    MainViewModel.getInstance().setJobStatusTextChanged();
                }
            };

            Message msg = uiHandler.obtainMessage();
            Bundle b = new Bundle();

            //msg.setData();
            uiHandler.sendMessage(msg);
        }
        break;


        case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_REMOVED: {
            Handler uiHandler = new Handler(Looper.getMainLooper()) {
                @Override
                public void handleMessage(Message message) {

                    UIDeletedJobDataModel uiDeletedJobDataModel = (UIDeletedJobDataModel) model;

                    List<Integer> deletedJobsIndex = new ArrayList<>();

                    //TODO: This looping can be moved to background thread
                    for (JobBriefViewModel jobVM : MainViewModel.getInstance().GridViewModel.getItemsSource()) {

                        for (Integer deletedJob : uiDeletedJobDataModel.getDeletedJobs()) {

                            boolean isMatched = ObjectCompararHelper.Equal(jobVM, deletedJob);

                            if (isMatched) {
                                Integer jobIndex = MainViewModel.getInstance().GridViewModel.getItemsSource().indexOf(jobVM);
                                deletedJobsIndex.add(jobIndex);
                            }
                        }
                    }

                    for (Integer index : deletedJobsIndex) {
                        MainViewModel.getInstance().GridViewModel.getItemsSource().remove(index);
                    }

                    MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    MainViewModel.getInstance().setJobStatusTextChanged();
                }
            };

            Message msg = uiHandler.obtainMessage();
            Bundle b = new Bundle();

            //msg.setData();
            uiHandler.sendMessage(msg);
        }
        break;

        case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_STATE_UPDATED: {

            Handler uiHandler = new Handler(Looper.getMainLooper()) {
                @Override
                public void handleMessage(Message message) {

                    UIJobStateChangedDataModel uiJobStateChangedDataModel = (UIJobStateChangedDataModel) model;

                    MainViewModel.getInstance().GridViewModel.getItemsSource().clear();

                    for (JobBriefViewModel jobBriefViewModel : uiJobStateChangedDataModel.getJobBriefViewModels()) {

                        jobBriefViewModel.ContextActivity = MainViewModel.getInstance().ContextActivity;
                        MainViewModel.getInstance().GridViewModel.getItemsSource().add(jobBriefViewModel);
                    }

                    MainViewModel.getInstance().GridViewModel.getDataAdaptor().notifyDataSetChanged();
                    MainViewModel.getInstance().setJobStatusTextChanged();
                }
            };

            Message msg = uiHandler.obtainMessage();
            Bundle b = new Bundle();

            //msg.setData();
            uiHandler.sendMessage(msg);

        }
    }
}
}