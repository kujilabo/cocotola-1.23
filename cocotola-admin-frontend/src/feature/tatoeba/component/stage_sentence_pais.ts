import type { TatoebaSentencePair } from "../model/sentence";

export class StageSentencePairs {
  map: Map<string, TatoebaSentencePair>;
  regex: RegExp = /<([^>]*)>/;

  constructor(map: Map<string, TatoebaSentencePair>) {
    this.map = new Map<string, TatoebaSentencePair>(map);
    this.regex = /<([^>]*)>/;
  }

  validate = (sentenceKey: string): string | null => {
    const stageSentencePair = this.map.get(sentenceKey);
    if (stageSentencePair === undefined) {
      return "problem is undefined";
    }

    if (
      !this.regex.test(stageSentencePair.src.text) ||
      !this.regex.test(stageSentencePair.dst.text)
    ) {
      return "Please mark at least one word";
    }
    return null;
  };

  get = (sentenceKey: string): TatoebaSentencePair | undefined => {
    return this.map.get(sentenceKey);
  };
  set = (sentenceKey: string, sentencePair: TatoebaSentencePair): void => {
    this.map.set(sentenceKey, sentencePair);
  };

  keys = (): string[] => {
    return Array.from(this.map.keys());
  };

  createWithNewSentencePair = (
    sentenceKey: string,
    sentencePair: TatoebaSentencePair,
  ): StageSentencePairs => {
    this.map.set(sentenceKey, sentencePair);
    return new StageSentencePairs(this.map);
  };

  clone = (): StageSentencePairs => {
    return new StageSentencePairs(this.map);
  };
}
