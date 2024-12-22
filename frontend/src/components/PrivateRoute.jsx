import React from 'react';
import { useLocation, Navigate, Outlet } from 'react-router-dom';
import { message } from 'antd';
import { useAuth } from '../contexts/AuthContext';

function PrivateRoute({ adminOnly = false }) {
  const { user, loading } = useAuth();
  const location = useLocation();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!user) {
    message.error('Please log in to access this page.');
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  if (adminOnly && !(user.role === 'admin' || user.role === 'super_admin')) {
    message.error('You do not have permission to access this page.');
    return <Navigate to="/user/info" replace />;
  }

  return <Outlet />;
}

export default PrivateRoute;