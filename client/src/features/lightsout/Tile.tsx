import { clickTile, selectStatus, selectTile } from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import FlagTwoToneIcon from "@mui/icons-material/FlagTwoTone";
import { ToggleButton } from "@mui/material";

interface TileProps {
  index: number;
}

export function Tile({ index }: TileProps) {
  const status = useAppSelector(selectStatus);
  const { lit, partOfSolution } = useAppSelector(selectTile(index));
  const dispatch = useAppDispatch();

  return (
    <ToggleButton
      value={index}
      size="large"
      fullWidth
      selected={lit}
      disabled={status !== "setup" && status !== "solution"}
      onChange={() => dispatch(clickTile(index))}
    >
      {status === "solution" && partOfSolution ? (
        <FlagTwoToneIcon color="secondary" fontSize="medium" />
      ) : (
        <>&nbsp;</>
      )}
    </ToggleButton>
  );
}
