package com.opensoach.hpft.ViewModels;

import android.widget.BaseAdapter;

import com.opensoach.hpft.Views.Interfaces.IGridView;
import com.opensoach.hpft.Views.Interfaces.IList;
import com.opensoach.hpft.Views.Miscellaneous.CardBriefViewAdaptor;

import java.util.Iterator;
import java.util.List;

/**
 * Created by Mandar on 30-07-2018.
 */

public class CardGridViewModel extends BaseViewModel implements IList<CardBriefViewModel>,IGridView {

    private CardBriefViewAdaptor dataAdapter;


    public CardGridViewModel(){
        dataAdapter = new CardBriefViewAdaptor();
    }

    @Override
    public BaseAdapter getDataAdaptor() {
        return dataAdapter;
    }

    @Override
    public BaseAdapter setItemsSource(Iterator item) {
        return null;
    }

    @Override
    public List<CardBriefViewModel> getItemsSource() {
        return   dataAdapter.getItemsSource();
    }

    @Override
    public void setItemsSource(List<CardBriefViewModel> source) {
        dataAdapter.setItemsSource(source);
    }
}
