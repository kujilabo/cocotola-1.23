import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

type State = {
  sentences: string[];
  error: string | null;
};
type Action = {
  getSentencds: () => Promise<string[]>;
};
export const useTodoStore = create<State & Action>()(
  devtools((set) => ({
    sentences: [],
    getSentencds: async (): Promise<string[]> => {
      return ["a", "b", "c"];
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
