import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { LightsOut } from "./features/lightsout/LightsOut";

function App() {
  const darkTheme = createTheme({
    palette: {
      mode: "dark",
    },
  });

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <LightsOut />
    </ThemeProvider>
  );
}

export default App;
