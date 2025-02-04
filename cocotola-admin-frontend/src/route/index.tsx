import { Suspense } from 'react';

import { Routes, Route, Outlet } from 'react-router-dom';

import { Dashboard }from '@/feature/dashboard/component/Dashboard';

export const AppRoutes = () => {
    return (
      <Routes>
        <Route path="/" element={<Dashboard />} />
        {/* <Route path="/app/auth/*" element={<AuthRoutes />} />
        <Route path="/app" element={<PrivateRoute element={<App />} />}>
          <Route path="/app/dis/*" element={<DiscussionsRoutes />} />
          <Route path="/app/workbook/*" element={<WorkbookRoutes />} />
          <Route path="/app" element={<Dashboard />} />
        </Route> */}
      </Routes>
    );
  };