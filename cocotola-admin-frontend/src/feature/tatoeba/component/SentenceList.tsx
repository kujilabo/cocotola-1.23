import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { Fragment, useCallback, useEffect, useState } from "react";

// <Pagination count={10} />
import { MainLayout } from "@/component/layout";
import {
  TatoebaSentencePair,
  newTatoebaSentenceWithText,
} from "@/feature/tatoeba/model/sentence";
import { useMySentenceListStore } from "@/feature/tatoeba/store/my_sentence_list";
import { useSentenceListStore } from "@/feature/tatoeba/store/sentence_list";
import Pagination from "@mui/material/Pagination";

// import useSWR, { preload } from 'swr'

// const fetcher = (url: string) => fetch(url).then((res) => res.json())

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

const createNewSentencePair = (
  selectedSentenceSrcDst: string,
  stageSentencePair: TatoebaSentencePair,
  start: number,
  end: number,
): TatoebaSentencePair => {
  if (selectedSentenceSrcDst === "src") {
    const newText = convertText(stageSentencePair.src.text, start, end);
    return new TatoebaSentencePair(
      newTatoebaSentenceWithText(stageSentencePair.src, newText),
      stageSentencePair.dst,
    );
  }
  if (selectedSentenceSrcDst === "dst") {
    const newText = convertText(stageSentencePair.dst.text, start, end);
    return new TatoebaSentencePair(
      stageSentencePair.src,
      newTatoebaSentenceWithText(stageSentencePair.dst, newText),
    );
  }
  throw new Error("error");
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
  const mySentencePairs = useMySentenceListStore((state) => state.sentences);
  const saveMySentencePair = useMySentenceListStore(
    (state) => state.saveSentencePair,
  );
  const deleteMySentencePair = useMySentenceListStore(
    (state) => state.deleteSentencePair,
  );
  const clearError = (sentenceKey: string) => {
    setErrors((errors) => {
      errors.delete(sentenceKey);
      return new Map(errors);
    });
  };

  useEffect(() => {
    getSentences();
  }, [getSentences]);

  useEffect(() => {
    const stageSentencePairs = new Map<string, TatoebaSentencePair>(
      sentencePairs.map((sentencePair) => {
        const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
        if (sentenceKey in mySentencePairs) {
          return [sentenceKey, mySentencePairs[sentenceKey]];
        }
        return [sentenceKey, sentencePair];
      }),
    );
    setStageSentencePairs(stageSentencePairs);
  }, [sentencePairs, mySentencePairs]);

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

    const sentenceKeyElement = findSentenceKeylement(
      selection.anchorNode.parentElement,
    );
    setSelection(selection);
    setSelectedSentenceKey(
      sentenceKeyElement?.getAttribute("data-sentence-key") ?? "",
    );
    setSelectedSentenceSrcDst(
      sentenceKeyElement?.getAttribute("data-sentence-src-dst") ?? "",
    );
  }, []);

  const onMarkClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    const sentenceKey = findSentenceKey(event.currentTarget);
    if (sentenceKey === null || sentenceKey === "") {
      return;
    }
    clearError(sentenceKey);

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

    let start: number;
    let end: number;
    if (selection.anchorOffset < selection.focusOffset) {
      start = selection.anchorOffset;
      end = selection.focusOffset;
    } else {
      start = selection.focusOffset;
      end = selection.anchorOffset;
    }

    const newSentencePair = createNewSentencePair(
      selectedSentenceSrcDst,
      stageSentencePair,
      start,
      end,
    );

    setStageSentencePairs((stageSentencePair) => {
      return new Map(stageSentencePair.set(sentenceKey, newSentencePair));
    });
  };

  const onSaveClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    const sentenceKey = findSentenceKey(event.currentTarget);
    if (sentenceKey === null || sentenceKey === "") {
      return;
    }
    clearError(sentenceKey);
    const stageSentencePair = stageSentencePairs.get(sentenceKey);
    if (stageSentencePair === undefined) {
      console.log("problem is undefined");
      return;
    }
    const regex = /<([^>]*)>/;
    if (
      !regex.test(stageSentencePair.src.text) ||
      !regex.test(stageSentencePair.dst.text)
    ) {
      setErrors((errors) => {
        return new Map(
          errors.set(sentenceKey, "Please mark at least one word"),
        );
      });
      return;
    }
    console.log("onSaveClick");
    saveMySentencePair(sentenceKey, stageSentencePair);
  };

  const onRemoveClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    const sentenceKey = findSentenceKey(event.currentTarget);
    if (sentenceKey === null || sentenceKey === "") {
      return;
    }
    deleteMySentencePair(sentenceKey);
    console.log("onRemoveClick");
  };

  const onExportClick = () => {
    const blob = new Blob([JSON.stringify(mySentencePairs)], {
      type: "application/json",
    });
    const objectUrl = URL.createObjectURL(blob);

    var a = document.createElement("a");
    document.body.appendChild(a);
    a.style = "display: none";
    a.href = objectUrl;
    a.download = "fileName.json";
    a.click();
    URL.revokeObjectURL(objectUrl);
    a.remove();
  };

  useEffect(() => {
    document.addEventListener("selectionchange", handleSelectionChange);

    return () => {
      document.removeEventListener("selectionchange", handleSelectionChange);
    };
  }, [handleSelectionChange]);

  const createCard = (sentencePair: TatoebaSentencePair) => {
    const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
    // const stagestageSentencePair = stageSentencePairs.get(sentenceKey);
    // if (stagestageSentencePair === undefined) {
    //   return <div>error</div>;
    // }
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
          <Button
            size="small"
            variant="outlined"
            sx={{ textTransform: "none" }}
            onClick={onMarkClick}
          >
            Mark / Unmark
          </Button>
          <Button
            size="small"
            variant="outlined"
            sx={{ textTransform: "none" }}
            onClick={onSaveClick}
          >
            Save
          </Button>
          <Button
            size="small"
            variant="outlined"
            sx={{ textTransform: "none" }}
            onClick={onRemoveClick}
          >
            Remove
          </Button>
        </CardActions>
      </Fragment>
    );
  };
  return (
    <MainLayout title="Sentence List">
      <Button
        size="small"
        variant="outlined"
        sx={{ textTransform: "none" }}
        onClick={onExportClick}
      >
        Export
      </Button>
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
