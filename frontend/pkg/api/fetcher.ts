interface APIProps {
  url: string;
  body: any;
}

export interface APIResponse<T = any> {
  success: boolean;
  data: T;
  error: Error | null;
  status: number;
}

export const BaseAPI = "http://localhost:8000/api";

export const GetAPI = async (url: string): Promise<APIResponse> => {
  let status = 0;
  try {
    const res = await fetch(BaseAPI + url, {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });
    status = res.status;
    const data = await res.json();
    return {
      success: true,
      status: status,
      data: data,
      error: null,
    };
  } catch (err) {
    console.log("An Error Occured",err)
    return {
      success: false,
      status: status,
      data: null,
      error: err as Error,
    };
  }
};

export const PostAPI = async ({
  url,
  body,
}: APIProps): Promise<APIResponse> => {
  let status = 0;
  try {
    const res = await fetch(BaseAPI + url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    status = res.status;
    const data = await res.json();
    return {
      success: true,
      status: status,
      data: data,
      error: null,
    };
  } catch (err: any) {
    return {
      success: false,
      status: status,
      data: null,
      error: err as Error,
    };
  }
};
