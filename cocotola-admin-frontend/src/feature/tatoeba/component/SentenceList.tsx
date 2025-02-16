import type { ChangeEvent } from "react";
import { memo, useCallback, useEffect, useState } from "react";
import { useSearchParams } from "react-router";
// <Pagination count={10} />

import { Pagination, TextField } from "@mui/material";

import Button from "@mui/material/Button";

import { MainLayout } from "@/component/layout";
import { SentencePairCardList } from "@/feature/tatoeba/component/SentencePairCardList";
import {
  TatoebaSentencePair,
  newTatoebaSentenceWithText,
} from "@/feature/tatoeba/model/sentence";
import { useSentenceListStore } from "@/feature/tatoeba/store/sentence_list";

import { StageSentencePairs } from "@/feature/tatoeba/component/stage_sentence_pairs";
import { useMySentencePairListStore } from "@/feature/tatoeba/store/my_sentence_pair_list";
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
  {
    const selectedText = text.substring(start - 1, end + 1);
    const first = selectedText.substring(0, 1);
    const last = selectedText.substring(selectedText.length - 1);
    if (first === "<" && last === ">") {
      // return text.substring(1, text.length - 1);
      const text1 = text.substring(0, start - 1);
      const text2 = selectedText.substring(1, selectedText.length - 1);
      const text3 = text.substring(end + 1);
      return text1 + text2 + text3;
    }
  }

  const selectedText = text.substring(start, end);
  const text1 = text.substring(0, start);
  const text2 = convertSelectedText(selectedText);
  const text3 = text.substring(end);
  return text1 + text2 + text3;
};

