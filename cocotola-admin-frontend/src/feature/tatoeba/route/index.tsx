import React, { useEffect } from "react";
import { Route, Routes } from "react-router";

// import { WorkbookView } from '@/features/workbook/components/WorkbookView';
import { useNavStore } from "@/feature/store/nav";
import { SentenceList } from "@/feature/tatoeba/component/SentenceList";
// import { useDispatch } from 'react-redux';
export const TatoebaRoutes = () => {
  const selectTab = useNavStore((state) => state.selectTab);
  // const dispatch = useDispatch();
  // selectTab('projects');
  useEffect(() => {
    selectTab("tatoeba");
  }, [selectTab]);
  return (
    <Routes>
      <Route path="" element={<SentenceList />} />
      {/* <Route path=":_workbookId" element={<WorkbookView />} /> */}
    </Routes>
  );
};
