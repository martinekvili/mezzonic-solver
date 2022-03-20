import {
  AppBar,
  Card,
  CardActions,
  CardContent,
  Container,
  Stack,
  Toolbar,
  Typography,
} from "@mui/material";
import { Actions } from "./Actions";
import { Board } from "./Board";
import { ErrorMessage } from "./ErrorMessage";
import Grid4x4Icon from "@mui/icons-material/Grid4x4";
import { Information } from "./Information";

export function LightsOut() {
  return (
    <>
      <Container maxWidth="sm">
        <AppBar position="static">
          <Toolbar variant="dense">
            <Grid4x4Icon fontSize="large" color="secondary" sx={{ mr: 2 }} />
            <Typography variant="h6">Mezzonic Solver</Typography>
          </Toolbar>
        </AppBar>
        <Card>
          <CardContent>
            <Stack
              direction="column"
              justifyContent="space-between"
              spacing={4}
            >
              <Information />
              <Board />
            </Stack>
          </CardContent>
          <CardActions>
            <Actions />
          </CardActions>
        </Card>
      </Container>
      <ErrorMessage />
    </>
  );
}
