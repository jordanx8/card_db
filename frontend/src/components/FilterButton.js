import React from 'react';
import ToggleButton from 'react-bootstrap/ToggleButton';

function FilterButton({ name, state, setState }) {
    function buttonHandler(e) {
        setState(e.currentTarget.checked)
    }

    return (
        <ToggleButton
            id={"toggle+" + name}
            type="checkbox"
            variant="outline-primary"
            checked={state}
            value="1"
            onChange={buttonHandler}>
            {name}
        </ToggleButton>
    )
}

export default FilterButton