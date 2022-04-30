import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { ErrorMessage } from "./features/errorhandling/ErrorMessage";
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
      <ErrorMessage />
    </ThemeProvider>
  );
}

export default App;
