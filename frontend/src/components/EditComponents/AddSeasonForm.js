import Button from 'react-bootstrap/Button';
import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import ListDropdown from './ListDropdown';
import * as playerData from '../../test_player_data.json';

function AddSeasonForm() {
    const [players, setPlayers] = useState(playerData.default);

    const [formState, setFormState] = useState({
        formPlayer: players[0].firstName + " " + players[0].lastName,
        formSeason: "",
    });

    function handleInputChange(event) {
        const value = event.target.value
        const name = event.target.id
        console.log(formState)
        setFormState({
            ...formState,
            [name]: value
        });
    }

    return (
        <>
            <h1>Add Season to Existing Player:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formPlayer">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select>
                        {players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formSeason">
                    <Form.Label>Season</Form.Label>
                    <Form.Control placeholder="Enter seasons to add to player (comma separated)" />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddSeasonForm