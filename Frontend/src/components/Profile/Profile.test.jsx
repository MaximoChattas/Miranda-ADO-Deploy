import { describe, it } from "vitest";
import { expect } from "vitest";
import { screen, render } from "@testing-library/react";
import { LoginContext, UserProfileContext   } from "../../App";
import { BrowserRouter as Router } from "react-router-dom";
import Profile from "./Profile";

describe('Profile', () => {

    const renderWithContext = (loggedIn = true, userProfile = 
        {id: 1, name: 'John', last_name: 'Doe', dni: '123456', email: 'johndoe@email.com', role: 'Customer'}) => {
        render(
            <UserProfileContext.Provider value={ {userProfile} }>
                <LoginContext.Provider value={{ loggedIn }}>
                    <Router>
                        <Profile />
                    </Router>
                </LoginContext.Provider>
            </UserProfileContext.Provider>
        );
    };

    it('should render not allowed message', () => {
        renderWithContext(false);
        expect(screen.getByText(/No puedes acceder a este sitio./i)).toBeInTheDocument();
    });

    it('should render profile data', () => {
        renderWithContext();
        expect(screen.getByText(/Perfil de Usuario/i)).toBeInTheDocument();
        expect(screen.getByText(/Nombre: John/i)).toBeInTheDocument();
        expect(screen.getByText(/Apellido: Doe/i)).toBeInTheDocument();
        expect(screen.getByText(/DNI: 123456/i)).toBeInTheDocument();
    });

    it('should display customer buttons when logged as customer', () => {
        renderWithContext();
        const myReservationsButton = screen.getByText(/Mis Reservas/i);
        expect(myReservationsButton).toBeInTheDocument();
        myReservationsButton.click();
        expect(window.location.pathname).toBe("/user/reservations/1");

        const myRangeReservationsButton = screen.getByText(/Reservas por Rango/i);
        expect(myRangeReservationsButton).toBeInTheDocument();
        myRangeReservationsButton.click();
        expect(window.location.pathname).toBe("/user/reservations/range");
    });

    it('should not display customer buttons when logged as admin', () => {
        renderWithContext(true,  
            {id: 1, name: 'John', last_name: 'Doe', dni: '123456', email: 'johndoe@email.com', role: 'Admin'});
        
        expect(screen.queryByText(/Mis Reservas/i)).not.toBeInTheDocument();
        expect(screen.queryByText(/Reservas por Rango/i)).not.toBeInTheDocument();
    });
})