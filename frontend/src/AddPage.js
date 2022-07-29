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
    return (
        <Form
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
                    <Search
                        id="search-1"
                        labelText="Search Players"
                        placeholder="Search Players"
                    />
                    <TextInput
                        id="one"
                        labelText="Season"
                    />
                    <TextInput
                        id="two"
                        labelText="Manufacturer"
                    />
                    <TextInput
                        id="three"
                        labelText="Set"
                    />
                    <TextInput
                        id="four"
                        labelText="Insert"
                    />
                    <TextInput
                        id="five"
                        labelText="Parallel"
                    />
                    <TextInput
                        id="six"
                        labelText="Notes"
                    />
                    <Button>
                        Submit
                    </Button>
                </Stack>
            </FormGroup>
        </Form>
    )
}

export default AddPage;