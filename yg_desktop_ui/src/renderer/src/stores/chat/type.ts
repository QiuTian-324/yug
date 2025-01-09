export interface SessionListResponse {
  list: SessionData[]
}

export interface SessionData {
  user_id: string
  friend_id: string
  nickname: string
  avatar: string
  unread_num: number
  last_msg: string
  last_msg_at: string
}

export interface ChatRecordResponse {
  list: ChatRecordData[]
}

export interface ChatRecordData {
  id: number;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
  is_deleted: boolean;
  sender_id: string;
  receiver_id: string;
  group_id: string;
  content: string;
  content_type: number;
  is_read: number;
  url: string;
  file_name: string;
  offline_message: boolean;
}
