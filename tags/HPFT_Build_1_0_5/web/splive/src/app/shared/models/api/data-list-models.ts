export class DataListRequest<T> {
    /**
     * Search parameters to be used for filtering the data retrieved.
     */
    filter: T;
    /**
     * The number of records to be displayed in the list
     */
    limit: number;

    /**
     * The name of the column on which data needs to be sorted.
     * The column name matches with the JSON tag used in the list of data returned.
     */
    orderby: string;

    /**
     * The direction in which the data needs to be sorted.
     * asc: Ascending
     * desc: Descending
     */
    orderdirection: string;

    /**
     * The page number.
     * Page number starts from 1.
     */
    page: number;
}

export class DataListResponse<T> {
    /**
     * Total number of records exists in the database.
     */
    totalrecords: number;
    /**
     * Mumber of records filtered based on the search criteria
     */
    filteredrecords: number;

    /**
     * array of records.
     */
    records: T[];
}
