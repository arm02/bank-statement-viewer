export const API_BASE =
  process.env.NEXT_PUBLIC_API_BASE || "http://localhost:8080/api";

export async function uploadCSV(file: File | null) {
  const formData = new FormData();
  if (!file) throw new Error("No file provided");
  formData.append("file", file);
  const res = await fetch(`${API_BASE}/upload`, {
    method: "POST",
    body: formData,
  });
  if (!res.ok) throw new Error("Upload failed");
  return res.json();
}

export async function getBalance() {
  const res = await fetch(`${API_BASE}/balance`);
  if (!res.ok) throw new Error("Failed to get balance");
  return res.json();
}

export async function getIssues(page = 1, limit = 10, sortBy = "timestamp", sortOrder = "asc") {
  const res = await fetch(`${API_BASE}/issues?page=${page}&limit=${limit}&sort_by=${sortBy}&sort_order=${sortOrder}`);
  if (!res.ok) throw new Error("Failed to get issues");
  return res.json();
}
