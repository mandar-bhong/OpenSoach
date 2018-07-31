package com.opensoach.hpft.Views.Miscellaneous;


import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentStatePagerAdapter;

import com.opensoach.hpft.Views.Fragment.MedicalDetailsFragment;
import com.opensoach.hpft.Views.Fragment.PatientDetailsFragment;

/**
 * Created by Mandar on 31-07-2018.
 */

public class TabPagerAdapter extends FragmentStatePagerAdapter {

    int mNumOfTabs;

    public TabPagerAdapter(FragmentManager fm, int NumOfTabs) {
        super(fm);
        this.mNumOfTabs = NumOfTabs;
    }

    @Override
    public Fragment getItem(int position) {

        switch (position) {
            case 0:
                PatientDetailsFragment tab1 = new PatientDetailsFragment();
                return tab1;
            case 1:
                MedicalDetailsFragment tab2 = new MedicalDetailsFragment();
                return tab2;
//            case 2:
//                TabFragment3 tab3 = new TabFragment3();
//                return tab3;
            default:
                return null;
        }
    }

    @Override
    public int getCount() {
        return mNumOfTabs;
    }
}
