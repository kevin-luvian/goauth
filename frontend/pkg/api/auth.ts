import { GetAPI, PostAPI } from "./utils";

export const GoogleLogin = async () => {
  const data = await PostAPI({
    url: "https://www.google.com/",
    body: {
      name: "bruhh",
    },
  });

  return data;
};

export const GoogleLoginGET = async () => {
  const data = await GetAPI("https://www.google.com/");
  return data;
};
