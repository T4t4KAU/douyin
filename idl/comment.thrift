namespace go comment

include "common.thrift"

struct Comment {
    1: required i64 id
    2: required common.User user
    3: optional string content
    4: required string create_date
}

struct CommentActionRequest {
    1: required i64 user_id
    2: required i64 video_id
    3: required i32 action_type
    4: required string comment_text
    5: optional i64 comment_id
}

struct CommentActionResponse {
    1: required i32 status_code
    2: required string status_msg
    3: optional Comment comment
}

struct CommentListRequest {
    1: required i64 user_id
    2: required i64 video_id
}

struct CommentListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Comment> comment_list
}

struct CommentCountRequest {
    1: required i64 video_id
}

struct CommentCountResponse {
    1: required i64 count
}

service CommentService {
    CommentActionResponse CommentAction(CommentActionRequest req)
    CommentListResponse CommentList(CommentListRequest req)
    CommentCountResponse CommentCount(CommentCountRequest req)
}