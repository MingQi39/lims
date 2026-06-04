import client from './client'
import type { LoginReq, LoginResp } from './types'

export const authApi = {
  login: (data: LoginReq) =>
    client.post<never, LoginResp>('/api/v1/auth/login', data),
}
