import { useState } from 'react'
import { Navigate } from 'react-router-dom'
import { useLogin } from '@/hooks/useAuth'
import { useAuthStore } from '@/stores/authStore'

export default function LoginPage() {
  const authed = useAuthStore((s) => s.isAuthenticated())
  const login = useLogin()
  const [username, setUsername] = useState('demo')
  const [password, setPassword] = useState('')

  if (authed) {
    return <Navigate to="/" replace />
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    login.mutate({ username: username.trim(), password })
  }

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <form
        onSubmit={handleSubmit}
        className="w-full max-w-sm rounded-xl border border-slate-200 bg-white p-8 shadow-sm"
      >
        <h1 className="text-xl font-semibold text-slate-900">LIMS Stack Demo</h1>
        <p className="mt-1 text-sm text-slate-500">演示登录 · 默认账号 demo / demo123</p>

        <label className="mt-6 block text-sm font-medium text-slate-700">
          用户名
          <input
            className="mt-1 w-full rounded-lg border border-slate-300 px-3 py-2 text-sm outline-none focus:border-slate-500"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            autoComplete="username"
          />
        </label>

        <label className="mt-4 block text-sm font-medium text-slate-700">
          密码
          <input
            type="password"
            className="mt-1 w-full rounded-lg border border-slate-300 px-3 py-2 text-sm outline-none focus:border-slate-500"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            autoComplete="current-password"
          />
        </label>

        <button
          type="submit"
          disabled={login.isPending}
          className="mt-6 w-full rounded-lg bg-slate-900 px-4 py-2 text-sm font-medium text-white hover:bg-slate-800 disabled:opacity-60"
        >
          {login.isPending ? '登录中…' : '登录'}
        </button>
      </form>
    </div>
  )
}
