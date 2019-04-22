package com.opensoach.hpft.Views.Notifier;

import com.opensoach.hpft.Model.View.PropChangeDataModel;

/**
 * Created by Mandar on 9/19/2017.
 */

public class NotifyPropChangeOnUIThread implements Runnable{

    PropChangeDataModel propChangeDataModel;
    @Override
    public void run() {
        propChangeDataModel.ChangeSupport.firePropertyChange(propChangeDataModel.PropName,propChangeDataModel.OldValue,propChangeDataModel.NewValue);
    }

    public NotifyPropChangeOnUIThread(PropChangeDataModel model){
        propChangeDataModel = model;
    }

}
