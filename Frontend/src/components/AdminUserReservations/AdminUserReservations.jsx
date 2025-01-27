import React, { useEffect, useState, useContext } from "react";
import { LoginContext, UserProfileContext } from '../../App';
import { Link } from "react-router-dom";
import Navbar from "../NavBar/NavBar";

const AdminUserReservations = () => {
  const [userReservations, setUserReservations] = useState({ reservations: [] });
  const [users, setUsers] = useState([]);
  const [error, setError] = useState(null);
  const [baseURL, setBaseURL] = useState('');
  const { userProfile } = useContext(UserProfileContext);
  const { loggedIn } = useContext(LoginContext);

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
    if (baseURL) {
      const fetchUserReservations = async () => {
        try {
          const response = await fetch(`${baseURL}/reservation`);
          if (response.ok) {
            const data = await response.json();
            setUserReservations({ reservations: data });

            const userResponse = await fetch(`${baseURL}/user`);
            if (userResponse.ok) {
              const userData = await userResponse.json();
              setUsers(userData);
            } else {
              const errorData = await userResponse.json();
              throw new Error(errorData.error);
            }
          } else {
            const errorData = await response.json();
            throw new Error(errorData.error);
          }
        } catch (error) {
          setError(error.message);
        }
      };

      fetchUserReservations();
    }
  }, [baseURL]);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!userReservations) {
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
      <h2>Reservas</h2>

      <div className="containerReservations">
      <ul className="list-group">
        {users.map(user => {
          const filteredReservations = userReservations.reservations || [];
          const userReservationsFiltered = filteredReservations.filter(
            reservation => reservation.user_id === user.id
          );
          if(user.role === "Admin") return null;
          return (
            <li key={user.id} className="list-group-item list-group-item-dark">
              <Link to={`/user/${user.id}`}>
                <h3>{user.name} {user.last_name}</h3>
              </Link>
              {userReservationsFiltered.length > 0 ? (
                <ul className="list-group">
                  {userReservationsFiltered.map(reservation => (
                    <li key={reservation.id} className="list-group-item">
                      <Link to={`/reservation/${reservation.id}`}>
                        <p>Nº Reserva: {reservation.id}</p>
                      </Link>
                      <p>Inicio: {reservation.start_date}</p>
                      <p>Fin: {reservation.end_date}</p>
                      <p>Costo: {reservation.amount}</p>
                      <Link to={`/hotel/${reservation.hotel_id}`}>
                        <p>Nº Hotel: {reservation.hotel_id}</p>
                      </Link>
                    </li>
                  ))}
                </ul>
              ) : (
                <p>Sin reservas.</p>
              )}
            </li>
          );
        })}
      </ul>
      </div>
    </>
  );
};

export default AdminUserReservations;
