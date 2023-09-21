namespace go favorite

include "common.thrift"

struct FavoriteActionRequest {
    1: required i64 user_id
    2: required i64 video_id
    3: required i32 action_type
}

struct FavoriteActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct FavoriteVideoListRequest {
    1: required i64 user_id
    2: required i64 to_user_id
}

struct FavoriteVideoListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.Video> video_list
}

struct FavoriteCountRequest {
    1: required i64 user_id
}

struct FavoriteCountResponse {
    1: required i64 favorite_count
    2: required i64 favorited_count
}

struct FavoriteExistRequest {
    1: required i64 user_id
    2: required i64 video_id
}

struct FavoriteExistResponse {
    1: required bool exist
}

struct FavoriteCountOfVideoRequest {
    1: required i64 user_id
    2: required i64 video_id
}

struct FavoriteCountOfVideoResponse {
    1: required i64 favorited_count
    3: required bool is_favorite
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteCountResponse FavoriteCount(1: FavoriteCountRequest req)
    FavoriteExistResponse FavoriteExist(1: FavoriteExistRequest req)
    FavoriteCountOfVideoResponse FavoriteCountOfVideo(1: FavoriteCountOfVideoRequest req)
    FavoriteVideoListResponse FavoriteVideoList(1: FavoriteVideoListRequest req)
}