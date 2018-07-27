package com.opensoach.hospital.ViewModels;

import android.databinding.BaseObservable;
import android.support.v7.app.AppCompatActivity;

/**
 * Created by Mandar on 8/25/2017.
 */

public class BaseViewModel extends BaseObservable {

    //public Context AppContext;
    public BaseViewModel Parent;
    //public FragmentManager FragmentMgr;
    public AppCompatActivity ContextActivity;

}
