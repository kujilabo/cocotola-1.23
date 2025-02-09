import type { AxiosRequestConfig } from "axios";

export const emptyFunction = () => {
  return;
};

export const jsonAccessTokenHeaders = (accessToken: string) => {
  return {
    "Content-Type": "application/json",
    Authorization: `Bearer ${accessToken}`,
  };
};

export const jsonAccessTokenRequestConfig = (
  accessToken: string,
): AxiosRequestConfig => {
  return {
    headers: jsonAccessTokenHeaders(accessToken),
  };
};

// export const jsonBasicAuthHeaders = (accessToken: string) => {
//   return {
//     'Content-Type': 'application/json',
//     Authorization: `Bearer ${accessToken}`,
//   };
// };

// export const jsonBasicAuthRequestConfig = (accessToken: string): AxiosRequestConfig => {
//   return {
//     headers: jsonBasicAuthHeaders(accessToken),
//   };
// };

export const blobRequestConfig = (accessToken: string): AxiosRequestConfig => {
  return {
    responseType: "blob",
    headers: { Authorization: `Bearer ${accessToken}` },
  };
};
