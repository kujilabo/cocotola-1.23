import "./App.css";
import { AppRoutes } from "@/route";
import { BrowserRouter } from "react-router";

function App() {
  return (
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  );
}

export default App;
