package com.opensoach.vst;

import android.app.Activity;
import android.app.Application;
import android.content.Context;
import android.content.Intent;
import android.databinding.DataBindingUtil;
import android.util.Log;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Communication.WebSocketConnector;
import com.opensoach.vst.DAL.DatabaseManager;
import com.opensoach.vst.Helper.AppHelper;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Helper.CommonHelper;
import com.opensoach.vst.Helper.ExceptionHelper;
import com.opensoach.vst.Model.AppNotificationModelBase;
import com.opensoach.vst.Model.Communication.PacketCardListConfigurationModel;
import com.opensoach.vst.Model.Communication.PacketModel;
import com.opensoach.vst.Model.Communication.PacketServiceCustomerDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceInstanceModel;
import com.opensoach.vst.Model.Communication.PacketServiceJobDetailsDataModel;
import com.opensoach.vst.Model.Communication.PacketServiceTaskItemDataModel;
import com.opensoach.vst.Model.Communication.PacketSimpleAckModel;
import com.opensoach.vst.Model.Communication.PacketVehicleDetailsModel;
import com.opensoach.vst.Model.DB.DBLocationTableRowModel;
import com.opensoach.vst.Model.DB.DBTokenTableRowModel;
import com.opensoach.vst.Model.View.ChartConfigModel;
import com.opensoach.vst.Model.View.DisplayChartDataModel;
import com.opensoach.vst.Scheduler.ScheduleManager;
import com.opensoach.vst.Utility.AppLogger;
import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.ViewModels.JobCustomerDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TokenItemViewModel;
import com.opensoach.vst.Views.Activity.TaskDetailsActivity;
import com.opensoach.vst.Views.Activity.TokenListActivity;
import com.opensoach.vst.Views.Activity.TokenSelectionActivity;
import com.opensoach.vst.Views.ChartActivity;
import com.opensoach.vst.Views.TimeChangeListner;
import com.opensoach.vst.Views.UpdateChartListner;
import com.opensoach.vst.Views.DataBinding.AppDataBindingComponent;

