import { Table } from "@chakra-ui/react";
import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

type State = {
  tab: string;
};
type Action = {
  selectTab: (tab: string) => void;
  getCurrentTab: () => string;
};
export const useNavStore = create<State & Action>()(
  devtools((set) => ({
    tab: "dashboard",
    selectTab: (tab: string): void => {
      set({ tab: tab });
    },
    getCurrentTab: (): string => {
      return "";
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
