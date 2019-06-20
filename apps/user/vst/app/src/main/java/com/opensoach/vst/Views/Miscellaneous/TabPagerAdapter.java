package com.opensoach.vst.Views.Miscellaneous;


import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentStatePagerAdapter;

import com.opensoach.vst.ViewModels.CardBriefViewModel;
import com.opensoach.vst.Views.Fragment.PatientDetailsFragment;
import com.opensoach.vst.Views.Fragment.TaskListFragment;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 31-07-2018.
 */

public class TabPagerAdapter extends FragmentStatePagerAdapter {

    int mNumOfTabs;
    private CardBriefViewModel cardBrief;
    private final List<Fragment> mFragmentList = new ArrayList<>();


    public TabPagerAdapter(FragmentManager fm, int NumOfTabs,
                           CardBriefViewModel vm) {
        super(fm);
        this.mNumOfTabs = NumOfTabs;
        cardBrief = vm;
    }

    @Override
    public Fragment getItem(int position) {

        switch (position) {
            case 0:
                PatientDetailsFragment patientDetailsFragment = new PatientDetailsFragment();
                patientDetailsFragment.DataContext = cardBrief;
                return patientDetailsFragment;
//            case 1:
//                MedicalDetailsFragment medicalDetailsFragment = new MedicalDetailsFragment();
//                medicalDetailsFragment.DataContext = cardBrief.getMedicalDetails();
//                return medicalDetailsFragment;
            case 1:
                TaskListFragment taskListFragment = new TaskListFragment();
                taskListFragment.DataContext = cardBrief.getTaskDetails();
                return taskListFragment;
//
//                TaskDetailsFragment taskListFragment = new TaskDetailsFragment();
//                taskListFragment.DataContext = cardBrief.getTaskDetails();
//                return taskListFragment;

            default:
                return null;
        }
    }

    @Override
    public int getCount() {
        return mNumOfTabs;
    }

    @Override
    public int getItemPosition(Object object) {
        // refresh all fragments when data set changed
        return TabPagerAdapter.POSITION_NONE;
    }
}
