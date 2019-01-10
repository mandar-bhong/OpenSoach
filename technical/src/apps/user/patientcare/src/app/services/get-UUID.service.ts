import { Injectable } from '@angular/core';
// import { v4 as uuid } from 'uuid';
// var uuid = require('uuid/v1');
var uuid = require("nativescript-uuid");
@Injectable()

export class GetUUIDService {

    uuid:string;
    
    getUUID() {
        this.uuid = uuid.getUUID();
        // console.log("UUID is " + this.uuid);
        return this.uuid;
    }
}