import { useEffect, useState } from "react";
import Layout from "../components/Layout";
import { useRouter } from "next/router";
import { useTheme } from "../context/ThemeContext";

const ImagePage = () => {
  const router = useRouter();
  const { darkMode } = useTheme();
  const { id } = router.query;
  const [image, setImage] = useState({});
  const [img, setImg] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!id) return;

    const fetchImage = async () => {
      try {
        const res = await fetch(`http://localhost:8081/images?id=${id}`);

        if (!res.ok) {
          throw new Error(`Fetch error: ${res.status} ${res.statusText}`);
        }

        const data = await res.json();
        if (!data || !data.value) {
          router.replace("/404");
        } else {
          setImage(data);
          setImg(data.value.substring(22));
        }
      } catch (error) {
        console.error(error);
        router.replace("/nf");
      } finally {
        setLoading(false);
      }
    };

    fetchImage();
  }, [id, router]);

  if (loading) return <div>Loading...</div>;

  return (
    <Layout title="Image | Phost">
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
      </div>
    </Layout>
  );
};

export default ImagePage;
