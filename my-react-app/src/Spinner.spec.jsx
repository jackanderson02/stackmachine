import {render, screen} from "@testing-library/react"
import Spinner from "./Spinner";

describe("Spinner", () => {
  it("shows reason as hidden accessibility text", async () => {
    render(<Spinner reason="reason for the delay" />);
    expect(await screen.getByRole("status")).toHaveTextContent(
      "reason for the delay"
    );

  });
});
