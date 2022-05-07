import { ToggleButton, useMediaQuery, useTheme } from "@mui/material";
import {
  clickTile,
  selectIsTileLit,
  selectIsTilePartOfSolution,
  selectStatus,
} from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import FlagTwoToneIcon from "@mui/icons-material/FlagTwoTone";
import { useMemo } from "react";

interface TileProps {
  index: number;
}

export function Tile({ index }: TileProps) {
  const status = useAppSelector(selectStatus);
  const dispatch = useAppDispatch();

  const [isLit, isPartOfSolution] = useTileState(index);
  const [buttonSize, iconFontSize] = useResponsiveSizes();

  return (
    <ToggleButton
      value={index}
      size={buttonSize}
      fullWidth
      selected={isLit}
      disabled={status !== "setup" && status !== "solution"}
      onChange={() => dispatch(clickTile(index))}
      disableTouchRipple
    >
      <FlagTwoToneIcon
        color="secondary"
        fontSize={iconFontSize}
        sx={{
          visibility:
            status === "solution" && isPartOfSolution ? "visible" : "hidden",
        }}
      />
    </ToggleButton>
  );
}

function useTileState(index: number): [boolean, boolean | undefined] {
  const isTileLitSelector = useMemo(() => selectIsTileLit(index), [index]);
  const isLit = useAppSelector(isTileLitSelector);

  const isTilePartOfSolutionSelector = useMemo(
    () => selectIsTilePartOfSolution(index),
    [index]
  );
  const isPartOfSolution = useAppSelector(isTilePartOfSolutionSelector);

  return [isLit, isPartOfSolution];
}

function useResponsiveSizes(): [
  "small" | "medium" | "large",
  "small" | "medium"
] {
  const theme = useTheme();
  const isSmall = useMediaQuery(theme.breakpoints.down("sm"));
  const isMedium = useMediaQuery(theme.breakpoints.between("sm", "md"));

  if (isSmall) {
    return ["small", "small"];
  } else if (isMedium) {
    return ["medium", "small"];
  } else {
    return ["large", "medium"];
  }
}
