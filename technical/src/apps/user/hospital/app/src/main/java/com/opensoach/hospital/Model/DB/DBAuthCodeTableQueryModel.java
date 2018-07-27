package com.opensoach.hospital.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;
import com.opensoach.hospital.DAL.IDBRowMapper;

/**
 * Created by Mandar on 8/26/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_AUTH_CODE)
public class DBAuthCodeTableQueryModel implements IDBRowMapper<DBAuthCodeTableRowModel> {

    public static final String SELECT_ALL_FILTER = "SELECT_ALL_FILTER";
    public static final String UDATE_ALL_FILTER = "UDATE_ALL_FILTER";

    @Override
    public DBAuthCodeTableRowModel Clone() {
        return new DBAuthCodeTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBAuthCodeTableRowModel dataModel) {
        dataModel.setAuthCode(cursor.getString(0));
    }


    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_AUTH_CODE_AUTHCODE};
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBAuthCodeTableRowModel dataModel, String filterName) {
        switch (filterName) {
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBAuthCodeTableRowModel dataModel, String filterName) {

        ContentValues values = new ContentValues();

        switch (filterName) {
            case UDATE_ALL_FILTER:
                values.put(DBConstants.TABLE_AUTH_CODE_AUTHCODE, dataModel.getAuthCodeJSON());
                break;
        }

        return values;
    }
}
