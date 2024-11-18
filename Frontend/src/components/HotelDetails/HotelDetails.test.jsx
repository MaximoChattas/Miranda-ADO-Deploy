import { render, screen, waitFor } from "@testing-library/react";
import { describe, it, vi } from "vitest";
import { LoginContext, UserProfileContext } from "../../App";
import { useParams, useNavigate, NavLink } from "react-router-dom";
import HotelDetails from "./HotelDetails";

vi.mock("react-router-dom", () => ({
    ...vi.importActual("react-router-dom"),
    useParams: () => ({ id: "1" }),
    useNavigate: () => vi.fn(),
    NavLink: () => vi.fn(),
}));
  
describe("HotelDetails", () => {
    const renderWithContext = (loggedIn = true, userProfile = {}) => {
      render(
        <LoginContext.Provider value={{ loggedIn }}>
          <UserProfileContext.Provider value={{ userProfile }}>
              <HotelDetails />
          </UserProfileContext.Provider>
        </LoginContext.Provider>
      );
    };
  
    beforeEach(() => {
      vi.restoreAllMocks();
    });
  
    it("should display hotel details when both fetches succeed", async () => {
      
      global.fetch = vi.fn()
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({ apiUrl: "https://apimock.com" }),
        })
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({
            id: 1,
            name: "Hotel 1",
            street_name: "Hotel St",
            street_number: 123,
            images: [{ id: "img1" }],
          }),
        });
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/Hotel 1/i)).toBeInTheDocument();
        expect(screen.getByText(/Hotel St 123/i)).toBeInTheDocument();
      });
    });
  
    it("should display an error when config fetch fails", async () => {

      global.fetch = vi.fn()
        .mockRejectedValueOnce(new Error("Failed to load configuration"));
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/Error: Failed to load configuration/i)).toBeInTheDocument();
      });
    });
  
    it("should display an error when hotel fetch fails", async () => {

      global.fetch = vi.fn()
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({ apiUrl: "https://apimock.com" }),
        })
        // Mock hotel fetch to fail
        .mockResolvedValueOnce({
          ok: false,
          json: async () => ({ error: "Hotel not found" }),
        });
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/Error: Hotel not found/i)).toBeInTheDocument();
      });
    });

    it("should display admin section when logged in as admin", async () => {
      
        global.fetch = vi.fn()
          .mockResolvedValueOnce({
            ok: true,
            json: async () => ({ apiUrl: "https://apimock.com" }),
          })
          .mockResolvedValueOnce({
            ok: true,
            json: async () => ({
              id: 1,
              name: "Hotel 1",
              street_name: "Hotel St",
              street_number: 123,
              images: [{ id: "img1" }],
            }),
          });
    
        renderWithContext(true, {id: 1, name: "John", last_name: "Doe", dni: "123456", email: "johndoe@email.com", role: "Admin"});
    
        await waitFor(() => {
          expect(screen.getByText(/Modificar Hotel/i)).toBeInTheDocument();
          expect(screen.getByText(/Borrar Hotel/i)).toBeInTheDocument();
        });
      });

});