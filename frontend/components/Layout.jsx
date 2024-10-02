import Head from "next/head";
import Header from "./Header";
import Footer from "./Footer";

export default function Layout({ children, title = "My Next.js Project" }) {
  return (
    <>
      <Head>
        <title>{title}</title>
        <link rel="icon" href="/favicon.ico" />
        <meta name="description" content="A sample Next.js project" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <div className="layout">
        <Header />
        <main className="">{children}</main>
        <Footer />
      </div>
    </>
  );
}
