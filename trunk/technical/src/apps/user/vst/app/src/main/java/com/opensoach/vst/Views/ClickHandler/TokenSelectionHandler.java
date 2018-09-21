package com.opensoach.vst.Views.ClickHandler;

import android.content.Intent;
import android.view.View;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.Constants.ApplicationConstants;
import com.opensoach.vst.Constants.CommandConstants;
import com.opensoach.vst.Constants.Constants;
import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Helper.PacketHelper;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.Communication.PacketHeaderModel;
import com.opensoach.vst.ViewModels.JobExeDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceDetailsViewModel;
import com.opensoach.vst.ViewModels.JobServiceItemViewModel;
import com.opensoach.vst.ViewModels.JobServiceListViewModel;
import com.opensoach.vst.ViewModels.JobServiceViewModel;
import com.opensoach.vst.ViewModels.TokenSelectionViewModel;
import com.opensoach.vst.Views.Activity.JobServiceDetailsActivity;
import com.opensoach.vst.Views.Activity.JobServiceTaskListActivity;

import java.util.ArrayList;

public class TokenSelectionHandler {

    public void onClick(View view, TokenSelectionViewModel vm) {

        switch (AppRepo.getInstance().getCurrentRunningMode()) {
            case Token: {

            }
            break;
            case JobCreation: {

                if (AppRepo.getInstance().getStore().containsKey(
                        vm.getTokenListViewModel().
                                getSelectedToken().
                                getDbTokenTableRowModel().
                                getVehicleno())) {

                } else {

                    JobServiceDetailsViewModel jobDetailsViewModel = new JobServiceDetailsViewModel();
                    jobDetailsViewModel.Parent = vm;
                    jobDetailsViewModel.ContextActivity = vm.ContextActivity;

                    jobDetailsViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
                    AppRepo.getInstance().setJobServiceDetailsViewModel(jobDetailsViewModel);

                    JobServiceListViewModel jobServiceListViewModel = new JobServiceListViewModel();
                    jobServiceListViewModel.Parent = jobDetailsViewModel;
                    jobServiceListViewModel.ContextActivity = vm.ContextActivity;
                    jobServiceListViewModel.setData(new ArrayList<JobServiceItemViewModel>());

                    JobServiceViewModel jobServiceViewModel = new JobServiceViewModel();
                    jobServiceViewModel.Parent = vm;
                    jobServiceViewModel.ContextActivity = vm.ContextActivity;

                    jobServiceViewModel.setJobServiceDetailsViewModel(jobDetailsViewModel);
                    jobServiceViewModel.setJobServiceListViewModel(jobServiceListViewModel);
                    jobServiceViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
                    AppRepo.getInstance().setJobServiceViewModel(jobServiceViewModel);

                    AppRepo.getInstance().getStore().put(
                            vm.getTokenListViewModel().
                                    getSelectedToken().
                                    getDbTokenTableRowModel().
                                    getVehicleno(),jobDetailsViewModel);


                }

                Intent i = new Intent(view.getContext(), JobServiceDetailsActivity.class);
                view.getContext().startActivity(i);

                SendPacketManager.Instance().send(AppAction.CLAIM_TOKEN, vm.getTokenListViewModel().getSelectedToken());
                SendPacketManager.Instance().send(AppAction.GET_VEHICLE_DETAILS, vm.getTokenListViewModel().getSelectedToken().getVehicleNo());

            }
            break;
            case JobExecution: {

                if (AppRepo.getInstance().getStore().containsKey(
                        vm.getTokenListViewModel().
                                getSelectedToken().
                                getDbTokenTableRowModel().
                                getVehicleno())) {

                } else {
                JobServiceDetailsViewModel jobDetailsViewModel = new JobServiceDetailsViewModel();
                jobDetailsViewModel.Parent = vm;
                jobDetailsViewModel.ContextActivity = vm.ContextActivity;

                jobDetailsViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
                AppRepo.getInstance().setJobServiceDetailsViewModel(jobDetailsViewModel);

                JobServiceListViewModel jobServiceListViewModel = new JobServiceListViewModel();
                jobServiceListViewModel.Parent = jobDetailsViewModel;
                jobServiceListViewModel.ContextActivity = vm.ContextActivity;
                jobServiceListViewModel.setData(new ArrayList<JobServiceItemViewModel>());

                JobServiceViewModel jobServiceViewModel = new JobServiceViewModel();
                jobServiceViewModel.Parent = vm;
                jobServiceViewModel.ContextActivity = vm.ContextActivity;

                jobServiceViewModel.setJobServiceDetailsViewModel(jobDetailsViewModel);
                jobServiceViewModel.setJobServiceListViewModel(jobServiceListViewModel);
                jobServiceViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());


                AppRepo.getInstance().setJobServiceViewModel(jobServiceViewModel);

                JobExeDetailsViewModel jobExeDetailsViewModel = new JobExeDetailsViewModel();
                jobExeDetailsViewModel.Parent = vm;
                jobExeDetailsViewModel.ContextActivity = vm.ContextActivity;

                jobExeDetailsViewModel.setTokenItemViewModel(vm.getTokenListViewModel().getSelectedToken());
                AppRepo.getInstance().setJobExeDetailsViewModel(jobExeDetailsViewModel);

                    AppRepo.getInstance().getStore().put(
                            vm.getTokenListViewModel().
                                    getSelectedToken().
                                    getDbTokenTableRowModel().
                                    getVehicleno(),jobDetailsViewModel);

                }

                Intent i = new Intent(view.getContext(), JobServiceTaskListActivity.class);
                view.getContext().startActivity(i);
                break;
            }
        }


    }
}
