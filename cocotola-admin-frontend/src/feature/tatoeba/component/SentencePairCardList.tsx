import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { Fragment, memo, useCallback, useEffect, useState } from "react";

// <Pagination count={10} />

import { MainLayout } from "@/component/layout";
import {
  type TatoebaSentencePair,
  newTatoebaSentenceWithText,
} from "@/feature/tatoeba/model/sentence";
import { useSentenceListStore } from "@/feature/tatoeba/store/sentence_list";

import type { StageSentencePairs } from "@/feature/tatoeba/component/stage_sentence_pais";
import { TatoebaSentenceFindParameter } from "@/feature/tatoeba/store/sentence_list";
import { TextField } from "@mui/material";
import { useMySentencePairListStore } from "../store/my_sentence_pair_list";

const formatText = (
  sentence: string,
  sentenceKey: string,
  sentenceSrcDst: string,
  stageSententcePairs: StageSentencePairs,
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
export type SentencePairCardListProps = {
  onMarkClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
  onSaveClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
  onRemoveClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
  errors: Map<string, string>;
  sentencePairs: TatoebaSentencePair[];
  stageSentencePairs: StageSentencePairs;
};

export const SentencePairCardList = memo(
  (props: SentencePairCardListProps): JSX.Element => {
    const {
      sentencePairs,
      stageSentencePairs,
      errors,
      onMarkClick,
      onSaveClick,
      onRemoveClick,
    } = props;

    const createCard = (sentencePair: TatoebaSentencePair) => {
      const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
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
      <div>
        {sentencePairs.map((sentencePair) => (
          <Card
            variant="outlined"
            key={`${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`}
            data-sentence-key={`${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`}
          >
            {createCard(sentencePair)}
          </Card>
        ))}
      </div>
    );
  },
);
