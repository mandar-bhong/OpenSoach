package com.opensoach.hospital.Model.DB;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;

/**
 * Created by Mandar on 9/5/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_ENGG_PART)
public class DBEnggPartTableRowModel {

    private Integer partId;
    private String partNo;
    private String partRevision;
    private String internalPartNo;
    private String process;
    private String toolJSON;

    public Integer getPartId() {
        return partId;
    }

    public void setPartId(Integer partId) {
        this.partId = partId;
    }

    public String getPartNo() {
        return partNo;
    }

    public void setPartNo(String partNo) {
        this.partNo = partNo;
    }

    public String getPartRevision() {
        return partRevision;
    }

    public void setPartRevision(String partRevision) {
        this.partRevision = partRevision;
    }

    public String getInternalPartNo() {
        return internalPartNo;
    }

    public void setInternalPartNo(String internalPartNo) {
        this.internalPartNo = internalPartNo;
    }

    public String getProcess() {
        return process;
    }

    public void setProcess(String process) {
        this.process = process;
    }

    public String getToolJSON() {
        return toolJSON;
    }

    public void setToolJSON(String toolJSON) {
        this.toolJSON = toolJSON;
    }
}
