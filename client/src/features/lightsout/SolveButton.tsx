import { selectBoard, selectStatus, solveAsync } from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import { LoadingButton } from "@mui/lab";
import RocketLaunchIcon from "@mui/icons-material/RocketLaunch";
import { useMemo } from "react";

export function SolveButton() {
  const status = useAppSelector(selectStatus);
  const board = useAppSelector(selectBoard);
  const dispatch = useAppDispatch();

  const enabled = useMemo(() => board.some((tile) => tile), [board]);

  return (
    <LoadingButton
      size="large"
      variant="contained"
      color="secondary"
      loadingPosition="start"
      startIcon={<RocketLaunchIcon />}
      disabled={!enabled}
      loading={status === "loading"}
      onClick={() => dispatch(solveAsync())}
    >
      Solve
    </LoadingButton>
  );
}
