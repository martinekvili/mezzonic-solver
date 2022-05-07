import { Box, Button, Stack } from "@mui/material";
import { modify, reset, selectStatus } from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import BuildIcon from "@mui/icons-material/Build";
import CheckIcon from "@mui/icons-material/Check";
import RestartAltIcon from "@mui/icons-material/RestartAlt";
import { SolveButton } from "./SolveButton";

export function Actions() {
  const status = useAppSelector(selectStatus);
  const dispatch = useAppDispatch();

  const showReset = status !== "solution" && status !== "done";
  const showSolve = status === "setup" || status === "loading";
  const showModify = status === "nosolution";
  const showDone = status === "solution" || status === "done";

  return (
    <Box sx={{ width: "100%" }}>
      <Stack direction="row" justifyContent="space-around">
        {showReset && (
          <Button
            size="large"
            variant="outlined"
            startIcon={<RestartAltIcon />}
            onClick={() => dispatch(reset())}
          >
            Reset
          </Button>
        )}
        {showSolve && <SolveButton />}
        {showModify && (
          <Button
            size="large"
            variant="outlined"
            color="secondary"
            startIcon={<BuildIcon />}
            onClick={() => dispatch(modify())}
          >
            Modify
          </Button>
        )}
        {showDone && (
          <Button
            size="large"
            variant={status === "done" ? "contained" : "outlined"}
            color="success"
            startIcon={<CheckIcon />}
            onClick={() => dispatch(reset())}
          >
            Done
          </Button>
        )}
      </Stack>
    </Box>
  );
}
