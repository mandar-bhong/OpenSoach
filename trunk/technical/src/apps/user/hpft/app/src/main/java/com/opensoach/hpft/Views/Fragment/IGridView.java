package com.opensoach.sme.Views.Interfaces;

import android.widget.BaseAdapter;

import java.util.Iterator;

/**
 * Created by Mandar on 8/25/2017.
 */

public interface IGridView {

    BaseAdapter getDataAdaptor();

    BaseAdapter setItemsSource(Iterator item);
}
