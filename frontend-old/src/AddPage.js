import {
    Stack,
    TextInput,
    FormGroup,
    Button,
} from '@carbon/react';
import './AddPage.scss';
import { useState } from 'react';
import { ADD_CARD } from './AddCardMutation';
import { useMutation } from '@apollo/client';

function AddPage() {
    function submitCardAddition() {
        console.log("function")
    }

    const [addCardHandler, { data, loading, error }] = useMutation(ADD_CARD)

    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [seasonsPlayed, setSeasonsPlayed] = useState("");
    const [season, setSeason] = useState("");
    const [manufacturer, setManufacturer] = useState("");
    const [set, setSet] = useState("");
    const [insert, setInsert] = useState("");
    const [parallel, setParallel] = useState("");
    const [cardNumber, setCardNumber] = useState("");
    const [notes, setNotes] = useState("");

    return (
        <form
            onSubmit={submitCardAddition}
            style={{
                background: 'black',
            }}>
            <FormGroup
                legendId="formgroup-legend-id"
                legendText="Add Card"
                style={{
                }}
            >
                <Stack gap={7}>
                    <TextInput
                        id="one"
                        labelText="FirstName"
                        onChange={(value) => setFirstName(value.target.value)}
                    />
                    <TextInput
                        id="two"
                        labelText="LastName"
                        onChange={(value) => setLastName(value.target.value)}
                    />
                    <TextInput
                        id="three"
                        labelText="Seasons Played"
                        onChange={(value) => setSeasonsPlayed(value.target.value)}
                    />
                    <TextInput
                        id="four"
                        labelText="Season"
                        onChange={(value) => setSeason(value.target.value)}
                    />
                    <TextInput
                        id="five"
                        labelText="Manufacturer"
                        onChange={(value) => setManufacturer(value.target.value)}
                    />
                    <TextInput
                        id="six"
                        labelText="Set"
                        onChange={(value) => setSet(value.target.value)}
                    />
                    <TextInput
                        id="seven"
                        labelText="Insert"
                        onChange={(value) => setInsert(value.target.value)}
                    />
                    <TextInput
                        id="eight"
                        labelText="Parallel"
                        onChange={(value) => setParallel(value.target.value)}
                    />
                    <TextInput
                        id="nine"
                        labelText="Card Number"
                        onChange={(value) => setCardNumber(value.target.value)}
                    />
                    <TextInput
                        id="ten"
                        labelText="Notes"
                        onChange={(value) => setNotes(value.target.value)}
                    />
                    <Button
                        onClick={() => addCardHandler({ variables: { firstName, lastName, seasonsPlayed, season, manufacturer, set, insert, parallel, cardNumber, notes } })}>
                        Submit
                    </Button>
                </Stack>
            </FormGroup>
        </form>
    )
}

export default AddPage;