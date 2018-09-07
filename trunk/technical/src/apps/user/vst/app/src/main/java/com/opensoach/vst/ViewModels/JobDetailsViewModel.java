package com.opensoach.vst.ViewModels;

import com.opensoach.vst.Model.DB.DBJobCreationDetailsTableRowModel;

public class JobDetailsViewModel extends BaseViewModel {


    private TokenItemViewModel tokenItemViewModel;
    private TokenSelectionViewModel tokenSelectionViewModel;
    private DBJobCreationDetailsTableRowModel dbJobCreationDetailsTableRowModel;

    public DBJobCreationDetailsTableRowModel getDbJobCreationDetailsTableRowModel() {
        return dbJobCreationDetailsTableRowModel;
    }

    public void setDbJobCreationDetailsTableRowModel(DBJobCreationDetailsTableRowModel dbJobCreationDetailsTableRowModel) {
        this.dbJobCreationDetailsTableRowModel = dbJobCreationDetailsTableRowModel;
    }

    public TokenSelectionViewModel getTokenSelectionViewModel() {
        return tokenSelectionViewModel;
    }

    public void setTokenSelectionViewModel(TokenSelectionViewModel tokenSelectionViewModel) {
        this.tokenSelectionViewModel = tokenSelectionViewModel;
    }

    public TokenItemViewModel getTokenItemViewModel() {
        return tokenItemViewModel;
    }

    public void setTokenItemViewModel(TokenItemViewModel tokenItemViewModel) {
        this.tokenItemViewModel = tokenItemViewModel;
    }



}
