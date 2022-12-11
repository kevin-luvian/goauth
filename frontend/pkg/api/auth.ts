import { BaseAPI, GetAPI, PostAPI } from "./fetcher";

export const GetGoogleLoginURL = () => `${BaseAPI}/v1/auth/login/google`;

export const GoogleLogin = () => GetAPI("/v1/auth/login/google");

export const GoogleLoginGET = async () => {
  const data = await GetAPI("https://www.google.com/");
  return data;
};
