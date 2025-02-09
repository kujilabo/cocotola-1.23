export const extractErrorMessage = (e: Error): string => {
  const err: {
    response?: {
      data?: {
        message?: string;
      };
      statusText?: string;
    };
  } = e as {
    response?: {
      data?: {
        message?: string;
      };
      statusText?: string;
    };
  };
  console.log("err", err);
  if (err.response) {
    if (err.response.data) {
      if (err.response.data.message) {
        return err.response.data.message;
      }
    }
    if (err.response.statusText) {
      return err.response.statusText;
    }
  }
  return "Error";
};
