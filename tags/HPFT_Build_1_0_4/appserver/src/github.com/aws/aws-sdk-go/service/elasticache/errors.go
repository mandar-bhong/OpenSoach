// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

const (

	// ErrCodeAPICallRateForCustomerExceededFault for service response error code
	// "APICallRateForCustomerExceeded".
	//
	// The customer has exceeded the allowed rate of API calls.
	ErrCodeAPICallRateForCustomerExceededFault = "APICallRateForCustomerExceeded"

	// ErrCodeAuthorizationAlreadyExistsFault for service response error code
	// "AuthorizationAlreadyExists".
	//
	// The specified Amazon EC2 security group is already authorized for the specified
	// cache security group.
	ErrCodeAuthorizationAlreadyExistsFault = "AuthorizationAlreadyExists"

	// ErrCodeAuthorizationNotFoundFault for service response error code
	// "AuthorizationNotFound".
	//
	// The specified Amazon EC2 security group is not authorized for the specified
	// cache security group.
	ErrCodeAuthorizationNotFoundFault = "AuthorizationNotFound"

	// ErrCodeCacheClusterAlreadyExistsFault for service response error code
	// "CacheClusterAlreadyExists".
	//
	// You already have a cluster with the given identifier.
	ErrCodeCacheClusterAlreadyExistsFault = "CacheClusterAlreadyExists"

	// ErrCodeCacheClusterNotFoundFault for service response error code
	// "CacheClusterNotFound".
	//
	// The requested cluster ID does not refer to an existing cluster.
	ErrCodeCacheClusterNotFoundFault = "CacheClusterNotFound"

	// ErrCodeCacheParameterGroupAlreadyExistsFault for service response error code
	// "CacheParameterGroupAlreadyExists".
	//
	// A cache parameter group with the requested name already exists.
	ErrCodeCacheParameterGroupAlreadyExistsFault = "CacheParameterGroupAlreadyExists"

	// ErrCodeCacheParameterGroupNotFoundFault for service response error code
	// "CacheParameterGroupNotFound".
	//
	// The requested cache parameter group name does not refer to an existing cache
	// parameter group.
	ErrCodeCacheParameterGroupNotFoundFault = "CacheParameterGroupNotFound"

	// ErrCodeCacheParameterGroupQuotaExceededFault for service response error code
	// "CacheParameterGroupQuotaExceeded".
	//
	// The request cannot be processed because it would exceed the maximum number
	// of cache security groups.
	ErrCodeCacheParameterGroupQuotaExceededFault = "CacheParameterGroupQuotaExceeded"

	// ErrCodeCacheSecurityGroupAlreadyExistsFault for service response error code
	// "CacheSecurityGroupAlreadyExists".
	//
	// A cache security group with the specified name already exists.
	ErrCodeCacheSecurityGroupAlreadyExistsFault = "CacheSecurityGroupAlreadyExists"

	// ErrCodeCacheSecurityGroupNotFoundFault for service response error code
	// "CacheSecurityGroupNotFound".
	//
	// The requested cache security group name does not refer to an existing cache
	// security group.
	ErrCodeCacheSecurityGroupNotFoundFault = "CacheSecurityGroupNotFound"

	// ErrCodeCacheSecurityGroupQuotaExceededFault for service response error code
	// "QuotaExceeded.CacheSecurityGroup".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of cache security groups.
	ErrCodeCacheSecurityGroupQuotaExceededFault = "QuotaExceeded.CacheSecurityGroup"

	// ErrCodeCacheSubnetGroupAlreadyExistsFault for service response error code
	// "CacheSubnetGroupAlreadyExists".
	//
	// The requested cache subnet group name is already in use by an existing cache
	// subnet group.
	ErrCodeCacheSubnetGroupAlreadyExistsFault = "CacheSubnetGroupAlreadyExists"

	// ErrCodeCacheSubnetGroupInUse for service response error code
	// "CacheSubnetGroupInUse".
	//
	// The requested cache subnet group is currently in use.
	ErrCodeCacheSubnetGroupInUse = "CacheSubnetGroupInUse"

	// ErrCodeCacheSubnetGroupNotFoundFault for service response error code
	// "CacheSubnetGroupNotFoundFault".
	//
	// The requested cache subnet group name does not refer to an existing cache
	// subnet group.
	ErrCodeCacheSubnetGroupNotFoundFault = "CacheSubnetGroupNotFoundFault"

	// ErrCodeCacheSubnetGroupQuotaExceededFault for service response error code
	// "CacheSubnetGroupQuotaExceeded".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of cache subnet groups.
	ErrCodeCacheSubnetGroupQuotaExceededFault = "CacheSubnetGroupQuotaExceeded"

	// ErrCodeCacheSubnetQuotaExceededFault for service response error code
	// "CacheSubnetQuotaExceededFault".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of subnets in a cache subnet group.
	ErrCodeCacheSubnetQuotaExceededFault = "CacheSubnetQuotaExceededFault"

	// ErrCodeClusterQuotaForCustomerExceededFault for service response error code
	// "ClusterQuotaForCustomerExceeded".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of clusters per customer.
	ErrCodeClusterQuotaForCustomerExceededFault = "ClusterQuotaForCustomerExceeded"

	// ErrCodeInsufficientCacheClusterCapacityFault for service response error code
	// "InsufficientCacheClusterCapacity".
	//
	// The requested cache node type is not available in the specified Availability
	// Zone.
	ErrCodeInsufficientCacheClusterCapacityFault = "InsufficientCacheClusterCapacity"

	// ErrCodeInvalidARNFault for service response error code
	// "InvalidARN".
	//
	// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
	ErrCodeInvalidARNFault = "InvalidARN"

	// ErrCodeInvalidCacheClusterStateFault for service response error code
	// "InvalidCacheClusterState".
	//
	// The requested cluster is not in the available state.
	ErrCodeInvalidCacheClusterStateFault = "InvalidCacheClusterState"

	// ErrCodeInvalidCacheParameterGroupStateFault for service response error code
	// "InvalidCacheParameterGroupState".
	//
	// The current state of the cache parameter group does not allow the requested
	// operation to occur.
	ErrCodeInvalidCacheParameterGroupStateFault = "InvalidCacheParameterGroupState"

	// ErrCodeInvalidCacheSecurityGroupStateFault for service response error code
	// "InvalidCacheSecurityGroupState".
	//
	// The current state of the cache security group does not allow deletion.
	ErrCodeInvalidCacheSecurityGroupStateFault = "InvalidCacheSecurityGroupState"

	// ErrCodeInvalidParameterCombinationException for service response error code
	// "InvalidParameterCombination".
	//
	// Two or more incompatible parameters were specified.
	ErrCodeInvalidParameterCombinationException = "InvalidParameterCombination"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValue".
	//
	// The value for a parameter is invalid.
	ErrCodeInvalidParameterValueException = "InvalidParameterValue"

	// ErrCodeInvalidReplicationGroupStateFault for service response error code
	// "InvalidReplicationGroupState".
	//
	// The requested replication group is not in the available state.
	ErrCodeInvalidReplicationGroupStateFault = "InvalidReplicationGroupState"

	// ErrCodeInvalidSnapshotStateFault for service response error code
	// "InvalidSnapshotState".
	//
	// The current state of the snapshot does not allow the requested operation
	// to occur.
	ErrCodeInvalidSnapshotStateFault = "InvalidSnapshotState"

	// ErrCodeInvalidSubnet for service response error code
	// "InvalidSubnet".
	//
	// An invalid subnet identifier was specified.
	ErrCodeInvalidSubnet = "InvalidSubnet"

	// ErrCodeInvalidVPCNetworkStateFault for service response error code
	// "InvalidVPCNetworkStateFault".
	//
	// The VPC network is in an invalid state.
	ErrCodeInvalidVPCNetworkStateFault = "InvalidVPCNetworkStateFault"

	// ErrCodeNoOperationFault for service response error code
	// "NoOperationFault".
	//
	// The operation was not performed because no changes were required.
	ErrCodeNoOperationFault = "NoOperationFault"

	// ErrCodeNodeGroupNotFoundFault for service response error code
	// "NodeGroupNotFoundFault".
	//
	// The node group specified by the NodeGroupId parameter could not be found.
	// Please verify that the node group exists and that you spelled the NodeGroupId
	// value correctly.
	ErrCodeNodeGroupNotFoundFault = "NodeGroupNotFoundFault"

	// ErrCodeNodeGroupsPerReplicationGroupQuotaExceededFault for service response error code
	// "NodeGroupsPerReplicationGroupQuotaExceeded".
	//
	// The request cannot be processed because it would exceed the maximum allowed
	// number of node groups (shards) in a single replication group. The default
	// maximum is 15
	ErrCodeNodeGroupsPerReplicationGroupQuotaExceededFault = "NodeGroupsPerReplicationGroupQuotaExceeded"

	// ErrCodeNodeQuotaForClusterExceededFault for service response error code
	// "NodeQuotaForClusterExceeded".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of cache nodes in a single cluster.
	ErrCodeNodeQuotaForClusterExceededFault = "NodeQuotaForClusterExceeded"

	// ErrCodeNodeQuotaForCustomerExceededFault for service response error code
	// "NodeQuotaForCustomerExceeded".
	//
	// The request cannot be processed because it would exceed the allowed number
	// of cache nodes per customer.
	ErrCodeNodeQuotaForCustomerExceededFault = "NodeQuotaForCustomerExceeded"

	// ErrCodeReplicationGroupAlreadyExistsFault for service response error code
	// "ReplicationGroupAlreadyExists".
	//
	// The specified replication group already exists.
	ErrCodeReplicationGroupAlreadyExistsFault = "ReplicationGroupAlreadyExists"

	// ErrCodeReplicationGroupNotFoundFault for service response error code
	// "ReplicationGroupNotFoundFault".
	//
	// The specified replication group does not exist.
	ErrCodeReplicationGroupNotFoundFault = "ReplicationGroupNotFoundFault"

	// ErrCodeReservedCacheNodeAlreadyExistsFault for service response error code
	// "ReservedCacheNodeAlreadyExists".
	//
	// You already have a reservation with the given identifier.
	ErrCodeReservedCacheNodeAlreadyExistsFault = "ReservedCacheNodeAlreadyExists"

	// ErrCodeReservedCacheNodeNotFoundFault for service response error code
	// "ReservedCacheNodeNotFound".
	//
	// The requested reserved cache node was not found.
	ErrCodeReservedCacheNodeNotFoundFault = "ReservedCacheNodeNotFound"

	// ErrCodeReservedCacheNodeQuotaExceededFault for service response error code
	// "ReservedCacheNodeQuotaExceeded".
	//
	// The request cannot be processed because it would exceed the user's cache
	// node quota.
	ErrCodeReservedCacheNodeQuotaExceededFault = "ReservedCacheNodeQuotaExceeded"

	// ErrCodeReservedCacheNodesOfferingNotFoundFault for service response error code
	// "ReservedCacheNodesOfferingNotFound".
	//
	// The requested cache node offering does not exist.
	ErrCodeReservedCacheNodesOfferingNotFoundFault = "ReservedCacheNodesOfferingNotFound"

	// ErrCodeServiceLinkedRoleNotFoundFault for service response error code
	// "ServiceLinkedRoleNotFoundFault".
	//
	// The specified service linked role (SLR) was not found.
	ErrCodeServiceLinkedRoleNotFoundFault = "ServiceLinkedRoleNotFoundFault"

	// ErrCodeSnapshotAlreadyExistsFault for service response error code
	// "SnapshotAlreadyExistsFault".
	//
	// You already have a snapshot with the given name.
	ErrCodeSnapshotAlreadyExistsFault = "SnapshotAlreadyExistsFault"

	// ErrCodeSnapshotFeatureNotSupportedFault for service response error code
	// "SnapshotFeatureNotSupportedFault".
	//
	// You attempted one of the following operations:
	//
	//    * Creating a snapshot of a Redis cluster running on a cache.t1.micro cache
	//    node.
	//
	//    * Creating a snapshot of a cluster that is running Memcached rather than
	//    Redis.
	//
	// Neither of these are supported by ElastiCache.
	ErrCodeSnapshotFeatureNotSupportedFault = "SnapshotFeatureNotSupportedFault"

	// ErrCodeSnapshotNotFoundFault for service response error code
	// "SnapshotNotFoundFault".
	//
	// The requested snapshot name does not refer to an existing snapshot.
	ErrCodeSnapshotNotFoundFault = "SnapshotNotFoundFault"

	// ErrCodeSnapshotQuotaExceededFault for service response error code
	// "SnapshotQuotaExceededFault".
	//
	// The request cannot be processed because it would exceed the maximum number
	// of snapshots.
	ErrCodeSnapshotQuotaExceededFault = "SnapshotQuotaExceededFault"

	// ErrCodeSubnetInUse for service response error code
	// "SubnetInUse".
	//
	// The requested subnet is being used by another cache subnet group.
	ErrCodeSubnetInUse = "SubnetInUse"

	// ErrCodeTagNotFoundFault for service response error code
	// "TagNotFound".
	//
	// The requested tag was not found on this resource.
	ErrCodeTagNotFoundFault = "TagNotFound"

	// ErrCodeTagQuotaPerResourceExceeded for service response error code
	// "TagQuotaPerResourceExceeded".
	//
	// The request cannot be processed because it would cause the resource to have
	// more than the allowed number of tags. The maximum number of tags permitted
	// on a resource is 50.
	ErrCodeTagQuotaPerResourceExceeded = "TagQuotaPerResourceExceeded"

	// ErrCodeTestFailoverNotAvailableFault for service response error code
	// "TestFailoverNotAvailableFault".
	//
	// The TestFailover action is not available.
	ErrCodeTestFailoverNotAvailableFault = "TestFailoverNotAvailableFault"
)
