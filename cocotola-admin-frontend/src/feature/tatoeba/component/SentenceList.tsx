import { MainLayout } from "@/component/layout";
import {
  TatoebaSentence,
  TatoebaSentencePair,
  newTatoebaSentenceWithText,
} from "@/feature/tatoeba/model/sentence";
import { useSentenceListStore } from "@/feature/tatoeba/store/sentence_list";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import CardActions from "@mui/material/CardActions";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import ListSubheader from "@mui/material/ListSubheader";
import Typography from "@mui/material/Typography";
import { Fragment, useCallback, useEffect, useState } from "react";

import { StreamTwoTone } from "@mui/icons-material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";

const convertSelectedText = (text: string): string => {
  const first = text.substring(0, 1);
  const last = text.substring(text.length - 1);
  if (first === "<" && last === ">") {
    return text.substring(1, text.length - 1);
  }
  if (text.indexOf("<") !== -1 || text.indexOf(">") !== -1) {
    return text.replace("<", "").replace(">", "");
  }
  return `<${text}>`;
};

const convertText = (text: string, start: number, end: number): string => {
  console.log("convertText", text, start, end);
  const selectedText = text.substring(start, end);
  const text1 = text.substring(0, start);
  const text2 = convertSelectedText(selectedText);
  const text3 = text.substring(end);
  return text1 + text2 + text3;
};

const findSentenceKeylement = (
  element: HTMLElement | null,
): HTMLElement | null => {
  let key = null;
  let target = element;
  for (let i = 0; i < 5; i++) {
    if (target === null) {
      break;
    }

    if (target instanceof HTMLElement) {
      key = target.getAttribute("data-sentence-key");
      if (key !== null) {
        return target;
      }
    }
    target = target?.parentElement;
  }
  console.log("not found");
  return null;
};

const findSentenceKey = (element: HTMLElement | null): string => {
  const sentenceKeyElement = findSentenceKeylement(element);
  return sentenceKeyElement?.getAttribute("data-sentence-key") ?? "";
};

