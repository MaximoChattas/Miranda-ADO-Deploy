import { describe, it } from "vitest";
import Navbar from "./NavBar";
import { expect } from "vitest";
import { screen, render } from "@testing-library/react";
import { LoginContext, UserProfileContext   } from "../../App";
import { BrowserRouter as Router } from "react-router-dom";

describe('NavBar', () => {

    const renderWithContext = (loggedIn = true, userProfile = {id: 1, name: 'John'}) => {
        render(
            <UserProfileContext.Provider value={ {userProfile} }>
                <LoginContext.Provider value={{ loggedIn }}>
                    <Router>
                        <Navbar />
                    </Router>
                </LoginContext.Provider>
            </UserProfileContext.Provider>
        );
    };

    it('should render the title', () => {
        renderWithContext();
        expect(screen.getByText(/MIRANDA/i)).toBeInTheDocument();
    });

    it('should render availability button', () => {
        renderWithContext();
        expect(screen.getByText(/Ver Disponibilidad/i)).toBeInTheDocument();
    });

    it('should show the user name when logged in', () => {
        renderWithContext(true, { id: 1, name: 'John' });
        expect(screen.getByText(/Hola John/i)).toBeInTheDocument();
    });
    
    it('should show "Iniciar sesion" button when not logged in', () => {
        renderWithContext(false);
        expect(screen.getByText(/Iniciar sesion/i)).toBeInTheDocument();
    });

    it('should have a link to the profile page when logged in', () => {
        renderWithContext(true, { id: 1, name: 'John' });
        const profileLink = screen.getByText(/Hola John/i).closest('a');
        expect(profileLink).toHaveAttribute('href', '/profile');
    });
    
    it('should have a link to the login page when not logged in', () => {
        renderWithContext(false);
        const loginLink = screen.getByText(/Iniciar sesion/i).closest('a');
        expect(loginLink).toHaveAttribute('href', '/login');
    });

})