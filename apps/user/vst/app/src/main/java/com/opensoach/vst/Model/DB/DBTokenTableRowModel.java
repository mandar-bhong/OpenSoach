package com.opensoach.vst.Model.DB;

import com.opensoach.vst.Constants.DBTableConstants;
import com.opensoach.vst.DAL.DBTableSchema;

import java.util.Date;

@DBTableSchema(TableName = DBTableConstants.TABLE_TOKEN_DATA)
public class DBTokenTableRowModel {

    private int   id;
    private int    tokenno;
    private  String vehicleno;
    private String mapping;
    private int state;
    private Date generatedon;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getTokenno() {
        return tokenno;
    }

    public void setTokenno(int tokenno) {
        this.tokenno = tokenno;
    }

    public String getVehicleno() {
        return vehicleno;
    }

    public void setVehicleno(String vehicleno) {
        this.vehicleno = vehicleno;
    }

    public String getMapping() {
        return mapping;
    }

    public void setMapping(String mapping) {
        this.mapping = mapping;
    }

    public int getState() {
        return state;
    }

    public void setState(int state) {
        this.state = state;
    }

    public Date getGeneratedon() {
        return generatedon;
    }

    public void setGeneratedon(Date generatedon) {
        this.generatedon = generatedon;
    }
}
