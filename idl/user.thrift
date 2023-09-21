namespace go user

include "common.thrift"

struct UserRegisterRequest {
    1: required string username
    2: required string password
}

struct UserRegisterResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserLoginRequest {
    1: required string username
    2: required string password

}

struct UserLoginResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserInfoRequest {
    1: required i64 current_user_id
    2: required i64 user_id
}

struct UserInfoResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required common.User user;
}

struct UserExistRequest {
    1: required i64 user_id
}

struct UserExistResponse {
    1: required bool exist
}

service UserService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    UserInfoResponse UserInfo(1: UserInfoRequest req)
    UserExistResponse UserExist(1: UserExistRequest req)
}