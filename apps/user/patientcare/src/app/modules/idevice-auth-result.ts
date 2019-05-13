export interface IDeviceAuthResult {
    onDeviceAuthSuccess(userid: number): void;
    onDeviceAuthError(error: any): void;
    onSubmitDiscarded(): void;
}