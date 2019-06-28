package com.opensoach.hpft.Model.View;

import java.beans.PropertyChangeSupport;

/**
 * Created by Mandar on 9/19/2017.
 */

public class PropChangeDataModel {

    public Object OldValue;
    public Object NewValue;
    public  String PropName;
    public  PropertyChangeSupport ChangeSupport;

    public PropChangeDataModel(PropertyChangeSupport changeSupport,String propName,Object oldValue,Object newValue){
        ChangeSupport = changeSupport;
        PropName = propName;
        OldValue = oldValue;
        NewValue= newValue;
    }
}
