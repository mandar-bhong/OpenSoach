import { Component } from "@angular/core";
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent {
    constructor(private databaseSchemaService: DatabaseSchemaService) {
        this.databaseSchemaService.setOfflineDB();
    }

}
