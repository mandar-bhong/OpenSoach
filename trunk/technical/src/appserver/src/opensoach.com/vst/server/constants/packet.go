package constants

import (
	pccont "opensoach.com/prodcore/constants"
)

const DEVICE_CMD_CAT_DEVICE_REG int = 1
const DEVICE_CMD_DEVICE_REGISTRATION int = 1

const DEVICE_CMD_CAT_CONFIG int = 2
const DEVICE_CMD_CONFIG_DEVICE_SYNC_COMPLETED int = 1
const DEVICE_CMD_CONFIG_DEVICE_SERVICE_POINTS int = 2
const DEVICE_CMD_CONFIG_SERVICE_POINTS_AUTH_CODE int = 5
const DEVCIE_CMD_CONFIG_SERVICE_POINTS_SERV_CONF int = 7
const DEVCIE_CMD_CONFIG_SERVER_SYNC_COMPLETED int = 10
const DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_ASSOCIATED int = 8
const DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_DEASSOCIATED int = 9
const DEVICE_CMD_CONFIG_SERVICE_POINTS_FIELD_OPERATOR_ADDED int = 11
const DEVICE_CMD_CONFIG_SERVICE_POINTS_TOKEN int = 12
const DEVICE_CMD_CONFIG_DEVICE_TOKEN_LIST int = 13

const DEVICE_CMD_CAT_DATA int = 3
const DEVICE_CMD_SERVICE_INST_DATA int = 1
const DEVICE_CMD_CONFIG_PART_DATA int = 6

const DEVICE_CMD_COMPLAINT_DATA int = 256
const DEVICE_CMD_FEEDBACK_DATA int = 257
const DEVICE_CMD_DEVICE_STATE_BATTERY_LEVEL_DATA int = 2

const DEVICE_CMD_VEHICLE_TOKEN_DATA int = 300
const DEVICE_CMD_VEHICLE_DETAILS_DATA int = 301
const DEVICE_CMD_TOKEN_GENERATION_DATA int = 302
const DEVICE_CMD_JOB_CREATION_DATA int = 303
const DEVICE_CMD_JOB_EXEC_DATA int = 304
const DEVICE_CMD_TOKEN_GENERATION_CLAIM_DATA int = 305
const DEVICE_CMD_JOB_EXEC_CLAIM_DATA int = 306

const DEVICE_CMD_CAT_ACK int = pccont.DEVICE_CMD_CAT_ACK
const DEVICE_CMD_CAT_ACK_DEFAULT int = pccont.DEVICE_CMD_CAT_ACK_DEFAULT
const DEVICE_CMD_CAT_ACK_CHART_DATA int = 1
