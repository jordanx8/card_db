import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';
import ListDropdown from './ListDropdown';
import { useQuery, useMutation } from "@apollo/client";
import { GET_PLAYERS_QUERY } from '../../util/queries';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { filterPlayerListDownTo } from '../../util/util'

function AddCardForm() {
    const [formState, setFormState] = useState({
        firstName: "",
        lastName: "",
        season: "",
        manufacturer: "",
        set: "",
        parallel: "",
        cardNumber: "",
        imageLink: "",
    });

    const { data: playersData, loading: playersLoading, error: playersError } = useQuery(GET_PLAYERS_QUERY, {
        onCompleted: (data) => {
            console.log(data)
            setSelectedPlayer(data.players[0].firstName + " " + data.players[0].lastName)
            setFormState({
                ...formState,
                "firstName": data.players[0].firstName,
                "lastName": data.players[0].lastName,
                "season": data.players[0].seasons[0],
            });
        }
    });

    function handleInputChange(event) {
        const value = event.target.value
        const name = event.target.id
        if (name === "playerName") {
            const names = value.split(" ");
            setFormState({
                ...formState,
                "firstName": names[0],
                "lastName": names[1],
                "season": playersData.players.filter(filterPlayerListDownTo(value))[0].seasons[0]
            });
        } else {
            setFormState({
                ...formState,
                [name]: value
            });
        }
    }
    const [selectedPlayer, setSelectedPlayer] = useState("");
    const [wChecked, setWChecked] = useState(false);
    const [numChecked, setNumChecked] = useState(false);


    function PlayerDropdownHandler(event) {
        setSelectedPlayer(event.target.value);
    }

    function numberHandler(event) {
        setNumChecked(!numChecked)
    }

    function wHandler(event) {
        setWChecked(!wChecked)
    }

    if (playersLoading) return "Loading...";
    if (playersError) return <pre>{playersError.message}</pre>

    return (
        <>
            <ToastContainer
                position="top-right"
                draggable
                autoClose={5000}
                closeOnClick
            />
            <h1>Add New Card:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="playerName">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select value={selectedPlayer} onChange={PlayerDropdownHandler}>
                        {playersData.players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="season">
                    <Form.Label>Select a Season</Form.Label>
                    <Form.Select>
                        {/* TODO: make dropdown not stay empty if select first player */}
                        {playersData.players.filter(filterPlayerListDownTo(selectedPlayer))[0].seasons.map(season => (
                            <ListDropdown text={season} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="manufacturer">
                    <Form.Label>Manufacturer</Form.Label>
                    <Form.Control placeholder="probably Panini" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="set">
                    <Form.Label>Set</Form.Label>
                    <Form.Control placeholder="Enter card's set" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="parallel">
                    <Form.Label>Parallel</Form.Label>
                    <Form.Control placeholder="Enter parallel" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="cardNumber">
                    <Form.Label>Card Number</Form.Label>
                    <Form.Control placeholder="Enter card number" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="imageLink">
                    <Form.Label>Image Link</Form.Label>
                    <Form.Control placeholder="Enter image link" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formNotes">
                    <Form.Check inline type="checkbox" label="RC" />
                    <Form.Check inline type="checkbox" label="Sticker" />
                    <Form.Check inline type="checkbox" label="College" />
                    <Form.Check inline type="checkbox" label="Pre-RC" />
                    <Form.Check inline type="checkbox" label="Auto" />
                    <Form.Check inline type="checkbox" label="Patch" />
                    <Form.Check value={wChecked} onChange={wHandler} type="checkbox" label="w/ Others" />
                    {wChecked &&
                        <Form.Control placeholder="who? (comma-separated)" />}
                    <Form.Check value={numChecked} onChange={numberHandler} type="checkbox" label="#'d" />
                    {numChecked &&
                        <Form.Control placeholder="Enter serial number" />}
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddCardForm