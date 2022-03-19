import {
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Container,
  Stack,
} from "@mui/material";
import { Actions } from "./Actions";
import { Board } from "./Board";
import Grid4x4Icon from "@mui/icons-material/Grid4x4";
import { Information } from "./Information";

export function LightsOut() {
  return (
    <Container maxWidth="sm">
      <Card>
        <CardHeader
          title="Mezzonic Solver"
          titleTypographyProps={{ variant: "h6" }}
          avatar={<Grid4x4Icon fontSize="large" color="secondary" />}
        />
        <CardContent>
          <Stack direction="column" justifyContent="space-between" spacing={4}>
            <Information />
            <Board />
          </Stack>
        </CardContent>
        <CardActions>
          <Actions />
        </CardActions>
      </Card>
    </Container>
  );
}
