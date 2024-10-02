import Link from "next/link";
import Layout from "../components/Layout";
import { useTheme } from "../context/ThemeContext";

const NotFoundPage = () => {
  const { darkMode } = useTheme();

  return (
    <Layout title="404 - Sayfa Bulunamadı">
      <div
        className={`${
          darkMode ? "bg-gray-800 text-white" : "bg-white text-gray-800"
        } flex flex-col items-center justify-center min-h-screen`}
      >
        <h1 className="text-6xl font-bold mb-4">404</h1>
        <p className="text-2xl mb-8">
          Sorry, the image you are looking for was not found
        </p>
        <Link href="/">
          <p className="text-lg text-blue-500 hover:underline">
            Click here to return to the home page
          </p>
        </Link>
      </div>
    </Layout>
  );
};

export default NotFoundPage;
