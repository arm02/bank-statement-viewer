"use client";

import React, { useEffect, useState, useCallback } from "react";
import FileUploader from "../components/FileUploader";
import BalanceCard from "../components/BalanceCard";
import IssuesTable from "../components/IssuesTable";
import { getBalance, getIssues } from "../services/api";
import { Transaction } from "../types";

export default function HomePage() {
    const [balance, setBalance] = useState<number>(0);
    const [issues, setIssues] = useState<Transaction[]>([]);
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [sortBy, setSortBy] = useState("timestamp");
    const [sortOrder, setSortOrder] = useState<"asc" | "desc">("asc");

    const refreshData = useCallback(async () => {
        try {
            const b = await getBalance();
            setBalance(b.balance || 0);

            const i = await getIssues(page, 5, sortBy, sortOrder);
            setIssues(i.data || []);
            setTotalPages(i.meta?.total_pages || 1);
        } catch (err) {
            console.error("Failed to fetch data:", err);
        }
    }, [page, sortBy, sortOrder]);

    useEffect(() => {
        refreshData();
    }, [refreshData]);

    const handleSortChange = (field: string, order: string) => {
        setSortBy(field);
        setSortOrder(order as "asc" | "desc");
        setPage(1); // reset ke page 1 pas sort berubah
    };

    return (
        <main>
            <h1>Bank Statement Viewer</h1>
            <div className="container">
                <FileUploader onUploadSuccess={refreshData} />
                <BalanceCard balance={balance} />
            </div>
            <IssuesTable
                issues={issues}
                page={page}
                totalPages={totalPages}
                onPageChange={setPage}
                onSortChange={handleSortChange}
            />
        </main>
    );
}
