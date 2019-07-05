package dbquery

const QUERY_DOCUMENT_TABLE_SELECT_BY_UUID = `Select id,uuid,name,doctype,location,location_type,persisted From spl_hpft_document_tbl Where uuid = ?`
