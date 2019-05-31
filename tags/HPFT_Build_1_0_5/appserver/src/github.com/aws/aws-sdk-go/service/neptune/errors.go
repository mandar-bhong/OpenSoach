// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

const (

	// ErrCodeAuthorizationNotFoundFault for service response error code
	// "AuthorizationNotFound".
	//
	// Specified CIDRIP or EC2 security group is not authorized for the specified
	// DB security group.
	//
	// Neptune may not also be authorized via IAM to perform necessary actions on
	// your behalf.
	ErrCodeAuthorizationNotFoundFault = "AuthorizationNotFound"

	// ErrCodeCertificateNotFoundFault for service response error code
	// "CertificateNotFound".
	//
	// CertificateIdentifier does not refer to an existing certificate.
	ErrCodeCertificateNotFoundFault = "CertificateNotFound"

	// ErrCodeDBClusterAlreadyExistsFault for service response error code
	// "DBClusterAlreadyExistsFault".
	//
	// User already has a DB cluster with the given identifier.
	ErrCodeDBClusterAlreadyExistsFault = "DBClusterAlreadyExistsFault"

	// ErrCodeDBClusterNotFoundFault for service response error code
	// "DBClusterNotFoundFault".
	//
	// DBClusterIdentifier does not refer to an existing DB cluster.
	ErrCodeDBClusterNotFoundFault = "DBClusterNotFoundFault"

	// ErrCodeDBClusterParameterGroupNotFoundFault for service response error code
	// "DBClusterParameterGroupNotFound".
	//
	// DBClusterParameterGroupName does not refer to an existing DB Cluster parameter
	// group.
	ErrCodeDBClusterParameterGroupNotFoundFault = "DBClusterParameterGroupNotFound"

	// ErrCodeDBClusterQuotaExceededFault for service response error code
	// "DBClusterQuotaExceededFault".
	//
	// User attempted to create a new DB cluster and the user has already reached
	// the maximum allowed DB cluster quota.
	ErrCodeDBClusterQuotaExceededFault = "DBClusterQuotaExceededFault"

	// ErrCodeDBClusterRoleAlreadyExistsFault for service response error code
	// "DBClusterRoleAlreadyExists".
	//
	// The specified IAM role Amazon Resource Name (ARN) is already associated with
	// the specified DB cluster.
	ErrCodeDBClusterRoleAlreadyExistsFault = "DBClusterRoleAlreadyExists"

	// ErrCodeDBClusterRoleNotFoundFault for service response error code
	// "DBClusterRoleNotFound".
	//
	// The specified IAM role Amazon Resource Name (ARN) is not associated with
	// the specified DB cluster.
	ErrCodeDBClusterRoleNotFoundFault = "DBClusterRoleNotFound"

	// ErrCodeDBClusterRoleQuotaExceededFault for service response error code
	// "DBClusterRoleQuotaExceeded".
	//
	// You have exceeded the maximum number of IAM roles that can be associated
	// with the specified DB cluster.
	ErrCodeDBClusterRoleQuotaExceededFault = "DBClusterRoleQuotaExceeded"

	// ErrCodeDBClusterSnapshotAlreadyExistsFault for service response error code
	// "DBClusterSnapshotAlreadyExistsFault".
	//
	// User already has a DB cluster snapshot with the given identifier.
	ErrCodeDBClusterSnapshotAlreadyExistsFault = "DBClusterSnapshotAlreadyExistsFault"

	// ErrCodeDBClusterSnapshotNotFoundFault for service response error code
	// "DBClusterSnapshotNotFoundFault".
	//
	// DBClusterSnapshotIdentifier does not refer to an existing DB cluster snapshot.
	ErrCodeDBClusterSnapshotNotFoundFault = "DBClusterSnapshotNotFoundFault"

	// ErrCodeDBInstanceAlreadyExistsFault for service response error code
	// "DBInstanceAlreadyExists".
	//
	// User already has a DB instance with the given identifier.
	ErrCodeDBInstanceAlreadyExistsFault = "DBInstanceAlreadyExists"

	// ErrCodeDBInstanceNotFoundFault for service response error code
	// "DBInstanceNotFound".
	//
	// DBInstanceIdentifier does not refer to an existing DB instance.
	ErrCodeDBInstanceNotFoundFault = "DBInstanceNotFound"

	// ErrCodeDBParameterGroupAlreadyExistsFault for service response error code
	// "DBParameterGroupAlreadyExists".
	//
	// A DB parameter group with the same name exists.
	ErrCodeDBParameterGroupAlreadyExistsFault = "DBParameterGroupAlreadyExists"

	// ErrCodeDBParameterGroupNotFoundFault for service response error code
	// "DBParameterGroupNotFound".
	//
	// DBParameterGroupName does not refer to an existing DB parameter group.
	ErrCodeDBParameterGroupNotFoundFault = "DBParameterGroupNotFound"

	// ErrCodeDBParameterGroupQuotaExceededFault for service response error code
	// "DBParameterGroupQuotaExceeded".
	//
	// Request would result in user exceeding the allowed number of DB parameter
	// groups.
	ErrCodeDBParameterGroupQuotaExceededFault = "DBParameterGroupQuotaExceeded"

	// ErrCodeDBSecurityGroupNotFoundFault for service response error code
	// "DBSecurityGroupNotFound".
	//
	// DBSecurityGroupName does not refer to an existing DB security group.
	ErrCodeDBSecurityGroupNotFoundFault = "DBSecurityGroupNotFound"

	// ErrCodeDBSnapshotAlreadyExistsFault for service response error code
	// "DBSnapshotAlreadyExists".
	//
	// DBSnapshotIdentifier is already used by an existing snapshot.
	ErrCodeDBSnapshotAlreadyExistsFault = "DBSnapshotAlreadyExists"

	// ErrCodeDBSnapshotNotFoundFault for service response error code
	// "DBSnapshotNotFound".
	//
	// DBSnapshotIdentifier does not refer to an existing DB snapshot.
	ErrCodeDBSnapshotNotFoundFault = "DBSnapshotNotFound"

	// ErrCodeDBSubnetGroupAlreadyExistsFault for service response error code
	// "DBSubnetGroupAlreadyExists".
	//
	// DBSubnetGroupName is already used by an existing DB subnet group.
	ErrCodeDBSubnetGroupAlreadyExistsFault = "DBSubnetGroupAlreadyExists"

	// ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs for service response error code
	// "DBSubnetGroupDoesNotCoverEnoughAZs".
	//
	// Subnets in the DB subnet group should cover at least two Availability Zones
	// unless there is only one Availability Zone.
	ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs = "DBSubnetGroupDoesNotCoverEnoughAZs"

	// ErrCodeDBSubnetGroupNotFoundFault for service response error code
	// "DBSubnetGroupNotFoundFault".
	//
	// DBSubnetGroupName does not refer to an existing DB subnet group.
	ErrCodeDBSubnetGroupNotFoundFault = "DBSubnetGroupNotFoundFault"

	// ErrCodeDBSubnetGroupQuotaExceededFault for service response error code
	// "DBSubnetGroupQuotaExceeded".
	//
	// Request would result in user exceeding the allowed number of DB subnet groups.
	ErrCodeDBSubnetGroupQuotaExceededFault = "DBSubnetGroupQuotaExceeded"

	// ErrCodeDBSubnetQuotaExceededFault for service response error code
	// "DBSubnetQuotaExceededFault".
	//
	// Request would result in user exceeding the allowed number of subnets in a
	// DB subnet groups.
	ErrCodeDBSubnetQuotaExceededFault = "DBSubnetQuotaExceededFault"

	// ErrCodeDBUpgradeDependencyFailureFault for service response error code
	// "DBUpgradeDependencyFailure".
	//
	// The DB upgrade failed because a resource the DB depends on could not be modified.
	ErrCodeDBUpgradeDependencyFailureFault = "DBUpgradeDependencyFailure"

	// ErrCodeDomainNotFoundFault for service response error code
	// "DomainNotFoundFault".
	//
	// Domain does not refer to an existing Active Directory Domain.
	ErrCodeDomainNotFoundFault = "DomainNotFoundFault"

	// ErrCodeEventSubscriptionQuotaExceededFault for service response error code
	// "EventSubscriptionQuotaExceeded".
	//
	// You have exceeded the number of events you can subscribe to.
	ErrCodeEventSubscriptionQuotaExceededFault = "EventSubscriptionQuotaExceeded"

	// ErrCodeInstanceQuotaExceededFault for service response error code
	// "InstanceQuotaExceeded".
	//
	// Request would result in user exceeding the allowed number of DB instances.
	ErrCodeInstanceQuotaExceededFault = "InstanceQuotaExceeded"

	// ErrCodeInsufficientDBClusterCapacityFault for service response error code
	// "InsufficientDBClusterCapacityFault".
	//
	// The DB cluster does not have enough capacity for the current operation.
	ErrCodeInsufficientDBClusterCapacityFault = "InsufficientDBClusterCapacityFault"

	// ErrCodeInsufficientDBInstanceCapacityFault for service response error code
	// "InsufficientDBInstanceCapacity".
	//
	// Specified DB instance class is not available in the specified Availability
	// Zone.
	ErrCodeInsufficientDBInstanceCapacityFault = "InsufficientDBInstanceCapacity"

	// ErrCodeInsufficientStorageClusterCapacityFault for service response error code
	// "InsufficientStorageClusterCapacity".
	//
	// There is insufficient storage available for the current action. You may be
	// able to resolve this error by updating your subnet group to use different
	// Availability Zones that have more storage available.
	ErrCodeInsufficientStorageClusterCapacityFault = "InsufficientStorageClusterCapacity"

	// ErrCodeInvalidDBClusterSnapshotStateFault for service response error code
	// "InvalidDBClusterSnapshotStateFault".
	//
	// The supplied value is not a valid DB cluster snapshot state.
	ErrCodeInvalidDBClusterSnapshotStateFault = "InvalidDBClusterSnapshotStateFault"

	// ErrCodeInvalidDBClusterStateFault for service response error code
	// "InvalidDBClusterStateFault".
	//
	// The DB cluster is not in a valid state.
	ErrCodeInvalidDBClusterStateFault = "InvalidDBClusterStateFault"

	// ErrCodeInvalidDBInstanceStateFault for service response error code
	// "InvalidDBInstanceState".
	//
	// The specified DB instance is not in the available state.
	ErrCodeInvalidDBInstanceStateFault = "InvalidDBInstanceState"

	// ErrCodeInvalidDBParameterGroupStateFault for service response error code
	// "InvalidDBParameterGroupState".
	//
	// The DB parameter group is in use or is in an invalid state. If you are attempting
	// to delete the parameter group, you cannot delete it when the parameter group
	// is in this state.
	ErrCodeInvalidDBParameterGroupStateFault = "InvalidDBParameterGroupState"

	// ErrCodeInvalidDBSecurityGroupStateFault for service response error code
	// "InvalidDBSecurityGroupState".
	//
	// The state of the DB security group does not allow deletion.
	ErrCodeInvalidDBSecurityGroupStateFault = "InvalidDBSecurityGroupState"

	// ErrCodeInvalidDBSnapshotStateFault for service response error code
	// "InvalidDBSnapshotState".
	//
	// The state of the DB snapshot does not allow deletion.
	ErrCodeInvalidDBSnapshotStateFault = "InvalidDBSnapshotState"

	// ErrCodeInvalidDBSubnetGroupStateFault for service response error code
	// "InvalidDBSubnetGroupStateFault".
	//
	// The DB subnet group cannot be deleted because it is in use.
	ErrCodeInvalidDBSubnetGroupStateFault = "InvalidDBSubnetGroupStateFault"

	// ErrCodeInvalidDBSubnetStateFault for service response error code
	// "InvalidDBSubnetStateFault".
	//
	// The DB subnet is not in the available state.
	ErrCodeInvalidDBSubnetStateFault = "InvalidDBSubnetStateFault"

	// ErrCodeInvalidEventSubscriptionStateFault for service response error code
	// "InvalidEventSubscriptionState".
	//
	// The event subscription is in an invalid state.
	ErrCodeInvalidEventSubscriptionStateFault = "InvalidEventSubscriptionState"

	// ErrCodeInvalidRestoreFault for service response error code
	// "InvalidRestoreFault".
	//
	// Cannot restore from vpc backup to non-vpc DB instance.
	ErrCodeInvalidRestoreFault = "InvalidRestoreFault"

	// ErrCodeInvalidSubnet for service response error code
	// "InvalidSubnet".
	//
	// The requested subnet is invalid, or multiple subnets were requested that
	// are not all in a common VPC.
	ErrCodeInvalidSubnet = "InvalidSubnet"

	// ErrCodeInvalidVPCNetworkStateFault for service response error code
	// "InvalidVPCNetworkStateFault".
	//
	// DB subnet group does not cover all Availability Zones after it is created
	// because users' change.
	ErrCodeInvalidVPCNetworkStateFault = "InvalidVPCNetworkStateFault"

	// ErrCodeKMSKeyNotAccessibleFault for service response error code
	// "KMSKeyNotAccessibleFault".
	//
	// Error accessing KMS key.
	ErrCodeKMSKeyNotAccessibleFault = "KMSKeyNotAccessibleFault"

	// ErrCodeOptionGroupNotFoundFault for service response error code
	// "OptionGroupNotFoundFault".
	//
	// The designated option group could not be found.
	ErrCodeOptionGroupNotFoundFault = "OptionGroupNotFoundFault"

	// ErrCodeProvisionedIopsNotAvailableInAZFault for service response error code
	// "ProvisionedIopsNotAvailableInAZFault".
	//
	// Provisioned IOPS not available in the specified Availability Zone.
	ErrCodeProvisionedIopsNotAvailableInAZFault = "ProvisionedIopsNotAvailableInAZFault"

	// ErrCodeResourceNotFoundFault for service response error code
	// "ResourceNotFoundFault".
	//
	// The specified resource ID was not found.
	ErrCodeResourceNotFoundFault = "ResourceNotFoundFault"

	// ErrCodeSNSInvalidTopicFault for service response error code
	// "SNSInvalidTopic".
	//
	// The SNS topic is invalid.
	ErrCodeSNSInvalidTopicFault = "SNSInvalidTopic"

	// ErrCodeSNSNoAuthorizationFault for service response error code
	// "SNSNoAuthorization".
	//
	// There is no SNS authorization.
	ErrCodeSNSNoAuthorizationFault = "SNSNoAuthorization"

	// ErrCodeSNSTopicArnNotFoundFault for service response error code
	// "SNSTopicArnNotFound".
	//
	// The ARN of the SNS topic could not be found.
	ErrCodeSNSTopicArnNotFoundFault = "SNSTopicArnNotFound"

	// ErrCodeSharedSnapshotQuotaExceededFault for service response error code
	// "SharedSnapshotQuotaExceeded".
	//
	// You have exceeded the maximum number of accounts that you can share a manual
	// DB snapshot with.
	ErrCodeSharedSnapshotQuotaExceededFault = "SharedSnapshotQuotaExceeded"

	// ErrCodeSnapshotQuotaExceededFault for service response error code
	// "SnapshotQuotaExceeded".
	//
	// Request would result in user exceeding the allowed number of DB snapshots.
	ErrCodeSnapshotQuotaExceededFault = "SnapshotQuotaExceeded"

	// ErrCodeSourceNotFoundFault for service response error code
	// "SourceNotFound".
	//
	// The source could not be found.
	ErrCodeSourceNotFoundFault = "SourceNotFound"

	// ErrCodeStorageQuotaExceededFault for service response error code
	// "StorageQuotaExceeded".
	//
	// Request would result in user exceeding the allowed amount of storage available
	// across all DB instances.
	ErrCodeStorageQuotaExceededFault = "StorageQuotaExceeded"

	// ErrCodeStorageTypeNotSupportedFault for service response error code
	// "StorageTypeNotSupported".
	//
	// StorageType specified cannot be associated with the DB Instance.
	ErrCodeStorageTypeNotSupportedFault = "StorageTypeNotSupported"

	// ErrCodeSubnetAlreadyInUse for service response error code
	// "SubnetAlreadyInUse".
	//
	// The DB subnet is already in use in the Availability Zone.
	ErrCodeSubnetAlreadyInUse = "SubnetAlreadyInUse"

	// ErrCodeSubscriptionAlreadyExistFault for service response error code
	// "SubscriptionAlreadyExist".
	//
	// This subscription already exists.
	ErrCodeSubscriptionAlreadyExistFault = "SubscriptionAlreadyExist"

	// ErrCodeSubscriptionCategoryNotFoundFault for service response error code
	// "SubscriptionCategoryNotFound".
	//
	// The designated subscription category could not be found.
	ErrCodeSubscriptionCategoryNotFoundFault = "SubscriptionCategoryNotFound"

	// ErrCodeSubscriptionNotFoundFault for service response error code
	// "SubscriptionNotFound".
	//
	// The designated subscription could not be found.
	ErrCodeSubscriptionNotFoundFault = "SubscriptionNotFound"
)
