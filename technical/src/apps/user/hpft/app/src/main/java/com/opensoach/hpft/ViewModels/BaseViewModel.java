package com.opensoach.hpft.ViewModels;

import android.databinding.BaseObservable;
import android.support.v7.app.AppCompatActivity;

/**
 * Created by Mandar on 27-06-2018.
 */

public class BaseViewModel extends BaseObservable {
    public BaseViewModel Parent;
    public AppCompatActivity ContextActivity;
}
