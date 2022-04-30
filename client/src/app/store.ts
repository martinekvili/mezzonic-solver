import { Action, ThunkAction, configureStore } from "@reduxjs/toolkit";
import errorHandlingReducer from "../features/errorhandling/errorHandlingSlice";
import lightsOutReducer from "../features/lightsout/lightsOutSlice";

export const store = configureStore({
  reducer: {
    lightsOut: lightsOutReducer,
    errorHandling: errorHandlingReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
