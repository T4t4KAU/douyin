# Simple Douyin

## Introduction

A minimalist tiktok, using a distributed microservice architecture based on **Kitex** and **Hertz**

### Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Message Queue
  - use **RabbitMQ** for asynchronous communication and module decoupling
- Memory Cache
  - use **Redis** to cache hot data

- Tracing
  - use **jaeger** to tracing
- Customized BoundHandler
  - achieve CPU utilization rate customized bound handler
- Service Discovery and Register
  - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### catalog introduce

| catalog        | introduce                |
| :------------- | :----------------------- |
| pkg/constants  | constant                 |
| pkg/bound      | customized bound handler |
| pkg/errno      | customized error number  |
| pkg/middleware | RPC middleware           |
| pkg/tracer     | init jaeger              |
| dal            | db operation             |
| pkg            | data pack                |
| service        | business logic           |

### code count

| Language     | files | blank | comment | code  |
| ------------ | ----- | ----- | ------- | ----- |
| Go           | 181   | 7355  | 421     | 53981 |
| Thrift       | 8     | 79    | 0       | 341   |
| Bourne Shell | 12    | 52    | 0       | 198   |
| XML          | 4     | 0     | 0       | 192   |
| YAML         | 7     | 7     | 0       | 86    |
| Markdown     | 1     | 26    | 0       | 75    |
| make         | 2     | 5     | 1       | 10    |
| JSON         | 1     | 0     | 0       | 1     |
| Text         | 1     | 0     | 0       | 1     |
| SUM:         | 217   | 7524  | 422     | 54885 |

## Quick Start

### 1. Setup Basic Dependence

```shell
docker-compose up
```

### 2. Run User RPC Server

```shell
cd cmd/user
sh build.sh
go run output/bin
```

### 3. Run Publish RPC Server

```shell
cd cmd/publish
sh build.sh
go run output/bin
```

### 4. Run Comment RPC Server

```shell
cd cmd/comment
sh build.sh
go run output/bin
```

### 5. Run Favorite RPC Server

```shell
cd cmd/favorite
sh build.sh
go run output/bin
```

### 6. Run Message RPC Server

```shell
cd cmd/message
sh build.sh
go run output/bin
```

### 7. Run Relation RPC Server

```shell
cd cmd/relation
sh build.sh
go run output/bin
```

### 8. Run API Server

```shell
cd cmd/api
chmod +x run.sh
go run api
```

### 9. Jaeger

visit `http://127.0.0.1:16686/` on browser.

## API requests

user register

```powershell
curl --location --request POST '/douyin/user/register/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
  "status_code": 0,
  "status_msg": "string",
  "user_id": 0,
  "token": "string"
}
```

user login

```powershell
curl --location --request POST '/douyin/user/login/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
```

infomation of user

```powershell
curl --location --request GET '/douyin/user/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user": {
        "id": 0,
        "name": "string",
        "follow_count": 0,
        "follower_count": 0,
        "is_follow": true,
        "avatar": "string",
        "background_image": "string",
        "signature": "string",
        "total_favorited": "string",
        "work_count": 0,
        "favorite_count": 0
    }
}
```

get video stream

```powershell
curl --location --request GET '/douyin/feed/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "next_time": 0,
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user selects video to upload

```powershell
curl --location --request POST '/douyin/publish/action/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'data=@""' \
--form 'token=""' \
--form 'title=""'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

list all videos contributed by the user

```powershell
curl --location --request GET '/douyin/publish/list/?token=&user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

all liked videos by user

```powershell
curl --location --request GET '/douyin/favorite/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

view all comments on video

```powershell
curl --location --request GET '/douyin/comment/list/?token=&video_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment_list": [
        {
            "id": 0,
            "user": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "content": "string",
            "create_date": "string"
        }
    ]
}
```

follow

```powershell
curl --location --request POST '/douyin/relation/action/?token=&to_user_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

follow list

```powershell
curl --location --request GET '/douyin/relation/follow/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

follower list

```powershell
curl --location --request GET '/douyin/relation/follower/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

friend list

```powershell
curl --location --request GET '/douyin/relation/friend/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

send message

