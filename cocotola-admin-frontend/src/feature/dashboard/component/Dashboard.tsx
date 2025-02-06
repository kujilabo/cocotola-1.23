import { ContentLayout } from "@/component/layout";
import { MainLayout } from "@/component/layout";

export const Dashboard = () => {
  return (
    <MainLayout title="Dashboard">
      <h1 className="text-xl mt-2">Welcome</h1>
      <p className="font-medium">In this application you can:</p>
    </MainLayout>
  );
};
