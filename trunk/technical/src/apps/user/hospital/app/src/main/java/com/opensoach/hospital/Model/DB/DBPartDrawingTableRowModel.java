package com.opensoach.hospital.Model.DB;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;

/**
 * Created by Mandar on 9/10/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_PART_DRAWING)
public class DBPartDrawingTableRowModel {

    private int drawingId;
    private int partId;
    private String path;

    public int getDrawingId() {
        return drawingId;
    }

    public void setDrawingId(int drawingId) {
        this.drawingId = drawingId;
    }

    public int getPartId() {
        return partId;
    }

    public void setPartId(int partId) {
        this.partId = partId;
    }

    public String getPath() {
        return path;
    }

    public void setPath(String path) {
        this.path = path;
    }
}
