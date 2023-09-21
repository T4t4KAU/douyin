namespace go feed

include "common.thrift"

struct FeedRequest {
    1: optional i64 latest_time
    2: optional i64 user_id
}

struct FeedResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<common.Video> video_list
    4: optional i64 next_time
}

service FeedService {
    FeedResponse Feed(1: FeedRequest req)
}

