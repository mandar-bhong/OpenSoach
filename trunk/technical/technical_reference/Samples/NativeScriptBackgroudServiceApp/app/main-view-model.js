"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var observable_1 = require("data/observable");
var observable_array_1 = require("data/observable-array");
var MainViewModel = /** @class */ (function (_super) {
    __extends(MainViewModel, _super);
    function MainViewModel() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(MainViewModel.prototype, "locations", {
        get: function () {
            if (!this._locations) {
                this._locations = new observable_array_1.ObservableArray();
            }
            return this._locations;
        },
        set: function (value) {
            if (this._locations !== value) {
                this._locations = value;
                this.notifyPropertyChange('locations', value);
            }
        },
        enumerable: true,
        configurable: true
    });
    return MainViewModel;
}(observable_1.Observable));
exports.MainViewModel = MainViewModel;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibWFpbi12aWV3LW1vZGVsLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsibWFpbi12aWV3LW1vZGVsLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7O0FBQ0EsOENBQTZDO0FBQzdDLDBEQUF3RDtBQUV4RDtJQUFtQyxpQ0FBVTtJQUE3Qzs7SUFnQkEsQ0FBQztJQWJHLHNCQUFXLG9DQUFTO2FBQXBCO1lBQ0ksRUFBRSxDQUFDLENBQUMsQ0FBQyxJQUFJLENBQUMsVUFBVSxDQUFDLENBQUMsQ0FBQztnQkFDbkIsSUFBSSxDQUFDLFVBQVUsR0FBRyxJQUFJLGtDQUFlLEVBQXdCLENBQUM7WUFDbEUsQ0FBQztZQUNELE1BQU0sQ0FBQyxJQUFJLENBQUMsVUFBVSxDQUFDO1FBQzNCLENBQUM7YUFFRCxVQUFxQixLQUE0QztZQUM3RCxFQUFFLENBQUMsQ0FBQyxJQUFJLENBQUMsVUFBVSxLQUFLLEtBQUssQ0FBQyxDQUFDLENBQUM7Z0JBQzVCLElBQUksQ0FBQyxVQUFVLEdBQUcsS0FBSyxDQUFDO2dCQUN4QixJQUFJLENBQUMsb0JBQW9CLENBQUMsV0FBVyxFQUFFLEtBQUssQ0FBQyxDQUFDO1lBQ2xELENBQUM7UUFDTCxDQUFDOzs7T0FQQTtJQVFMLG9CQUFDO0FBQUQsQ0FBQyxBQWhCRCxDQUFtQyx1QkFBVSxHQWdCNUM7QUFoQlksc0NBQWEiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgKiBhcyBnZW9sb2NhdGlvbiBmcm9tIFwibmF0aXZlc2NyaXB0LWdlb2xvY2F0aW9uXCI7XG5pbXBvcnQgeyBPYnNlcnZhYmxlIH0gZnJvbSBcImRhdGEvb2JzZXJ2YWJsZVwiO1xuaW1wb3J0IHsgT2JzZXJ2YWJsZUFycmF5IH0gZnJvbSBcImRhdGEvb2JzZXJ2YWJsZS1hcnJheVwiO1xuXG5leHBvcnQgY2xhc3MgTWFpblZpZXdNb2RlbCBleHRlbmRzIE9ic2VydmFibGUge1xuICAgIHByaXZhdGUgX2xvY2F0aW9uczogT2JzZXJ2YWJsZUFycmF5PGdlb2xvY2F0aW9uLkxvY2F0aW9uPjtcblxuICAgIHB1YmxpYyBnZXQgbG9jYXRpb25zKCk6IE9ic2VydmFibGVBcnJheTxnZW9sb2NhdGlvbi5Mb2NhdGlvbj4ge1xuICAgICAgICBpZiAoIXRoaXMuX2xvY2F0aW9ucykge1xuICAgICAgICAgICAgdGhpcy5fbG9jYXRpb25zID0gbmV3IE9ic2VydmFibGVBcnJheTxnZW9sb2NhdGlvbi5Mb2NhdGlvbj4oKTtcbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gdGhpcy5fbG9jYXRpb25zO1xuICAgIH1cblxuICAgIHB1YmxpYyBzZXQgbG9jYXRpb25zKHZhbHVlOiBPYnNlcnZhYmxlQXJyYXk8Z2VvbG9jYXRpb24uTG9jYXRpb24+KSB7XG4gICAgICAgIGlmICh0aGlzLl9sb2NhdGlvbnMgIT09IHZhbHVlKSB7XG4gICAgICAgICAgICB0aGlzLl9sb2NhdGlvbnMgPSB2YWx1ZTtcbiAgICAgICAgICAgIHRoaXMubm90aWZ5UHJvcGVydHlDaGFuZ2UoJ2xvY2F0aW9ucycsIHZhbHVlKTtcbiAgICAgICAgfVxuICAgIH1cbn0iXX0=