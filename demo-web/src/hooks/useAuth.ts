import { useMutation } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { authApi } from '@/api/auth'
import { getApiErrorMessage } from '@/api/client'
import { useAuthStore } from '@/stores/authStore'

export function useLogin() {
  const setSession = useAuthStore((s) => s.setSession)
  const navigate = useNavigate()

  return useMutation({
    mutationFn: authApi.login,
    onSuccess: (data) => {
      if (!data?.access_token) {
        throw new Error('登录响应异常')
      }
      setSession(data.access_token, data.user_id, data.user_name)
      navigate('/', { replace: true })
    },
    onError: (err) => {
      alert(getApiErrorMessage(err))
    },
  })
}
