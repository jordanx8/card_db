import React, { useState } from 'react';
import ToggleButton from 'react-bootstrap/ToggleButton';

function FilterButton({name, filter, data, setData}) {
    const [originalData, setOriginalData] = useState([])
    const [checked, setChecked] = useState(false);

    function buttonHandler(e) {
        setChecked(e.currentTarget.checked)
        if(e.currentTarget.checked){
            setOriginalData(data)
            setData(data.filter(filter))
        } else{
            setData(originalData)
        }
    }

    return (
        <ToggleButton
            id="toggle-check"
            type="checkbox"
            variant="outline-primary"
            checked={checked}
            value="1"
            onChange={buttonHandler}>
            {name}
        </ToggleButton>
    )
}

export default FilterButton