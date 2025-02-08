import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import { TatoebaSentence ,TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";

type State = {
  sentences: TatoebaSentencePair[];
  error: string | null;
};
type Action = {
  getSentences: () => Promise<void>;
};
export const useSentenceListStore = create<State & Action>()(
  devtools((set) => ({
    sentences: [],
    getSentences: async (): Promise<void> => {
      set({sentences: [
        new  TatoebaSentencePair(
          new TatoebaSentence(2, 2, "en", "I am a student", "tatoeba"),
          new TatoebaSentence(1, 1, "jp", "私は学生です", "tatoeba"),
        ),
      ]});
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
