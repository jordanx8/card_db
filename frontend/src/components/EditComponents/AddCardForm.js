import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';
import ListDropdown from './ListDropdown';
import * as playerData from '../../test_player_data.json';
import { filterPlayerListDownTo } from '../../util/util'

function AddCardForm() {
    const [players, setPlayers] = useState(playerData.default);
    const [selectedPlayer, setSelectedPlayer] = useState("");
    const [autoChecked, setAutoChecked] = useState(false);
    const [patchChecked, setPatchChecked] = useState(false);
    const [wChecked, setWChecked] = useState(false);
    const [numChecked, setNumChecked] = useState(false);


    function PlayerDropdownHandler(event) {
        setSelectedPlayer(event.target.value);
        console.log(players)
    }

    function AutoHandler(event) {
        setAutoChecked(!autoChecked)
    }

    function PatchHandler(event) {
        setPatchChecked(!patchChecked)
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
                <Form.Group className="mb-3">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select value={selectedPlayer} onChange={PlayerDropdownHandler}>
                        {players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Select a Season</Form.Label>
                    <Form.Select>
                        {/* TODO: make dropdown not stay empty if select first player */}
                        {(selectedPlayer !== "") && players.filter(filterPlayerListDownTo(selectedPlayer))[0].seasons.map(season => (
                            <ListDropdown text={season} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formManu">
                    <Form.Label>Manufacturer</Form.Label>
                    <Form.Control placeholder="probably Panini" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formSet">
                    <Form.Label>Set</Form.Label>
                    <Form.Control placeholder="Enter card's set" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formParallel">
                    <Form.Label>Parallel</Form.Label>
                    <Form.Control placeholder="Enter parallel" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formCardNumber">
                    <Form.Label>Card Number</Form.Label>
                    <Form.Control placeholder="Enter card number" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formCardNumber">
                    <Form.Label>Image Link</Form.Label>
                    <Form.Control placeholder="Enter image link" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formNotes">
                    <Form.Check inline type="checkbox" label="RC" />
                    <Form.Check inline type="checkbox" label="Sticker" />
                    <Form.Check inline type="checkbox" label="College" />
                    <Form.Check value={wChecked} onChange={wHandler} type="checkbox" label="w/ Others" />
                    {wChecked &&
                        <Form.Control placeholder="who? (comma-separated)" />}
                    <Form.Check value={numChecked} onChange={numberHandler} type="checkbox" label="#'d" />
                    {numChecked &&
                        <Form.Control placeholder="Enter serial number" />}
                    <Form.Check value={autoChecked} onChange={AutoHandler} type="checkbox" label="Auto" />
                    {autoChecked && <>
                        <Form.Check inline name="autogroup" type="radio" label="Sticker Auto" />
                        <Form.Check inline name="autogroup" type="radio" label="On-Card Auto" /></>}
                    <Form.Check value={patchChecked} onChange={PatchHandler} type="checkbox" label="Patch" />
                    {patchChecked && <>
                        <Form.Check inline name="autogroup" type="radio" label="Game-Worn" />
                        <Form.Check inline name="autogroup" type="radio" label="Player-Worn" /></>}
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddCardForm