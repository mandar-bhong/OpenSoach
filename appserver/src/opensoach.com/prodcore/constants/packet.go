package constants

const DEVICE_CMD_PRE_EXECUTOR string = "PreExecuteProcessor"

const DEVICE_CMD_EP_HANDLER_CONNECTED string = "EndPoint.Connected"
const DEVICE_CMD_EP_HANDLER_DISCONNECTED string = "EndPoint.DisConnected"
const DEVICE_CMD_EP_HANDLER_SEND_PACKET string = "EndPoint.SendPacket"

const DEVICE_CMD_CAT_DEVICE_VALIDATION int = 1
const DEVICE_CMD_DEVICE_AUTH int = 1

const DEVICE_CMD_CAT_ACK int = 6
const DEVICE_CMD_CAT_ACK_DEFAULT int = 0

//EP Command packets
const DEVICE_CMD_CAT_CONFIG int = 2

const DEVICE_CMD_CAT_DATA int = 3

const DEVICE_CMD_CAT_NOTIFICATION int = 5

const DEVICE_CMD_STORE_GET_SYNC int = 50
const DEVICE_CMD_STORE_APPLY_SYNC int = 51

const DEVICE_TYPE_SHARED_DEVICE = 0
const DEVICE_TYPE_USER_DEVICE = 1
const DEVICE_TYPE_NONE = 2

const SHARED_DEVICE_TOKEN_PREFIX = "Dev"
const USER_DEVICE_TOKEN_PREFIX = "DU"
