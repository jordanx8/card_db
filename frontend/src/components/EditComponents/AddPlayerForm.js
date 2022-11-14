import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

function AddPlayerForm() {
    return (
        <>
            <h1>Add New Player:</h1>
            <Form>
                <Form.Group className="mb-3" controlId="formFirstName">
                    <Form.Label>First Name</Form.Label>
                    <Form.Control placeholder="Enter first name" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formLastName">
                    <Form.Label>Last Name</Form.Label>
                    <Form.Control placeholder="Enter last name" />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formSeasonsPlayed">
                    <Form.Label>Season's Played</Form.Label>
                    <Form.Control placeholder="Enter season's played" />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddPlayerForm