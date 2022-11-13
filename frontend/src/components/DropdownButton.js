import Button from 'react-bootstrap/Button';

function DropdownButton({ state, setState }) {
    function useStateHandler() {
        setState(!state)
    }

    if (!state) {
        return (
            <Button onClick={useStateHandler}>Expand</Button>
        );
    } else {
        return (
            <Button onClick={useStateHandler}>Collapse</Button>
        );
    }

}

export default DropdownButton;