const formatText = (
  sentence: string,
  sentenceKey: string,
  sentenceSrcDst: string,
  stageSententcePairs: Map<string, TatoebaSentencePair>,
) => {
  const stageSentencePair = stageSententcePairs.get(sentenceKey);
  if (stageSentencePair === undefined) {
    return sentence;
  }
  // console.log("sentence, selections", sentence, stageSentencePair);
  if (sentenceSrcDst === "src") {
    return stageSentencePair.src.text;
  }
  if (sentenceSrcDst === "dst") {
    return stageSentencePair.dst.text;
  }
  return "ERROR";
};
export const SentenceList = () => {
  const [stageSentencePairs, setStageSentencePairs] = useState<
    Map<string, TatoebaSentencePair>
  >(new Map<string, TatoebaSentencePair>());
  const [errors, setErrors] = useState<Map<string, string>>(
    new Map<string, string>(),
  );
  const [selection, setSelection] = useState<Selection | null>(null);

  const [selectedSentenceKey, setSelectedSentenceKey] = useState<string>("");
  const [selectedSentenceSrcDst, setSelectedSentenceSrcDst] =
    useState<string>("");
  const sentencePairs = useSentenceListStore((state) => state.sentences);
  const getSentences = useSentenceListStore((state) => state.getSentences);

  // console.log("sentenceMap", sentenceMap);

  useEffect(() => {
    getSentences();
  }, [getSentences]);

  useEffect(() => {
    const stageSentencePairs = new Map<string, TatoebaSentencePair>(
      sentencePairs.map((sentencePair) => [
        `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`,
        sentencePair,
      ]),
    );
    setStageSentencePairs(stageSentencePairs);
  }, [sentencePairs]);

  const handleSelectionChange = useCallback(() => {
    const selection = document.getSelection();
    // console.log("selection", selection);
    if (
      selection === null ||
      selection?.anchorNode === null ||
      selection?.focusNode === null ||
      selection?.anchorNode !== selection?.focusNode
    ) {
      setSelection(null);
      setSelectedSentenceKey("");
      setSelectedSentenceSrcDst("");
      return;
    }

    const sentenceKeyElement = findSentenceKeylement(selection?.anchorNode);
    setSelection(selection);
    setSelectedSentenceKey(
      sentenceKeyElement?.getAttribute("data-sentence-key") ?? "",
    );
    setSelectedSentenceSrcDst(
      sentenceKeyElement?.getAttribute("data-sentence-src-dst") ?? "",
    );
  }, []);

  const handleCardClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    const sentenceKey = findSentenceKey(event.target);
    if (sentenceKey === null || sentenceKey === "") {
      return;
    }
    setErrors((errors) => {
      errors.delete(sentenceKey);
      return new Map(errors);
    });

    // const cardElement = findCardElement(event.target);
    if (selectedSentenceKey === "" || selectedSentenceSrcDst === "") {
      return;
    }
    if (selection === null || selection.toString() === "") {
      return;
    }
    if (sentenceKey !== selectedSentenceKey) {
      return;
    }

    const stageSentencePair = stageSentencePairs.get(sentenceKey);
    if (stageSentencePair === undefined) {
      console.log("problem is undefined");
      return;
    }

    const spacePos = selection.toString().indexOf(" ");
    if (spacePos !== -1) {
      setErrors((errors) => {
        return new Map(errors.set(sentenceKey, "Please select one word"));
      });
      return "";
    }

    const anchorOffset = selection.anchorOffset;
    const focusOffset = selection.focusOffset;

    let newSentencePair: TatoebaSentencePair;
    if (selectedSentenceSrcDst === "src") {
      const newText = convertText(
        stageSentencePair.src.text,
        anchorOffset,
        focusOffset,
      );
      newSentencePair = new TatoebaSentencePair(
        newTatoebaSentenceWithText(stageSentencePair.src, newText),
        stageSentencePair.dst,
      );
    } else if (selectedSentenceSrcDst === "dst") {
      const newText = convertText(
        stageSentencePair.dst.text,
        anchorOffset,
        focusOffset,
      );
      newSentencePair = new TatoebaSentencePair(
        stageSentencePair.src,
        newTatoebaSentenceWithText(stageSentencePair.dst, newText),
      );
    } else {
      console.log("error", selectedSentenceKey);
      return;
    }

    setStageSentencePairs((stageSentencePair) => {
      return new Map(stageSentencePair.set(sentenceKey, newSentencePair));
    });
  };

  useEffect(() => {
    document.addEventListener("selectionchange", handleSelectionChange);

    return () => {
      document.removeEventListener("selectionchange", handleSelectionChange);
    };
  }, [handleSelectionChange]);

  const createCard = (sentencePair: TatoebaSentencePair) => {
    const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
    const stagestageSentencePair = stageSentencePairs.get(sentenceKey);
    if (stagestageSentencePair === undefined) {
      return <div>error</div>;
    }
    const error = errors.get(sentenceKey);
    return (
      <Fragment>
        <CardContent>
          <Typography
            sx={{ color: "text.primary", mb: 1.5 }}
            data-sentence-key={sentenceKey}
            data-sentence-src-dst={"src"}
          >
            {formatText(
              sentencePair.src.text,
              sentenceKey,
              "src",
              stageSentencePairs,
            )}
          </Typography>
          <Typography
            sx={{ color: "text.secondary", mb: 1.5 }}
            data-sentence-key={`${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`}
            data-sentence-src-dst={"dst"}
          >
            {formatText(
              sentencePair.dst.text,
              sentenceKey,
              "dst",
              stageSentencePairs,
            )}
          </Typography>
          {error !== "" ? <Typography>{error}</Typography> : <></>}
        </CardContent>
        <CardActions>
          <Button size="small" variant="outlined" onClick={handleCardClick}>
            Mark
          </Button>
        </CardActions>
      </Fragment>
    );
  };
  return (
    <MainLayout title="Sentence List">
      {sentencePairs.map((sentencePair) => (
        <Card
          variant="outlined"
          key={`${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`}
          data-sentence-key={`${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`}
        >
          {createCard(sentencePair)}
        </Card>
      ))}
    </MainLayout>
  );
};
