package spl.hkt.opensoach.splapp.dal;

import android.content.ContentValues;
import android.content.Context;
import android.database.Cursor;
import android.database.SQLException;
import android.database.sqlite.SQLiteDatabase;

import java.lang.reflect.Field;
import java.lang.reflect.Modifier;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.List;

import spl.hkt.opensoach.splapp.util.CommonUtility;


/**
 * Created by samir.s.bukkawar on 2/26/2017.
 */

public class DatabaseManager {

    private static DatabaseHandler mDbHelper;
    static HashMap<String, ArrayList<Field>> _classFields;

    static {
        _classFields = new HashMap<>();
    }

    public void Init(Context context) {

        mDbHelper = new DatabaseHandler(context);
    }

    public static SQLiteDatabase openDatabase() throws SQLException {

        return mDbHelper.getWritableDatabase();
    }

    public static void closeDatabase() {
        mDbHelper.close();
    }


    public static long InsertRow(Object dbModel) {

        String tableName = DBTableName(dbModel);

        ContentValues values = new ContentValues();

        ArrayList<Field> fields = DiscoverClassField(dbModel);

        for (Field f : fields) {
            f.setAccessible(true);
            String fieldType = f.getType().getName();
            String fieldName = f.getName();

            if(f.isAnnotationPresent(DBFieldSchema.class)) {
                DBFieldSchema fS = f.getAnnotation(DBFieldSchema.class);
                if(fS.IsAutoIncreament())
                    continue;
            }

            try {

                switch (fieldType) {
                    case "boolean":
                        if ((boolean) f.get(dbModel) == true) {
                            values.put(fieldName, 1);
                        } else {
                            values.put(fieldName, 0);
                        }
                        break;
                    case "int":
                        values.put(fieldName, (int) f.get(dbModel));
                        break;
                    case "java.lang.String":
                        values.put(fieldName, (String) f.get(dbModel));
                        break;
                    case "java.lang.Integer":
                        values.put(fieldName, (Integer) f.get(dbModel));
                        break;
                    case "java.lang.Short":
                        values.put(fieldName, (Short) f.get(dbModel));
                        break;
                    case "java.lang.Long":
                        values.put(fieldName, (Long) f.get(dbModel));
                        break;
                    case "java.lang.Float":
                        values.put(fieldName, (Float) f.get(dbModel));
                        break;
                    case "java.lang.Double":
                        values.put(fieldName, (Double) f.get(dbModel));
                        break;
                    case "java.util.Date":
                        values.put(fieldName,(long) ((Date) f.get(dbModel)).getTime());
                        break;
                }

            } catch (Exception ex) {

            }
        }

        SQLiteDatabase db = DatabaseManager.openDatabase();
        long count = db.insertOrThrow(tableName, null, values);
        DatabaseManager.closeDatabase();

        return count;
    }

    public static <T extends IDBRowMapper, K extends Object> List<K> SelectAll(T queryModel, K dataModel) {
        List selectList = new ArrayList<K>();

        SQLiteDatabase db = DatabaseManager.openDatabase();

        Cursor cursor = db.query(DBTableName(queryModel), queryModel.SelectColumn(), null, null, null, null, null);

        if (cursor != null && cursor.moveToFirst()) {
            do {
                K newDataModel = (K) queryModel.Clone();
                queryModel.PrepareModel(cursor, newDataModel);
                selectList.add(newDataModel);
            } while (cursor.moveToNext());
        }
        cursor.close();
        DatabaseManager.closeDatabase();
        return selectList;
    }

    public static <T extends IDBRowMapper, K extends Object> int GetRowCount(T queryModel, K dataModel, String filter) {

        List selectList = new ArrayList<K>();
        String tableName = DBTableName(queryModel);
        SQLiteDatabase db = DatabaseManager.openDatabase();

        Cursor cursor = db.query(tableName, queryModel.SelectColumn(), queryModel.WhereFilter(filter),
                queryModel.FilterArgs(dataModel, filter), null, null, null, null);

        int rowCount = 0;
        if (cursor != null && cursor.moveToFirst()) {
            do {
                rowCount++;
            } while (cursor.moveToNext());
        }
        cursor.close();
        DatabaseManager.closeDatabase();

        return rowCount;
    }

    public static <T extends IDBRowMapper, K extends Object> List<K> SelectByFilter(T queryModel, K dataModel, String filter) {

        List selectList = new ArrayList<K>();
        String tableName = DBTableName(queryModel);
        SQLiteDatabase db = DatabaseManager.openDatabase();

        Cursor cursor = db.query(tableName, queryModel.SelectColumn(), queryModel.WhereFilter(filter),
                queryModel.FilterArgs(dataModel, filter), null, null, null, null);

        if (cursor != null && cursor.moveToFirst()) {
            do {
                K newDataModel = (K) queryModel.Clone();
                queryModel.PrepareModel(cursor, newDataModel);
                selectList.add(newDataModel);
            } while (cursor.moveToNext());
        }
        cursor.close();
        DatabaseManager.closeDatabase();

        return selectList;
    }

    public static <T extends IDBRowMapper, T1 extends Object> int UpdateRow(T queryModel, T1 dataModel, String filter) {

        String tableName = DBTableName(queryModel);
        SQLiteDatabase db = DatabaseManager.openDatabase();

        int rowAffected = db.update(tableName, queryModel.UpdateFieldSet(dataModel, filter), queryModel.WhereFilter(filter),
                queryModel.FilterArgs(dataModel, filter));

        DatabaseManager.closeDatabase();

        return rowAffected;
    }

    public static <T extends IDBRowMapper, T1 extends Object> int DeleteByFilter(T queryModel, T1 dataModel, String filter) {

        String tableName = DBTableName(queryModel);
        SQLiteDatabase db = DatabaseManager.openDatabase();

        int affectedRows = db.delete(tableName, queryModel.WhereFilter(filter),
                queryModel.FilterArgs(dataModel, filter));

        //TODO: DB close or DatabaseManager.closeDatabase()
        DatabaseManager.closeDatabase();
        return affectedRows;
    }

    private static ArrayList<Field> DiscoverClassField(Object model) {

        String className = model.getClass().getName();
        ArrayList<Field> classFields;
        if (_classFields.containsKey(className)) {
            classFields = _classFields.get(className);

        } else {
            Field[] cf = model.getClass().getDeclaredFields();
            ArrayList<Field> cfList = new ArrayList<Field>();

            for (Field f : cf) {
                if (f.getModifiers() == Modifier.PRIVATE) {
                    cfList.add(f);
                }

            }
            classFields = cfList;

            _classFields.put(className, cfList);

        }

        return classFields;
    }

    private static String DBTableName(Object model) {

        String dbTableName = "";
        if (model.getClass().isAnnotationPresent(DBTableSchema.class)) {

            DBTableSchema clsAnn = (DBTableSchema) model.getClass().getAnnotation(DBTableSchema.class);
            dbTableName = clsAnn.TableName();
        }

        return dbTableName;
    }

}
