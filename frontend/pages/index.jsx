import Layout from "../components/Layout";
import { useState, useEffect } from "react";
import { RiDeleteBin7Fill } from "react-icons/ri";
import { useTheme } from "../context/ThemeContext";
import { LuInspect } from "react-icons/lu";
import { FaRegCopy } from "react-icons/fa";
export default function Home() {
  const { darkMode } = useTheme();
  const [photo, setPhoto] = useState(null);
  const [imageURL, setImageURL] = useState("");

  const handleFileUpload = (event) => {
    const file = event.target.files[0];

    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setPhoto({
          value: reader.result,
          name: file.name,
        });
      };
      reader.readAsDataURL(file);
    }
  };

  const uploadImage = async () => {
    if (!photo) return;

    try {
      const response = await fetch("http://localhost:8081/images", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ value: photo.value }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log("Resim yüklendi:", data);
        setPhoto(null);
        setImageURL(`http://localhost:3000/${data.id}`);
      } else {
        console.error("Yükleme hatası:", response.statusText);
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(imageURL);
  };
  const goToImage = () => {
    window.open(imageURL, "_blank");
  };
  return (
    <Layout title="Phost">
      <div
        className={`min-h-screen w-full ${
          darkMode ? "bg-gray-800" : "bg-slate-200"
        } transition duration-300 ease-in-out flex flex-col justify-start items-center`}
      >
        {!photo ? (
          <div className="flex flex-col items-center mt-48">
            <div
              className={`mb-6 border p-4 ${
                darkMode ? "border-gray-700" : "border-black"
              } rounded-md shadow-md w-full max-w-md`}
            >
              <label
                className={`block font-semibold text-2xl ${
                  darkMode ? "text-gray-300" : "text-gray-700"
                } m-2 p-4 rounded-md`}
              >
                Upload Your Image
              </label>
              <input
                type="file"
                accept="image/*"
                onChange={handleFileUpload}
                className={`block w-full text-sm ${
                  darkMode
                    ? "text-gray-300 bg-gray-600"
                    : "text-gray-500 bg-blue-50"
                } file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold hover:file:bg-blue-100 m-2 rounded-[9999px]`}
              />
            </div>
          </div>
        ) : (
          <div className="flex flex-col mt-48 items-center">
            <div className="relative mb-4">
              <img
                src={photo.value}
                alt={photo.name}
                className={`w-full h-48 object-cover rounded-lg shadow-lg border ${
                  darkMode ? "border-white" : "border-black"
                }`}
              />
              <div className="w-full flex justify-center items-center mt-4 gap-x-2">
                <button
                  onClick={() => setPhoto(null)}
                  className="text-red-500 hover:text-red-700 transition duration-200"
                >
                  <RiDeleteBin7Fill size={24} />
                </button>
                <p
                  className={`text-sm text-center ${
                    darkMode ? "text-gray-200" : "text-gray-700"
                  }`}
                >
                  {photo.name}
                </p>
              </div>
            </div>
            <div>
              <button
                onClick={uploadImage}
                className={`${
                  darkMode
                    ? "bg-blue-700 hover:bg-blue-800"
                    : "bg-blue-600 hover:bg-blue-700"
                } px-12 py-2 text-white rounded-md transition-all ease-linear`}
              >
                Upload
              </button>
            </div>
          </div>
        )}

        {imageURL && (
          <div className="mt-8 flex flex-col items-center">
            <label
              className={`block font-semibold text-xl ${
                darkMode ? "text-gray-300" : "text-gray-700"
              } mb-2`}
            >
              Image URL:
            </label>
            <div className="flex items-center">
              <input
                type="text"
                value={imageURL}
                readOnly
                className={`px-4 py-2 rounded-md border ${
                  darkMode
                    ? "bg-gray-700 text-white border-gray-600"
                    : "bg-white text-gray-700 border-gray-300"
                } w-80`}
              />
              <button
                onClick={copyToClipboard}
                className={`ml-2 w-24 h-10 rounded-md flex justify-center items-center ${
                  darkMode
                    ? "bg-blue-700 hover:bg-blue-800 text-white"
                    : "bg-blue-600 hover:bg-blue-700 text-white"
                } transition duration-300`}
              >
                <FaRegCopy className="text-2xl" />
              </button>
              <button
                onClick={goToImage}
                className={`ml-2 w-24 h-10 rounded-md flex justify-center items-center ${
                  darkMode
                    ? "bg-green-700 hover:bg-green-800 text-white"
                    : "bg-green-600 hover:bg-green-700 text-white"
                } transition duration-300`}
              >
                <LuInspect className="text-2xl" />
              </button>
            </div>
          </div>
        )}
      </div>
    </Layout>
  );
}
