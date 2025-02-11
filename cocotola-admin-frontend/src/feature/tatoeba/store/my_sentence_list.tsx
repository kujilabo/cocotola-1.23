import axios from "axios";
import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

import type { TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";

type State = {
  sentences: { [key: string]: TatoebaSentencePair };
  //Map<string, TatoebaSentencePair>;
  //   error: string | null;
};
type Action = {
  saveSentencePair: (
    sentenceKey: string,
    sentencepair: TatoebaSentencePair,
  ) => void;
  deleteSentencePair: (sentenceKey: string) => void;
};

export const useMySentenceListStore = create<State & Action>()(
  devtools(
    persist(
      (set, get) => ({
        sentences: {},
        saveSentencePair: (
          sentenceKey: string,
          sentencepair: TatoebaSentencePair,
        ): void => {
          const sentencePairs = get().sentences;
          set({
            sentences: {
              ...sentencePairs,
              [sentenceKey]: sentencepair,
            },
          });
        },
        deleteSentencePair: (sentenceKey: string): void => {
          const sentencePairs = get().sentences;
          delete sentencePairs[sentenceKey];
          set({
            sentences: {
              ...sentencePairs,
            },
          });
        },
      }),
      {
        name: "my-sentence-list",

        partialize: (state) => ({
          sentences: state.sentences,
        }),
      },
    ),
  ),
);
