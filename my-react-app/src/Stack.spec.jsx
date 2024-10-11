import * as React from "react";
import { render, screen } from "@testing-library/react";
import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";
import Stack from "./components/Stack.jsx";

const stubStackApiRoutes = [
    http.get("https://stackmachine.com/stack", () => {
        return HttpResponse.json(
            {
                "frame": {
                    "stackframe_id":0,
                    "stack_number":10
                },
            }
        )
    })

];

const stubQuoteApi = setupServer(...stubStackApiRoutes)

describe("Stack", () => {
    stubQuoteApi.listen()

    it("fetches and renders stack frames", async () => {
        render(<Stack></Stack>)
        screen.debug()
        // expect(await screen.findByText(10)).toBeIntheDocment();
    });
})