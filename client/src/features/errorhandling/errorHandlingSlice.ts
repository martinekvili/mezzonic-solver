import { RootState } from "../../app/store";
import { createSlice } from "@reduxjs/toolkit";

export interface ErrorHandlingState {
  showErrorMessage: boolean;
}

const initialState: ErrorHandlingState = {
  showErrorMessage: false,
};

export const errorHandlingSlice = createSlice({
  name: "errorHandling",
  initialState,
  reducers: {
    showErrorMessage: (state) => {
      state.showErrorMessage = true;
    },
    closeErrorMessage: (state) => {
      state.showErrorMessage = false;
    },
  },
});

export const { showErrorMessage, closeErrorMessage } =
  errorHandlingSlice.actions;

export const selectShowErrorMessage = (state: RootState) =>
  state.errorHandling.showErrorMessage;

export default errorHandlingSlice.reducer;
