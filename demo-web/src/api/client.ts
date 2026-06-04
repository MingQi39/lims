import axios, { type AxiosError } from 'axios'
import type { ApiResponse } from './types'
import { useAuthStore } from '@/stores/authStore'

function getBaseURL() {
  const raw = import.meta.env.VITE_API_BASE_URL?.trim()
  return raw || 'http://localhost:8080'
}

const client = axios.create({
  baseURL: getBaseURL(),
  timeout: 30_000,
})

client.interceptors.request.use((config) => {
  const token = useAuthStore.getState().accessToken
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

client.interceptors.response.use(
  (res) => {
    const body = res.data as ApiResponse
    if (body?.code !== 0) {
      return Promise.reject(new Error(body.message || '请求失败'))
    }
    return body.data as never
  },
  (error: AxiosError<ApiResponse>) => {
    const msg =
      error.response?.data?.message ||
      (error.response?.status === 401 ? '请重新登录' : '网络错误')
    if (error.response?.status === 401 && !error.config?.url?.includes('/auth/login')) {
      useAuthStore.getState().logout()
      window.location.href = '/login'
    }
    return Promise.reject(new Error(msg))
  },
)

export function getApiErrorMessage(err: unknown): string {
  if (err instanceof Error) return err.message
  return '请求失败'
}

export default client
