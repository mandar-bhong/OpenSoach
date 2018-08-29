package com.opensoach.vst.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.vst.DAL.IDBRowMapper;

public class DBTokenTableQueryModel implements IDBRowMapper<DBTokenTableRowModel> {
    @Override
    public DBTokenTableRowModel Clone() {
        return new DBTokenTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBTokenTableRowModel dataModel) {
       
    }

    @Override
    public String[] SelectColumn() {
        return new String[0];
    }

    @Override
    public String WhereFilter(String filterName) {
        return null;
    }

    @Override
    public String[] FilterArgs(DBTokenTableRowModel dataModel, String filterName) {
        return new String[0];
    }

    @Override
    public ContentValues UpdateFieldSet(DBTokenTableRowModel dataModel, String filterName) {
        return null;
    }
}
