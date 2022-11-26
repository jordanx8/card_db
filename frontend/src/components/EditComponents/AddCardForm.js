import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';
import ListDropdown from './ListDropdown';
import * as playerData from '../../test_player_data.json';
import { filterPlayerListDownTo } from '../../util/util'

function AddCardForm() {
    const [players, setPlayers] = useState(playerData.default);
    const [formState, setFormState] = useState({
        playerName: players[0].firstName + " " + players[0].lastName,
        season: players[0].seasons[0],
        manufacturer: "",
        set: "",
        parallel: "",
        cardNumber: "",
        imageLink: "",
    });
    function handleInputChange(event) {
        console.log(formState)
        const value = event.target.value
        const name = event.target.id
        if(name === "playerName")
        {
            setFormState({
                ...formState,
                [name]: value,
                ["season"]: players.filter(filterPlayerListDownTo(value))[0].seasons[0]
            });
        } else {
            setFormState({
                ...formState,
                [name]: value
            });
        }
    }
    const [selectedPlayer, setSelectedPlayer] = useState(players[0].firstName + " " + players[0].lastName);
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

    return (
        <>
            <h1>Add New Card:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="playerName">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select value={selectedPlayer} onChange={PlayerDropdownHandler}>
                        {players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="season">
                    <Form.Label>Select a Season</Form.Label>
                    <Form.Select>
                        {/* TODO: make dropdown not stay empty if select first player */}
                        {players.filter(filterPlayerListDownTo(selectedPlayer))[0].seasons.map(season => (
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