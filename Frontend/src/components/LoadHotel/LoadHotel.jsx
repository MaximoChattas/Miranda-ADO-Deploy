import React, { useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { LoginContext, UserProfileContext } from '../../App';
import Navbar from '../NavBar/NavBar';
import './LoadHotel.css';

function LoadHotel() {
    const [name, setName] = useState('');
    const [street_name, setStreet_name] = useState('');
    const [street_number, setStreet_number] = useState('');
    const [room_amount, setRoom_amount] = useState('');
    const [rate, setRate] = useState('');
    const [description, setDescription] = useState('');
    const [amenities, setAmenities] = useState([]);
    const [selectedAmenities, setSelectedAmenities] = useState([]);
    const [images, setImages] = useState([]);

    const [hotelId, setHotelId] = useState('');
    const [isLoaded, setIsLoaded] = useState(false);

    const { loggedIn } = useContext(LoginContext);
    const { userProfile } = useContext(UserProfileContext);

    const [error, setError] = useState('');

    const navigate = useNavigate();

    const [baseURL, setBaseURL] = useState('');

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

    const handleLoadHotel = async (e) => {
        e.preventDefault();
        setError('');

        try {
            if (!name || !street_name || !street_number || !room_amount || !rate) {
                throw new Error('Complete todos los campos requeridos');
            }

            const response = await fetch(`${baseURL}/hotel`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name,
                    street_name,
                    street_number: parseInt(street_number),
                    room_amount: parseInt(room_amount),
                    rate: parseFloat(rate),
                    description,
                    amenities: selectedAmenities,
                }),
            });

            if (response.status === 201) {
                setIsLoaded(true);
                const data = await response.json();
                setHotelId(data.id);
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

    const handleImagesUpload = async (e) => {
        e.preventDefault();

        try {
            const formData = new FormData();
            images.forEach((image) => {
                formData.append('images', image);
            });

            const response = await fetch(`${baseURL}/hotel/${hotelId}/images`, {
                method: 'POST',
                body: formData,
            });

            if (response.ok) {
                navigate('/');
            } else {
                const errorData = await response.json();
                throw new Error(errorData.error);
            }
        } catch (error) {
            console.error(error);
            setError(error.message);
        }
    };

    const handleAmenityChange = (e, amenityName) => {
        if (e.target.checked) {
            setSelectedAmenities([...selectedAmenities, amenityName]);
        } else {
            setSelectedAmenities(selectedAmenities.filter((name) => name !== amenityName));
        }
    };

    useEffect(() => {
        const fetchAmenities = async () => {
            try {
                const response = await fetch(`${baseURL}/amenity`);
                if (response.ok) {
                    const data = await response.json();
                    setAmenities(data);
                } else {
                    throw new Error('Error fetching amenities');
                }
            } catch (error) {
                console.error(error);
                setError('Error fetching amenities');
            }
        };
        if (baseURL) {
            fetchAmenities();
        }
    }, [baseURL]);

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
            <div className="contenedorLoadHotel">
                <h2>Cargar Hotel</h2>
                <form onSubmit={handleLoadHotel}>
                    <div>
                        <label>Nombre:</label>
                        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
                    </div>
                    <div>
                        <label>Calle:</label>
                        <input type="text" value={street_name} onChange={(e) => setStreet_name(e.target.value)} />
                    </div>
                    <div>
                        <label>Número:</label>
                        <input type="number" value={street_number} onChange={(e) => setStreet_number(e.target.value)} />
                    </div>
                    <div>
                        <label>Cantidad de habitaciones:</label>
                        <input type="number" value={room_amount} onChange={(e) => setRoom_amount(e.target.value)} />
                    </div>
                    <div>
                        <label>Rate:</label>
                        <input type="number" step="0.1" value={rate} onChange={(e) => setRate(e.target.value)} />
                    </div>
                    <div>
                        <label>Descripción:</label>
                        <textarea value={description} onChange={(e) => setDescription(e.target.value)}></textarea>
                    </div>
                    <div>
                        <label>Amenities:</label>
                        {amenities.map((amenity, index) => (
                            <div key={index}>
                                <input
                                    type="checkbox"
                                    onChange={(e) => handleAmenityChange(e, amenity.name)}
                                />
                                {amenity.name}
                            </div>
                        ))}
                    </div>

                    {error && <p className="error-message">{error}</p>}
                    <button type="submit">Cargar Hotel</button>
                </form>

                {isLoaded && (
                    <form onSubmit={handleImagesUpload}>
                        <div>
                            <label>Subir Imágenes:</label>
                            <input
                                type="file"
                                multiple
                                onChange={(e) => setImages([...e.target.files])}
                            />
                        </div>
                        <button type="submit">Cargar Imágenes</button>
                    </form>
                )}
            </div>
        </>
    );
}

export default LoadHotel;
