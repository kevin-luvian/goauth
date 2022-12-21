import { ReactNode } from "react";
import { Layout as AntdLayout, Menu } from "antd";
import {
  HomeOutlined,
  HeartOutlined,
  LockOutlined,
  DoubleRightOutlined,
} from "@ant-design/icons";
import Link from "next/link";

interface LayoutProps {
  children: ReactNode;
  menuID: string;
}

export default function Layout({ children, menuID }: LayoutProps) {
  return (
    <AntdLayout style={{ minHeight: "100vh" }}>
      <AntdLayout.Sider collapsible style={{ paddingTop: "1rem" }}>
        <Menu theme="dark" mode="inline" selectedKeys={[menuID]}>
          <Menu.Item key="home">
            <HomeOutlined />
            <span>Home</span>
            <Link href="/" />
          </Menu.Item>

          <Menu.Item key="login">
            <HeartOutlined />
            <span>Login</span>
            <Link href="/auth/login" />
          </Menu.Item>

          <Menu.Item key="test">
            <HeartOutlined />
            <span>Test</span>
            <Link href="/test" />
          </Menu.Item>

          <Menu.SubMenu
            level={1}
            title={
              <>
                <LockOutlined />
                <span>Protected</span>
              </>
            }
          >
            <Menu.Item key="protected-one">
              <DoubleRightOutlined />
              <span>Item</span>
              <Link href="/protected/one" />
            </Menu.Item>
          </Menu.SubMenu>
        </Menu>
      </AntdLayout.Sider>

      <AntdLayout>{children}</AntdLayout>
    </AntdLayout>
  );
}
