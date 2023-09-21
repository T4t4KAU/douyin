namespace go relation

include "common.thrift"

struct RelationActionRequest {
    1: required i64 current_user_id  // 用户鉴权token
    2: required i64 to_user_id       // 对方用户id
    3: required i32 action_type = 3; // 1-关注 2-取消关注
}

struct RelationActionResponse {
    1: required i32 status_code      // 状态码 0-成功 other-失败
    2: required string status_msg    // 状态描述
}

struct RelationFollowListRequest {
    1: required i64 user_id
    2: required i64 current_user_id
}

struct RelationFollowListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.User> user_list   // 用户列表
}

struct RelationFollowerListRequest {
    1: required i64 user_id
    2: required i64 current_user_id
}

struct RelationFollowerListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.User> user_list
}

struct RelationFriendListRequest {
    1: required i64 user_id;
    2: required i64 current_user_id;
}

struct RelationFriendListResponse {
    1: i32 status_code
    2: required string status_msg
    3: required list<common.User> user_list   // 用户列表
}

struct FriendUser {
    1: optional string message
    2: required i64 msgType
}

struct RelationCountRequest {
    1: required i64 user_id
}

struct RelationCountResponse {
    1: required i64 follow_count
    2: required i64 follower_count
}

struct RelationExistRequest {
    1: required i64 current_user_id
    2: required i64 user_id
}

struct RelationExistResponse {
    1: required bool follow_exist
    2: required bool followed_exist
}


service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req)
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req)
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest req)
    RelationCountResponse RelationCount(1: RelationCountRequest req)
    RelationExistResponse RelationExist(1: RelationExistRequest req)
}