import React, { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { LoginContext, UserProfileContext } from '../../App';
import { useParams } from "react-router-dom";
import "./UserDetails.css"
import Navbar from "../NavBar/NavBar";

const UserDetails = () => {
  const { id } = useParams();
  const [user, setUser] = useState(null);
  const [baseURL, setBaseURL] = useState('');
  const [error, setError] = useState(null);
  const { loggedIn } = useContext(LoginContext);
  const { userProfile } = useContext(UserProfileContext);
  const navigate = useNavigate()

  useEffect(() => {
    fetch('/config.json')
        .then(response => response.json())
        .then(data => {
          setBaseURL(data.apiUrl);
        })
        .catch(error => {
          console.error('Error loading config:', error);
          setError('Failed to load configuration');
        });
  }, []);

  useEffect(() => {
    const fetchUserDetails = async () => {
      if (baseURL) {
        try {
          const response = await fetch(`${baseURL}/user/${id}`);
          if (response.ok) {
            const data = await response.json();
            setUser(data);
          } else {
            const errorData = await response.json();
            throw new Error(errorData.error);
          }
        } catch (error) {
          setError(error.message);
        }
      }
    };

    fetchUserDetails();
  }, [id, baseURL]);

  if (error) {
    return (
        <>
          <Navbar />
          <p className="fullscreen">{error}</p>
        </>
    );
  }

  if (!user) {
    return <div>Loading...</div>;
  }

  if (!loggedIn || userProfile.role !== "Admin") {
    return (
      <>
        <Navbar />
        <p className="fullscreen">No puedes acceder a este sitio.</p>
      </>
    );
  }

  return (
    <>
      <Navbar />
      <div className="UserDetail">
        <h3>Perfil de Usuario</h3>
        <p>Nombre: {user.name}</p>
        <p>Apellido: {user.last_name}</p>
        <p>DNI: {user.dni}</p>
        <p>Email: {user.email}</p>
        <p>Número de usuario: {user.id}</p>
        <button onClick={() => navigate(`/user/reservations/${user.id}`)}>Ver Reservas</button>
      </div>
    </>
  );
};

export default UserDetails;