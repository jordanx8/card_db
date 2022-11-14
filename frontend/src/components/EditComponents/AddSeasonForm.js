import Button from 'react-bootstrap/Button';
import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import ListDropdown from './ListDropdown';
import * as playerData from '../../test_player_data.json';

function AddSeasonForm() {
    const [players, setPlayers] = useState(playerData.default);

    return (
        <>
            <h1>Add Season to Existing Player:</h1>
            <Form>
                <Form.Group className="mb-3">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select>
                        {players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formSeason">
                    <Form.Label>Season</Form.Label>
                    <Form.Control placeholder="Enter season to add to player" />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddSeasonForm