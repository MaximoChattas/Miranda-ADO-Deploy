import React, { useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { LoginContext, UserProfileContext } from '../../App';
import Navbar from '../NavBar/NavBar';
import './LoadAmenity.css';

function LoadAmenity() {
    const [name, setName] = useState('');

    const { loggedIn } = useContext(LoginContext);
    const { userProfile } = useContext(UserProfileContext);

    const [error, setError] = useState('');
    const navigate = useNavigate();

    const baseURL = import.meta.env.VITE_base_url

    const handleLoadAmenity = async (e) => {
        e.preventDefault();
        setError('');

        try {
            if (!name) {
                throw new Error('Complete todos los campos requeridos');
            }

            const response = await fetch(`${baseURL}/amenity`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name
                }),
            });

            if (response.status === 201) {
                navigate('/');
            } else {
                const data = await response.json();
                const errorMessage = data.error || 'Error';
                throw new Error(errorMessage);
            }
        } catch (error) {
            console.error(error);
            setError(error.message);
        }
    };


    if (!loggedIn || userProfile.role !== 'Admin') {
        return (
            <>
                <Navbar />
                <p>No puedes acceder a este sitio.</p>
            </>
        );
    }

    return (
        <>
            <Navbar />
            <div className="contenedorLoad">
                <h2>Cargar Amenity</h2>
                <form onSubmit={handleLoadAmenity}>
                    <div>
                        <label>Nombre:</label>
                        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
                    </div>

                    {error && <p className="error-message">{error}</p>}
                    <button type="submit">Cargar Amenity</button>
                </form>
            </div>
        </>
    );
}

export default LoadAmenity;
