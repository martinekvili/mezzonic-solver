import { Box, Button, Stack } from "@mui/material";
import {
  modify,
  reset,
  selectBoard,
  selectStatus,
  solveAsync,
} from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import BuildIcon from "@mui/icons-material/Build";
import CheckIcon from "@mui/icons-material/Check";
import { LoadingButton } from "@mui/lab";
import RestartAltIcon from "@mui/icons-material/RestartAlt";
import RocketLaunchIcon from "@mui/icons-material/RocketLaunch";
import { useMemo } from "react";

export function Actions() {
  const status = useAppSelector(selectStatus);
  const board = useAppSelector(selectBoard);
  const dispatch = useAppDispatch();

  const showReset = status !== "solution" && status !== "done";
  const showSolve = status === "setup" || status === "loading";
  const showModify = status === "nosolution";
  const showDone = status === "solution" || status === "done";

  const solveEnabled = useMemo(() => board.some((tile) => tile), [board]);

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
        {showSolve && (
          <LoadingButton
            size="large"
            variant="contained"
            color="secondary"
            loadingPosition="start"
            startIcon={<RocketLaunchIcon />}
            disabled={!solveEnabled}
            loading={status === "loading"}
            onClick={() => dispatch(solveAsync())}
          >
            Solve
          </LoadingButton>
        )}
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
