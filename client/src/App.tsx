import { CssBaseline, ThemeProvider, createTheme } from "@mui/material";
import { ErrorMessage } from "./features/errorhandling/ErrorMessage";
import { LightsOut } from "./features/lightsout/LightsOut";

function App() {
  const darkTheme = createTheme({
    palette: {
      mode: "dark",
    },
    breakpoints: {
      values: {
        xs: 0,
        sm: 320,
        md: 380,
        lg: 600,
        xl: 900,
      },
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
