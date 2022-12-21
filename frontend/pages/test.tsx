import { ReactElement, useState } from "react";
import Layout from "../common/layout/MainLayout";
import { NextPageWithLayout } from "./_app";
import api from "../pkg/api";
import { Button, Space } from "antd";

const Page: NextPageWithLayout = () => {
  const [accessToken, setAccessToken] = useState("");
  const [reLogin, setReLogin] = useState(false);

  const refreshClick = async () => {
    const res = await api.Auth.RefreshToken();
    if (!res.reLogin) {
      setAccessToken(res.token);
    }
    setReLogin(res.reLogin);
  };

  return (
    <div className="m-3">
      <p>access token got: {accessToken}</p>
      <p>got relogin: {`${reLogin}`}</p>
      <Button onClick={refreshClick}>Refresh</Button>
    </div>
  );
};

Page.getLayout = function getLayout(page: ReactElement) {
  return <Layout menuID="test">{page}</Layout>;
};

export default Page;
