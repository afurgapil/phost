import Layout from "../components/Layout";
import React, { useState } from "react";
import { useTheme } from "../context/ThemeContext";
import { RiDeleteBin7Fill } from "react-icons/ri";
import { FaRegCopy } from "react-icons/fa";
import { LuInspect } from "react-icons/lu";
import { FaUpload } from "react-icons/fa";
const Home = () => {
  const { darkMode } = useTheme();
  const [file, setFile] = useState(null);
  const [uploadedImage, setUploadedImage] = useState(null);
  const [imageURL, setImageURL] = useState(null);
  const [isUploaded, setIsUploaded] = useState(false);

  const handleFileUpload = (event) => {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
      setFile(selectedFile);
      const reader = new FileReader();
      reader.onload = () => {
        setUploadedImage(reader.result);
      };
      reader.readAsDataURL(selectedFile);
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(imageURL);
  };

  const goToImage = () => {
    window.open(imageURL, "_blank");
  };

  const uploadImage = async () => {
    if (!file) return;

    try {
      const response = await fetch("http://localhost:8081/images", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ value: uploadedImage }),
      });
      if (response.ok) {
        const data = await response.json();
        console.log("Image uploaded:", data);
        setImageURL(`http://localhost:3000/${data.id}`);
        setIsUploaded(true);
      } else {
        console.error("Upload error:", response.statusText);
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  return (
    <Layout title="Phost">
      <main>
        <div
          data-testid="home-page"
          className={`min-h-screen w-full transition duration-300 ease-in-out flex flex-col justify-start items-center ${
            darkMode ? "bg-gray-800" : "bg-slate-200"
          }`}
        >
          <div className="flex flex-col items-center mt-48">
            {uploadedImage ? (
              <div>
                <div className="mt-4 flex flex-col justify-center items-center gap-y-2">
                  <img
                    src={uploadedImage}
                    alt="Uploaded Preview"
                    className="max-w-full h-auto border rounded-md"
                  />
                  <div
                    className={`${
                      isUploaded ? "hidden" : "flex"
                    } flex-row justify-center items-center gap-x-2`}
                  >
                    <button
                      onClick={() => setUploadedImage(null)}
                      data-testid="remove-button"
                      className={`${
                        darkMode
                          ? "bg-red-700 hover:bg-red-800"
                          : "bg-red-600 hover:bg-red-700"
                      } px-12 py-2 text-white rounded-md transition-all ease-linear`}
                    >
                      <RiDeleteBin7Fill data-testid="delete-icon" />
                    </button>

                    <button
                      onClick={uploadImage}
                      className={`${
                        darkMode
                          ? "bg-blue-700 hover:bg-blue-800"
                          : "bg-blue-600 hover:bg-blue-700"
                      } px-12 py-2 text-white rounded-md transition-all ease-linear`}
                    >
                      <FaUpload />
                    </button>
                  </div>
                </div>
              </div>
            ) : (
              <div
                className={`${
                  darkMode ? "border-white" : "border-black"
                } mb-6 border px-16 py-8 rounded-md shadow-md w-full max-w-md`}
              >
                <label
                  className={`${
                    darkMode ? "text-slate-200" : "text-gray-700"
                  } block font-semibold text-2xl m-2 p-4 rounded-md`}
                  htmlFor="imageUpload"
                >
                  Upload Your Image
                </label>
                <input
                  id="imageUpload"
                  accept="image/*"
                  className="block w-full text-sm text-gray-500 bg-slate-300 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold hover:file:bg-blue-100 m-2 rounded-[9999px]"
                  type="file"
                  onChange={handleFileUpload}
                />
              </div>
            )}
            {isUploaded && (
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
        </div>
      </main>
    </Layout>
  );
};

export default Home;
