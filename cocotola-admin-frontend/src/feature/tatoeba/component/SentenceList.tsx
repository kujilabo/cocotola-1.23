import { MainLayout } from "@/component/layout";
import { useSentenceListStore } from "@/feature/store/sentence_list";
import { useEffect } from "react";
export const SentenceList = () => {

  const sentences = useSentenceListStore((state) => state.sentences);
  const getSentences = useSentenceListStore((state) => state.getSentences);

  useEffect(() => {
    getSentences();
  }, []);
  return (<MainLayout title="Sentence List">
    <div>
      {sentences.map((sentence) => (
        <div key={sentence.english.id+"-"+sentence.translation.id}>
          <div>{sentence.english.text}</div>
          <div>{sentence.translation.text}</div>
        </div>
      ))}
    </div>

  </MainLayout>);
};
