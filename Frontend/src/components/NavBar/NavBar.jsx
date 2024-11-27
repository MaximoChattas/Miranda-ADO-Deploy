import React, { useContext, useEffect } from "react";
import { Link, NavLink } from "react-router-dom";
import { LoginContext, UserProfileContext } from '../../App';
import'./NavBar.css';

function Navbar() {
  const { loggedIn, setLoggedIn } = useContext(LoginContext);
  const { userProfile, setUserProfile } = useContext(UserProfileContext);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const userProfileData = localStorage.getItem('userProfile');

    if (token && userProfileData) {
      setLoggedIn(true);
      setUserProfile(JSON.parse(userProfileData));
    }
  }, []);

  return (
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA HOTELS</h1>
        </NavLink>
        <div className="contenedorBotones">
           <NavLink className="nav-link" to="/hotel/availability">
              <button className="boton">Ver Disponibilidad</button>
            </NavLink>
          {loggedIn ? (
            <NavLink className="nav-link" to="/profile">
              <button className="boton">Hola {userProfile.name}</button>
            </NavLink>
          ) : (
            <NavLink className="nav-link" to="/login">
              <button className="boton">Iniciar sesion</button>
            </NavLink>
          )}
        </div>
      </div>
    </header>
  );
}


export default Navbar;
