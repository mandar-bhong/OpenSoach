package com.opensoach.sme.Views.Interfaces;

/**
 * Created by Mandar on 8/25/2017.
 */

public interface IFragment<T> {

    T getDataContext();

    void setDataContext(T viewModel);
}

