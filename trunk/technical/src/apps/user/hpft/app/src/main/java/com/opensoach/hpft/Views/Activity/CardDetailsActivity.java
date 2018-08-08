package com.opensoach.hpft.Views.Activity;

import android.net.Uri;
import android.support.design.widget.TabLayout;
import android.support.v4.view.PagerAdapter;
import android.support.v4.view.ViewPager;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.support.v7.widget.Toolbar;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.view.WindowManager;

import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.Views.Fragment.HeaderFragment;
import com.opensoach.hpft.Views.Fragment.MedicalDetailsFragment;
import com.opensoach.hpft.Views.Fragment.PatientDetailsFragment;
import com.opensoach.hpft.Views.Fragment.TaskDetailsFragment;
import com.opensoach.hpft.Views.Fragment.TaskListFragment;
import com.opensoach.hpft.Views.Miscellaneous.TabPagerAdapter;

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
                }else{
                    MainViewModel.getInstance().getHeaderViewModel().setUploadVisiable(false);
                }

                if (AppRepo.getInstance().getSelectedTaskDataViewModels().size() > 0 ){
                    MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(true);
                }else{
                    MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(false);
                }

                //TODO: Calculate time and select value
                AppRepo.getInstance().getActiveCard().getTaskDetails().getTaskTimeDataViewModel().setSelectedTimeTaskItem();
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
