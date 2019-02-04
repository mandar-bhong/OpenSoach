export interface IDeviceAuthResult {
    onDeviceAuthSuccess();
    onDeviceAuthError(error: any);
    onSubmitDiscarded();
}