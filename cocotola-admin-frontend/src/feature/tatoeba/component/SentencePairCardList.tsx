import { Fragment, memo } from "react";

import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";

import type { TatoebaSentencePair } from "@/feature/tatoeba/model/sentence";

import type { StageSentencePairs } from "@/feature/tatoeba/component/stage_sentence_pairs";

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
  sentencePairStatuses: Map<string, string>;
  stageSentencePairs: StageSentencePairs;
};

export const SentencePairCardList = memo(
  (props: SentencePairCardListProps): JSX.Element => {
    const {
      sentencePairs,
      sentencePairStatuses,
      stageSentencePairs,
      errors,
      onMarkClick,
      onSaveClick,
      onRemoveClick,
    } = props;

    const createCard = (sentencePair: TatoebaSentencePair) => {
      const sentenceKey = `${sentencePair.src.sentenceNumber}-${sentencePair.dst.sentenceNumber}`;
      const error = errors.get(sentenceKey);
      const status = sentencePairStatuses.get(sentenceKey);
      const color =
        status === "saved"
          ? "orange"
          : status === "staged"
            ? "green"
            : "text.primary";
      return (
        <Fragment>
          <CardContent>
            <Typography
              // sx={{ color: "text.primary", mb: 1.5 }}
              sx={{ color: { color }, mb: 1.5 }}
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
              sx={{ color: { color }, mb: 1.5 }}
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
              disabled={status !== "staged"}
            >
              Save
            </Button>
            <Button
              size="small"
              variant="outlined"
              sx={{ textTransform: "none" }}
              onClick={onRemoveClick}
              disabled={status !== "saved"}
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
