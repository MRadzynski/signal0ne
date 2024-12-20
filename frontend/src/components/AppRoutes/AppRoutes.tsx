import { Route, Routes } from 'react-router-dom';
import { ROUTES } from '../../data/routes';
import NotFoundPage from '../../pages/NotFoundPage/NotFoundPage';
import ProtectedRoute from '../ProtectedRoute/ProtectedRoute';

const AppRoutes = () => (
  <Routes>
    {ROUTES.map(({ Component, isDisabled, path, redirectTo, unAuthed }) => (
      <Route
        element={
          <ProtectedRoute
            isDisabled={isDisabled}
            redirectTo={redirectTo}
            unAuthed={unAuthed}
          >
            <Component />
          </ProtectedRoute>
        }
        key={path}
        path={path}
      />
    ))}
    <Route element={<NotFoundPage />} path="*" />
  </Routes>
);

export default AppRoutes;