import org.apache.log4j.chainsaw.Main;

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

        //Setting global unhandled exception
        new ExceptionHelper(this);

        (new ScheduleManager()).startScheduler(this, 1, 30, 1);

        DataBindingUtil.setDefaultComponent(new AppDataBindingComponent());
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

            case ApplicationConstants.UI_PROCESSING_STATERGY_TOKEN_CREATED: {

                final DBTokenTableRowModel tokenItem = (DBTokenTableRowModel) model.Data;
                final TokenItemViewModel viewModel = new TokenItemViewModel(tokenItem);

                MainViewModel.getInstance().ContextActivity.runOnUiThread(new Runnable() {
                    @Override
                    public void run() {

                        MainViewModel.getInstance().getTokenListViewModel().getTokensDataAdapter().addItem(viewModel);

                        if (MainViewModel.getInstance().getCreateTokenViewModel() != null) {
                            MainViewModel.getInstance().getCreateTokenViewModel().setDbTokenTableRowModel(tokenItem);
                        }
                    }
                });
            }
            break;
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

                if (packetAuthCodeDataModel == null) {
                    AppRepo.getInstance().setAuthCodeList(new ArrayList<String>());
                } else {
                    AppRepo.getInstance().setAuthCodeList(packetAuthCodeDataModel);
                }
            }

            case ApplicationConstants.UI_PROCESSING_STATERGY_CARD_LIST_DATA: {

                PacketCardListConfigurationModel packetCardListConfigurationModel = (PacketCardListConfigurationModel) model.Data;

                String configType = packetCardListConfigurationModel.ConfTypeCode;
                Intent i = null;

                AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_SERVICE_INST_ID, packetCardListConfigurationModel.SerInID);
                AppRepo.getInstance().getStore().put(ApplicationConstants.APP_STORE_SERVICE_CONFIG_ID, packetCardListConfigurationModel.ServConfID);


                switch (configType) {
                    case "TOKEN_GENERATION":
                        i = new Intent(MainViewModel.getInstance().ContextActivity, TokenListActivity.class);
                        AppRepo.getInstance().setCurrentRunningMode(ApplicationConstants.AppRunningMode.Token);
                        break;
                    case "JOB_CREATION":
                        i = new Intent(MainViewModel.getInstance().ContextActivity, TokenSelectionActivity.class);
                        AppRepo.getInstance().setCurrentRunningMode(ApplicationConstants.AppRunningMode.JobCreation);
                        break;
                    case "JOB_EXECUTION":
                        i = new Intent(MainViewModel.getInstance().ContextActivity, TokenSelectionActivity.class);
                        AppRepo.getInstance().setCurrentRunningMode(ApplicationConstants.AppRunningMode.JobExecution);
                        break;
                }

                //MainViewModel.getInstance().ContextActivity.finish();
                MainViewModel.getInstance().ContextActivity.startActivity(i);

            }
            break;


            case ApplicationConstants.UI_PROCESSING_STATERGY_TOKEN_LIST: {
                final ArrayList<DBTokenTableRowModel> tokenList = (ArrayList<DBTokenTableRowModel>) model.Data;

                MainViewModel.getInstance().ContextActivity.runOnUiThread(new Runnable() {

                    @Override
                    public void run() {

                        MainViewModel.getInstance().getTokenListViewModel().getData().clear();

                        for (DBTokenTableRowModel tokenModel : tokenList) {

                            TokenItemViewModel tokenItemViewModel = new TokenItemViewModel(tokenModel);

                            tokenItemViewModel.ContextActivity = MainViewModel.getInstance().ContextActivity;
                            tokenItemViewModel.Parent = MainViewModel.getInstance().getTokenListViewModel();

                            MainViewModel.getInstance().getTokenListViewModel().getData().add(tokenItemViewModel);
                        }

                        MainViewModel.getInstance().getTokenListViewModel().getTokensDataAdapter().notifyDataSetChanged();
                    }
                });

            }
            break;

            case ApplicationConstants.UI_PROCESSING_STATERGY_VEHICLE_DETAILS: {
                final PacketServiceCustomerDetailsDataModel custDetails = (PacketServiceCustomerDetailsDataModel) model.Data;

                MainViewModel.getInstance().ContextActivity.runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        //JobServiceDetailsViewModel
                        JobServiceDetailsViewModel jobServiceDetailsVM = AppRepo.getInstance().getJobServiceViewModel().getJobServiceDetailsViewModel();
                        jobServiceDetailsVM.setFirstName(custDetails.FirstName);
                        jobServiceDetailsVM.setLastName(custDetails.LastName);
                        jobServiceDetailsVM.setMobileNo(custDetails.MobileNo);
                    }
                });
                break;
            }
            case ApplicationConstants.UI_PROCESSING_STATERGY_JOB_SERVICE_DETAILS: {
                final PacketServiceJobDetailsDataModel jobDetails = (PacketServiceJobDetailsDataModel) model.Data;

                if (jobDetails.ServiceConfig.size() < 1) {
                    return;
                }

                Gson gson = new GsonBuilder().setDateFormat(ApplicationConstants.PACKET_DATE_FORMAT).create();

                Type packetType = new TypeToken<PacketVehicleDetailsModel>() {
                }.getType();

                try{
                    final PacketVehicleDetailsModel configTasks = gson.fromJson(jobDetails.ServiceConfig.get(0).txndata, packetType);
                    //final PacketVehicleDetailsModel exeTasks = gson.fromJson(jobDetails.ServiceExeConfig.get(0).txndata, packetType);

                    MainViewModel.getInstance().ContextActivity.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {

                            JobServiceListViewModel jobListVM = AppRepo.getInstance().getJobServiceViewModel().getJobServiceListViewModel();
                            for (PacketServiceTaskItemDataModel model : configTasks.Tasks) {

                                JobServiceItemViewModel vm = new JobServiceItemViewModel();
                                vm.Parent = jobListVM;
                                vm.ContextActivity = jobListVM.ContextActivity;
                                vm.setComment(model.Comment);
                                vm.setCost(model.Cost);
                                vm.setTaskName(model.taskName);

                                jobListVM.getJobServiceDataAdapter().addItem(vm);
                            }

                            //TODO: Implement the completed task checkbox
//                            for (PacketServiceTaskItemDataModel model : exeTasks) {
//                                for (JobServiceItemViewModel vm : jobListVM.getData()) {
//                                    if (model.taskName.equals( vm.getTaskName())) {
//                                        vm.setTaskCompleted(true);
//                                        break;
//                                    }
//                                }
//                            }
                        }
                    });

                }catch (Exception ex){
                    AppLogger.getInstance().Log(ex);
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