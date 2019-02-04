export interface IDeviceAuthResult {
    onDeviceAuthSuccess(userid: number);
    onDeviceAuthError(error: any);
    onSubmitDiscarded();
}