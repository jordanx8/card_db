import React, { useState } from 'react';
import ToggleButton from 'react-bootstrap/ToggleButton';

function CheckBoxButton(props) {

    function buttonHandler(e) {
        props.setChecked(e.currentTarget.checked)
        if(e.currentTarget.checked){
            props.setOriginalData(props.data)
            props.setData(props.data.filter(props.filter))
        } else{
            props.setData(props.originalData)
        }
    }

    return (
        <ToggleButton
            id="toggle-check"
            type="checkbox"
            variant="outline-primary"
            checked={props.checked}
            value="1"
            onChange={buttonHandler}>
            {props.name}
        </ToggleButton>
    )
}

export default CheckBoxButton