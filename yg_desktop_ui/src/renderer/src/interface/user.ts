export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface UserInfoResponse {
  username: string;
  avatar: string;
  email: string;
  phone: string;
  sex: string;
  birthday: string;
  signature: string;
  address: string;
  friends: number;
  groups: number;
  messages: number;
  groups_list: Array<any>;
  friends_list: Array<any>;
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