```powershell
curl --location --request POST '/douyin/message/action/?token=&to_user_id=&action_type=&content=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

message list

```powershell
curl --location --request GET '/douyin/message/chat/?token=&to_user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "message_list": [
        {
            "id": 0,
            "to_user_id": 0,
            "from_user_id": 0,
            "content": "string",
            "create_time": 0
        }
    ]
}
```

## Screenshot

<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image1.png?raw=true" alt="image1.png" style="width:30%; height:auto;">
<br>
<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image2.png?raw=true" alt="image2.png" style="width:30%; height:auto;">

## Give a star! ⭐

If you think this project is interesting, or helpful to you, please give a star!Simple Douyin

## Introduction

A minimalist tiktok, using a distributed microservice architecture based on **Kitex** and **Hertz**

### Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Message Queue
  - use **RabbitMQ** for asynchronous communication and module decoupling
- Memory Cache
  - use **Redis** to cache hot data

- Tracing
  - use **jaeger** to tracing
- Customized BoundHandler
  - achieve CPU utilization rate customized bound handler
- Service Discovery and Register
  - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### catalog introduce

| catalog        | introduce                |
| :------------- | :----------------------- |
| pkg/constants  | constant                 |
| pkg/bound      | customized bound handler |
| pkg/errno      | customized error number  |
| pkg/middleware | RPC middleware           |
| pkg/tracer     | init jaeger              |
| dal            | db operation             |
| pkg            | data pack                |
| service        | business logic           |

### code count

| Language     | files | blank | comment | code  |
| ------------ | ----- | ----- | ------- | ----- |
| Go           | 181   | 7355  | 421     | 53981 |
| Thrift       | 8     | 79    | 0       | 341   |
| Bourne Shell | 12    | 52    | 0       | 198   |
| XML          | 4     | 0     | 0       | 192   |
| YAML         | 7     | 7     | 0       | 86    |
| Markdown     | 1     | 26    | 0       | 75    |
| make         | 2     | 5     | 1       | 10    |
| JSON         | 1     | 0     | 0       | 1     |
| Text         | 1     | 0     | 0       | 1     |
| SUM:         | 217   | 7524  | 422     | 54885 |

## Quick Start

### 1. Setup Basic Dependence

```shell
docker-compose up
```

### 2. Run User RPC Server

```shell
cd cmd/user
sh build.sh
go run output/bin
```

### 3. Run Publish RPC Server

```shell
cd cmd/publish
sh build.sh
go run output/bin
```

### 4. Run Comment RPC Server

```shell
cd cmd/comment
sh build.sh
go run output/bin
```

### 5. Run Favorite RPC Server

```shell
cd cmd/favorite
sh build.sh
go run output/bin
```

### 6. Run Message RPC Server

```shell
cd cmd/message
sh build.sh
go run output/bin
```

### 7. Run Relation RPC Server

```shell
cd cmd/relation
sh build.sh
go run output/bin
```

### 8. Run API Server

```shell
cd cmd/api
chmod +x run.sh
go run api
```

### 9. Jaeger

visit `http://127.0.0.1:16686/` on browser.

## API requests

user register

```powershell
curl --location --request POST '/douyin/user/register/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
```

user login

```powershell
curl --location --request POST '/douyin/user/login/?username=&password=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user_id": 0,
    "token": "string"
}
```

infomation of user

```powershell
curl --location --request GET '/douyin/user/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "user": {
        "id": 0,
        "name": "string",
        "follow_count": 0,
        "follower_count": 0,
        "is_follow": true,
        "avatar": "string",
        "background_image": "string",
        "signature": "string",
        "total_favorited": "string",
        "work_count": 0,
        "favorite_count": 0
    }
}
```

get video stream

```powershell
curl --location --request GET '/douyin/feed/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "next_time": 0,
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user selects video to upload

```powershell
curl --location --request POST '/douyin/publish/action/' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--form 'data=@""' \
--form 'token=""' \
--form 'title=""'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

list all videos contributed by the user

```powershell
curl --location --request GET '/douyin/publish/list/?token=&user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

all liked videos by user

```powershell
curl --location --request GET '/douyin/favorite/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "video_list": [
        {
            "id": 0,
            "author": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "play_url": "string",
            "cover_url": "string",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": true,
            "title": "string"
        }
    ]
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

logged in user to comment on video

```powershell
curl --location --request POST '/douyin/comment/action/?token=&video_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment": {
        "id": 0,
        "user": {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        },
        "content": "string",
        "create_date": "string"
    }
}
```

view all comments on video

