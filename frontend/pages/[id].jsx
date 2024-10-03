import React from "react";
import { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { useRouter } from "next/router";
import { useTheme } from "../context/ThemeContext";
import Link from "next/link";

const ImagePage = () => {
  const router = useRouter();
  const { darkMode } = useTheme();
  const { id } = router.query;
  const [image, setImage] = useState({});
  const [img, setImg] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    if (!id) return;

    const fetchImage = async () => {
      try {
        const res = await fetch(`http://localhost:8081/images?id=${id}`);

        if (!res.ok) {
          setError(true);
          throw new Error(`Fetch error: ${res.status} ${res.statusText}`);
        }

        const data = await res.json();
        console.log(data);

        if (!data || !data.value) {
          setError(true);
        } else {
          setImage(data);
          setImg(data.value.substring(22));
        }
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchImage();
  }, [id, router]);

  if (loading) return <div>Loading...</div>;

  return (
    <Layout title="Image | Phost">
      {error ? (
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
      ) : (
        <div
          className={`${
            darkMode ? "bg-gray-800" : "bg-slate-200"
          } flex flex-col items-center justify-start min-h-screen pt-12`}
        >
          <img
            unoptimized
            src={`data:image/jpeg;base64,${img}`}
            alt={image.id}
            className="max-w-11/12 md:max-w-full md:h-[70vh] rounded-lg shadow-lg"
          />

          <a
            href={`data:image/jpeg;base64,${img}`}
            download={`image-${image.id}.jpg`}
            className="mt-4 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
          >
            Download Image
          </a>
        </div>
      )}
    </Layout>
  );
};

export default ImagePage;
