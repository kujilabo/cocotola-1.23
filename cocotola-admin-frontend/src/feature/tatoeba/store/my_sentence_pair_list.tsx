import type { TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";
import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import { immer } from "zustand/middleware/immer";

type State = {
  //   sentencePairs: Map<string, TatoebaSentencePair>;
  sentencePairs: { [key: string]: TatoebaSentencePair };
};
type Action = {
  addSentencePair: (
    sentenceKey: string,
    sentencePair: TatoebaSentencePair,
  ) => void;
  removeSentencePair: (sentenceKey: string) => void;
};

export const useMySentencePairListStore = create<State & Action>()(
  // immer(
  devtools(
    persist(
      (set, get) => ({
        sentencePairs: {},
        addSentencePair: (
          sentenceKey: string,
          sentencePair: TatoebaSentencePair,
        ): void => {
          set((state) => {
            // state.sentencePairs[sentenceKey] = sentencePair;
            return {
              sentencePairs: {
                ...state.sentencePairs,
                [sentenceKey]: sentencePair,
              },
            };
          });
        },
        removeSentencePair: (sentenceKey: string): void => {
          set((state) => {
            delete state.sentencePairs[sentenceKey];
            return {
              sentencePairs: { ...state.sentencePairs },
            };
          });
        },
      }),
      {
        name: "my-sentence-pairs",

        partialize: (state) => ({
          sentencePairs: state.sentencePairs,
        }),
      },
    ),
  ),
  //   ),
);
