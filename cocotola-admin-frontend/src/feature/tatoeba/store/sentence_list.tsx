import axios from "axios";
import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

import { backendTatoebaUrl } from "@/config/config";
import type { TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";
import { extractErrorMessage } from "@/lib/base";

type State = {
  sentences: TatoebaSentencePair[];
  error: string | null;
};
type Action = {
  getSentences: () => Promise<void>;
};

type SentencePairFindParamter = {
  pageNo: number;
  pageSize: number;
  keyword: string;
  random: boolean;
};

type SentenceFindResponse = {
  totalCount: number;
  results: TatoebaSentencePair[];
};
export const useSentenceListStore = create<State & Action>()(
  devtools((set) => ({
    sentences: [],
    getSentences: async (): Promise<void> => {
      const param: SentencePairFindParamter = {
        pageNo: 1,
        pageSize: 10,
        keyword: "",
        random: false,
      };
      await axios
        .post(`${backendTatoebaUrl}/api/v1/user/sentence_pair/find`, param, {
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
          set({ sentences: data.results });
        })
        .catch((err: Error) => {
          console.log("callback err");
          const errorMessage = extractErrorMessage(err);
          console.log(errorMessage);
          //   arg.postFailureProcess(errorMessage);
          //   return thunkAPI.rejectWithValue(errorMessage);
          set({ error: errorMessage });
          return "";
        });
      // set({sentences: [
      //   new  TatoebaSentencePair(
      //     new TatoebaSentence(2, 2, "en", "I am a student", "tatoeba"),
      //     new TatoebaSentence(1, 1, "jp", "私は学生です", "tatoeba"),
      //   ),
      // ]});
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
