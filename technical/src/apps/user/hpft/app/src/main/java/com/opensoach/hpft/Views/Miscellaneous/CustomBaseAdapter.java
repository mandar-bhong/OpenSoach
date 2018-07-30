package com.opensoach.hpft.Views.Miscellaneous;

import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.GridView;

import com.opensoach.hpft.Views.Interfaces.IList;

import java.util.List;

/**
 * Created by Mandar on 8/25/2017.
 */

public abstract class CustomBaseAdapter<T, E> extends BaseAdapter implements IList<E> {

    public AppCompatActivity ContextActivity;
    public GridView GridViewContainer;

    List<E> ItemsSource;

    @Override
    public int getCount() {

        return ItemsSource.size();
    }

    @Override
    public Object getItem(int position) {
        return ItemsSource.get(position);
    }

    @Override
    public long getItemId(int position) {
        return position;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {

        //Need to comment below to stop re-creation of view when scrolling. It will maintain the grid order.
       /* if (convertView != null)
            return convertView;*/

        E dataModel = ItemsSource.get(position);
        return getItemView(dataModel, position);
    }

    @Override
    public List getItemsSource() {
        return ItemsSource;
    }

    @Override
    public void setItemsSource(List source) {
        ItemsSource = source;
    }

    protected abstract View getItemView(E dataModel, int position);

}
