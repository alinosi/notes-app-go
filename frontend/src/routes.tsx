import { createBrowserRouter, Navigate } from 'react-router';
import { ProtectedRoute } from './components/ProtectedRoute';
import { LoginPage } from './pages/LoginPage';
import { RegisterPage } from './pages/RegisterPage';
import { ForgotPasswordPage } from './pages/ForgotPasswordPage';
import { ResetPasswordPage } from './pages/ResetPasswordPage';
import { NotesPage } from './pages/NotesPage';
import { CreateNotePage } from './pages/CreateNotePage';
import { EditNotePage } from './pages/EditNotePage';
import { SettingsPage } from './pages/SettingsPage';
import { NotFoundPage } from './pages/NotFoundPage';
import { useAuth } from './contexts/AuthContext';

// Redirect to notes if already authenticated
const AuthRedirect: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (isAuthenticated) {
    return <Navigate to="/notes" replace />;
  }

  return <>{children}</>;
};

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Navigate to="/notes" replace />,
  },
  {
    path: '/login',
    element: (
      <AuthRedirect>
        <LoginPage />
      </AuthRedirect>
    ),
  },
  {
    path: '/register',
    element: (
      <AuthRedirect>
        <RegisterPage />
      </AuthRedirect>
    ),
  },
  {
    path: '/forgot-password',
    element: (
      <AuthRedirect>
        <ForgotPasswordPage />
      </AuthRedirect>
    ),
  },
  {
    path: '/reset-password',
    element: (
      <AuthRedirect>
        <ResetPasswordPage />
      </AuthRedirect>
    ),
  },
  {
    path: '/notes',
    element: (
      <ProtectedRoute>
        <NotesPage />
      </ProtectedRoute>
    ),
  },
  {
    path: '/notes/new',
    element: (
      <ProtectedRoute>
        <CreateNotePage />
      </ProtectedRoute>
    ),
  },
  {
    path: '/notes/:id/edit',
    element: (
      <ProtectedRoute>
        <EditNotePage />
      </ProtectedRoute>
    ),
  },
  {
    path: '/settings',
    element: (
      <ProtectedRoute>
        <SettingsPage />
      </ProtectedRoute>
    ),
  },
  {
    path: '*',
    element: <NotFoundPage />,
  },
]);
