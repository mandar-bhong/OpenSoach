package spl.hkt.opensoach.splapp.view;


import spl.hkt.opensoach.splapp.viewModels.CellViewModel;

/**
 * Created by samir.s.bukkawar on 3/14/2017.
 */

public interface ICellClick {
    CellViewModel getCellViewModel();

    void onCellClick(CellViewModel chartViewModel);
}
