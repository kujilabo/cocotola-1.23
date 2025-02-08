export class TatoebaSentence{
    id: number;
    number: number;
    lang2: string;
    text: string;
    author: string;
    constructor(id: number, number: number, lang2: string, text: string, author: string){
        this.id = id;
        this.number = number;
        this.lang2 = lang2;
        this.text = text;
        this.author = author;
    }
}

export class TatoebaSentencePair{
    english: TatoebaSentence;
    translation: TatoebaSentence;
    constructor(english: TatoebaSentence, translation: TatoebaSentence){
        this.english = english;
        this.translation = translation;
    }
}
