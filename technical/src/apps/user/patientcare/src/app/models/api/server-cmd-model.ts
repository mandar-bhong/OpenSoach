export class CmdModel {
    header: Header;
    payload: any;
}

export class Header {
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

export class tokenModel {
    token: string;
}