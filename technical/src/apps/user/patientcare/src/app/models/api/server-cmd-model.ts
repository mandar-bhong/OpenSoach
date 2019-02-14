export class CmdModel {
    header: CmdHeader;
    payload: any;
}

export class CmdHeader {
    crc: string;
    category: number;
    commandid: number;
    seqid: number;
}

export class GetSyncRequestModel {
    storename: string;
    updatedon: Date;
}

export class ApplySyncRequestModel {
    storename: string;
    storedata: any[];
}

export class AuthTokenModel {
    authtoken: string;
}