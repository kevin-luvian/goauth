import { ReactElement } from "react";
import Layout from "../common/layout/MainLayout";
import { NextPageWithLayout } from "./_app";

const Page: NextPageWithLayout = () => {
  return <div>404 Not Found</div>;
};

Page.getLayout = function getLayout(page: ReactElement) {
  return <Layout menuID="">{page}</Layout>;
};

export default Page;
