import { describe, it, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { LoginContext, UserProfileContext } from "../../App";
import { BrowserRouter as Router } from "react-router-dom";
import HotelList from "./HotelList";

describe("HotelList", () => {
    
    const renderWithContext = (loggedIn = false, userProfile = {}) => {
        render(
            <UserProfileContext.Provider value={ {userProfile} }>
                <LoginContext.Provider value={{ loggedIn }}>
                    <Router>
                        <HotelList />
                    </Router>
                </LoginContext.Provider>
            </UserProfileContext.Provider>
        );
    };
  
    beforeEach(() => {
      vi.restoreAllMocks();
    });
  
    it("should display hotels when data is fetched successfully", async () => {
      
      global.fetch = vi.fn()
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({ apiUrl: "https://apimock.com" }),
        })
        .mockResolvedValueOnce({
          ok: true,
          json: async () => [
            { id: 1, name: "Hotel A", street_name: "Hotel-A St", street_number: 123, rate: 150, images: [{ id: "img1" }] },
          ],
        });
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/Hoteles/i)).toBeInTheDocument();
        expect(screen.getByText(/Hotel A/i)).toBeInTheDocument();
        expect(screen.getByText(/Hotel-A St 123/i)).toBeInTheDocument();
        expect(screen.getByText(/\$150/i)).toBeInTheDocument();
      });
    });
  
    it("should display an error when the config fetch fails", async () => {
      
      global.fetch = vi.fn()
        .mockRejectedValueOnce(new Error("Failed to load configuration"));
  
      renderWithContext();

      await waitFor(() => {
        expect(screen.getByText(/Error: Failed to load configuration/i)).toBeInTheDocument();
      });
    });
  
    it("should display an error when the hotel fetch fails", async () => {
      
      global.fetch = vi.fn()
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({ apiUrl: "https://apimock.com" }),
        })
        .mockRejectedValueOnce(new Error("Failed to fetch hotels"));
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/Error: Failed to fetch hotels/i)).toBeInTheDocument();
      });
    });
  
    it("should display 'No hay hoteles disponibles' when no hotels are returned", async () => {
      
      global.fetch = vi.fn()
        .mockResolvedValueOnce({
          ok: true,
          json: async () => ({ apiUrl: "https://example.com" }),
        })
        .mockResolvedValueOnce({
          ok: true,
          json: async () => {},
        });
  
      renderWithContext();
  
      await waitFor(() => {
        expect(screen.getByText(/No hay hoteles disponibles/i)).toBeInTheDocument();
      });
    });

  });