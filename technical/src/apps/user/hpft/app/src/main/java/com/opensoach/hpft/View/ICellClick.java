package com.opensoach.hpft.View;


import com.opensoach.hpft.ViewModels.CellViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface ICellClick {
    CellViewModel getCellViewModel();

    void onCellClick(CellViewModel chartViewModel);
}
