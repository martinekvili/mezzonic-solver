import { columnCount, rowCount } from "./constants";
import { Stack } from "@mui/material";
import { Tile } from "./Tile";

export function Board() {
  return (
    <Stack
      direction="column"
      justifyContent="space-evenly"
      alignItems="stretch"
      spacing={1}
    >
      {Array.from(Array(rowCount)).map((_, row) => (
        <Stack
          key={`row_${row}`}
          direction="row"
          justifyContent="space-evenly"
          alignItems="stretch"
          spacing={1}
        >
          {Array.from(Array(columnCount)).map((_, column) => (
            <Tile
              key={`tile_${row * columnCount + column}`}
              index={row * columnCount + column}
            />
          ))}
        </Stack>
      ))}
    </Stack>
  );
}
