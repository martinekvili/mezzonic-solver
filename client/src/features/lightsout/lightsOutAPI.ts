import axios from "axios";
import { boardSize } from "./constants";

export interface Solution {
  hasSolution: boolean;
  solution?: number[];
}

export async function solve(board: boolean[]) {
  const response = await axios.get<Solution>(
    `solutions/${getParamForBoard(board)}`,
    { baseURL: process.env.REACT_APP_API_BASE_URL, timeout: 30000 }
  );

  return response.data;
}

function getParamForBoard(board: boolean[]) {
  var result = 0;
  var bitMask = 1;

  for (var i = 0; i < boardSize; i++) {
    if (board[i]) {
      result |= bitMask;
    }

    bitMask <<= 1;
  }

  return result.toString(32);
}
