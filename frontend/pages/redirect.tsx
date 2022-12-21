import { useRouter } from "next/router";
import { ReactElement, useEffect, useMemo } from "react";
import Layout from "../common/layout/MainLayout";
import { NextPageWithLayout } from "./_app";
import empty from "is-empty";

const Page: NextPageWithLayout = () => {
  const router = useRouter();
  const query = useMemo(
    () => ({
      token: router.query["token"],
      err: router.query["err"],
    }),
    [router.query]
  );

  useEffect(() => {
    if (!empty(query.token)) {
      console.log("MToken", query.token);
    }

    if (!empty(query.err)) {
      console.log("MErr", query.err);
    }
  }, [query]);

  return (
    <div className="m-3">
      authenticating redirect...
      <p>{query.err}</p>
    </div>
  );
};

Page.getLayout = function getLayout(page: ReactElement) {
  return <Layout menuID="redirect">{page}</Layout>;
};

export default Page;
