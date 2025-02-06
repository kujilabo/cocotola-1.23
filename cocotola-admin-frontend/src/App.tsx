import { useState } from "react";
import viteLogo from "/vite.svg";
import reactLogo from "./assets/react.svg";
import "./App.css";
import { Provider } from "@/components/ui/provider";
import { AppRoutes } from "@/route";
import {
  Button,
  ChakraProvider,
  createSystem,
  defaultConfig,
  defineConfig,
} from "@chakra-ui/react";
import { BrowserRouter } from "react-router";
const config = defineConfig({
  theme: {
    semanticTokens: {
      colors: {
        brand: {
          900: { value: "#00458e" },
          800: { value: "#0e61ad" },
          700: { value: "#1672bf" },
          600: { value: "#2083d1" },
          500: { value: "#2791de" },
          400: { value: "#46a0e2" },
          300: { value: "#65b0e6" },
          200: { value: "#8fc7ee" },
          100: { value: "#badcf4" },
          50: { value: "#e3f1fa" },
        },
      },
    },
  },
  // globalCss: {
  //   "div": {
  //     color: "red.500",
  //   },
  // },
});

const system = createSystem(defaultConfig, config);
function App() {
  const [count, setCount] = useState(0);

  return (
    <BrowserRouter>
      <ChakraProvider value={system}>
        <AppRoutes />
      </ChakraProvider>
    </BrowserRouter>
  );
}

export default App;
