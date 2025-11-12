"use client";
import React, { useRef, useState } from "react";
import { uploadCSV } from "../services/api";
import "../styles/file-uploader.css"

interface Props {
    onUploadSuccess: () => void;
}

export default function FileUploader({ onUploadSuccess }: Props) {
    const [file, setFile] = useState<File | null>(null);
    const [isUploading, setIsUploading] = useState(false);
    const fileInputRef = useRef<HTMLInputElement>(null);

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            setFile(e.target.files[0]);
        }
    };

    async function handleUpload() {
        if (!file) return;
        setIsUploading(true);
        try {
            await uploadCSV(file);
            onUploadSuccess();
            alert("Upload success!");
            setFile(null);
            if (fileInputRef.current) {
                fileInputRef.current.value = "";
            }
        } catch (err) {
            alert("Upload failed!");
        } finally {
            setIsUploading(false);
        }
    }

    return (
        <div className="uploader-card">
            <h3>Upload Bank Statement</h3>

            <div className="uploader-input">
                <label htmlFor="fileInput" className="file-label">
                    Choose File
                </label>
                <input
                    id="fileInput"
                    type="file"
                    accept=".csv"
                    onChange={handleFileChange}
                    ref={fileInputRef}
                    className="hidden-input"
                />
                {file && <span className="file-name">{file.name}</span>}
            </div>

            <button
                className="upload-btn"
                onClick={handleUpload}
                disabled={!file || isUploading}
            >
                {isUploading ? "Uploading..." : "Upload File"}
            </button>
        </div>
    );
}
