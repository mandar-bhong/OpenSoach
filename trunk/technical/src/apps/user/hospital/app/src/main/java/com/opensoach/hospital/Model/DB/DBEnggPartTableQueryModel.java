package com.opensoach.hospital.Model.DB;

import android.content.ContentValues;
import android.database.Cursor;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;
import com.opensoach.hospital.DAL.IDBRowMapper;

/**
 * Created by Mandar on 9/5/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_ENGG_PART)
public class DBEnggPartTableQueryModel implements IDBRowMapper<DBEnggPartTableRowModel> {

    public static final String SELECT_ALL_FILTER = "SELECT_ALL_FILTER";
    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";


    @Override
    public DBEnggPartTableRowModel Clone() {
        return new DBEnggPartTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBEnggPartTableRowModel dataModel) {
        dataModel.setPartId(cursor.getInt(0));
        dataModel.setPartNo(cursor.getString(1));
        dataModel.setPartRevision(cursor.getString(2));
        dataModel.setInternalPartNo(cursor.getString(3));
        dataModel.setProcess(cursor.getString(4));
        dataModel.setToolJSON(cursor.getString(5));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_ENGG_PART_PART_ID,
                DBConstants.TABLE_ENGG_PART_PART_NO,
                DBConstants.TABLE_ENGG_PART_PART_REVISION,
                DBConstants.TABLE_ENGG_PART_INTERNAL_PART_NO,
                DBConstants.TABLE_ENGG_PART_PROCESS,
                DBConstants.TABLE_ENGG_PART_TOOL_JSON
        };
    }

    @Override
    public String WhereFilter(String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return DBConstants.TABLE_ENGG_PART_PART_ID + "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBEnggPartTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
                return new String[]{String.valueOf(dataModel.getPartId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBEnggPartTableRowModel dataModel, String filterName) {
        return null;
    }
}
