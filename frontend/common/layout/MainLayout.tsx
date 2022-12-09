import { ReactNode } from "react";
import Image from "next/image";
import { Layout as AntdLayout, Menu } from "antd";
import { HomeOutlined, HeartOutlined } from "@ant-design/icons";
import Link from "next/link";

interface LayoutProps {
  children: ReactNode;
  menuID: string;
}

export default function Layout({ children, menuID }: LayoutProps) {
  return (
    <AntdLayout style={{ minHeight: "100vh" }}>
      <AntdLayout.Sider collapsible>
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
          <Menu.Item key="wishlists">
            <HeartOutlined />
            <span>Wishlists</span>
            <Link href="/wishlists" />
          </Menu.Item>
        </Menu>
      </AntdLayout.Sider>

      <AntdLayout>{children}</AntdLayout>
    </AntdLayout>
  );
}
