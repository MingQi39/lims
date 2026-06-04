export interface ApiResponse<T = unknown> {
  code: number
  message: string
  detail?: string
  data?: T
}

export interface PageData<T> {
  list: T[]
  total: number
  page: number
  page_size: number
}

export interface LoginReq {
  username: string
  password: string
}

export interface LoginResp {
  access_token: string
  user_id: number
  user_name: string
}

export interface Note {
  id: number
  user_id: number
  title: string
  content: string
  created_at: string
  updated_at: string
}

export interface CreateNoteReq {
  title: string
  content?: string
}

export interface UpdateNoteReq {
  title: string
  content?: string
  updated_at: string
}
