interface APIProps {
  url: string;
  body: any;
}

interface APIResponse {
  success: boolean;
  data: any;
  error: Error | null;
}

export const BaseAPI = "http://localhost:8000/api";

export const GetAPI = async (url: string): Promise<APIResponse> =>
  fetch(BaseAPI + url, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  })
    .then((res) => res.json())
    .then((data) => ({
      success: true,
      data: data,
      error: null,
    }))
    .catch((err) => ({
      success: false,
      data: null,
      error: err as Error,
    }));

export const PostAPI = async ({
  url,
  body,
}: APIProps): Promise<APIResponse> => {
  try {
    const res = await fetch(BaseAPI + url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    const data = await res.json();
    return {
      success: true,
      data: data,
      error: null,
    };
  } catch (err: any) {
    return {
      success: false,
      data: null,
      error: err as Error,
    };
  }
};
