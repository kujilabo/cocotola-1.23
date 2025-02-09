import { Suspense } from "react";

import { Outlet, Route, Routes } from "react-router";

import { DashboardRoutes } from "@/feature/dashboard/route";
import { TatoebaRoutes } from "@/feature/tatoeba/route";

export const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<DashboardRoutes />} />
      <Route path="/tatoeba/*" element={<TatoebaRoutes />} />
      {/* <Route path="/app/auth/*" element={<AuthRoutes />} />
        <Route path="/app" element={<PrivateRoute element={<App />} />}>
          <Route path="/app/dis/*" element={<DiscussionsRoutes />} />
          <Route path="/app/workbook/*" element={<WorkbookRoutes />} />
          <Route path="/app" element={<Dashboard />} />
        </Route> */}
    </Routes>
  );
};
