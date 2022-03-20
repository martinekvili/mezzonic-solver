import { Alert, Slide, SlideProps, Snackbar, Typography } from "@mui/material";
import { closeErrorMessage, selectShowErrorMessage } from "./lightsOutSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";

export function ErrorMessage() {
  const showErrorMessage = useAppSelector(selectShowErrorMessage);
  const dispatch = useAppDispatch();

  const handleClose = (_?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }

    dispatch(closeErrorMessage());
  };

  return (
    <Snackbar
      anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      open={showErrorMessage}
      autoHideDuration={6000}
      onClose={handleClose}
      TransitionComponent={TransitionRight}
    >
      <Alert onClose={handleClose} severity="error" sx={{ width: "100%" }}>
        <Typography variant="body2">
          An error happened while calculating your solution. Please try again
          later!
        </Typography>
      </Alert>
    </Snackbar>
  );
}

type TransitionProps = Omit<SlideProps, "direction">;

function TransitionRight(props: TransitionProps) {
  return <Slide {...props} direction="left" />;
}