const findSentenceKeylement = (element: Node | null): HTMLElement | null => {
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

type SearchButtonProp = {
  onClick: () => void;
};

const SearchButton = memo(({ onClick }: SearchButtonProp) => {
  console.log("reset: SearchButton");
  return (
    <Button
      size="small"
      variant="outlined"
      sx={{ textTransform: "none" }}
      onClick={onClick}
    >
      Search
    </Button>
  );
});

type ExportButtonProp = {
  onClick: () => void;
};

const ExportButton = memo(({ onClick }: ExportButtonProp) => {
  console.log("reset: ExportButton");
  return (
    <Button
      size="small"
      variant="outlined"
      sx={{ textTransform: "none" }}
      onClick={onClick}
    >
      Export
    </Button>
  );
});

const pageNoStrToInt = (pageNoStr: string): number => {
  const pageNoInt = Number.parseInt(pageNoStr);
  if (!Number.isNaN(pageNoInt)) {
    return pageNoInt;
  }
  return 1;
};

const downloadObjectAsJson = (
  document: Document,
  fileName: string,
  mySentencePairs: { [key: string]: TatoebaSentencePair },
) => {
  const blob = new Blob([JSON.stringify(mySentencePairs)], {
    type: "application/json",
  });
  const objectUrl = URL.createObjectURL(blob);

  const a = document.createElement("a");
  document.body.appendChild(a);
  a.style = "display: none";
  a.href = objectUrl;
  a.download = fileName;
  a.click();
  URL.revokeObjectURL(objectUrl);
  a.remove();
};
// Reset.displayName = "Reset";

export const SentenceList = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [stageSentencePairs, setStageSentencePairs] =
    useState<StageSentencePairs>(
      new StageSentencePairs(new Map<string, TatoebaSentencePair>()),
    );

  const [errors, setErrors] = useState<Map<string, string>>(
    new Map<string, string>(),
  );
  const [selection, setSelection] = useState<Selection | null>(null);

  const [selectedSentenceKey, setSelectedSentenceKey] = useState<string>("");
  const [selectedSentenceSrcDst, setSelectedSentenceSrcDst] =
    useState<string>("");

  const sentencePairs = useSentenceListStore((state) => state.sentences);
  const totalSentencePairs = useSentenceListStore(
    (state) => state.totalSentencePairs,
  );
  const getSentences = useSentenceListStore((state) => state.getSentences);
  const [sentencePairStatuses, setSentencePairStatuses] = useState<
    Map<string, string>
  >(new Map<string, string>());

  const isLoading = useSentenceListStore((state) => state.loading);
  const mySentencePairs = useMySentencePairListStore(
    (state) => state.sentencePairs,
  );
  const addSentencePair = useMySentencePairListStore(
    (state) => state.addSentencePair,
  );
  const removeSentencePair = useMySentencePairListStore(
    (state) => state.removeSentencePair,
  );
  const [keyword, setKeyword] = useState<string>("");
  const [pageNo, setPageNo] = useState<number>(1);

  const numPages = Math.floor(totalSentencePairs / 10);
  console.log("numPages", numPages);

  const onKeywordChange = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => setKeyword(event.target.value),
    [],
  );

  const getQueryParam = useCallback(
    (key: string) => searchParams.get(key) || "",
    [searchParams],
  );

  const setError = useCallback((sentenceKey: string, error: string) => {
    setErrors((errors) => {
      if (errors.get(sentenceKey) === error) {
        return errors;
      }
      return new Map(errors.set(sentenceKey, error));
    });
  }, []);

  const clearError = useCallback((sentenceKey: string) => {
    setErrors((errors) => {
      errors.delete(sentenceKey);
      return errors;
    });
  }, []);

  useEffect(() => {
    const keyword = getQueryParam("keyword");
    setKeyword(keyword);

    const pageNo = pageNoStrToInt(getQueryParam("pageNo"));
    setPageNo(pageNo);

    getSentences({
      pageNo: pageNo,
      pageSize: 10,
      keyword: keyword,
      srcLang2: "en",
      dstLang2: "ja",
      random: false,
    });
  }, [getSentences, getQueryParam]);

  useEffect(() => {
    const newStageSentencePairs = new Map<string, TatoebaSentencePair>();

    for (const sentencePair of sentencePairs) {
      const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
      if (sentenceKey in mySentencePairs) {
        newStageSentencePairs.set(sentenceKey, mySentencePairs[sentenceKey]);
        setSentencePairStatuses((sentencePairStatuses) =>
          sentencePairStatuses.set(sentenceKey, "saved"),
        );
      } else {
        newStageSentencePairs.set(sentenceKey, sentencePair);
      }
    }

    const stageSentencePairs = new Map<string, TatoebaSentencePair>(
      newStageSentencePairs,
    );
    setStageSentencePairs(new StageSentencePairs(stageSentencePairs));
  }, [sentencePairs, mySentencePairs]);

  useEffect(() => {
    if (pageNo !== 1) {
      setSearchParams({ pageNo: pageNo.toString() });
    } else {
      setSearchParams((searchParams) => {
        searchParams.delete("pageNo");
        return searchParams;
      });
    }
  }, [pageNo, setSearchParams]);

  const onSearchClick = useCallback(() => {
    if (keyword !== "") {
      setSearchParams({ keyword: keyword });
    } else {
      setSearchParams((searchParams) => {
        searchParams.delete("keyword");
        return searchParams;
      });
    }

    getSentences({
      pageNo: pageNo,
      pageSize: 10,
      keyword: keyword,
      srcLang2: "en",
      dstLang2: "ja",
      random: false,
    });
  }, [getSentences, setSearchParams, pageNo, keyword]);

  const handleSelectionChange = useCallback(() => {
    const currSelection = document.getSelection();
    // console.log("selection", selection);
    if (
      currSelection === null ||
      currSelection?.anchorNode === null ||
      currSelection?.focusNode === null ||
      currSelection?.anchorNode !== currSelection?.focusNode ||
      currSelection?.toString() === ""
    ) {
      setSelection(null);
      setSelectedSentenceKey("");
      setSelectedSentenceSrcDst("");
      return;
    }
    console.log("selection 2", currSelection);

    const sentenceKeyElement = findSentenceKeylement(currSelection?.anchorNode);
    const sentenceKey =
      sentenceKeyElement?.getAttribute("data-sentence-key") ?? "";
    if (sentenceKey === "") {
      setSelection(null);
      setSelectedSentenceKey("");
      setSelectedSentenceSrcDst("");
      return;
    }

    setSelection(currSelection);
    setSelectedSentenceKey(
      sentenceKeyElement?.getAttribute("data-sentence-key") ?? "",
    );
    setSelectedSentenceSrcDst(
      sentenceKeyElement?.getAttribute("data-sentence-src-dst") ?? "",
    );
  }, []);

  const onMarkClick = useCallback(
    (event: React.MouseEvent<HTMLButtonElement>) => {
      const sentenceKey = findSentenceKey(event.currentTarget);
      if (sentenceKey === null || sentenceKey === "") {
        return;
      }
      clearError(sentenceKey);

      if (selection === null || selection.toString() === "") {
        setError(sentenceKey, "Please select one word");
        return;
      }
      // const cardElement = findCardElement(event.target);
      if (selectedSentenceKey === "" || selectedSentenceSrcDst === "") {
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
        setError(sentenceKey, "Please select one word");
        return "";
      }

      const anchorOffset = selection.anchorOffset;
      const focusOffset = selection.focusOffset;

      const start = Math.min(anchorOffset, focusOffset);
      const end = Math.max(anchorOffset, focusOffset);
      const newSentencePair = createNewSentencePair(
        selectedSentenceSrcDst,
        stageSentencePair,
        start,
        end,
      );

      setStageSentencePairs((stageSentencePairs) =>
        stageSentencePairs.createWithNewSentencePair(
          sentenceKey,
          newSentencePair,
        ),
      );
      setSentencePairStatuses((sentencePairStatuses) =>
        sentencePairStatuses.set(sentenceKey, "staged"),
      );
    },
    [
      setError,
      clearError,
      selection,
      stageSentencePairs,
      selectedSentenceKey,
      selectedSentenceSrcDst,
    ],
  );

  const onSaveClick = useCallback(
    (event: React.MouseEvent<HTMLButtonElement>) => {
      const sentenceKey = findSentenceKey(event.currentTarget);
      if (sentenceKey === null || sentenceKey === "") {
        return;
      }
      clearError(sentenceKey);
      const error = stageSentencePairs.validate(sentenceKey);
      if (error !== null) {
        setError(sentenceKey, error);
        return;
      }
      const sentencePair = stageSentencePairs.get(sentenceKey);
      if (sentencePair === undefined) {
        return;
      }
      console.log("sentencePair", sentencePair);
      addSentencePair(sentenceKey, sentencePair);
    },
    [setError, clearError, stageSentencePairs, addSentencePair],
  );

  const onRemoveClick = useCallback(
    (event: React.MouseEvent<HTMLButtonElement>) => {
      const sentenceKey = findSentenceKey(event.currentTarget);
      if (sentenceKey === null || sentenceKey === "") {
        return;
      }
      removeSentencePair(sentenceKey);
      setSentencePairStatuses((sentencePairStatuses) => {
        sentencePairStatuses.delete(sentenceKey);
        return sentencePairStatuses;
      });
      console.log("onRemoveClick");
    },
    [removeSentencePair],
  );

  const onExportClick = useCallback(() => {
    downloadObjectAsJson(document, "fileName.json", mySentencePairs);
  }, [mySentencePairs]);

  const onEmptyClick = useCallback(
    (event: React.MouseEvent<HTMLButtonElement>) => {},
    [],
  );

  useEffect(() => {
    document.addEventListener("selectionchange", handleSelectionChange);

    return () => {
      document.removeEventListener("selectionchange", handleSelectionChange);
    };
  }, [handleSelectionChange]);

  const onPageChange = (event: React.ChangeEvent<unknown>, value: number) => {
    console.log("onPageChange", value);
    setPageNo(value);
  };

  return (
    <MainLayout title="Sentence List">
      <ExportButton onClick={onExportClick} />
      <TextField
        id="keyword"
        label="Keyword"
        variant="standard"
        value={keyword}
        onChange={onKeywordChange}
      />
      <SearchButton onClick={onSearchClick} />

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <SentencePairCardList
          errors={errors}
          sentencePairs={sentencePairs}
          sentencePairStatuses={sentencePairStatuses}
          stageSentencePairs={stageSentencePairs}
          onMarkClick={onMarkClick}
          onSaveClick={onSaveClick}
          onRemoveClick={onRemoveClick}
          // onMarkClick={onEmptyClick}
          // onSaveClick={onEmptyClick}
          // onRemoveClick={onEmptyClick}
        />
      )}
      <Pagination count={numPages} onChange={onPageChange} />
    </MainLayout>
  );
};
