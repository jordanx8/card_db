import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';
import Container from 'react-bootstrap/Container';

function Navigation() {
    return (
        <Navbar bg="dark" variant="dark" expand="lg">
            <Container>
                <Navbar.Brand >Card-DB</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href="/">New Orleans Pelicans</Nav.Link>
                        <Nav.Link href="/tigers">LSU Tigers</Nav.Link>
                        <NavDropdown title="Edit" id="basic-nav-dropdown">
                            <NavDropdown.Item href="/add/card">Add new card</NavDropdown.Item>
                            <NavDropdown.Item href="/add/player">Add new player</NavDropdown.Item>
                        </NavDropdown>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}

export default Navigation