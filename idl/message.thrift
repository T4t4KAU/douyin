namespace go message

struct MessageListRequest {
    1: required i64 user_id
    2: required i64 to_user_id
    3: required i64 pre_msg_time
}

struct MessageListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: list<Message> message_list
}

struct Message {
    1: required i64 id
    2: required i64 user_id
    3: required i64 to_user_id
    4: required i64 from_user_id
    5: required string content
    6: required string create_time
}

struct MessageActionRequest {
    1: required i64 user_id
    2: required i64 to_user_id
    3: required i32 action_type
    4: required string content
}

struct MessageActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

service MessageService {
    MessageListResponse MessageList(1: MessageListRequest req)
    MessageActionResponse MessageAction(1: MessageActionRequest req)
}
