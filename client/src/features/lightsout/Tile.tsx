import { ToggleButton, useMediaQuery, useTheme } from "@mui/material";
import { clickTile, selectStatus, selectTile } from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import FlagTwoToneIcon from "@mui/icons-material/FlagTwoTone";

interface TileProps {
  index: number;
}

export function Tile({ index }: TileProps) {
  const status = useAppSelector(selectStatus);
  const { lit, partOfSolution } = useAppSelector(selectTile(index));
  const dispatch = useAppDispatch();

  const [buttonSize, iconFontSize] = useResponsiveSizes();

  return (
    <ToggleButton
      value={index}
      size={buttonSize}
      fullWidth
      selected={lit}
      disabled={status !== "setup" && status !== "solution"}
      onChange={() => dispatch(clickTile(index))}
    >
      <FlagTwoToneIcon
        color="secondary"
        fontSize={iconFontSize}
        sx={{
          visibility:
            status === "solution" && partOfSolution ? "visible" : "hidden",
        }}
      />
    </ToggleButton>
  );
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
