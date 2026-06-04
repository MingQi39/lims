import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface AuthState {
  accessToken: string | null
  userId: number | null
  userName: string | null
  setSession: (token: string, userId: number, userName: string) => void
  logout: () => void
  isAuthenticated: () => boolean
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      accessToken: null,
      userId: null,
      userName: null,
      setSession: (accessToken, userId, userName) =>
        set({ accessToken, userId, userName }),
      logout: () => set({ accessToken: null, userId: null, userName: null }),
      isAuthenticated: () => Boolean(get().accessToken),
    }),
    { name: 'lims-demo-auth' },
  ),
)
