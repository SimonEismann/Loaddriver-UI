import { API_ROOT } from "../env.js";
import TextFile from "./text-file.js";

export const uploadIntensityFile = async (fileName, content) => {
  try {
    await fetch(`${API_ROOT}/intensities`, {
      body: JSON.stringify(new TextFile(fileName, content)),
      headers: {
        "Content-type": "application/json"
      },
      method: "POST",
      mode: "cors"
    });
  } catch (error) {
    throw error;
  }
};