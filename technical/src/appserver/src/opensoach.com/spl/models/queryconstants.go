package models

const QUERY_GET_CONFIGURATION = "SELECT param_key,category,value FROM spl_master_config where category = ?"
