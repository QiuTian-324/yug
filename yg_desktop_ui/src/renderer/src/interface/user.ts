export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  user_id: string;
  extra: {
    token: string;
  }
}

export interface UserInfoResponse {
  user_id: string;
  username: string;
  nickname: string;
  avatar: string;
  email: string;
  phone: string;
  avatar_url: string;
  bio: string;
  online: number;
  status: number;
  last_login_at: string;
}

export interface AddFriendRequest {
  username: string;
}

export interface FriendListResponse {
  user_id: number;
  username: string;
  nickname: string;
  email: string;
  phone: string;
  avatar_url: string;
  bio: string;
  online: number;
}

export interface UserInfoRequest {
  user_id: string;
}

export interface UserInfoResponse {
  user_id: string;
  username: string;
  nickname: string;
  avatar: string;
  email: string;
  phone: string;
  avatar_url: string;
  bio: string;
  online: number;
  status: number;
  last_login_at: string;
}
