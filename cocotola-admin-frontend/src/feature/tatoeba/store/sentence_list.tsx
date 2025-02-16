import axios from "axios";
import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

import { backendTatoebaUrl } from "@/config/config";
import type { TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";
import { extractErrorMessage } from "@/lib/base";
export type TatoebaSentenceFindParameter = {
  pageNo: number;
  pageSize: number;
  srcLang2: string;
  dstLang2: string;
  keyword: string;
  random: boolean;
};

type State = {
  sentences: TatoebaSentencePair[];
  totalSentencePairs: number;
  loading: boolean;
  error: string | null;
};
type Action = {
  getSentences: (param: TatoebaSentenceFindParameter) => Promise<void>;

  setSentences: (sentences: TatoebaSentencePair[]) => void;
};

type SentenceFindResponse = {
  totalCount: number;
  results: TatoebaSentencePair[];
};
export const useSentenceListStore = create<State & Action>()(
  devtools((set) => ({
    sentences: [],
    totalSentencePairs: 0,
    getSentences: async (
      param: TatoebaSentenceFindParameter,
    ): Promise<void> => {
      set({ loading: true });
      await axios
        .get(`${backendTatoebaUrl}/api/v1/user/sentence_pair/find`, {
          params: param,
          auth: {
            username: "username",
            password: "password",
          },
        })
        .then((resp) => {
          console.log("callback then");
          const data = resp.data as SentenceFindResponse;
          // for (const sentencePair of data.results) {
          //   console.log(sentencePair);
          // }
          set({
            sentences: data.results,
            totalSentencePairs: data.totalCount,
          });
          set({ loading: false });
        })
        .catch((err: Error) => {
          console.log("callback err");
          const errorMessage = extractErrorMessage(err);
          console.log(errorMessage);
          //   arg.postFailureProcess(errorMessage);
          //   return thunkAPI.rejectWithValue(errorMessage);
          set({ error: errorMessage });
          set({ loading: false });
          return "";
        });
      // set({sentences: [
      //   new  TatoebaSentencePair(
      //     new TatoebaSentence(2, 2, "en", "I am a student", "tatoeba"),
      //     new TatoebaSentence(1, 1, "jp", "私は学生です", "tatoeba"),
      //   ),
      // ]});
    },
    setSentences: (sentences: TatoebaSentencePair[]): void => {
      set({ sentences: sentences });
    },

    // addTodo: (todoText) =>
    //     set((state) => ({
    //     todos: [
    //         ...state.todos,
    //         {
    //         text: todoText,
    //         id: uid(`${todoText}-${state.todos.length}`),
    //         isCompleted: false
    //         }
    //     ]
    //     })),
    // deleteTodo: (todoId) =>
    //     set((state) => ({
    //     todos: state.todos.filter((todo) => todo.id !== todoId)
    //     })),
    // completeTodo: (todoId) =>
    //     set((state) => ({
    //     todos: state.todos.map((todo) => {
    //         if (todo.id === todoId) {
    //         return {
    //             ...todo,
    //             isCompleted: true
    //         };
    //         }

    //         return todo;
    //     })
    //     })),
  })),
);
