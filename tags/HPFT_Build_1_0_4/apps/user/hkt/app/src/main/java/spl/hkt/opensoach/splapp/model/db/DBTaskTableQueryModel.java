package spl.hkt.opensoach.splapp.model.db;

import android.content.ContentValues;
import android.database.Cursor;

import spl.hkt.opensoach.splapp.dal.DBConstants;
import spl.hkt.opensoach.splapp.dal.DBTableSchema;
import spl.hkt.opensoach.splapp.dal.IDBRowMapper;

/**
 * Created by Mandar on 3/12/2017.
 */
@DBTableSchema(TableName = DBConstants.TABLE_TASKS)
public class DBTaskTableQueryModel implements IDBRowMapper<DBTaskTableRowModel> {

    public static final String SELECT_ID_FILTER = "SELECT_ID_FILTER";
    public static final String UPDATE_INFO_FILTER_BY_ID = "UPDATE_INFO_FILTER_BY_ID";

    @Override
    public DBTaskTableRowModel Clone() {
        return new DBTaskTableRowModel();
    }

    @Override
    public void PrepareModel(Cursor cursor, DBTaskTableRowModel dataModel) {
        dataModel.setTaskId(cursor.getInt(0));
        dataModel.setTaskName(cursor.getString(1));
        dataModel.setTaskOrder(cursor.getInt(2));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_TASKS_TASK_ID,
                DBConstants.TABLE_TASKS_TASK_NAME,
                DBConstants.TABLE_TASKS_TASK_ORDER
        };
    }

    @Override
    public String WhereFilter(String filterName) {

        switch (filterName) {
            case UPDATE_INFO_FILTER_BY_ID:
            case SELECT_ID_FILTER:
                return DBConstants.TABLE_TASKS_TASK_ID + "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBTaskTableRowModel dataModel, String filterName) {
        switch (filterName) {
            case SELECT_ID_FILTER:
            case UPDATE_INFO_FILTER_BY_ID:
                return new String[]{String.valueOf(dataModel.getTaskId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBTaskTableRowModel dataModel, String filterName) {
        ContentValues values = new ContentValues ();

        switch (filterName) {
            case UPDATE_INFO_FILTER_BY_ID:
                values.put(DBConstants.TABLE_TASKS_TASK_NAME, dataModel.getTaskName());
                values.put(DBConstants.TABLE_TASKS_TASK_ORDER, dataModel.getTaskOrder());
                return values;
        }

        return values;
    }
}
