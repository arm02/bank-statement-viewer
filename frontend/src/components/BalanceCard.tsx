import React from "react";

interface Props {
    balance: number;
}

export default function BalanceCard({ balance }: Props) {
    return (
        <div className="card">
            <h3>End Balance</h3>
            <p className="balance">Rp {Number(balance).toLocaleString("id-ID")}</p>
        </div >
    );
}
