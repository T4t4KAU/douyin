namespace go publish

include "common.thrift"

struct PublishActionRequest {
    1: required i64 user_id
    2: required binary data
    3: required string title
}

struct PublishActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct PublishListRequest {
    1: required i64 current_user_id
    2: required i64 user_id
}

struct PublishListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.Video> video_list
}

struct PublishVideoListRequest {
    1: required list<i64> video_ids
}

struct PublishVideoListResponse {
    1: required list<common.Video> video_list
}

struct PublishCountRequest {
    1: required i64 user_id
}

struct PublishCountResponse {
    1: required i64 work_count
}

struct PublishExistRequest {
    1: required i64 video_id
}

struct PublishExistResponse {
    1: required bool exist
}

struct PublishInfoRequest {
    1: required i64 current_user_id
    2: required i64 video_id
}

struct PublishInfoResponse {
    1: required common.Video video
}

struct PublishListByLastTimeRequest {
    1: required i64 user_id
    2: required string last_time
}

struct PublishListByLastTimeResponse {
    1: required list<common.Video> video_list
}

struct FeedActionRequest {
    1: optional i64 latest_time
    2: optional i64 user_id
}

struct FeedActionResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.Video> video_list
    4: optional i64 next_time
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishCountResponse PublishCount(1: PublishCountRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    PublishExistResponse PublishExist(1: PublishExistRequest req)
    PublishInfoResponse PublishInfo(1 :PublishInfoRequest req)
    PublishVideoListResponse PublishVideoList(1: PublishVideoListRequest req)
    PublishListByLastTimeResponse PublishListByLastTime(1: PublishListByLastTimeRequest req)
    FeedActionResponse FeedAction(1: FeedActionRequest req)
}