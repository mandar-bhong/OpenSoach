package com.opensoach.vst.Views.Activity;

import android.net.Uri;
import android.os.Handler;
import android.support.design.widget.TabLayout;
import android.support.v4.view.PagerAdapter;
import android.support.v4.view.ViewPager;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.RecyclerView;
import android.view.Menu;
import android.view.MenuItem;

import com.opensoach.vst.AppRepo.AppRepo;
import com.opensoach.vst.R;
import com.opensoach.vst.ViewModels.MainViewModel;
import com.opensoach.vst.ViewModels.TaskDetailsViewModel;
import com.opensoach.vst.ViewModels.TaskTimeItemViewModel;
import com.opensoach.vst.Views.ClickHandler.TaskTimeClickHandler;
import com.opensoach.vst.Views.Fragment.HeaderFragment;
import com.opensoach.vst.Views.Fragment.MedicalDetailsFragment;
import com.opensoach.vst.Views.Fragment.PatientDetailsFragment;
import com.opensoach.vst.Views.Fragment.TaskDetailsFragment;
import com.opensoach.vst.Views.Fragment.TaskListFragment;
import com.opensoach.vst.Views.Miscellaneous.TabPagerAdapter;

import java.util.Date;

public class CardDetailsActivity extends AppCompatActivity
        implements PatientDetailsFragment.OnFragmentInteractionListener,
        MedicalDetailsFragment.OnFragmentInteractionListener,
        TaskListFragment.OnFragmentInteractionListener,
        TaskDetailsFragment.OnFragmentInteractionListener,
        HeaderFragment.OnFragmentInteractionListener{


    final String CONST_TAB_CHECK_LIST = "Check List";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_card_details);
//        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
//        setSupportActionBar(toolbar);



        TabLayout tabLayout = (TabLayout) findViewById(R.id.tab_layout);
        tabLayout.addTab(tabLayout.newTab().setText("Details"));
        //tabLayout.addTab(tabLayout.newTab().setText("Medical Details"));
        tabLayout.addTab(tabLayout.newTab().setText(CONST_TAB_CHECK_LIST));

        final ViewPager viewPager = (ViewPager) findViewById(R.id.pager);
        final PagerAdapter adapter = new TabPagerAdapter (getSupportFragmentManager(),
                tabLayout.getTabCount(),
                AppRepo.getInstance().getActiveCard());

        viewPager.setAdapter(adapter);
        viewPager.setOffscreenPageLimit(3);
        viewPager.addOnPageChangeListener(new TabLayout.TabLayoutOnPageChangeListener(tabLayout));
        tabLayout.addOnTabSelectedListener(new TabLayout.OnTabSelectedListener() {
            @Override
            public void onTabSelected(TabLayout.Tab tab) {
                viewPager.setCurrentItem(tab.getPosition());

                if (tab.getText() == CONST_TAB_CHECK_LIST){
                    MainViewModel.getInstance().getHeaderViewModel().setUploadVisiable(true);


                    TaskDetailsViewModel taskDetailsViewModel =  AppRepo.getInstance().getActiveCard().getTaskDetails();
                    TaskTimeItemViewModel taskTimeItemViewModel = taskDetailsViewModel.getSelectedItem();

                    if (taskTimeItemViewModel != null){

                        RecyclerView rvTime = (RecyclerView)findViewById(R.id.time_recycler_view);
                        rvTime.getLayoutManager().scrollToPosition(
                                taskTimeItemViewModel.getTaskTimeDataModel().getIndex()
                        );

                        new Handler(getMainLooper()).post(new Runnable() {
                            @Override
                            public void run() {
                                new TaskTimeClickHandler().onClick(null, AppRepo.getInstance().getActiveCard().getTaskDetails().getSelectedItem());
                            }
                        });
                    }

                }else{
                    MainViewModel.getInstance().getHeaderViewModel().setUploadVisiable(false);
                }

                if (AppRepo.getInstance().getSelectedTaskDataViewModels().size() > 0 ){
                    MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(true);
                }else{
                    MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(false);
                }
            }

            @Override
            public void onTabUnselected(TabLayout.Tab tab) {

            }

            @Override
            public void onTabReselected(TabLayout.Tab tab) {

            }
        });

        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        //TODO: Calculate time and select value
        //AppRepo.getInstance().getActiveCard().getTaskDetails().getTaskTimeDataViewModel().setSelectedTimeTaskItem();

        Date d = new Date();

        for (TaskTimeItemViewModel taskTimeDataModel : AppRepo.getInstance().getActiveCard().getTaskDetails().getTaskTimeDataViewModel().getData()){

            if (taskTimeDataModel.getTaskTimeDataModel().getStartTime().getTime() < d.getTime() &&
                    taskTimeDataModel.getTaskTimeDataModel().getEndTime().getTime() >    d.getTime() ){
                AppRepo.getInstance().getActiveCard().getTaskDetails().setSelectedItem(taskTimeDataModel);
                break;
            }
        }

    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        int id = item.getItemId();
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    protected  void onDestroy(){
        super.onDestroy();
        MainViewModel.getInstance().getHeaderViewModel().setUploadVisiable(false);
    }


    @Override
    public void onFragmentInteraction(Uri uri) {

    }

}
