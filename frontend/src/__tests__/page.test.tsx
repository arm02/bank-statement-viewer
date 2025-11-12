import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import BalanceCard from "../components/BalanceCard";

test("renders balance correctly", () => {
    render(<BalanceCard balance={11750000} />);
    expect(screen.getByText(/End Balance/i)).toBeInTheDocument();
    expect(screen.getByText(/11.750.000/)).toBeInTheDocument();
});
