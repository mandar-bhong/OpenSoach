package spl.hkt.opensoach.splapp.dal;

import android.content.ContentValues;
import android.database.Cursor;

/**
 * Created by Mandar on 3/11/2017.
 */

public interface IDBRowMapper<T> {

    T Clone();
    void PrepareModel(Cursor cursor,T dataModel);
    String[] SelectColumn();
    String WhereFilter(String filterName);
    String[] FilterArgs(T dataModel,String filterName);
    ContentValues UpdateFieldSet(T dataModel,String filterName);
}
