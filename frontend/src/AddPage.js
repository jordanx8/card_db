import {
    Stack,
    TextInput,
    FormGroup,
    Form,
    Button,
    Search,
    Grid,
    Row,
    Column
} from '@carbon/react';
import './AddPage.scss';

function AddPage() {
    function submitCardAddition(){
        console.log("function")
    }

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
                        onChange={console.log("")}
                    />
                    <TextInput
                        id="two"
                        labelText="LastName"
                    />
                    <TextInput
                        id="three"
                        labelText="Seasons Played"
                    />
                    <TextInput
                        id="four"
                        labelText="Season"
                    />
                    <TextInput
                        id="five"
                        labelText="Manufacturer"
                    />
                    <TextInput
                        id="six"
                        labelText="Set"
                    />
                    <TextInput
                        id="seven"
                        labelText="Insert"
                    />
                    <TextInput
                        id="eight"
                        labelText="Parallel"
                    />
                    <TextInput
                        id="nine"
                        labelText="Notes"
                    />
                    <Button type="submit">
                        Submit
                    </Button>
                </Stack>
            </FormGroup>
        </form>
    )
}

export default AddPage;