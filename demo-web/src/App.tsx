import { Navigate, Route, Routes } from 'react-router-dom'
import { useAuthStore } from '@/stores/authStore'
import LoginPage from '@/pages/Login'
import NotesPage from '@/pages/Notes'

function RequireAuth({ children }: { children: React.ReactNode }) {
  const authed = useAuthStore((s) => s.isAuthenticated())
  if (!authed) {
    return <Navigate to="/login" replace />
  }
  return <>{children}</>
}

export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<LoginPage />} />
      <Route
        path="/"
        element={
          <RequireAuth>
            <NotesPage />
          </RequireAuth>
        }
      />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}
