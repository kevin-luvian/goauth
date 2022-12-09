import { Button, Space } from "antd";
import { ReactElement } from "react";
import Layout from "../../common/layout/MainLayout";
import { NextPageWithLayout } from "../_app";
import { GoogleOutlined, DownloadOutlined } from "@ant-design/icons";
import API from "../../pkg/api";

const Page: NextPageWithLayout = () => {
  const handleGoogleLogin = async () => {
    console.log("Google Login Button Clicked!!");
    // const res = await API.Auth.GoogleLogin();
    const res = await API.Auth.GoogleLoginGET();
    console.log("GoogleLogin", res);
    if (res.error) {
      console.log("GoogleLogin Error Msg", res.error.message);
    }
  };

  return (
    <div className="m-3">
      <p>hello world</p>

      <Space direction="horizontal">
        <Button type="primary" shape="round" icon={<DownloadOutlined />}>
          Download
        </Button>

        <Button
          type="primary"
          shape="round"
          icon={<GoogleOutlined />}
          size="middle"
          onClick={handleGoogleLogin}
        >
          Google
        </Button>
      </Space>
    </div>
  );
};

Page.getLayout = function getLayout(page: ReactElement) {
  return <Layout menuID="login">{page}</Layout>;
};

export default Page;
