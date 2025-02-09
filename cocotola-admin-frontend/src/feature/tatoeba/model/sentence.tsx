export class TatoebaSentence {
  sentenceNumber: number;
  lang2: string;
  text: string;
  author: string;
  constructor(
    sentenceNumber: number,
    lang2: string,
    text: string,
    author: string,
  ) {
    this.sentenceNumber = sentenceNumber;
    this.lang2 = lang2;
    this.text = text;
    this.author = author;
  }
}

export const newTatoebaSentenceWithText = (
  obj: TatoebaSentence,
  text: string,
): TatoebaSentence => {
  return new TatoebaSentence(obj.sentenceNumber, obj.lang2, text, obj.author);
};

export class TatoebaSentencePair {
  src: TatoebaSentence;
  dst: TatoebaSentence;
  constructor(src: TatoebaSentence, dst: TatoebaSentence) {
    this.src = src;
    this.dst = dst;
  }
}
