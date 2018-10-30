"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var router_1 = require("nativescript-angular/router");
var DetailsComponent = /** @class */ (function () {
    function DetailsComponent(routerExtensions) {
        this.routerExtensions = routerExtensions;
        this.selectedfirst = true;
        this.selectedsecond = false;
        this.tabSelectedIndex = 0;
    }
    DetailsComponent.prototype.ngOnInit = function () { };
    DetailsComponent.prototype.changeTab = function () {
        if (this.tabSelectedIndex === 0) {
            this.tabSelectedIndex = 1;
        }
        else if (this.tabSelectedIndex === 1) {
            this.tabSelectedIndex = 2;
        }
        else if (this.tabSelectedIndex === 2) {
            this.tabSelectedIndex = 0;
        }
    };
    DetailsComponent.prototype.goBackPage = function () {
        this.routerExtensions.navigate(["/list"], { clearHistory: true });
    };
    DetailsComponent.prototype.firstTab = function () {
        this.selectedfirst = true;
        this.selectedsecond = false;
    };
    DetailsComponent.prototype.secondTab = function () {
        this.selectedfirst = false;
        this.selectedsecond = true;
    };
    DetailsComponent.prototype.test = function () {
        var coll = document.getElementsByClassName("collapsible");
        var i;
        for (i = 0; i < coll.length; i++) {
            coll[i].addEventListener("click", function () {
                this.classList.toggle("active");
                var content = this.nextElementSibling;
                if (content.style.display === "block") {
                    content.style.display = "none";
                }
                else {
                    content.style.display = "block";
                }
            });
        }
    };
    DetailsComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'details',
            templateUrl: './details.component.html',
            styleUrls: ['./details.component.css']
        }),
        __metadata("design:paramtypes", [router_1.RouterExtensions])
    ], DetailsComponent);
    return DetailsComponent;
}());
exports.DetailsComponent = DetailsComponent;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZGV0YWlscy5jb21wb25lbnQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJkZXRhaWxzLmNvbXBvbmVudC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUFBLHNDQUFrRDtBQUNsRCxzREFBK0Q7QUFXL0Q7SUFPQywwQkFBb0IsZ0JBQWtDO1FBQWxDLHFCQUFnQixHQUFoQixnQkFBZ0IsQ0FBa0I7UUFKdEQsa0JBQWEsR0FBRyxJQUFJLENBQUM7UUFDckIsbUJBQWMsR0FBRyxLQUFLLENBQUM7UUFJdEIsSUFBSSxDQUFDLGdCQUFnQixHQUFHLENBQUMsQ0FBQztJQUMzQixDQUFDO0lBRUQsbUNBQVEsR0FBUixjQUFhLENBQUM7SUFDZCxvQ0FBUyxHQUFUO1FBQ0MsRUFBRSxDQUFDLENBQUMsSUFBSSxDQUFDLGdCQUFnQixLQUFLLENBQUMsQ0FBQyxDQUFDLENBQUM7WUFDakMsSUFBSSxDQUFDLGdCQUFnQixHQUFHLENBQUMsQ0FBQztRQUMzQixDQUFDO1FBQUMsSUFBSSxDQUFDLEVBQUUsQ0FBQyxDQUFDLElBQUksQ0FBQyxnQkFBZ0IsS0FBSyxDQUFDLENBQUMsQ0FBQyxDQUFDO1lBQ3hDLElBQUksQ0FBQyxnQkFBZ0IsR0FBRyxDQUFDLENBQUM7UUFDM0IsQ0FBQztRQUFDLElBQUksQ0FBQyxFQUFFLENBQUMsQ0FBQyxJQUFJLENBQUMsZ0JBQWdCLEtBQUssQ0FBQyxDQUFDLENBQUMsQ0FBQztZQUN4QyxJQUFJLENBQUMsZ0JBQWdCLEdBQUcsQ0FBQyxDQUFDO1FBQzNCLENBQUM7SUFDRixDQUFDO0lBQ0QscUNBQVUsR0FBVjtRQUNDLElBQUksQ0FBQyxnQkFBZ0IsQ0FBQyxRQUFRLENBQUMsQ0FBQyxPQUFPLENBQUMsRUFBRSxFQUFFLFlBQVksRUFBRSxJQUFJLEVBQUUsQ0FBQyxDQUFDO0lBQ25FLENBQUM7SUFDRCxtQ0FBUSxHQUFSO1FBQ0MsSUFBSSxDQUFDLGFBQWEsR0FBRyxJQUFJLENBQUM7UUFDMUIsSUFBSSxDQUFDLGNBQWMsR0FBRyxLQUFLLENBQUM7SUFDN0IsQ0FBQztJQUNELG9DQUFTLEdBQVQ7UUFDQyxJQUFJLENBQUMsYUFBYSxHQUFHLEtBQUssQ0FBQztRQUMzQixJQUFJLENBQUMsY0FBYyxHQUFHLElBQUksQ0FBQztJQUM1QixDQUFDO0lBQ0QsK0JBQUksR0FBSjtRQUNDLElBQUksSUFBSSxHQUFHLFFBQVEsQ0FBQyxzQkFBc0IsQ0FBQyxhQUFhLENBQUMsQ0FBQztRQUMxRCxJQUFJLENBQUMsQ0FBQztRQUVOLEdBQUcsQ0FBQyxDQUFDLENBQUMsR0FBRyxDQUFDLEVBQUUsQ0FBQyxHQUFHLElBQUksQ0FBQyxNQUFNLEVBQUUsQ0FBQyxFQUFFLEVBQUUsQ0FBQztZQUNsQyxJQUFJLENBQUMsQ0FBQyxDQUFDLENBQUMsZ0JBQWdCLENBQUMsT0FBTyxFQUFFO2dCQUNqQyxJQUFJLENBQUMsU0FBUyxDQUFDLE1BQU0sQ0FBQyxRQUFRLENBQUMsQ0FBQztnQkFDaEMsSUFBSSxPQUFPLEdBQUcsSUFBSSxDQUFDLGtCQUFrQixDQUFDO2dCQUN0QyxFQUFFLENBQUMsQ0FBQyxPQUFPLENBQUMsS0FBSyxDQUFDLE9BQU8sS0FBSyxPQUFPLENBQUMsQ0FBQyxDQUFDO29CQUN2QyxPQUFPLENBQUMsS0FBSyxDQUFDLE9BQU8sR0FBRyxNQUFNLENBQUM7Z0JBQ2hDLENBQUM7Z0JBQUMsSUFBSSxDQUFDLENBQUM7b0JBQ1AsT0FBTyxDQUFDLEtBQUssQ0FBQyxPQUFPLEdBQUcsT0FBTyxDQUFDO2dCQUNqQyxDQUFDO1lBQ0YsQ0FBQyxDQUFDLENBQUM7UUFDSixDQUFDO0lBRUYsQ0FBQztJQWhEVyxnQkFBZ0I7UUFQNUIsZ0JBQVMsQ0FBQztZQUNWLFFBQVEsRUFBRSxNQUFNLENBQUMsRUFBRTtZQUNuQixRQUFRLEVBQUUsU0FBUztZQUNuQixXQUFXLEVBQUUsMEJBQTBCO1lBQ3ZDLFNBQVMsRUFBRSxDQUFDLHlCQUF5QixDQUFDO1NBQ3RDLENBQUM7eUNBU3FDLHlCQUFnQjtPQVAxQyxnQkFBZ0IsQ0F1RDVCO0lBQUQsdUJBQUM7Q0FBQSxBQXZERCxJQXVEQztBQXZEWSw0Q0FBZ0IiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBDb21wb25lbnQsIE9uSW5pdCB9IGZyb20gJ0Bhbmd1bGFyL2NvcmUnO1xuaW1wb3J0IHsgUm91dGVyRXh0ZW5zaW9ucyB9IGZyb20gXCJuYXRpdmVzY3JpcHQtYW5ndWxhci9yb3V0ZXJcIjtcbmltcG9ydCB7IFNlbGVjdGVkSW5kZXhDaGFuZ2VkRXZlbnREYXRhIH0gZnJvbSBcInRucy1jb3JlLW1vZHVsZXMvdWkvdGFiLXZpZXdcIjtcbmltcG9ydCB7IFNlZ21lbnRlZEJhciwgU2VnbWVudGVkQmFySXRlbSB9IGZyb20gXCJ0bnMtY29yZS1tb2R1bGVzL3VpL3NlZ21lbnRlZC1iYXJcIjtcblxuQENvbXBvbmVudCh7XG5cdG1vZHVsZUlkOiBtb2R1bGUuaWQsXG5cdHNlbGVjdG9yOiAnZGV0YWlscycsXG5cdHRlbXBsYXRlVXJsOiAnLi9kZXRhaWxzLmNvbXBvbmVudC5odG1sJyxcblx0c3R5bGVVcmxzOiBbJy4vZGV0YWlscy5jb21wb25lbnQuY3NzJ11cbn0pXG5cbmV4cG9ydCBjbGFzcyBEZXRhaWxzQ29tcG9uZW50IGltcGxlbWVudHMgT25Jbml0IHtcblx0cHVibGljIHRhYlNlbGVjdGVkSW5kZXg6IG51bWJlcjtcblx0YjE7XG5cdHNlbGVjdGVkZmlyc3QgPSB0cnVlO1xuXHRzZWxlY3RlZHNlY29uZCA9IGZhbHNlO1xuXHRwdWJsaWMgU2VsZWN0ZWRJbmRleDogbnVtYmVyO1xuXG5cdGNvbnN0cnVjdG9yKHByaXZhdGUgcm91dGVyRXh0ZW5zaW9uczogUm91dGVyRXh0ZW5zaW9ucykge1xuXHRcdHRoaXMudGFiU2VsZWN0ZWRJbmRleCA9IDA7XG5cdH1cblxuXHRuZ09uSW5pdCgpIHsgfVxuXHRjaGFuZ2VUYWIoKSB7XG5cdFx0aWYgKHRoaXMudGFiU2VsZWN0ZWRJbmRleCA9PT0gMCkge1xuXHRcdFx0dGhpcy50YWJTZWxlY3RlZEluZGV4ID0gMTtcblx0XHR9IGVsc2UgaWYgKHRoaXMudGFiU2VsZWN0ZWRJbmRleCA9PT0gMSkge1xuXHRcdFx0dGhpcy50YWJTZWxlY3RlZEluZGV4ID0gMjtcblx0XHR9IGVsc2UgaWYgKHRoaXMudGFiU2VsZWN0ZWRJbmRleCA9PT0gMikge1xuXHRcdFx0dGhpcy50YWJTZWxlY3RlZEluZGV4ID0gMDtcblx0XHR9XG5cdH1cblx0Z29CYWNrUGFnZSgpIHtcblx0XHR0aGlzLnJvdXRlckV4dGVuc2lvbnMubmF2aWdhdGUoW1wiL2xpc3RcIl0sIHsgY2xlYXJIaXN0b3J5OiB0cnVlIH0pO1xuXHR9XG5cdGZpcnN0VGFiKCkge1xuXHRcdHRoaXMuc2VsZWN0ZWRmaXJzdCA9IHRydWU7XG5cdFx0dGhpcy5zZWxlY3RlZHNlY29uZCA9IGZhbHNlO1xuXHR9XG5cdHNlY29uZFRhYigpIHtcblx0XHR0aGlzLnNlbGVjdGVkZmlyc3QgPSBmYWxzZTtcblx0XHR0aGlzLnNlbGVjdGVkc2Vjb25kID0gdHJ1ZTtcblx0fVxuXHR0ZXN0KCkge1xuXHRcdHZhciBjb2xsID0gZG9jdW1lbnQuZ2V0RWxlbWVudHNCeUNsYXNzTmFtZShcImNvbGxhcHNpYmxlXCIpO1xuXHRcdHZhciBpO1xuXG5cdFx0Zm9yIChpID0gMDsgaSA8IGNvbGwubGVuZ3RoOyBpKyspIHtcblx0XHRcdGNvbGxbaV0uYWRkRXZlbnRMaXN0ZW5lcihcImNsaWNrXCIsIGZ1bmN0aW9uICgpIHtcblx0XHRcdFx0dGhpcy5jbGFzc0xpc3QudG9nZ2xlKFwiYWN0aXZlXCIpO1xuXHRcdFx0XHR2YXIgY29udGVudCA9IHRoaXMubmV4dEVsZW1lbnRTaWJsaW5nO1xuXHRcdFx0XHRpZiAoY29udGVudC5zdHlsZS5kaXNwbGF5ID09PSBcImJsb2NrXCIpIHtcblx0XHRcdFx0XHRjb250ZW50LnN0eWxlLmRpc3BsYXkgPSBcIm5vbmVcIjtcblx0XHRcdFx0fSBlbHNlIHtcblx0XHRcdFx0XHRjb250ZW50LnN0eWxlLmRpc3BsYXkgPSBcImJsb2NrXCI7XG5cdFx0XHRcdH1cblx0XHRcdH0pO1xuXHRcdH1cblxuXHR9XG5cblxuXG5cblxuXHRcbn0iXX0=