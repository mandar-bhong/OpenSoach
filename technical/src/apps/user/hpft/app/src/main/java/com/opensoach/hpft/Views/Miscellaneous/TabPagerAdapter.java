package com.opensoach.hpft.Views.Miscellaneous;


import android.databinding.DataBindingUtil;
import android.os.Handler;
import android.os.Looper;
import android.support.annotation.MainThread;
import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentStatePagerAdapter;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.LinearLayout;

import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.Views.Activity.CardListActivity;
import com.opensoach.hpft.Views.ClickHandler.CardItemClickHandler;
import com.opensoach.hpft.Views.Fragment.MedicalDetailsFragment;
import com.opensoach.hpft.Views.Fragment.PatientDetailsFragment;
import com.opensoach.hpft.Views.Fragment.TaskListFragment;
import com.opensoach.hpft.Views.Notifier.NotifyPropChangeOnUIThread;
import com.opensoach.hpft.databinding.FragmentPatientDetailsBinding;

/**
 * Created by Mandar on 31-07-2018.
 */

public class TabPagerAdapter extends FragmentStatePagerAdapter {

    int mNumOfTabs;
    private CardBriefViewModel cardBrief;


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
                patientDetailsFragment.DataContext = cardBrief.getPatientDetails();
                return patientDetailsFragment;
            case 1:
                MedicalDetailsFragment medicalDetailsFragment = new MedicalDetailsFragment();
                return medicalDetailsFragment;
            case 2:
                TaskListFragment taskListFragment = new TaskListFragment();
                return taskListFragment;
            default:
                return null;
        }
    }

    @Override
    public int getCount() {
        return mNumOfTabs;
    }

}
