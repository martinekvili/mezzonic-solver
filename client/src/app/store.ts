import { Action, ThunkAction, configureStore } from "@reduxjs/toolkit";
import lightsOutReducer from "../features/lightsout/lightsOutSlice";

export const store = configureStore({
  reducer: {
    lightsOut: lightsOutReducer,
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
