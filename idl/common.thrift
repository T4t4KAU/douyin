namespace go common

struct User {
  1: required i64 id; // user id
  2: required string name; // user name
  3: optional i64 follow_count; // total number of people the user follows
  4: optional i64 follower_count; // total number of fans
  5: required bool is_follow; // whether the currently logged-in user follows this user
  6: optional string avatar; // user avatar URL
  7: optional string background_image; // image at the top of the user's personal page
  8: optional string signature; // user profile
  9: optional i64 total_favorited; // number of likes for videos published by user
  10: optional i64 work_count; // number of videos published by user
  11: optional i64 favorite_count; // number of likes by this user
}

struct Video {
  1: required i64 id, // video id
  2: required User author, // author information
  3: required string play_url, // video playback URL
  4: required string cover_url, // video cover URL
  5: required i64 favorite_count, // total number of likes for the video
  6: required i64 comment_count, // total number of comments on the video
  7: required bool is_favorite, // true-Liked, false-did not like
  8: required string title // video title
}