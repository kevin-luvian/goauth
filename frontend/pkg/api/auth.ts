import { APIResponse, BaseAPI, GetAPI, PostAPI } from "./fetcher";

export const GetGoogleLoginURL = () => `${BaseAPI}/v1/auth/login/google`;
export const GetGoogleSignupURL = () => `${BaseAPI}/v1/auth/signup/google`;

export const RefreshToken = async () => {
  let token = "";

  const res = await GetAPI(`${BaseAPI}/v1/auth/refresh-token`);
  if (res.success) {
    token = res.data["token"];
  }

  const reLogin = !res.success || res.error != null;

  return {
    reLogin: reLogin,
    token: token,
  };
};
