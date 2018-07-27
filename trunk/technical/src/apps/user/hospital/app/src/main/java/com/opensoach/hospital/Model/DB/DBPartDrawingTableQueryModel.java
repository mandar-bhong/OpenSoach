package com.opensoach.hospital.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;
import com.opensoach.hospital.DAL.IDBRowMapper;

/**
 * Created by Mandar on 9/10/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_PART_DRAWING)
public class DBPartDrawingTableQueryModel implements IDBRowMapper<DBPartDrawingTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String SELECT_ALL_FILTER = "SELECT_ALL_FILTER";
    public static final String SELECT_BY_PART_ID_FILTER = "SELECT_BY_PART_ID_FILTER";


    @Override
    public DBPartDrawingTableRowModel Clone() {
        return new DBPartDrawingTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBPartDrawingTableRowModel dataModel) {
        dataModel.setDrawingId(cursor.getInt(0));
        dataModel.setPartId(cursor.getInt(1));
        dataModel.setPath(cursor.getString(2));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_PART_DRAWING_ID,
                DBConstants.TABLE_PART_DRAWING_PART_ID,
                DBConstants.TABLE_PART_DRAWING_PATH
        };
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return DBConstants.TABLE_PART_DRAWING_ID + "=?";
            case SELECT_BY_PART_ID_FILTER:
                return DBConstants.TABLE_PART_DRAWING_PART_ID + "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBPartDrawingTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getDrawingId())};
            case SELECT_BY_PART_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getPartId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBPartDrawingTableRowModel dataModel, String filterName) {
        return null;
    }
}
