package com.opensoach.hospital.Views.ClickHandler;

import android.view.View;

import com.opensoach.hospital.ViewModels.JobBriefViewModel;
import com.opensoach.hospital.ViewModels.MainViewModel;

import java.util.List;

/**
 * Created by Mandar on 9/7/2017.
 */

public class TestButtonClickHander {

    public void onClick(View view, MainViewModel vm) {

        //DatabaseManager.SelectAll()


//        JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
//        //jobBriefViewModel.AppContext = this.getBaseContext();
//        jobBriefViewModel.ContextActivity = vm.ContextActivity;
//        jobBriefViewModel.setPart("Part: " + "MNO");
//        jobBriefViewModel.setCustomer("Customer: " + "XYZ");
//        jobBriefViewModel.setJob("kjhhj");
//
//        int colorInt = new Random().nextInt(3);
//
//        //vm.GridViewModel.getItemsSource().add(jobBriefViewModel );
//
//        List<JobBriefViewModel> models =vm.GridViewModel.getItemsSource();
//
//        //models.add(jobBriefViewModel);
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));
//
//        //vm.GridViewModel.getItemsSource().remove(0);
//
//        //vm.GridViewModel.setItemsSource(models);
//        //vm.GridViewModel.notifyPropertyChanged(R.layout.fragment_job_brief_list);vm
//
//        //vm.GridViewModel.notify();
//
////        synchronized(vm.GridViewModel){
////            vm.GridViewModel.notify();
////        }
//
//        vm.GridViewModel.getDataAdaptor().notifyDataSetChanged();


//        if(AppRepo.getInstance().getCurrentLocationId() == 1){
//            AppRepo.getInstance().setCurrentLocationId(2);
//        }else if (AppRepo.getInstance().getCurrentLocationId() == 2){
//            AppRepo.getInstance().setCurrentLocationId(3);
//        }else{
//
//            InsertJoBDataIntoDatabase();
//            AppRepo.getInstance().setCurrentLocationId(1);
//        }
//
//        List<JobBriefViewModel> models =vm.GridViewModel.getItemsSource();
//
//        //models.add(jobBriefViewModel);
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity,1).get(0));

        //vm.GridViewModel.getItemsSource().remove(0);

        //vm.GridViewModel.setItemsSource(models);
        //vm.GridViewModel.notifyPropertyChanged(R.layout.fragment_job_brief_list);vm

        //vm.GridViewModel.notify();

//        synchronized(vm.GridViewModel){
//            vm.GridViewModel.notify();
//        }

//        vm.GridViewModel.getDataAdaptor().notifyDataSetChanged();

        List<JobBriefViewModel> models = vm.GridViewModel.getItemsSource();

        //models.add(jobBriefViewModel);
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity, 1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity, 1).get(0));
//        vm.GridViewModel.getItemsSource().add(TestDataHelper.GenerateData(vm.ContextActivity, 1).get(0));

//        ArrayList<String> locList = new ArrayList<>();
//        locList.add("Location 1");
//        locList.add("Location 2");
//        vm.setLocationList(locList);

        //vm.GridViewModel.getItemsSource().remove(0);

        //vm.GridViewModel.setItemsSource(models);
        //vm.GridViewModel.notifyPropertyChanged(R.layout.fragment_job_brief_list);vm

        //vm.GridViewModel.notify();

//        synchronized(vm.GridViewModel){
//            vm.GridViewModel.notify();
//        }

//        TestDataHelper.InsertTestDataIntoDatabase();

        vm.GridViewModel.getDataAdaptor().notifyDataSetChanged();
    }
}
