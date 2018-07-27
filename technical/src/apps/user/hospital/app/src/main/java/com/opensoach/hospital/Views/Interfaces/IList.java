package com.opensoach.hospital.Views.Interfaces;

import java.util.List;

/**
 * Created by Mandar on 8/25/2017.
 */

public interface IList<T> {

    List<T> getItemsSource();

    void setItemsSource(List<T> source);


}