```powershell
curl --location --request GET '/douyin/comment/list/?token=&video_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string",
    "comment_list": [
        {
            "id": 0,
            "user": {
                "id": 0,
                "name": "string",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": true,
                "avatar": "string",
                "background_image": "string",
                "signature": "string",
                "total_favorited": "string",
                "work_count": 0,
                "favorite_count": 0
            },
            "content": "string",
            "create_date": "string"
        }
    ]
}
```

follow

```powershell
curl --location --request POST '/douyin/relation/action/?token=&to_user_id=&action_type=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

follow list

```powershell
curl --location --request GET '/douyin/relation/follow/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

follower list

```powershell
curl --location --request GET '/douyin/relation/follower/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

friend list

```powershell
curl --location --request GET '/douyin/relation/friend/list/?user_id=&token=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "user_list": [
        {
            "id": 0,
            "name": "string",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": true,
            "avatar": "string",
            "background_image": "string",
            "signature": "string",
            "total_favorited": "string",
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```

send message

```powershell
curl --location --request POST '/douyin/message/action/?token=&to_user_id=&action_type=&content=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": 0,
    "status_msg": "string"
}
```

message list

```powershell
curl --location --request GET '/douyin/message/chat/?token=&to_user_id=' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'
```

response

```json
{
    "status_code": "string",
    "status_msg": "string",
    "message_list": [
        {
            "id": 0,
            "to_user_id": 0,
            "from_user_id": 0,
            "content": "string",
            "create_time": 0
        }
    ]
}
```

## Deploy with docker

### Setup Basic Dependence

```powershell
docker-compose up
```

### Get Default Network Gateway Ip

`docker-compose up` will create a default bridge network for mysql,etcd and jaeger. Get the gateway ip of this default network to reach three components.

```shell
docker inspect douyin
```

### Replace ip in Dockerfile

Example:

```dockerfile
FROM golang:1.20

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

ENV MYSQL_DSN="douyin:123456@tcp(172.21.0.6:3306)/douyin?charset=utf8&parseTime=true"
ENV REDIS_ADDR="your Mysql IP"
ENV ETCD_ADDR="your Etcd IP"
ENV MINIO_ENDPOINT="your Minio IP"
ENV RABBIT_MQ_URI="your RabbitMQ IP"

ENV JAEGER_AGENT_HOST="172.21.0.7"
ENV JAEGER_DISABLED=false
ENV JAEGER_SAMPLER_TYPE="const"
ENV JAEGER_SAMPLER_PARAM=1
ENV JAEGER_REPORTER_LOG_SPANS=true
ENV JAEGER_AGENT_PORT=6831

WORKDIR $GOPATH/src/douyin
COPY .. $GOPATH/src/douyin
WORKDIR $GOPATH/src/douyin/cmd/comment
RUN ["sh", "build.sh"]
EXPOSE 8888
ENTRYPOINT ["./output/bin/commentservice"]
```

### Build images from Dockerfile

Build image of commentservice

```powershell
docker build -t douyin/comment -f cmd/comment/script/Dockerfile .
```

Build image of favoriteservice:

```powershell
docker build -t douyin/favorite -f cmd/favorite/script/Dockerfile .
```

Build image of relationservice:

```powershell
docker build -t douyin/relation -f cmd/relation/script/Dockerfile .
```

Build image of messageservice:

```powershell
docker build -t douyin/message -f cmd/message/script/Dockerfile .
```

Build image of publishservice:

```powershell
docker build -t douyin/publish -f cmd/publish/script/Dockerfile .
```

Build image of userservice:

```powershell
docker build -t douyin/user -f cmd/user/script/Dockerfile .
```

### Run containers

- Create bridge network for these  services.

  ```shell
  docker network create -d bridge douyin
  ```

- Run contains in douyin network

  ```shell
  docker run -d --name favorite --network douyin douyin/favorite
  docker run -d --name relation --network douyin douyin/relation
  docker run -d --name comment --network douyin douyin/comment
  docker run -d --name publish --network douyin douyin/publish
  docker run -d --name message --network douyin douyin/message
  docker run -d --name user --network douyin douyin/user
  ```

## Screenshot

<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image1.png?raw=true" alt="image1.png" style="width:30%; height:auto;">
<br>
<img src="https://github.com/T4t4KAU/douyin/blob/main/image/image2.png?raw=true" alt="image2.png" style="width:30%; height:auto;">

## Give a star! ⭐

If you think this project is interesting, or helpful to you, please give a star!