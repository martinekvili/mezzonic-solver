import { PayloadAction, createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { Solution, solve } from "./lightsOutAPI";
import { boardSize, columnCount, rowCount } from "./constants";
import { RootState } from "../../app/store";
import { showErrorMessage } from "../errorhandling/errorHandlingSlice";

export type LightsOutStatus =
  | "setup"
  | "loading"
  | "solution"
  | "nosolution"
  | "done";

export interface LightsOutState {
  status: LightsOutStatus;
  board: boolean[];
  solution?: boolean[];
}

const initialState: LightsOutState = {
  status: "setup",
  board: new Array<boolean>(boardSize).fill(false),
};

export const solveAsync = createAsyncThunk<
  Solution,
  void,
  { state: RootState }
>("lightsOut/solve", async (_, { dispatch, getState }) => {
  try {
    const board = selectBoard(getState());
    return await solve(board);
  } catch (error) {
    dispatch(showErrorMessage());
    throw error;
  }
});

export const lightsOutSlice = createSlice({
  name: "lightsOut",
  initialState,
  reducers: {
    reset: (state) => {
      state.status = "setup";
      state.board = initialState.board;
      state.solution = undefined;
    },
    modify: (state) => {
      state.status = "setup";
    },
    clickTile: (state, action: PayloadAction<number>) => {
      if (state.status === "setup") {
        state.board[action.payload] = !state.board[action.payload];
      } else if (state.status === "solution" && state.solution) {
        state.board[action.payload] = !state.board[action.payload];
        if (action.payload > columnCount - 1) {
          state.board[action.payload - columnCount] =
            !state.board[action.payload - columnCount];
        }
        if (action.payload < (rowCount - 1) * columnCount) {
          state.board[action.payload + columnCount] =
            !state.board[action.payload + columnCount];
        }
        if (action.payload % columnCount > 0) {
          state.board[action.payload - 1] = !state.board[action.payload - 1];
        }
        if (action.payload % columnCount < columnCount - 1) {
          state.board[action.payload + 1] = !state.board[action.payload + 1];
        }

        state.solution[action.payload] = !state.solution[action.payload];

        if (state.board.every((tile) => !tile)) {
          state.status = "done";
        }
      }
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(solveAsync.pending, (state) => {
        state.status = "loading";
      })
      .addCase(solveAsync.fulfilled, (state, action) => {
        if (!action.payload.hasSolution || !action.payload.solution) {
          state.status = "nosolution";
        } else {
          state.status = "solution";

          const solution = new Array<boolean>(boardSize).fill(false);
          for (let index of action.payload.solution) {
            solution[index] = true;
          }
          state.solution = solution;
        }
      })
      .addCase(solveAsync.rejected, (state) => {
        state.status = "setup";
      });
  },
});

export const { reset, modify, clickTile } = lightsOutSlice.actions;

export const selectStatus = (state: RootState) => state.lightsOut.status;
export const selectBoard = (state: RootState) => state.lightsOut.board;
export const selectTile = (index: number) => (state: RootState) => ({
  lit: state.lightsOut.board[index],
  partOfSolution: state.lightsOut.solution
    ? state.lightsOut.solution[index]
    : undefined,
});

export default lightsOutSlice.reducer;
