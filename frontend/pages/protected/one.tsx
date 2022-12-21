import { ReactElement } from "react";
import Layout from "../../common/layout/MainLayout";
import { NextPageWithLayout } from "../_app";

const Page: NextPageWithLayout = () => {
  return (
    <div className="m-3">
      <p>protected</p>
    </div>
  );
};

Page.getLayout = function getLayout(page: ReactElement) {
  return <Layout menuID="protected-one">{page}</Layout>;
};

export default Page;
