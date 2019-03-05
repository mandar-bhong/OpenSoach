import { Component } from "@angular/core";
import { Page } from "ui/page";

@Component({
    moduleId: module.id,
    selector: "startup",
    templateUrl: "startup.component.html"
})
export class StartupComponent {
    constructor(private page: Page) {
        page.actionBarHidden = true;
    }
}