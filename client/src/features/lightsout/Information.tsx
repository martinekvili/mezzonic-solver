import { Alert, AlertTitle, Box, Typography } from "@mui/material";
import { LightsOutStatus, selectStatus } from "./lightsOutSlice";
import FlagTwoToneIcon from "@mui/icons-material/FlagTwoTone";
import { useAppSelector } from "../../app/hooks";

export function Information() {
  const status = useAppSelector(selectStatus);

  const severity = getSeverity(status);
  const title = getTitle(status);
  const information = getInformation(status);

  return (
    <Box sx={{ minHeight: "8em" }}>
      <Alert severity={severity}>
        <AlertTitle>{title}</AlertTitle>
        <Typography variant="body2" align="justify">
          {information}
        </Typography>
      </Alert>
    </Box>
  );
}

function getSeverity(status: LightsOutStatus) {
  switch (status) {
    case "setup":
    case "loading":
      return "info";
    case "nosolution":
      return "warning";
    case "solution":
    case "done":
      return "success";
  }
}

function getTitle(status: LightsOutStatus) {
  switch (status) {
    case "setup":
    case "loading":
      return "Set up your board";
    case "nosolution":
      return "This board has no solution";
    case "solution":
      return "Your solution is ready";
    case "done":
      return "Completed";
  }
}

function getInformation(status: LightsOutStatus) {
  switch (status) {
    case "setup":
    case "loading":
      return (
        <>
          Set up the board as you see it in the game and then click the{" "}
          <strong>Solve</strong> button!
        </>
      );
    case "nosolution":
      return (
        <>
          Click the <strong>Modify</strong> button and make sure you replicated
          the board exactly as it is in the game!
        </>
      );
    case "solution":
      return (
        <>
          Click the tiles marked with{" "}
          <FlagTwoToneIcon color="secondary" fontSize="small" /> to solve the
          puzzle!
        </>
      );
    case "done":
      return "You successfully solved the puzzle. Thanks for using this tool, see you again soon!";
  }
}
