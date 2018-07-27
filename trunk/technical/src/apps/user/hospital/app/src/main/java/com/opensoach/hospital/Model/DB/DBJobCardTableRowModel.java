package com.opensoach.hospital.Model.DB;

import com.opensoach.hospital.DAL.DBConstants;
import com.opensoach.hospital.DAL.DBTableSchema;

import java.util.Date;


/**
 * Created by Mandar on 9/5/2017.
 */

@DBTableSchema(TableName = DBConstants.TABLE_JOB_CARD)
public class DBJobCardTableRowModel {

    private Integer jobCardId;
    private Integer locationId;
    private String customer;
    private Integer partId;
    private Integer partCount;
    private String jobCode;
    private Date startDate;
    private Date endDate;
    private Date actualStartDate;
    private Date actualEndDate;
    private Integer state;
    private String comments;
    private Integer completedCount;
    private String jobConfig;

    public Integer getJobCardId() {
        return jobCardId;
    }

    public void setJobCardId(Integer jobCardId) {
        this.jobCardId = jobCardId;
    }

    public Integer getLocationId() {
        return locationId;
    }

    public void setLocationId(Integer locationId) {
        this.locationId = locationId;
    }

    public String getCustomer() {
        return customer;
    }

    public void setCustomer(String customer) {
        this.customer = customer;
    }

    public Integer getPartId() {
        return partId;
    }

    public void setPartId(Integer partId) {
        this.partId = partId;
    }


    public Integer getPartCount() {
        return partCount;
    }

    public void setPartCount(Integer partCount) {
        this.partCount = partCount;
    }

    public String getJobCode() {
        return jobCode;
    }

    public void setJobCode(String jobCode) {
        this.jobCode = jobCode;
    }

    public Date getStartDate() {
        return startDate;
    }

    public void setStartDate(Date startDate) {
        this.startDate = startDate;
    }

    public Date getEndDate() {
        return endDate;
    }

    public void setEndDate(Date endDate) {
        this.endDate = endDate;
    }

    public Date getActualStartDate() {
        return actualStartDate;
    }

    public void setActualStartDate(Date actualStartDate) {
        this.actualStartDate = actualStartDate;
    }

    public Date getActualEndDate() {
        return actualEndDate;
    }

    public void setActualEndDate(Date actualEndDate) {
        this.actualEndDate = actualEndDate;
    }

    public Integer getState() {
        return state;
    }

    public void setState(Integer state) {
        this.state = state;
    }

    public String getComments() {
        return comments;
    }

    public void setComments(String comments) {
        this.comments = comments;
    }

    public Integer getCompletedCount() {
        return completedCount;
    }

    public void setCompletedCount(Integer completedCount) {
        this.completedCount = completedCount;
    }

    public String getJobConfig() {
        return jobConfig;
    }

    public void setJobConfig(String jobConfig) {
        this.jobConfig = jobConfig;
    }
}
