import { describe, it } from "vitest";
import AdminPanel from "./AdminPanel";
import { expect } from "vitest";
import { screen, render } from "@testing-library/react";
import { LoginContext, UserProfileContext   } from "../../App";
import { BrowserRouter as Router } from "react-router-dom";

describe('AdminPanel', () => {

    const renderWithContext = (loggedIn = true, userProfile = 
        {id: 1, name: 'John', last_name: 'Doe', dni: '123456', email: 'johndoe@email.com', role: 'Admin'}) => {
        render(
            <UserProfileContext.Provider value={ {userProfile} }>
                <LoginContext.Provider value={{ loggedIn }}>
                    <Router>
                        <AdminPanel />
                    </Router>
                </LoginContext.Provider>
            </UserProfileContext.Provider>
        );
    };

    it('should render the Admin Panel when logged in as Admin', () => {
        renderWithContext();

        expect(screen.getByText(/Panel de Administración/i)).toBeInTheDocument();
        expect(screen.getByText(/Nuevo Hotel/i)).toBeInTheDocument();
        expect(screen.getByText(/Nuevo Amenity/i)).toBeInTheDocument();
        expect(screen.getByText(/Ver reservas por Hotel/i)).toBeInTheDocument();
        expect(screen.getByText(/Ver reservas por usuario/i)).toBeInTheDocument();
    });

    it('should not render the Admin Panel when not logged in', () => {
        renderWithContext(false);

        expect(screen.queryByText(/Panel de Administración/i)).not.toBeInTheDocument();
    });

    it('should not render the Admin Panel when logged in as non-admin', () => {
        renderWithContext(true, { role: 'User' });

        expect(screen.queryByText(/Panel de Administración/i)).not.toBeInTheDocument();
        expect(screen.queryByText(/Nuevo Hotel/i)).not.toBeInTheDocument();
        expect(screen.queryByText(/Nuevo Amenity/i)).not.toBeInTheDocument();
        expect(screen.queryByText(/Ver reservas por Hotel/i)).not.toBeInTheDocument();
        expect(screen.queryByText(/Ver reservas por usuario/i)).not.toBeInTheDocument();
    });

})