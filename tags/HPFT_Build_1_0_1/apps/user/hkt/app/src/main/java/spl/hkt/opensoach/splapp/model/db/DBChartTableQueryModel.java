package spl.hkt.opensoach.splapp.model.db;

import android.content.ContentValues;
import android.database.Cursor;

import java.util.Date;

import spl.hkt.opensoach.splapp.dal.DBConstants;
import spl.hkt.opensoach.splapp.dal.DBTableSchema;
import spl.hkt.opensoach.splapp.dal.IDBRowMapper;

/**
 * Created by Mandar on 3/12/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_CHART)
public class DBChartTableQueryModel implements IDBRowMapper<DBChartTableRowModel> {

    public static final String SELECT_CHART_ID_FILTER = "SELECT_CHART_ID_FILTER";
    public static final String SELECT_LOCATION_ID_FILTER = "SELECT_LOCATION_ID_FILTER";
    public static final String SELECT_CHART_PAYLOAD_FILTER = "SELECT_CHART_PAYLOAD_FILTER";
    public static final String SELECT_SERVER_CHART_ID_FILTER = "SELECT_SERVER_CHART_ID_FILTER";

    @Override
    public DBChartTableRowModel Clone() {
        return new DBChartTableRowModel();
    }


    @Override
    public void PrepareModel(Cursor cursor,DBChartTableRowModel dataModel) {
        dataModel.setChartId(cursor.getInt(0));
        dataModel.setServerChartId(cursor.getInt(1));
        dataModel.setLocationId(cursor.getInt(2));
        dataModel.setChartPayload(cursor.getString(3));
        dataModel.setChartDispStartDate(new Date( cursor.getLong(4)));
        dataModel.setChartDispEndDate(new Date( cursor.getLong(5)));
    }

    @Override
    public String[] SelectColumn() {
        return new String[]{DBConstants.TABLE_CHART_CHART_ID,
                DBConstants.TABLE_CHART_SERVER_CHART_ID,
                DBConstants.TABLE_CHART_LOCATION_ID,
                DBConstants.TABLE_CHART_CHART_PAYLOAD,
                DBConstants.TABLE_CHART_CHART_DISP_START_DATE,
                DBConstants.TABLE_CHART_CHART_DISP_END_DATE
        };
    }

    @Override
    public String WhereFilter(String filterName) {

        switch (filterName) {
            case SELECT_CHART_ID_FILTER:
                return DBConstants.TABLE_CHART_CHART_ID + "=?";
            case SELECT_LOCATION_ID_FILTER:
                return DBConstants.TABLE_CHART_LOCATION_ID + "=?";
            case SELECT_CHART_PAYLOAD_FILTER:
                return DBConstants.TABLE_CHART_CHART_PAYLOAD+ "=?";
            case SELECT_SERVER_CHART_ID_FILTER:
                return DBConstants.TABLE_CHART_SERVER_CHART_ID+ "=?";
        }
        return "";
    }

    @Override
    public String[] FilterArgs(DBChartTableRowModel dataModel,String filterName) {
        switch (filterName) {
            case SELECT_CHART_ID_FILTER:
                return  new String[]{String.valueOf(dataModel.getChartId())};
            case SELECT_LOCATION_ID_FILTER:
                return  new String[]{String.valueOf(dataModel.getLocationId())};
            case SELECT_CHART_PAYLOAD_FILTER:
                return  new String[]{String.valueOf(dataModel.getChartPayload())};
            case SELECT_SERVER_CHART_ID_FILTER:
                return  new String[]{String.valueOf(dataModel.getServerChartId())};
        }

        return new String[]{};
    }

    @Override
    public ContentValues UpdateFieldSet(DBChartTableRowModel dataModel,String filterName) {
        ContentValues values = new ContentValues ();

        switch (filterName) {
            case SELECT_SERVER_CHART_ID_FILTER:
                values.put(DBConstants.TABLE_CHART_CHART_PAYLOAD, dataModel.getChartPayload());
                return values;
        }

        return values;
    }
}
