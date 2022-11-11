import DropdownButton from "./DropdownButton";
import CardRow from "./CardRow";
import React, { useState } from 'react';
import { v4 as uuidv4 } from 'uuid';

function SeasonRow(props) {
    const [dropdown, setDropdown] = useState(false);
    return (
        <React.Fragment key={uuidv4()}>
            <tr>
                <th colSpan={5}>{props.season}</th>
                <th><DropdownButton state={dropdown} setState={setDropdown} /></th>
            </tr>
            {dropdown &&
        <React.Fragment key={uuidv4()}>
                    <tr>
                        <th>Manufacturer</th>
                        <th>Set</th>
                        <th>Insert</th>
                        <th>Parallel</th>
                        <th>Card Number</th>
                        <th>Notes</th>
                    </tr>
                    {props.cardData.map(card => (
                        <CardRow name={card.playerName}
                            season={card.season}
                            manufacturer={card.manufacturer}
                            set={card.set}
                            insert={card.insert}
                            parallel={card.parallel}
                            cardNumber={card.cardNumber}
                            notes={card.notes}
                        />
                    ))
                    }
        </React.Fragment>
            }
        </React.Fragment>
    );
}

export default SeasonRow;
