import { Route, Routes } from "react-router";

import { Dashboard } from "@/feature/dashboard/component/Dashboard";
import { useNavStore } from "@/feature/store/nav";
// import { WorkbookView } from '@/features/workbook/components/WorkbookView';
import { useEffect } from "react";
export const DashboardRoutes = () => {
  const selectTab = useNavStore((state) => state.selectTab);
  //   selectTab('members');
  useEffect(() => {
    selectTab("dashboard");
  }, [selectTab]);
  return (
    <Routes>
      <Route path="*" element={<Dashboard />} />
      {/* <Route path=":_workbookId" element={<WorkbookView />} /> */}
    </Routes>
  );
};
