import React, { useState } from "react";
import { Transaction } from "../types";

interface Props {
    issues: Transaction[];
    page: number;
    totalPages: number;
    onPageChange: (page: number) => void;
    onSortChange?: (sortBy: string, sortOrder: string) => void;
}

export default function IssuesTable({
    issues,
    page,
    totalPages,
    onPageChange,
    onSortChange,
}: Props) {
    const [sortBy, setSortBy] = useState<string>("timestamp");
    const [sortOrder, setSortOrder] = useState<"asc" | "desc">("asc");

    const handleSort = (field: string) => {
        let newOrder: "asc" | "desc" = "asc";
        if (field === sortBy) {
            newOrder = sortOrder === "asc" ? "desc" : "asc";
        }
        setSortBy(field);
        setSortOrder(newOrder);
        onSortChange?.(field, newOrder);
    };

    const renderSortIcon = (field: string) => {
        if (sortBy !== field) return "↕️";
        return sortOrder === "asc" ? "▲" : "▼";
    };

    return (
        <div className="table-container">
            <h3>Transaction Issues</h3>
            <table>
                <thead>
                    <tr>
                        <th onClick={() => handleSort("timestamp")} style={{ cursor: "pointer" }}>
                            Date {renderSortIcon("timestamp")}
                        </th>
                        <th onClick={() => handleSort("name")} style={{ cursor: "pointer" }}>
                            Name {renderSortIcon("name")}
                        </th>
                        <th onClick={() => handleSort("type")} style={{ cursor: "pointer" }}>
                            Type {renderSortIcon("type")}
                        </th>
                        <th onClick={() => handleSort("amount")} style={{ cursor: "pointer" }}>
                            Amount {renderSortIcon("amount")}
                        </th>
                        <th onClick={() => handleSort("status")} style={{ cursor: "pointer" }}>
                            Status {renderSortIcon("status")}
                        </th>
                        <th>Description</th>
                    </tr>
                </thead>
                <tbody>
                    {issues.map((t, idx) => (
                        <tr key={idx} className={t.status.toLowerCase()}>
                            <td>{new Date(t.timestamp).toLocaleDateString("id-ID")}</td>
                            <td>{t.name}</td>
                            <td>{t.type}</td>
                            <td>Rp {Number(t.amount).toLocaleString("id-ID")}</td>
                            <td className={`status ${t.status.toLowerCase()}`}>{t.status}</td>
                            <td>{t.description}</td>
                        </tr>
                    ))}
                </tbody>
            </table>

            <div className="pagination">
                <button disabled={page <= 1} onClick={() => onPageChange(page - 1)}>
                    Prev
                </button>
                <span>
                    Page {page} / {totalPages}
                </span>
                <button disabled={page >= totalPages} onClick={() => onPageChange(page + 1)}>
                    Next
                </button>
            </div>
        </div>
    );
}
